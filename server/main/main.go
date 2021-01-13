package main

import(
	"fmt"
	"net"
	"time"
	"chatroom/server/model"
	_"io"
	_"errors"
	_"encoding/json"
	_"chatroom/common/message"
	_"encoding/binary"
	"redigo/redis"
)

//process communicate with client
func process(conn net.Conn)  {

	//read the message fo client
	defer conn.Close()//结束这个协程
	processor:=&Processor{
		Conn:conn,
	}	
	err:=processor.process2()
	if err==nil{
		fmt.Println("processor.process2 err=",err)
		return
	}
}

//init user data access object
func initUserDao(pool *redis.Pool)  {
	model.MyUserDao=model.NewUserDao(pool)
}

func main()  {

	//初始化redis连接池
	initPool("0.0.0.0:6379",16,0,300*time.Second)
	//初始化user 数据库操作
	
	initUserDao(pool)
	
	//返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络：
	fmt.Println("服务器监听端口:8889")
	listen,err:=net.Listen("tcp","0.0.0.0:8889")	
	
	defer listen.Close()
	
	if err!=nil{
		fmt.Println("net listen err=",err)
		return
	}

	for{
		fmt.Println("等待客户端连接服务器...")

		//Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("listen.accept err=",err)
		}
		//启动一个协程
		go process(conn)
	}
}
