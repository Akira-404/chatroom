package process

import(
	"fmt"
	"net"
	"os"
	"encoding/json"
	"chatroom/common/message"
	"chatroom/client/utils"
	"encoding/binary"
)

type UserProcess struct{

}

func (this *UserProcess)Register(userId int,userPwd string,userName string) (err error) {

	//link server
	conn,err:=net.Dial("tcp","0.0.0.0:8889")
	if err!=nil{
		fmt.Println("net.Dial err=",err)
		return
	}
	
	defer conn.Close()

	//send message to server

	//message struct
	var mes message.Message
	mes.Type=message.RegisterMesType

	//registermestype struct
	var registerMes message.RegisterMes
	registerMes.User.UserId=userId
	registerMes.User.UserPwd=userPwd
	registerMes.User.UserName=userName

	//register序列化
	data,err:=json.Marshal(registerMes)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//赋值message data
	mes.Data=string(data)
	
	//message序列化
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	tf:=&utils.Transfer{
		Conn:conn,
	}

	err=tf.WritePkg(data)
	if err!=nil{
		fmt.Println("tf.Write err=",err)
		return
	}
	
	mes,err=tf.ReadPkg()
	if err!=nil{
		fmt.Println("tf.ReadPkg err=",err)
		return
	}

	//将mesresmessage 反序列化
	var registerResMes message.RegisterResMes
	err=json.Unmarshal([]byte(mes.Data),&registerResMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	if registerResMes.Code==200{
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	}else{
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return
}
	
func (this *UserProcess)Login(userId int,userPwd string) (err error) {
	
	//connnect server
	//在网络network上连接地址address，并返回一个Conn接口 err
	conn,err:=net.Dial("tcp","0.0.0.0:8889")
	if err!=nil{
		fmt.Println("net dial err=",err)
		return
	}

	defer conn.Close()
	
	//创建信息总体结构体
	var mes message.Message
	//赋值信息类型
	mes.Type=message.LoginMesType

	//创建登录信息结构体
	var loginMes message.LoginMes
	//赋值
	loginMes.UserId=userId
	loginMes.UserPwd=userPwd

	//封装信息------------------------
	//登录信息结构体序列化
	data,err:=json.Marshal(loginMes)
	if err!=nil{
		fmt.Println("json marshal err=",err)
		return
	}

	//赋值信息内容
	mes.Data=string(data)
	//信息总体结构体序列化
	data,err=json.Marshal(mes)
	if err!=nil{
		fmt.Println("json marshal err=",err)
		return
	}

	var pkgLen uint32
	pkgLen=uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	//封装信息------------------------
	
	//send len of data
	//func (c *IPConn) Write(b []byte) (int, error)
	n,err:=conn.Write(buf[:4])
	if err!=nil||n!=4{
		fmt.Println("conn.Write err",err)
		return
	}
	
	fmt.Println("信息长度发送完毕")
	fmt.Printf("长度=%d,数据内容=%s\n",len(data),string(data))

	//send data
	_,err=conn.Write(data)
	if err!=nil{
		fmt.Println("conn.Write err",err)
		return
	}
	//Processing messages returned by the server
	tf:=&utils.Transfer{
		Conn:conn,
	}	
	mes,err=tf.ReadPkg()	
	if err!=nil{
		fmt.Println("readPkg err",err)
		return
	}
	
	//将mse的data部分反系列化成LoginResMes
	var loginResMes message.LoginResMes
	err=json.Unmarshal([]byte(mes.Data),&loginResMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	
	//查看回馈消息code
	if loginResMes.Code==200{
		//init curuser
		CurUser.Conn=conn
		CurUser.UserId=userId
		CurUser.UserStatus=message.UserOnline

		fmt.Println("登录成功")
		//show online users list
		fmt.Println("在线人员列表:")
		for _,v:=range loginResMes.UsersId{
			if v==userId{
				continue
			}
			fmt.Println("用户id \t",v)

			user:=&message.User{
				UserId:v,
				UserStatus:message.UserOnline,
			}
			onlineUsers[v]=user
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)
		for{
			ShowMenu()
		}
	}else if loginResMes.Code==500{
		fmt.Println(loginResMes.Error)
	}
	
	return
}