package process2

import(
	"fmt"
	"net"
	"encoding/json"
	"chatroom/common/message"
	"chatroom/server/utils"
)

type SmsProcess struct{

}

func (this *SmsProcess)SendGroupMes(mes *message.Message)  {

	var smsMes message.SmsMes
	err:=json.Unmarshal([]byte(mes.Data),&smsMes)
	
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
	}	
	
	data,err:=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
	}	
	
	for id,up:=range userMgr.onlineUsers{
		//过滤自己，不给自己发信息
		if id==smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess)SendMesToEachOnlineUser(data []byte,conn net.Conn)  {
	
	tf:=&utils.Transfer{
		Conn:conn,
	}
	err:=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("转发失败 err=",err)
	} 
}