package model

import(
	"errors"
)

//自定义错误类型
var(
	ERROR_USER_NOTEXISTS=errors.New("user is not exists")
	ERROR_USER_EXISTS=errors.New("user exists")
	ERROR_USER_PWD=errors.New("password is incorrect")
)