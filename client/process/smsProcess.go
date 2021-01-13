package process

import(
	"fmt"
	"chatroom/common/message"
	"encoding/json"
	"chatroom/client/utils"
)

type SmsProcess struct{

}

func (this *SmsProcess)SendGroupMes(content string)(err error){

	//create a message(type,data)
	var mes message.Message
	mes.Type=message.SmsMesType

	//create a smsmes
	var smsMes message.SmsMes
	smsMes.Content=content
	smsMes.UserId=CurUser.UserId
	smsMes.UserStatus=CurUser.UserStatus

	//封装数据
	data,err:=json.Marshal(smsMes)
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
	
	//发送数据
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