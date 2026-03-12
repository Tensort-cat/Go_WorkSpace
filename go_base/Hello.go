package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
)

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func RandValue(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
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

	var ver uint64
	fmt.Println("ver:", ver)
}

func simplifyFormat(s string) string {
	if idx := strings.LastIndex(s, "-"); idx != -1 {
		return s[:idx] + ".txt"
	}
	return s
}
