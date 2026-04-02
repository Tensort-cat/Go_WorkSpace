package main

import (
	"log"
	"net/http"
)

func main() {
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
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

/*
	2026/04/01 16:58:55 Error fetching url www.baidu.com : Get "www.baidu.com": unsupported protocol scheme ""
	2026/04/01 16:58:55 Status Code for https://www.baidu.com : 200 OK
*/
