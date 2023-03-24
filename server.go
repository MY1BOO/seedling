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

//Test PreHandle
func (this *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Call Router PreHandle")

	_, err := request.GetConnection().(*net.Connection).GetTCPConnection().Write(append(request.GetData(), []byte("before ping ....\n")...))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test Handle
func (this *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().(*net.Connection).GetTCPConnection().Write(append(request.GetData(), []byte("ping...ping...ping\n")...))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().(*net.Connection).GetTCPConnection().Write(append(request.GetData(), []byte("After ping .....\n")...))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	//创建一个server句柄
	s := net.NewServer()

	s.AddRouter(&PingRouter{})

	//2 开启服务
	s.Serve()
}
