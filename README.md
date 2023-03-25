# seedling

### seedling v2.0 雏形版本
介绍：

1.定义server接口和connection接口及其实现类

2.connection类可以将socket连接与业务处理方法HandFunc绑定在一起

3.处理流程为——main创建server对象，server.Start()，main阻塞，
开启一个Goroutine作为监听连接的listener，然后for循环监听连接，
阻塞在accept上，当有连接来到时，创建connection对象，将HandFunc类型的函数和该socket
连接传入，开启一个业务处理Goroutine执行connection.Start()，
在这个业务处理Goroutine中再开启一个Goroutine执行StartReader，
然后利用select配合chan阻塞等待StartReader完成，StartReader其实
是真正执行读写数据的方法，但在处理前要defer stop()确保正常退出

总结：单独的监听者Goroutine会开启很多个业务处理Goroutine，
每个业务处理Goroutine会开启一个执行HandFunc的Goroutine

### seedling v4.0
介绍：

1.定义request接口和router接口及其实现类

2.router其实是处理业务的方法集合，包含前置处理、处理、后置处理，
request包含一个socket连接和从客户端接收到的数据，作为router中方法的参数

3.server实例初始化时要传入一个用户自定义好的router，然后在listener协程创建connection
的时候，也要将router传入，代替v2.0中的HandFunc进行业务处理。

4.小技巧：先定义一个BaseRouter基类，只不过当中的三个方法全为空，然后在自定义router
的时候，都先继承这个BaseRouter，这样就更灵活，不用全部实现所有方法也可以是router接口的实现类

5.通过读文件+json解析实现全局配置

### seedling v6.0
介绍：

1.定义message接口、datapack接口和msghandler接口

2.message其实就是对[]byte数据的封装，包含数据本身、数据长度和ID，
datapack可以理解为一个工具类，提供数据拆包、数据封包两个方法，目的是
解决TCP粘包问题，msghandler为了实现多路由模式，包含一个map，key为消息ID，
value为具体的router，这样seedling就可以根据消息ID的不同选择合适的router处理业务

3.拆包过程是先用datapack的unpack方法读出消息长度和ID，然后根据长度再次从socket里读出具体消息，
也就是输入[]byte，输出message，封包过程相反，输入message，然后分段，输出[]byte。然后connection
还添加一个sendmsg方法，封包接着发送

4.将之前server实例中的router换成msghandler，也就是一个包含着多个router的map，可以不停的
add自定义的router进去
