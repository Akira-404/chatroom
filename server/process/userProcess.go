package process2

import(
	"fmt"
	"net"
	"encoding/json"
	"chatroom/common/message"
	"chatroom/server/utils"
	"chatroom/server/model"
)

type UserProcess struct{
	Conn net.Conn
	UserId int
}

func (this *UserProcess)NotifyOthersOnlineUser(userId int)  {

	for id,up:=range userMgr.onlineUsers{
		if id==userId{
			continue
		}

		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess)NotifyMeOnline(userId int)  { 

	var mes message.Message
	mes.Type=message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId=userId
	notifyUserStatusMes.Status=message.UserOnline

	data,err:=json.Marshal(notifyUserStatusMes)
	if err!=nil{
		fmt.Println("json.Marshal")
		return
	}
	
	mes.Data=string(data)
	
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal")
		return
	}
	
	tf:=&utils.Transfer{
		Conn:this.Conn,
	}

	//server send the registerresmes to client
	err=tf.WritePkg(data)

	if err!=nil{
		fmt.Println("notifymeonline err=",err)
		return
	}
	return
}

func (this *UserProcess)ServerProcessRegister(mes *message.Message)(err error)  {

	/*
	Send:
		message(tyep,data)
		type:register response 
		data:register response message:user exists or not
	*/

	//create a message struct
	var resMes message.Message
	//message type
	resMes.Type=message.RegisterResMesType
	//create registerresmess struct
	var registerResMes message.RegisterResMes

	//反序列化得到 registermessage struct	
	var registerMes message.RegisterMes
	err=json.Unmarshal([]byte(mes.Data),&registerMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	//检验用户是否已经存在
	err=model.MyUserDao.Register(&registerMes.User)
	if err!=nil{

		if err==model.ERROR_USER_EXISTS{
			fmt.Println("用户已存在")
			registerResMes.Code=505
			registerResMes.Error=model.ERROR_USER_EXISTS.Error()
		}else {
			registerResMes.Code=506
			registerResMes.Error="somewhere havs broken..."
		}

	}else{

		fmt.Println("新用户")
		registerResMes.Code=200
	}

	//序列化register response message
	data,err:=json.Marshal(registerResMes)
	if err!=nil{
		fmt.Println("json.marshal err=",err)
		return
	}

	//填入message data
	resMes.Data=string(data)

	//序列化message
	data,err=json.Marshal(resMes)
	if err!=nil{
		fmt.Println("json.marshal err=",err)
		return
	}

	tf:=&utils.Transfer{
		Conn:this.Conn,
	}

	//server send the registerresmes to client
	err=tf.WritePkg(data)
	fmt.Println("注册响应信息")
	return
}

func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error)  {
	
	//get the data from mes.data and 反序列化
	var loginMes message.LoginMes
	err=json.Unmarshal([]byte(mes.Data),&loginMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}	
	//定义消息体
	var resMes message.Message
	//消息类型
	resMes.Type=message.LoginResMesType
	//消息内容
	var loginResMes message.LoginResMes

	user,err:=model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	
	if err!=nil{

		if err==model.ERROR_USER_NOTEXISTS{
			loginResMes.Code=500
			loginResMes.Error="用户不存在"
		}else if err==model.ERROR_USER_PWD{
			loginResMes.Code=403
			loginResMes.Error="密码错误"
		}else{
			loginResMes.Code=505
			loginResMes.Error="服务器错误...."
		}

	}else{
		loginResMes.Code=200

		this.UserId=loginMes.UserId
		userMgr.AddOnlineUser(this)

		this.NotifyOthersOnlineUser(loginMes.UserId)

		for id,_:=range userMgr.onlineUsers{
			loginResMes.UsersId=append(loginResMes.UsersId,id)
		}

		fmt.Println(user,"登录成功")
	}

	//序列化消息体内容
	data,err:=json.Marshal(loginResMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	resMes.Data=string(data)

	//序列化消息
	data,err=json.Marshal(resMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	//send data
	tf:=&utils.Transfer{
		Conn:this.Conn,
	}
	err=tf.WritePkg(data)
	fmt.Println("发送登录响应信息")
	return
}
