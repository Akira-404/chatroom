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
	
	fmt.Println("serverprocessmes mes=",mes)
	//根据消息类型处理信息
	switch mes.Type {
		case message.LoginMesType://登录信息
			up:=&process2.UserProcess{
				Conn:this.Conn,
			}
			err=up.ServerProcessLogin(mes)
		case message.RegisterMesType://注册信息
			up:=&process2.UserProcess{
				Conn:this.Conn,
			}
			err=up.ServerProcessRegister(mes)
		case message.SmsMesType:
			// up:=&process2.UserProcess{
			// 	Conn:this.Conn,
			// }
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
				fmt.Println("client quit and server quit")
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