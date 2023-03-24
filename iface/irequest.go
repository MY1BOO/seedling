package iface

/*
	IRequest 接口：
	实际上是把客户端请求的连接信息和请求的数据包装到了 Request里
*/
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
}