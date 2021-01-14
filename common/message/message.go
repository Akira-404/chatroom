package message

//type constant
const(
	LoginMesType="LoginMes"
	LoginResMesType="LoginResMes"
	RegisterMesType="RegisterMes"
	RegisterResMesType="RegisterResMes"
	NotifyUserStatusMesType="NotifyUserStatusMes"
	SmsMesType="SmsMes"
	PersonalMesType="PersonalMesType"
	PersonalResMesType="PersonalResMesType"
)

const(
	UserOnline=iota
	UserOffline
	UserBusyStatus
)

//信息总体:1 信息类型，2 信息内容
type Message struct{
	Type string `json:"type"`
	Data string`json:"data"`
}

//登录信息
type LoginMes struct{
	UserId int`json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

//登录反馈信息
type LoginResMes struct{
	Code int`json:"code"` //500:no register,200:login success
	Error string`json:"error"`
	UsersId []int
}

//注册信息
type RegisterMes struct{
	User `json:"user"`
}

//注册反馈信息
type RegisterResMes struct{
	Code int`json:"code"` //500:no register,200:login success
	Error string`json:"error"`
}

//服务器推送用户状态变换信息
type NotifyUserStatusMes struct{
	UserId int`json:"userid"` 
	Status int `json:"status"`
}

//send message
type SmsMes struct{
	Content string`json:"content"`
	User
}

type PersonalMes struct{
	
	SUserId int`json:"suserid"`
	Content string `json:"content"`
	RUserId int `json:"ruserid"`
}

type PersonalResMes struct{
	
	Code int `json:"code"`//200:success,201:failed
	Error error `json:"error"`
}