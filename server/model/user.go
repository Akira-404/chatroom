package model

import(
)

type User struct{
	UserId int`json:"userid"`
	UserPwd string`json:"userpwd"`
	UserName string`json:"username"`
}