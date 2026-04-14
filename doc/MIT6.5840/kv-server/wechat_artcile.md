# MIT 6.5840 Lab 2：Key/Value Server

---

## 1. Lab 2 简介

在这个实验中，我们要做两件事：先实现一个完全没有并发安全的 **Key/Value Server**，再在它上面实现一个分布式锁 **lock** 来保证并发安全。全部做来下感觉是最简单的一个实验，要写的代码也很少（网上都说mapreduce 最简单，然而我觉得它真挺难的，堪比 lab3 Raft 的两个部分了）；它主要是实现了一个只能存取字符串的键值对数据库——涉及处理网络丢包、重试、**at-most-once **和 **linearizable** 这些东西。

lab2 很像是在重复问一句很烦但很现实的话：

> **客户端没收到回复，到底是请求没执行，还是执行了但回复丢了？**

这句话基本涵盖了实验的所有难点。

官方实验文档: [6.5840 Lab 2: Key/Value Server](http://nil.csail.mit.edu/6.5840/2025/labs/lab-kvsrv1.html)

我的完整实现: [MIT-6.5840/src/kvsrv1 at master · Tensort-cat/MIT-6.5840](https://github.com/Tensort-cat/MIT-6.5840/tree/master/src/kvsrv1)

|            名词            |     翻译     |                             语义                             |
| :------------------------: | :----------: | :----------------------------------------------------------: |
|      **at-most-once**      | 最多执行一次 |                   消息可能丢失，但不会重复                   |
|    **linearizability**     |  线性一致性  |            https://zhuanlan.zhihu.com/p/42239873             |
| **CAS (Compare-And-Swap)** |      -       | [阐述CAS的原理CAS（Compare-And-Swap）是一种 **无锁原子操作**，通过硬件指令直接支持多线程环境下 - 掘金](https://juejin.cn/post/7477381239410884617) |
|      **Exactly-Once**      |   精确一次   | 保证每条消息或数据在分布式系统中被精确处理一次，既不会丢失也不会重复（lab 没有实现） |



---

## 2. Lab 2 任务描述

官方给的 KV 接口不多，核心就两个 RPC：`Get(key)` 和 `Put(key, value, version)`。
`Get` 负责读取某个 key 当前的 value 和 version；`Put` 不是无脑覆盖，而是**只有当客户端传来的 version 和服务器当前 version 一致时，才允许写入**。如果 key 不存在，那么只有 `version == 0` 的 `Put` 才能创建这个 key；否则返回 `ErrNoKey`。如果 key 存在但是版本对不上，就返回 `ErrVersion`。这些错误码在官方给的 `rpc.go` 里都已经定义好了。

真正的难点在不可靠网络上。题目要求 Clerk 在 RPC 丢请求或者丢回复时要一直重试；但 `Put` 又是写操作，不能因为重试就重复执行。于是官方专门设计了一个很有意思的返回值：`ErrMaybe`。它的意思不是“失败”，而是**“我不确定这次写到底有没有成功”**。实验第二部分锁的实现，也是在这个带版本号的 KV 之上完成的。

---

## 3. 整体思路

### 3.1 先把 `Put` 看成一个迷你版 CAS

我觉得 lab2 最好理解的方式，就是把 `Put(key, value, version)` 看成一次简化版的 **CAS** (version 可以看作“期待值”)：

> **如果这个 key 现在的 version 还是我刚刚看到的那个版本，那我就改；否则我就不改。**

也就是说，version 在这里不是一个装饰字段，它其实是整套语义的关键。没有它，就没法处理并发写，也没法处理重传后的“最多执行一次”。

### 3.2 服务端思路：一个 map + 一把锁

我的 `server.go` 很直接：用一个 `map[string]Value` 存数据，其中 `Value` 里有两个字段，分别是字符串值和版本号。然后再用 `mutex` 把 `Get` 和 `Put` 包起来。这样每次 RPC 对这个单机 server 来说，都像是在串行执行，保证了系统的 **linearizability**

```go
type Value struct {
    Val string
    Ver rpc.Tversion
}

type KVServer struct {
    mu   sync.Mutex
    data map[string]Value
}
```

### 3.3 客户端负责扛住不可靠网络

服务端这边并没有额外维护“这个请求之前有没有执行过”的表，所以不可靠网络下的核心逻辑，主要落在 Clerk 上：

* `Get`：没收到回复就 sleep 一下再重试；
* `Put`：也是一直重试，但要额外记住自己是不是已经“重发过”了。

这个“是不是重发过”很关键，因为它决定了收到 `ErrVersion` 时，到底该返回 `ErrVersion` 还是 `ErrMaybe`。

### 3.4 锁其实就是一个存储在服务器上的特殊键值对

`lock.go` 的思路很巧妙：直接把锁名 `l` 当成 key，把 value 写成 `"clientId:LOCK"` 或 `"clientId:UNLOCK"`。每个 lock client 在 `MakeLock()` 时生成一个随机 id，所以 value 里天然就能表示“这把锁现在是谁的”。这样一来，`Acquire()` 本质上就是先读锁状态，再尝试做一次条件写。搞开发用过 Redis 的应该对这个技巧很熟悉，我以前写黑马点评的时候也用过类似的技巧。

---

## 4. 代码实现拆解

### 4.1 `Get()`：读当前值和版本号

`Get()` 是整个 lab 里最清爽的一部分。进入临界区后，先去 map 里找 key：

* 找不到：返回 `ErrNoKey`
* 找得到：把 `Value` 和 `Version` 都塞进 reply，返回 `OK`

没有额外分支，也没有什么隐藏状态。因为 `Get` 不修改数据，所以它在重传语义上也好处理得多。就算一个 `Get` 被服务端执行了两遍，结果也不会乱。

```go
func (ck *Clerk) Get(key string) (string, rpc.Tversion, rpc.Err) {
	// You will have to modify this function.
	args := rpc.GetArgs{
		Key: key,
	}
	reply := rpc.GetReply{}
	for {
		ok := ck.clnt.Call(ck.server, "KVServer.Get", &args, &reply)
		if ok {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return reply.Value, reply.Version, reply.Err
}
```



### 4.2 `Put()`：版本对得上才能写

`Put()` 的逻辑稍微绕一点，但代码其实并不长。大概就是这三种情况：

1. key 不存在，而且传进来的 version 不是 0：`ErrNoKey`
2. key 存在，但 version 不匹配：`ErrVersion`
3. 其他情况：写入成功，同时把 version 加 1

成功写入时直接把服务器里的 version 更新成 `args.Version + 1`。也就是说，客户端传来的 version 其实代表“我认知里的旧版本”，而服务器持久化的是“写完之后的新版本”。

```go
func (ck *Clerk) Put(key, value string, version rpc.Tversion) rpc.Err {
	args := rpc.PutArgs{
		Key:     key,
		Value:   value,
		Version: version,
	}
	reply := rpc.PutReply{}
	sentOnce := false

	for {
		ok := ck.clnt.Call(ck.server, "KVServer.Put", &args, &reply)
		if ok {
			if reply.Err == rpc.ErrVersion && sentOnce {
				reply.Err = rpc.ErrMaybe
			}
			break
		}
		sentOnce = true
		time.Sleep(100 * time.Millisecond)
	}

	return reply.Err
}
```

### 4.3 `client.go`：`Get` 可以傻重试，`Put` 不能傻判断

`Get()` 里只要 `Call()` 失败，就等 `100ms` 再发一次，直到收到回复为止。因为 `Get` 不改状态，这样写完全没问题。

但是 `Put()` 要复杂一点：

```go
sentOnce := false

for {
    ok := ck.clnt.Call(ck.server, "KVServer.Put", &args, &reply)
    if ok {
        if reply.Err == rpc.ErrVersion && sentOnce {
            reply.Err = rpc.ErrMaybe
        }
        break
    }
    sentOnce = true
    time.Sleep(100 * time.Millisecond)
}
```

这里 `sentOnce` 表示：**这次拿到的回复，是不是发生在“曾经有过一次无回复重试”之后？**

* 如果第一次 RPC 就收到了 `ErrVersion`，那说明这次 `Put` 肯定没执行成功，直接返回 `ErrVersion`
* 如果前面某次 RPC 没收到回复，后来重试时收到了 `ErrVersion`，那就不能下结论了，只能返回 `ErrMaybe`

我觉得这块是整个 lab2 最有味道的地方：**客户端不是在判断“成没成”，而是在判断“我有没有资格确信它没成”。**

### 4.4 为什么这个设计能做到 at-most-once？

这个实验很妙的一点是：它没有让服务端保存“请求 ID -> 是否执行过”的表，而是借助 version 本身，把“同一个 Put 最多成功一次”做出来了。

原因很简单。同一个 `Put` 重传时带的 version 是一样的：

* 如果第一次已经成功了，服务器里的 version 会加 1
* 那么第二次再收到同样的 `Put`，version 肯定对不上
* 于是服务器就会返回 `ErrVersion`，而不会再执行一次写入

所以它做不到 exactly-once，但它能做到：**这次写最多只会真正生效一次**。至于客户端能不能百分之百知道结果，那就是 `ErrMaybe` 要解决的事了。题目本身也明确说了，这一版并没有做到 exactly-once；想做到那一步，通常要在服务端为每个 Clerk 额外维护状态，这也是我觉得这个实验简单的一个原因。

### 4.5 `Acquire()`：发现锁空闲就抢锁

`lock.go` 里我给每个客户端生成了一个随机 id，然后把锁状态编码成：

* `id:LOCK`
* `id:UNLOCK`

`Acquire()` 的循环逻辑：

1. 先 `Get(lockName)` 读当前状态
2. 如果已经是自己的 `id:LOCK`，直接返回
3. 如果看到的是 `UNLOCK`，或者这个 key 还不存在，就尝试 `Put(lockName, want, ver)`

```go
// 锁在服务端的存储格式: lockName -> "id:LOCK"
func (lk *Lock) Acquire() {
	// Your code here
	lockName := lk.l
	want := lk.id + ":" + LOCK
	/*
		锁被LOCK，就不停get直到UNLOCK(某客户端释放锁)
		发现UNLOCK，尝试put锁，此时因为并发环境会导致竞争，竞争失败则继续不断get直到下次put
		第二个if中的val == ""是因为一开始服务器中没有锁，由第一个使用锁的客户端扮演创建锁的角色 (梦回操作系统的读者写者问题)
		注：golang 的零值真是一个伟大的设计
	*/
	for {
		val, ver, _ := lk.ck.Get(lockName)
		if val == want {
			return
		}
		if strings.HasSuffix(val, UNLOCK) || val == "" {
			lk.ck.Put(lockName, want, ver)
		}
	}
}
```

注意 `val == ""` 这个小细节。因为 key 一开始在服务端根本不存在，而 `Get` 在 `ErrNoKey` 的情况下，reply 里的 `Value` 和 `Version` 会自然保持 Go 的零值，也就是 `""` 和 `0`。于是“第一次创建这把锁”这件事，就被我顺手揉进统一逻辑里了。只能说，**golang 的零值有时候真挺香。**

### 4.6 `Release()`：先拿 version，再把状态改成 `UNLOCK`

`Release()` 就简单了。因为 `Put` 需要正确 version，所以还是得先 `Get()` 一次拿到当前版本号，然后再写回 `id:UNLOCK`：

```go
func (lk *Lock) Release() {
	// Your code here
	/*
		由于一个客户端的Release()一定发生在Acquire()之后, 现在锁一定是自己的
		此时Get("l")的结果一定是"id:LOCK", 仅需将其改为"id:UNLOCK"即可
		这个过程不会因为并发让锁被其他客户端抢占，他们都会卡在自己的Acquire()
	*/
	lockName := lk.l
	_, ver, _ := lk.ck.Get(lockName) // 没版本号没法put，只能先get再put
	lk.ck.Put(lockName, lk.id+":"+UNLOCK, ver)
}
```

由于`Release()` 一定发生在自己已经 `Acquire()` 成功之后，所以当前锁应该就在自己手上，别的客户端在看见 `LOCK` 之前也不会去乱写，所以这里的“先读后写”是安全的。

---

## 5. 容易踩坑的点

### 5.1 没收到回复，不等于服务器没执行

请求丢了，客户端没收到回复。
回复丢了，客户端也没收到回复。

从客户端视角看，这两种情况的体验是一样的；但从系统状态看，它们一个是“没执行”，一个是“已经执行了”。所以只要开始重传，你就必须面对“不确定性”。`ErrMaybe` 的存在，本质上就是在承认这种不确定性，而不是假装自己知道答案。

### 5.2 `ErrMaybe` 不是把 `ErrVersion` 换个名字

很多时候最容易想错的一点是：
“既然重传收到的是 `ErrVersion`，那不还是版本冲突吗？”

其实不是的。第一次就收到 `ErrVersion`，那才是真的版本冲突；说明服务器明确告诉你：版本号不匹配，这次写不进去。
但重传之后收到 `ErrVersion`，语义已经变了。它只说明**当前**版本不匹配，并不能说明你之前那次到底有没有成功。于是 Clerk 才必须把它翻译成 `ErrMaybe`，以前的某次没收到回复的 `Put` 操作可能是已经被服务器执行了的。

---

## 6. 总结与收获

![](D:\dev_soft\Go_WorkSpace\doc\MIT6.5840\kv-server\passed.png)

​														受益匪浅~MIT牛逼

虽然这个 lab 实现的键值对服务器功能非常单一和简陋，但他涉及到了很多在分布式系统中常见的不确定事件的处理。lab 难点不是存取键值对本身，而是**在网络不可靠、消息会丢、结果不确定的前提下，保证客户端请求的每个指令至多只执行一次 (at-most-once) 和 线性一致性 (linearizability) **
