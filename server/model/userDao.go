package model

import(
	"fmt"
	_"net"
	"encoding/json"
	"redigo/redis"
	"chatroom/common/message"
)

var(
	MyUserDao *UserDao
)

type UserDao struct{
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool)(userDao *UserDao)  {

	userDao=&UserDao{
		pool:pool,
	}

	return
}

func (this *UserDao)getUserById(conn redis.Conn,id int)(user *User,err error)  {
	
	res,err:=redis.String(conn.Do("hget","users",id))
	if err!=nil{
		if err==redis.ErrNil{
			err=ERROR_USER_NOTEXISTS
		}
		return
	}

	user=&User{}
	//反序列化
	err=json.Unmarshal([]byte(res),user)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)	
	}

	return//什么错误都没有-->用户存在
}

func (this *UserDao)Login(userId int,userPwd string)(user *User,err error)  {
	
	conn:=this.pool.Get()
	
	defer conn.Close()
	
	user,err=this.getUserById(conn,userId)	
	
	if err!=nil{
		return
	}

	if user.UserPwd!=userPwd{
		err=ERROR_USER_PWD
		return
	}

	return	
}

func (this *UserDao)Register(user *message.User)(err error)  {
	
	conn:=this.pool.Get()
	
	defer conn.Close()
	
	//查询用户是否已经存在
	_,err=this.getUserById(conn,user.UserId)	
	if err==nil{
		err=ERROR_USER_EXISTS
		return
	}

	//序列化用户信息
	data,err:=json.Marshal(user)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//写入数据库
	_,err=conn.Do("hset","users",user.UserId,string(data))
	if err!=nil{
		fmt.Println("conn.Do err=",err)
		return
	}
	
	return	
}