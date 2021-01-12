package main

import(
	"fmt"
	"os"
	"chatroom/client/process"
)
var userId int
var userPwd string
var userName string
func main()  {
	
	var key int
	var loop=true
	
	for loop{
		fmt.Println(" ==welcome  to chatroom==")
		fmt.Println("\t 1 登录")
		fmt.Println("\t 2 注册")
		fmt.Println("\t 3 退出")
		fmt.Println("\t 选择1-3")
		
		fmt.Scanf("%d\n",&key)
		switch key {
			case 1:
				fmt.Println("login")
				fmt.Println("enter your id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("enter your password")
				fmt.Scanf("%s\n",&userPwd)

				up:=&process.UserProcess{}
				up.Login(userId,userPwd)
				// loop=false
			case 2:
				fmt.Println("register")
				fmt.Println("enter your id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("enter your password")
				fmt.Scanf("%s\n",&userPwd)
				fmt.Println("enter your name")
				fmt.Scanf("%s\n",&userName)

				up:=&process.UserProcess{}
				up.Register(userId,userPwd,userName)
			case 3:
				fmt.Println("quit")
				// loop=false
				os.Exit(0)
			default:
				fmt.Println("enter error try again")
		}
	}
}