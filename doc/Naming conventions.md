太棒了 👍！下面是为你准备的 **《Go 命名规范速查表》Markdown 版**，
内容简洁实用，适合直接放进你的笔记或项目 README 中👇

---

````markdown
# 🧭 Go 命名规范速查表

Go 的命名风格追求：**简洁、清晰、语义明确**  
它没有复杂的访问修饰符（public/private），而是通过**首字母大小写**来控制可见性。

---

## ✨ 一、可见性规则

| 首字母 | 可见性 | 说明 |
|---------|---------|------|
| **大写字母开头** | 导出（Public） | 可以被其他包访问 |
| **小写字母开头** | 未导出（Private） | 只能在当前包内访问 |

```go
func Add(a, b int) int { return a + b } // 可导出
func subtract(a, b int) int { return a - b } // 仅包内可用
````

---

## 📦 二、包名（package）

* 全部小写，不使用下划线或驼峰。
* 名称应简短、语义明确。
* 包名通常和所在文件夹名相同。

✅ 推荐：

```go
package httpserver
package database
package user
```

❌ 不推荐：

```go
package UserService
package user_service
```

---

## 📄 三、文件名（file）

* 全部小写，单词间以下划线分隔。
* 反映文件主要内容。

✅ 示例：

```
user_service.go
jwt_token.go
config_loader.go
```

---

## 🧩 四、变量与函数

### 变量

* 使用 **小驼峰（camelCase）**。
* 简洁、语义清楚；不要重复类型或包名。

✅ 推荐：

```go
var userCount int
var fileReader io.Reader
```

❌ 不推荐：

```go
var countOfUsers int
var readerIO io.Reader
```

---

### 函数

* 使用 **小驼峰**（包内）或 **大驼峰**（导出）。
* 函数名应为动词或动宾短语。

✅ 示例：

```go
func readConfig() {}
func LoadConfig() {}
```

❌ 不推荐：

```go
func DoReadingOfConfigFile() {}
```

---

## 🧱 五、结构体与字段

* 结构体名用 **大驼峰（PascalCase）**。
* 字段名首字母决定可见性。

```go
type User struct {
    Name string // 导出
    age  int    // 仅包内可见
}
```

结构体方法命名：

```go
func (u *User) GetName() string { return u.Name }
func (u *User) setAge(a int) { u.age = a }
```

---

## ⚙️ 六、接口（interface）

* 通常使用“**动词 + er**”命名，表示行为能力。
* 单方法接口可用动词命名。

✅ 示例：

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}
```

---

## 🧮 七、常量（const）

### 普通常量

* 小驼峰（包内）或大驼峰（导出）。

```go
const defaultTimeout = 30
const MaxConnections = 100
```

### 枚举常量

* 一般用前缀 + 大写：

```go
const (
    StatusOK    = 200
    StatusError = 500
)
```

---

## 🌿 八、命名风格建议

✅ 推荐的风格：

* 简短、有意义（prefer clarity over abbreviation）
* 避免冗余（包名已表达语义时不要重复）
* 保持一致（整个项目遵循统一风格）

❌ 不推荐的风格：

* 过度缩写（除非众所周知，如 `err`, `cfg`, `db`）
* 驼峰和下划线混用
* 含糊不清的变量名（如 `data`, `info`, `tmp`）

---

## 💡 九、小结

| 类型  | 命名规则     | 示例                            |
| --- | -------- | ----------------------------- |
| 包名  | 小写       | `package user`                |
| 文件名 | 小写+下划线   | `user_service.go`             |
| 变量  | 小驼峰      | `userCount`                   |
| 常量  | 小驼峰/大写   | `maxSize`, `StatusOK`         |
| 函数  | 小/大驼峰    | `getUser()`, `LoadUser()`     |
| 结构体 | 大驼峰      | `type User struct {}`         |
| 字段  | 首字母控制可见性 | `Name`, `age`                 |
| 接口  | 动词+er    | `Reader`, `Writer`, `Handler` |

---

## 🪄 记忆口诀

> 包小写，文件下划线；
> 变量小驼峰，类型大写显；
> 首字母定可见，接口加 “er” 连。

---

```

---

是否希望我帮你加上一个「对比示例」，展示同一功能分别用**非规范**与**规范命名**写出来的版本？  
这样你可以更直观地看到命名规范的作用。
```
