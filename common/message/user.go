package message

import(
)

//用户信息
type User struct{
	UserId int`json:"userid"`
	UserPwd string`json:"userpwd"`
	UserName string`json:"username"`
	UserStatus int `json:"userstatus"`
}