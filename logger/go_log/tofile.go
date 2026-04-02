package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	SetupLogger()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}

func SetupLogger() { // 将日志输出到指定文件中
	logFileLocation, _ := os.OpenFile("./tmp/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
	return
}
