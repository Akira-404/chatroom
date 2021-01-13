package process

import(
	"fmt"
	"os"
	"net"
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
)

func ShowMenu()  {
	fmt.Println("=========登录成功=========")	 
	fmt.Println("=========1 显示在线用户列表 =========")	 
	fmt.Println("=========2 对所有在线用户发送信息=========")	 
	fmt.Println("=========3 信息列表=========")	 
	fmt.Println("=========4 退出=========")
	
	var key int
	var content string
	smsProcess:=&SmsProcess{}
	fmt.Scanf("%d\n",&key)
	
	switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Println("=========2 对所有在线用户发送信息=========")
			fmt.Scanf("%s\n",&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("=========3 信息列表=========")	 
		case 4:
			fmt.Println("=========4 退出=========")
			os.Exit(0)
		default:
			fmt.Println("重试")
	}
 }

 func serverProcessMes(conn net.Conn)  {
	
	tf:=&utils.Transfer{
		Conn:conn,
	} 
	
	for{
		fmt.Println("客户端等待信息...\r",)
		mes,err:=tf.ReadPkg()
		if err!=nil{
			fmt.Println("tf.ReadPkg err=",err)
			return
		}
		
		switch mes.Type {

			case message.NotifyUserStatusMesType :
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
				updataUserStatus(&notifyUserStatusMes)

			case message.SmsMesType :
				fmt.Println(mes)
				outputGroupMes(&mes)

			default:fmt.Println("没有找到这种类型")
			
		}
	} 
 }