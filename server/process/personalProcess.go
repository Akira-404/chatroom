package process2

import(
	"fmt"
	"net"
	"encoding/json"
	"chatroom/common/message"
	"chatroom/server/utils"
	"chatroom/server/model"
)

type PersonalMesProcess struct{
	Conn net.Conn
}

func (this *PersonalMesProcess)SendMesTo(mes *message.Message)  {

	//获取信息
	var personalMes message.PersonalMes
	err:=json.Unmarshal([]byte(mes.Data),&personalMes)
	
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
	}	
	
	data,err:=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
	}
	//服务器转发
	var flag=false
	for id,up:=range userMgr.onlineUsers{
		if id==personalMes.RUserId{
			this.ForwardInfo(data,up.Conn)
			this.PersonalMesRes(200)
			fmt.Println("转发完成,转发用户:",personalMes.RUserId)
			flag=true
		}
		if flag==true{
			break
		}
	}
	if flag==false{
		this.PersonalMesRes(201)
		fmt.Println("当前用户不在线")
	}
	
}

//服务器转发信息
func (this *PersonalMesProcess)ForwardInfo(data []byte,conn net.Conn)  {

	tf:=&utils.Transfer{
		Conn:conn,
	}
	err:=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("转发失败 err=",err)
	}
}

//服务器反馈转发结果
func (this *PersonalMesProcess)PersonalMesRes(code int)  {
	
	var mes message.Message
	mes.Type=message.PersonalResMesType

	var respersonalmes message.PersonalResMes
	respersonalmes.Code=code

	if code==200{
		respersonalmes.Error=nil
	}
	if code==201{
		respersonalmes.Error=model.ERROR_USER_ONTONLINE
	}
	
	//序列化消息体内容
	data,err:=json.Marshal(respersonalmes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	mes.Data=string(data)
	
	//序列化消息
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	//send data
	tf:=&utils.Transfer{
		Conn:this.Conn,
	}
	err=tf.WritePkg(data)
	fmt.Println("发送转发响应信息->",mes)

}