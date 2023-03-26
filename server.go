package main

import (
	"fmt"
	"github.com/MY1BOO/seedling/iface"
	"github.com/MY1BOO/seedling/net"
)

//ping test 自定义路由
type PingRouter struct {
	net.BaseRouter //一定要先基础BaseRouter
}

////Test PreHandle
//func (this *PingRouter) PreHandle(request iface.IRequest) {
//	fmt.Println("Call Router PreHandle")
//
//	_, err := request.GetConnection().(*net.Connection).GetTCPConnection().Write(append(request.GetData(), []byte("before ping ....\n")...))
//	if err != nil {
//		fmt.Println("call back ping ping ping error")
//	}
//}

//Test Handle
func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

////Test PostHandle
//func (this *PingRouter) PostHandle(request iface.IRequest) {
//	fmt.Println("Call Router PostHandle")
//	_, err := request.GetConnection().(*net.Connection).GetTCPConnection().Write(append(request.GetData(), []byte("After ping .....\n")...))
//	if err != nil {
//		fmt.Println("call back ping ping ping error")
//	}
//}

//HelloSeedlingRouter Handle
type HelloSeedlingRouter struct {
	net.BaseRouter
}

func (this *HelloSeedlingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call HelloSeedlingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Seedling Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

//创建连接的时候执行
func DoConnectionBegin(conn iface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")

	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name", "qiujun")
	conn.SetProperty("Home", "qingdao")

	err := conn.SendBuffMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

//连接断开的时候执行
func DoConnectionLost(conn iface.IConnection) {
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}

	fmt.Println("DoConneciotnLost is Called ... ")
}

func main() {
	//创建一个server句柄
	s := net.NewServer()

	//注册连接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloSeedlingRouter{})

	//2 开启服务
	s.Serve()
}
