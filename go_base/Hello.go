package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

type Reply struct {
	Success bool
}

func main() {
	// 我是单行注释
	/* 	我是
		多行
	 	注释
	*/
	fmt.Println("Hello World!")
	fmt.Println("Hello Again!")

	// 字符串拼接
	fmt.Println("Hello" + " " + "World!")

	a := 10
	if a > 2 {
		fmt.Println("a is greater than 2")
	}

	var code = 123
	var date = "2021-01-01"
	var url = "Code = %d & endDate = %s"
	var target = fmt.Sprintf(url, code, date)
	fmt.Println(target) // 输出结果：Code = 123 & endDate = 2021-01-01

	num := rand.Int()
	fmt.Printf("Random number: %d\nnumber type: %T\n", num, num) // 输出随机数

	fmt.Println(ihash("A") % 10)

	m := make(map[string]string)
	m["hello"] = "1"
	m["world"] = "1"
	m["bitch"] = "1"
	fmt.Println(m)

	files, err := filepath.Glob(fmt.Sprintf("he*")) // 读取属于该任务的reduce任务文件
	if err != nil {
		fmt.Println("err: %v", err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	fmt.Println(simplifyFormat("100-250-114514.txt"))

	rep := Reply{}
	fmt.Println(rep.Success)

	// os.Rename("hello.txt", "fuck.txt")
	bytes, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("err: %v", err)
	} else {
		fmt.Println(string(bytes))
	}

	fmt.Printf("%v", true)
}

func simplifyFormat(s string) string {
	if idx := strings.LastIndex(s, "-"); idx != -1 {
		return s[:idx] + ".txt"
	}
	return s
}
