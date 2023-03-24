package iface

/*
	路由接口，这里面路由是使用框架者给该连接自定的处理业务方法
	路由里的IRequest则包含用该连接的连接信息和该连接的请求数据信息
*/
type IRouter interface {
	PreHandle(request IRequest)  //在处理conn业务之前的钩子方法
	Handle(request IRequest)     //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务之后的钩子方法
}
