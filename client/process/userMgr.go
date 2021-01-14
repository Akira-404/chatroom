package process

import(
	"fmt"
	"chatroom/common/message"
	"chatroom/client/model"
)

var onlineUsers map[int]*message.User=make(map[int]*message.User, 10)
var CurUser model.CurUser

//显示在线列表
func outputOnlineUser()  {
	for id,_:=range onlineUsers{
		fmt.Printf("ID:%d在线\n",id)
	}	
}

//更新用户状态
func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes)  {
		
	user,ok:=onlineUsers[notifyUserStatusMes.UserId]
	if !ok{

		user=&message.User{
			UserId:notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus=notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId]=user

	outputOnlineUser()
}