package process2

import(
	"fmt"
)

var(
	userMgr *UserMgr
)
/*
map[UserId]=*UserProcess{Conn net.Conn,UserId int}
*/
type UserMgr struct{
	onlineUsers map[int]*UserProcess
}

func init()  {
	fmt.Println("初始化用户列表")
	userMgr=&UserMgr{
		onlineUsers:make(map[int]*UserProcess, 1024),
	}
}
//添加在线用户
func (this *UserMgr)AddOnlineUser(up *UserProcess)  {
	fmt.Println("添加用户到用户列表,id=",up.UserId)
	this.onlineUsers[up.UserId]=up
}

//删除在线用户
func (this *UserMgr)DelOnlineUser(userId int)  {
	fmt.Println("删除用户到用户id=",userId)
	delete(this.onlineUsers,userId)
}

//获取在线用户
func (this *UserMgr)GetOnlineUser()(map[int]*UserProcess)  {
	return this.onlineUsers
}

//获取在线用户id
func (this *UserMgr)GetOnlineUserById(userId int)(up *UserProcess,err error)  {

	up,ok:=this.onlineUsers[userId]
	if !ok{
		err=fmt.Errorf("user%d not find\n",userId)
		return
	}
	
	return
}

