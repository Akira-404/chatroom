package model

import(
	"net"
	"chatroom/common/message"
)

//当前在线用用户
type CurUser struct{
	Conn net.Conn
	message.User
}