# 			MIT 6.5840 Lab 1：MapReduce

---

## 1. Lab 1 简介

在这个实验中，我们会构建一个 **MapReduce** 系统。我们将会实现 **worker** 进程来完成对 **Map** 任务和 **Reduce** 任务的处理，以及 **coordinator** 进程来为 **worker** 分配任务并处理执行失败的 **worker **，就像 **MapReduce** 论文中做的那样。（本实验的 **coordinator** 对应论文中的 **master** ）

实验官方文档：[6.5840 Lab 1: MapReduce](http://nil.csail.mit.edu/6.5840/2025/labs/lab-mr.html)

我对完整实现：[MIT-6.5840/src/mr at master · Tensort-cat/MIT-6.5840](https://github.com/Tensort-cat/MIT-6.5840/tree/master/src/mr)

![](D:\dev_soft\Go_WorkSpace\doc\MIT6.5840\mapreduce\全局工作原理.png)

---

## 2. Lab 1 任务描述

整个实验可以依次归为以下步骤：

* 创建 Map 任务：每个输入文件对应一个 map task；
* Map 任务分发：让 coordinator 把 map 任务发给空闲 worker；
* 处理 Map 任务：worker 最终会根据分配到的原文件生成 reduce task 
* 阶段切换：当 map 全部完成，才能进入 reduce 阶段；
* 创建 Reduce 任务：每一个 map task 的输出文件对应一个 reduce 任务；
* Reduce 任务分发：让 coordinator 把 reduce 任务发给空闲 worker；
* 处理 Reduce 任务：worker 根据分配到的 reduce 任务生成最终结果

lab 文档给出的运行模型也很清楚：系统中只有一个 coordinator，但可以有一个或多个 worker；worker 会不断通过 RPC 请求任务，执行 map / reduce 后再回来汇报完成情况；如果 coordinator 发现某个任务长时间没完成，就应该把它交给别的 worker 继续做。

简单总结一下，其实就是：

1. coordinator 启动时先创建所有 map 任务；
2. worker 不断拉取任务；
3. map worker 读取原始文件，执行 `mapf`，按 `ihash(key) % NReduce` 分桶；
4. 所有 map 结束后，coordinator 才进入 reduce 阶段；
5. reduce worker 收集属于自己分区的中间文件，排序聚合后调用 `reducef`；
6. 所有 reduce 完成后，整个作业结束。 

---

## 3. 整体思路

**coordinator 维护任务状态和任务队列，worker 负责“拉任务—执行—回报”，整个系统严格按 Map -> Reduce -> Finished 三个阶段推进。**

### 3.1 把整个系统分成三个阶段

在 `coordinator.go` 里，先定义了三个状态：

```go
const (
	MapStatus = iota
	ReduceStatus
	Finished
)
```

这三个状态基本就把整个系统的生命周期说完了：先跑 map，再跑 reduce，最后结束。这样做的好处是，很多本来容易混乱的问题，一旦被“阶段化”以后就会清晰很多。比如 reduce 什么时候能发？ **只有所有 map 任务都完成之后，系统状态切到 `ReduceStatus`，reduce 任务才能被创建和分配。** 

### 3.2 用 channel 当任务池，用 map 记录任务状态

coordinator 里有两类核心数据：

* `MapChan` / `ReduceChan`：保存还没被取走的任务；
* `MapFinished` / `ReduceFinished`：记录任务是否已经完成。

也就是说，**channel 负责“还有哪些活没发出去”，状态表负责“这些活现在做完多少了”。**

```go
type Coordinator struct {
	// Your definitions here.
	Status         int          // 当前系统状态
	MapChan        chan *Task   // map任务通道
	ReduceChan     chan *Task   // reduce任务通道
	Nmap           int          // map任务数(实际就是原文本文件数)
	MapCount       int          // 已完成的map任务数
	Nreduce        int          // reduce任务数
	ReduceCount    int          // 已完成的reduce任务数
	Mutex          sync.Mutex   // 经典大锁
	MapAssign      map[int]bool // map任务是否被分配(taskId -> bool)
	MapFinished    map[int]bool // map任务是否完成(taskId -> bool)
	ReduceAssign   map[int]bool // reduce任务是否被分配(taskId -> bool)
	ReduceFinished map[int]bool // reduce任务是否完成(taskId -> bool)
}
```



### 3.3 worker 采用 pull 模式，而不是 push 模式

```go
// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {
	// Your worker implementation here.
	for {
		task := CallPullTask() // 请求任务

		switch task.Type {
		case Map:
			doMap(task, mapf)

		case Reduce:
			doReduce(task, reducef)

		case Wait:
			time.Sleep(5 * time.Second)

		default:
			log.Printf("a worker end\n")
			return
		}
	}
}
```

worker 不会傻等 coordinator 主动推任务，而是间歇的来问：“现在有没有活干？”

这种 **pull 模式** 的好处是 coordinator 比较省心，不需要维护一堆“谁当前空闲、谁该被唤醒”的复杂状态，worker 只要在没拿到任务时睡一会儿再来问就行。`rpc.go` 里也专门定义了 `Map`、`Reduce`、`Wait`、`End` 这几种任务类型，用来表达不同阶段下 worker 该做的事：

```go
// rpc.go
const (
	Map = iota
	Reduce
	Wait
	End
)

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type Task struct {
	Id       int    // 任务id
	Type     int    // 任务类型
	FileName string // 文件名
	Nreduce  int    // reduce任务数
}

type TaskReq struct {
	Success bool
}

type TaskReply struct { // 没什么吊用，单纯填它的参数
	Success bool
}

type FinishedReq struct {
	Id int // 任务id
}

type Reply struct { // 没什么吊用，单纯填它的参数
	Success bool
}
```



---

## 4. 代码实现拆解

### 4.1 coordinator 启动：先把所有 map 任务准备好

在 `MakeCoordinator()` 里，系统初始状态被设成 `MapStatus`，然后调用 `createMapTasks(files)` 把每个输入文件变成一个 map task，塞进 `MapChan` 里：

```go
func (c *Coordinator) createMapTasks(files []string) {
	for i, file := range files {
		mapTask := Task{
			Id:       i,
			Type:     Map,
			FileName: file,
			Nreduce:  c.Nreduce,
		}
		c.MapChan <- &mapTask
		c.MapAssign[i] = false
		c.MapFinished[i] = false
	}
}
```

这里一个输入文件对应一个 map task，也正符合 lab1 文档里“每个输入文件就是一个 split”的意思。至于 reduce task，则不是一开始就创建，而是在所有 map 任务都完成后，才通过 `createReduceTasks()` 统一生成。 

### 4.2 AssignTask()：当前阶段能发什么任务，就发什么任务

worker 每次请求任务，都会通过rpc调用到 coordinator 的 `AssignTask()`：

```go
switch c.Status {
case MapStatus:
	if len(c.MapChan) > 0 {
		tmp := <-c.MapChan
		*reply = *tmp
		c.MapAssign[reply.Id] = true
		go c.timeoutHandler(reply) // 超时处理
	} else {
		reply.Type = Wait
	}
case ReduceStatus:
	if len(c.ReduceChan) > 0 {
		tmp := <-c.ReduceChan
		*reply = *tmp
		c.ReduceAssign[reply.Id] = true
		go c.timeoutHandler(reply)
	} else {
		reply.Type = Wait
	}
case Finished:
	reply.Type = End
}
```

这段逻辑背后的意思其实很简单：

* 如果当前是 map 阶段，就只可能发 map 任务；
* 如果当前是 reduce 阶段，就只可能发 reduce 任务；
* 如果当前阶段没有可派发任务，但系统还没结束，就让 worker 先 `Wait`；
* 如果整个系统结束了，就直接让 worker 退出。

### 4.3 doMap()：读取文件、执行 map、按 reduce 编号分桶

map worker 真正干活的地方在 `doMap()`。

它先读原始文件，然后调用应用层的 `mapf()` 生成 `[]KeyValue`，接着按下面这行代码把中间结果分给不同的 reduce 分区：

```go
rid := ihash(kv.Key) % task.Nreduce
```

这一步非常关键。因为 reduce 任务本质上是“按 key 聚合”，所以同一个 key 必须稳定地落到同一个 reduce 分区里，不然同一个单词的统计结果就会被拆散，后面根本没法正确汇总。

分桶之后，我给每个 reduce 任务分别生成一个文件。这样 reduce 阶段很好读：reduce worker 只要去拿“属于我这个 reduce id 的所有文件”就够了。

```go
tmpName := fmt.Sprintf("mr-%d-%d-%d.txt", prefix, rid, suffix)
...
newName := oldName[:strings.LastIndex(oldName, "-")] + "-reduce.txt"
os.Rename(oldName, newName)
```

这里我还加了一层“先写临时文件，再 rename 转正”的处理。这个思路和 lab 文档的提示是一致的：**不要让别人看到半写完的文件。** 只有文件已经完整写好，才把它变成正式中间结果。 

### 4.4 TaskFinishHandler()：用完成计数推进阶段切换

worker 做完任务后，会调用 `TaskFinishHandler()` 告诉 coordinator：“这份作业我做完了。”

在 map 阶段里，每收到一个未完成的 map 任务回报，就把 `MapCount` 加一；当 `MapCount == Nmap` 时，说明所有 map 都完成了，系统正式切到 `ReduceStatus`，并创建全部 reduce 任务。

```go
// worker 每完成一个任务就通过rpc调用该函数
func (c *Coordinator) TaskFinishHandler(req *FinishedReq, reply *Reply) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	switch c.Status {
	case MapStatus:
		if c.MapFinished[req.Id] {
			reply.Success = false
		} else {
			c.MapFinished[req.Id] = true
			reply.Success = true
			c.MapCount++
			if c.MapCount == c.Nmap {
				// 所有 Map 完成，进入 Reduce 阶段
				c.Status = ReduceStatus
				c.createReduceTasks()
			}
		}
	case ReduceStatus:
		if c.ReduceFinished[req.Id] {
			reply.Success = false
		} else {
			c.ReduceFinished[req.Id] = true
			reply.Success = true
			c.ReduceCount++
			if c.ReduceCount == c.Nreduce {
				c.Status = Finished
			}
		}
	default:
		reply.Success = false
	}
	return nil
}
```

同理，reduce 阶段也是一样的逻辑：所有 reduce 完成后，`Status = Finished`，整个系统结束。

### 4.5 doReduce()：把属于自己的中间文件全读出来，再排序聚合

```go
files, err := filepath.Glob(fmt.Sprintf("*%d-reduce.txt", task.Id))
```

这句代码的意思是：先找到所有属于当前 reduce 编号的中间文件。
随后把这些文件逐行读出来，构造成 `countMap[word] = []string{...}` 这样的聚合结构，再按 key 排序，最后调用 `reducef(word, list)` 生成结果并写入 `mr-out-X`。

完整的 doReduce() ：

```go
func doReduce(task *Task, reducef func(string, []string) string) {
	/*
		1. 通过读取reduce任务文件，调用reducef函数，得到结果
		2. 将结果写mr-out-X.txt文件中，X为reduce任务ID
	*/
	files, err := filepath.Glob(fmt.Sprintf("*%d-reduce.txt", task.Id)) // 读取属于该任务的reduce任务文件
	if err != nil {
		log.Printf("Read dir failed. taskID: %d, err: %v", task.Id, err)
		return
	}

	countMap := make(map[string][]string)
	for _, fileName := range files { // 处理每一个reduce文件
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("Open file failed. fileName: %s, err: %v", fileName, err)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			params := strings.Split(line, " ")
			word, count := params[0], params[1]
			countMap[word] = append(countMap[word], count)
		}
		file.Close()
	}

	// 收集所有 key
	keys := make([]string, 0, len(countMap))
	for k := range countMap {
		keys = append(keys, k)
	}

	// 按字典序排序
	sort.Strings(keys)
	suffix := rand.Int()
	// 与map任务相同，要考虑超时情况，因此要先将数据写入临时文件，然后再根据coordinator的回复，决定是删除临时文件还是重命名为最终结果文件
	tmpFileName := fmt.Sprintf("mr-out-%d-%d.txt", task.Id, suffix)
	tmpOutFile, err := os.OpenFile(tmpFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Open file failed. fileName: %s, err: %v", tmpFileName, err)
		return
	}
	// 按排序后的顺序输出
	for _, word := range keys {
		list := countMap[word]
		wc := reducef(word, list)
		line := word + " " + wc + "\n"
		_, err = tmpOutFile.WriteString(line)
		if err != nil {
			log.Printf("Write file failed. fileName: %s, err: %v",
				tmpFileName, err)
			return
		}
	}
	tmpOutFile.Close()

	// 最后通知coordinator完成了reduce任务
	req := FinishedReq{
		Id: task.Id,
	}
	reply := Reply{
		Success: true,
	}
	// 让临时文件转正
	newName := fmt.Sprintf("mr-out-%d.txt", task.Id)
	os.Rename(tmpFileName, newName)

	// 最后通知coordinator完成了reduce任务
	CallTaskFinished(&req, &reply)

}
```

因为测试脚本最终会检查输出格式和内容，它有个标准答案，单词是按字典序排的，所以必须排序不然测试会报错。最后生成结果时，我同样用了先写临时文件、再 `rename` 成正式 `mr-out-X` 的方式，避免别人读到不完整结果。

### 4.6 timeoutHandler()：worker 超时处理

lab1 最有“分布式味道”的地方，其实不是 map 和 reduce 本身，而是 **失败恢复**。

我的实现是，任务被分配出去，就会起一个超时 goroutine：

```go
func (c *Coordinator) timeoutHandler(task *Task) {
	time.Sleep(10 * time.Second)
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	switch task.Type {
	case Map:
		// 如果该 Map 任务在超时后仍未完成，重新放回队列
		if !c.MapFinished[task.Id] {
			log.Printf("Map 任务 %d (文件 %s) 超时，重新调度\n", task.Id, task.FileName)
			c.MapChan <- task
			c.MapAssign[task.Id] = false
		}
	case Reduce:
		// 如果 Reduce 任务在超时后仍未完成，重新放回队列
		if !c.ReduceFinished[task.Id] {
			log.Printf("Reduce 任务 %d 超时，重新调度\n", task.Id)
			c.ReduceChan <- task
			c.ReduceAssign[task.Id] = false
		}
	}
}
```

如果 10 秒后发现这个任务还没有完成，就默认这个 worker 已经挂了、卡住了，或者慢得不值得再等，于是把任务重新塞回任务队列，交给别的 worker。

---

## 5. 容易踩坑的点

### 5.1 Map 和 Reduce 不能混着跑

这是整个实现里最重要的一条边界。
如果 map 还没全部结束，就提前让 reduce 开始，那 reduce 读到的中间文件就会不完整，最终会出错。

所以这个 lab 里最核心的“全局约束”其实不是某个复杂算法，而是一个很朴素的条件：

> **只有所有 map 完成，reduce 才能开始。**

### 5.2 “没任务可发” 不等于 “系统已经结束”

worker 拉任务时，如果队列暂时空了，很可能只是因为别的 worker 还在干活，不代表整个作业已经做完。
因此我这里用 `Wait` 和 `End` 两种任务类型专门区分：

* `Wait`：过一会儿再来问；
* `End`：这次是真的结束了。

如果把这两种情况混成一种，worker 要么会过早退出，要么会在作业完成后一直空转。  

### 5.3 超时重试不关心 worker 的状态

coordinator 其实不需要、也做不到精准判断 worker 到底是 crash、网络问题、还是单纯太慢。它只需要：

> **过了规定时间还没交结果，就重新分配任务。**

---

## 6. 总结与收获

![](D:\dev_soft\Go_WorkSpace\doc\MIT6.5840\mapreduce\pass.png)

​														通过！受益匪浅，MIT牛逼

本人本科是大数据专业，课程经常接触到 mapreduce ，但我一直疑惑为什么要学习已经被淘汰的技术，因为 Google 早就没有在用mapreduce了，现在也没有软件在用 mapreduce 作为自己的核心计算模型。我接触到这个实验的时候震惊：居然连mit也在教十几年前就不用了的技术？？？

**但是在做完这个实验后我便理解了 mapreduce 的伟大之处**

它小到可以让学生短时间亲手做出来，**麻雀虽小，五脏俱全**。MapReduce 的伟大之处不在于它快，而在于它用两个逻辑极其简单的函数（Map 和 Reduce），第一次向全世界证明了：**即使不是分布式系统的专家，也能套用单机逻辑的，容易理解的编程模型，写出跑在数千台机器上且充分利用计算资源的程序**。现在工业界常用的 Spark、Flink、Ray 等框架，它们的底层核心逻辑依然是：**“把数据切分，分发到节点，在本地计算，然后聚合结果。”** 这便是 MapReduce 的遗产，他确实是最适合作为高校分布式系统课程教学的计算模型。
