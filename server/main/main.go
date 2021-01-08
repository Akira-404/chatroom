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
)

//process communicate with client
func process(conn net.Conn)  {

	//read the message fo client
	defer conn.Close()
	processor:=&Processor{
		Conn:conn,
	}	
	err:=processor.process2()
	if err==nil{
		fmt.Println("processor.process2 err=",err)
		return
	}
}

func initUserDao()  {
	model.MyUserDao=model.NewUserDao(pool)
}

func main()  {

	initPool("0.0.0.0:6379",16,0,300*time.Second)
	initUserDao()
	//监听端口
	fmt.Println("server listening 8889")
	listen,err:=net.Listen("tcp","0.0.0.0:8889")	
	
	defer listen.Close()
	
	if err!=nil{
		fmt.Println("net listen err=",err)
		return
	}

	for{
		fmt.Println("Wait for the client to link to the server...")
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("listen.accept err=",err)
		}
		go process(conn)
	}
}
