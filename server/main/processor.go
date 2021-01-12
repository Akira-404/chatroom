package main

import(
	"fmt"
	"net"
	"io"
	"chatroom/common/message"
	"chatroom/server/utils"
	"chatroom/server/process"
)

type Processor struct{
	Conn net.Conn
}

func (this *Processor)serverProcessMES(mes *message.Message) (err error) {
	
	fmt.Println("服务器接收到的信息=",mes)
	//根据消息类型处理信息
	switch mes.Type {
		case message.LoginMesType://登录信息
			fmt.Println("属于登录类型")
			up:=&process2.UserProcess{
				Conn:this.Conn,
			}
			err=up.ServerProcessLogin(mes)
		case message.RegisterMesType://注册信息
			fmt.Println("属于注册类型")
			up:=&process2.UserProcess{
				Conn:this.Conn,
			}
			err=up.ServerProcessRegister(mes)
		case message.SmsMesType://群发信息
			fmt.Println("属于群发类型")
			smsProcess:=&process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default:
			fmt.Println("type is not exited")	
	}	
	return
}

func (this *Processor)process2() (err error) {
	
	tf:=&utils.Transfer{
		Conn:this.Conn,
	}
	for{
		//读取数据
		mes,err:=tf.ReadPkg()
		if err!=nil{
			if err==io.EOF{
				fmt.Println("服务器和客户端结束")
				return err
			}else{
				fmt.Println("err=",err)
				return err
			}
		}

		//处理信息
		err=this.serverProcessMES(&mes)
		if err!=nil{
			fmt.Println("serverProcessMES err=",err)
			return err
		}
	}
}