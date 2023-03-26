# seedling

### seedling v2.0

介绍：

1.定义server接口和connection接口及其实现类

2.connection类可以将socket连接与业务处理方法HandFunc绑定在一起

3.处理流程为——main创建server对象，server.Start()，main阻塞，开启一个Goroutine作为监听连接listener，然后for循环监听连接，阻塞在accept上，当有连接来到时，创建connection对象，将HandFunc类型的函数和该socket连接传入，开启一个业务处理Goroutine执行connection.Start()，在这个业务处理Goroutine中再开启一个Goroutine执行StartReader，然后利用select配合chan阻塞等待StartReader完成，StartReader其实是真正执行读写数据的方法，但在处理前要defer stop()确保正常退出

总结：单独的监听者Goroutine会开启很多个业务处理Goroutine，每个业务处理Goroutine会开启一个执行HandFunc的Goroutine

### seedling v4.0

介绍：

1.定义request接口和router接口及其实现类

2.router其实是处理业务的方法集合，包含前置处理、处理、后置处理，request包含一个socket连接和从客户端接收到的数据，作为router中方法的参数

3.server实例初始化时要传入一个用户自定义好的router，然后在listener协程创建connection的时候，也要将router传入，代替v2.0中的HandFunc进行业务处理。

4.小技巧：先定义一个BaseRouter基类，只不过当中的三个方法全为空，然后在自定义router的时候，都先继承这个BaseRouter，这样就更灵活，不用全部实现所有方法也可以是router接口的实现类

5.通过读文件+json解析实现全局配置

### seedling v6.0

介绍：

1.定义message接口、datapack接口和msghandler接口及其实现类

2.message其实就是对[]byte数据的封装，包含数据本身、数据长度和ID，datapack可以理解为一个工具类，提供数据拆包、数据封包两个方法，目的是解决TCP粘包问题，msghandler为了实现多路由模式，包含一个map，key为消息ID，value为具体的router，这样seedling就可以根据消息ID的不同选择合适的router处理业务

3.拆包过程是先用datapack的unpack方法读出消息长度和ID，然后根据长度再次从socket里读出具体消息，也就是输入[]byte，输出message，封包过程相反，输入message，然后分段，输出[]byte。然后connection还添加一个sendmsg方法，封包接着发送

4.将之前server实例中的router换成msghandler，也就是一个包含着多个router的map，可以不停的add自定义的router进去

### seedling v8.0

介绍：

1.connection中本来只有go StartReader() --> 读取数据 --> go DoMsgHandler() --> SendMsg()写回数据，现在将其改造成读写分离模式，即启动一个connection时，go StartReader()+go StartWriter()同时启动，SendMsg()方法不再直接写回数据，而是将数据发送到channel当中，writer Goroutine检测到channel中有数据就会立马将其发送出去

2.msgHandler新增worker池和消息队列，每个worker带一个消息队列，每个消息队列其实是一个装request的channel，在server启动的最开始worker Goroutine就已经按照配置文件中的数量启动好，并for select不断的等待队列中的消息，一有消息就DoMsgHandler()。以前connection是直接go DoMsgHandler()处理业务，而现在则是SendMsgToTaskQueue()发往一个消息队列中，由其所对应的worker Goroutine处理。这样可以通过worker的数量来限定处理业务的固定goroutine数量，而不是无限制的开辟Goroutine。

### seedling v10.0

介绍：

1.定义ConnManager接口及其实现类

2.ConnManager为server实例中的一个新成员，作用为管理server中的所有connection，可以添加、删除连接，获取连接数量等，在连接到达配置文件中设置的数量之后，不再添加新连接

3.添加带缓冲的发包方法，其实就是connection之前只有一个无缓冲的channel，SendMsg()发送数据过去，现在给connection添加一个有缓冲的channel和一个新的方法SendBuffMsg()

4.添加钩子方法，可以分别在建立连接之后和销毁连接之前调用

5.给connection添加K-V键值对所代表的属性