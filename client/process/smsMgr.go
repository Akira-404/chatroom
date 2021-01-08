package process

import(
	"fmt"	
	"chatroom/common/message"	
	"encoding/json"	
)

func outputGroupMes(mes *message.Message)  {
	
	var smsMes message.SmsMes
	err:=json.Unmarshal([]byte(mes.Data),&smsMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err.Error())
		return
	}

	info:=fmt.Sprintf("user id:\t%d,say everybody:\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}