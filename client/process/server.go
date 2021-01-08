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
	fmt.Println("=========login success=========")	 
	fmt.Println("=========1 show online user list =========")	 
	fmt.Println("=========2 send message=========")	 
	fmt.Println("=========3 info list=========")	 
	fmt.Println("=========4 quit=========")
	fmt.Println("select 1-4")
	var key int
	var content string
	smsProcess:=&SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			// fmt.Println("=========1 show online user list =========")	 
			outputOnlineUser()
		case 2:
			fmt.Println("=========2 send message for everybody=========")
			fmt.Scanf("%s\n",&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("=========3 info list=========")	 
		case 4:
			fmt.Println("=========4 quit=========")
			os.Exit(0)
		default:
			fmt.Println("try again")
	}
 }

 func serverProcessMes(conn net.Conn)  {
	
	tf:=&utils.Transfer{
		Conn:conn,
	} 
	
	for{
		fmt.Println("client ip is waiting",)
		mes,err:=tf.ReadPkg()
		if err!=nil{
			fmt.Println("tf.ReadPkg err=",err)
			return
		}
		// fmt.Println("mes=",mes)
		switch mes.Type {
			case message.NotifyUserStatusMesType :
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
				updataUserStatus(&notifyUserStatusMes)
			case message.SmsMesType :
				fmt.Println(mes)
				outputGroupMes(&mes)
			default:fmt.Println("not find this type")
			
		}
	} 
 }