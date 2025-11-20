package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip          string
	Port        int
	OnlineUsers map[string]*User // 在线用户的列表
	mapLock     sync.RWMutex
	Message     chan string // 用于发送消息的管道
}

// 创建一个新的Server实例
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:          ip,
		Port:        port,
		OnlineUsers: make(map[string]*User),
		Message:     make(chan string),
	}
}

// 处理函数
func (server *Server) Handler(conn net.Conn) {
	// do something
	fmt.Println("connection create success")

	// 用户上线
	user := NewUser(conn, server)
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接受客户端发送的信息
	go func() {
		buffer := make([]byte, 4096)
		for {
			n, err := conn.Read(buffer)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}
			// 处理客户端发送的信息
			msg := string(buffer[:n-2])
			// 用户针对msg进行处理
			user.HandleMessage(msg)
			// 标记用户活跃
			isLive <- true
		}
	}()

	// 阻塞当前handler
	for {
		select {
		case <-isLive: // 用户活跃，继续监听客户端发送的信息
		case <-time.After(time.Second * 10):
			// 10秒没有收到客户端的消息，认为用户不活跃，关闭连接
			user.Offline()
			return // 退出整个Handler
		}
	}
}

// 广播消息的方法
func (server *Server) BroadCast(user *User, msg string) {
	// 构造最终发送的信息
	sendMsg := fmt.Sprintf("[%s] %s: %s", user.Addr, user.Name, msg)
	server.Message <- sendMsg // 发送消息到server的消息通道
}

// 监听广播信息channel的goroutine，一旦server的通道中有信息就发送给每一个在线的用户
func (server *Server) ListenMessage() { // 消费server的通道中的信息
	for {
		msg := <-server.Message
		server.mapLock.Lock() // 上锁
		for _, u := range server.OnlineUsers {
			u.C <- msg
		}
		server.mapLock.Unlock() // 解锁
	}
}

func (server *Server) Start() {
	// socket listen
	socket := fmt.Sprintf("%s:%d", server.Ip, server.Port)
	listener, err := net.Listen("tcp", socket)
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	fmt.Println("server start success")
	go server.ListenMessage() // 启动监听广播消息的goroutine
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// do handler
		go server.Handler(conn)
	}

}
