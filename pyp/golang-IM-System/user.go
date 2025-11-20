package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user
}

// 监听当前user channel的方法，一旦有消息，就发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线
func (u *User) Online() {
	u.server.mapLock.Lock()          // 上锁
	u.server.OnlineUsers[u.Name] = u // 添加到在线用户列表
	u.server.mapLock.Unlock()        // 解锁
	// 广播当前用户上线信息
	u.server.BroadCast(u, "I am online!")
}

// 用户下线，将用户从server的map中删除
func (u *User) Offline() {
	u.server.mapLock.Lock()              // 上锁
	delete(u.server.OnlineUsers, u.Name) // 从在线用户列表中删除
	u.server.mapLock.Unlock()            // 解锁
	u.conn.Close()                       // 关闭当前用户的连接
	// 广播当前用户下线信息
	u.server.BroadCast(u, "I am offline!")
}

// 用户发送消息
func (u *User) SendMessage(msg string) {
	u.C <- msg
}

// 用户处理消息
func (u *User) HandleMessage(msg string) {
	switch {
	case msg == "who":
		// 获取当前在线用户列表
		u.server.mapLock.Lock()
		fmt.Println(u.server.OnlineUsers)
		for _, onlineUser := range u.server.OnlineUsers {
			onlineMsg := fmt.Sprintf("%s is online", onlineUser.Name)
			u.SendMessage(onlineMsg)
		}
		u.server.mapLock.Unlock()

	case msg == "myself":
		// 发送自己的信息
		myselfMsg := fmt.Sprintf("My name is %s", u.Name)
		u.SendMessage(myselfMsg)

	case strings.HasPrefix(msg, "rename"):
		newName := strings.Split(msg, " ")[1]
		oldName := u.Name
		// 修改用户名
		u.server.mapLock.Lock()
		u.Name = newName
		delete(u.server.OnlineUsers, oldName)
		u.server.OnlineUsers[newName] = u
		u.server.mapLock.Unlock()
		// 返回成功信息
		renameMsg := fmt.Sprintf("%s changed name to %s", oldName, newName)
		u.SendMessage(renameMsg)

	default: // 默认广播
		u.server.BroadCast(u, msg)
	}
}
