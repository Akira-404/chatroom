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

	fmt.Println("=========> 登录成功")	 
	fmt.Println("=========> 1 在线用户列表")	 
	fmt.Println("=========> 2 群发")	 
	fmt.Println("=========> 3 私发")	 
	fmt.Println("=========> 4 退出")
	
	var key int
	var content string
	var sendUserId int 
	smsProcess:=&SmsProcess{}
	personalmessage:=&PersonalMessage{}
	fmt.Scanf("%d\n",&key)
	
	switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Println("=========> 输入群发信息")
			fmt.Scanf("%s\n",&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("=========> 当前在线用户有:")
			outputOnlineUser()	 

			fmt.Println("=========> 选择对象：")
			fmt.Scanf("%d\n",&sendUserId)

			fmt.Println("=========> 输入私发信息")
			fmt.Scanf("%s\n",&content)

			personalmessage.SendMessage(content,sendUserId)
		case 4:
			fmt.Println("=========> 退出")
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
		fmt.Println(mes)
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
			case message.PersonalResMesType :
				var personalResMes message.PersonalResMes
				json.Unmarshal([]byte(mes.Data),&personalResMes)
				if personalResMes.Code==200{
					fmt.Println("服务器反馈信息->",personalResMes)
					fmt.Println("服务器反馈->发送成功")
				}

				if personalResMes.Code==201{

					fmt.Println("服务器反馈->用户不在线")
				}
			case message.PersonalMesType:
				var personalMes message.PersonalMes
				json.Unmarshal([]byte(mes.Data),&personalMes)
				fmt.Printf("%d发来%s",personalMes.SUserId,personalMes.Content)

			default:fmt.Println("服务器反馈->没有找到这种类型")
			
		}
	} 
 }