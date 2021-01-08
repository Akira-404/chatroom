package process2

import(
	"fmt"
)

var(
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int]*UserProcess
}

func init()  {
	userMgr=&UserMgr{
		onlineUsers:make(map[int]*UserProcess, 1024),
	}
}

func (this *UserMgr)AddOnlineUser(up *UserProcess)  {
	this.onlineUsers[up.UserId]=up
}

func (this *UserMgr)DelOnlineUser(userId int)  {
	delete(this.onlineUsers,userId)
}

func (this *UserMgr)GetOnlineUser()(map[int]*UserProcess)  {
	return this.onlineUsers
}

func (this *UserMgr)GetOnlineUserById(userId int)(up *UserProcess,err error)  {

	up,ok:=this.onlineUsers[userId]
	if !ok{
		err=fmt.Errorf("user%d not find\n",userId)
		return
	}

	return
}

