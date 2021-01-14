package process

import(
	"fmt"
	"chatroom/common/message"
	"encoding/json"
	"chatroom/client/utils"
)

type PersonalMessage struct{

}

func (this *PersonalMessage)SendMessage(content string,ruserId int)(err error)  {


	var mes message.Message
	mes.Type=message.PersonalMesType

	var personalMes message.PersonalMes
	personalMes.SUserId=CurUser.UserId 
	personalMes.Content=content 
	personalMes.RUserId=ruserId

	//封装数据
	data,err:=json.Marshal(personalMes)
	if err!=nil{
		fmt.Println("json.Marshal(smsMes) err=",err.Error())
		return
	}
	mes.Data=string(data)
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal(mes) err=",err.Error())
		return
	}
	
	//发送数据到服务器
	tf:=&utils.Transfer{
		Conn:CurUser.Conn,
	}
	err=tf.WritePkg(data)
	
	if err!=nil{
		fmt.Println("tf.WriterPkg(data) err=",err.Error())
		return
	}	
	
	return
}