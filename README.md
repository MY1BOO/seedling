# seedling
###seedling v2.0 雏形版本
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

总结：单独的监听者Goroutine会开启很多个业务处理Goroutine，每个业务处理Goroutine会开启一个执行HandFunc的Goroutine