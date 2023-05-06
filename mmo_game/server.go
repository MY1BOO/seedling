package main

import (
	"fmt"
	"github.com/MY1BOO/seedling/iface"
	"github.com/MY1BOO/seedling/mmo_game/core"
	"github.com/MY1BOO/seedling/net"
)

//当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn iface.IConnection) {
	//创建一个玩家
	player := core.NewPlayer(conn)
	//同步当前的PlayerID给客户端， 走MsgID:1 消息
	player.SyncPid()
	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	player.BroadCastStartPosition()

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived ====")
}

func main() {
	//创建服务器句柄
	s := net.NewServer()
	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnecionAdd)
	//启动服务
	s.Serve()
}
