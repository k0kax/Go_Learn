package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex //同步机制

	//消息广播
	Message chan string
}

// 创建Server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给所有在线User
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		//将msg发送给所有在线User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// Handler 方法
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")

	user := NewUser(conn, this)
	////用户上线，将用户加入onlinemap表
	//this.mapLock.Lock()
	//this.OnlineMap[user.Name] = user
	//this.mapLock.Unlock()
	//
	////用户上线广播
	//this.BroadCast(user, "已上线")

	user.Online() //上线
	fmt.Println(user.Name, ":已上线")

	isLive := make(chan bool)
	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline() //下线
				fmt.Println(user.Name, ":已下线")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//提取用户的信息(去掉“\n”)
			msg := string(buf[:n-1])

			//用户消息处理
			user.DoMessage(msg)

			//用户的任意消息代表用户处于活跃状态
			isLive <- true
			fmt.Println(user.Addr, "|", user.Name, ":", msg)
		}
	}()
	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//当前用户活跃，应该重置定时器
			//不做任何事情，为了激活select,更新下面的定时器

		case <-time.After(time.Second * 100): //定时器
			//已经超时
			//强制User强制关闭
			user.SendMsg("你被踢了")

			//销毁用的资源
			close(user.C)

			//关闭连接
			conn.Close()

			//退出当前handler
			return

		}
	}

}

// 启动服务器接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//socket close
	defer listener.Close()

	//启动Message的goroutine
	go this.ListenMessage()

	//循环接受请求和监听
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			return
		}

		//do handler
		go this.Handler(conn)

	}
}
