package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string //channel
	conn net.Conn    //链接方法

	server *Server
}

// 创建一个用户的Api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() //拿到地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//启动当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//用户上线广播
	this.server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//用户上线广播
	this.server.BroadCast(this, "已下线")
}

// 给当前用户对应的客户端发消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {

	if msg == "who" {

		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onLineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线。。。\n"
			this.SendMsg(onLineMsg)
		}
		this.server.mapLock.Unlock()
		//重命名
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式：rename|张3
		newName := strings.Split(msg, "|")[1] //截取第一个元素,也就是用户的新名字

		//判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名已被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("你已经更新用户名成功" + this.Name + "\n")
		}

		//私聊模块
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式： to|zhang3|消息内容

		//1.获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式错误 请使用 to|zhang3|消息内容 ")
			return
		}

		//2.根据用户名，得到对方的User
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户不存在")
			return
		}

		//3.获取消息内容，通过对象的User对象将消息内容发送出去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发")
			return
		}
		remoteUser.SendMsg(this.Name + "对你说：" + content)
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channel的方法,一旦有消息，就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		_, err := this.conn.Write([]byte(msg + "\n"))
		if err != nil {
			return
		}
	}
}
