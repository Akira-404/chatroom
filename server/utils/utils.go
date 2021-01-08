package utils

import(
	"fmt"
	"net"
	_"io"
	_"errors"
	"encoding/json"
	"chatroom/common/message"
	"encoding/binary"
)

type Transfer struct{

	Conn net.Conn
	Buf [8096]byte
}

//读数据，反序列化
func (this *Transfer)ReadPkg()(mes message.Message,err error)  {

	fmt.Println("receive data of client to send....")
	// buf:=make([]byte, 8096)

	//func (c *IPConn) Read(b []byte) (int, error)
	_,err=this.Conn.Read(this.Buf[:4])
	if err!=nil{
		// err=errors.New("read pkg header error")	
		return
	}
	
	//求数据总长 buf[:4]->uint32
	var pkgLen uint32
	pkgLen=binary.BigEndian.Uint32(this.Buf[:4])

	//read the data
	n,err:=this.Conn.Read(this.Buf[:pkgLen])
	if n!=int(pkgLen)||err!=nil{
		// err=errors.New("read pkg body error")	
		return
	}

	//pkglen 反序列化->message.Message
	//func Unmarshal(data []byte, v interface{}) error
	err=json.Unmarshal(this.Buf[:pkgLen],&mes)	
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

//写数据，序列化
func (this *Transfer)WritePkg(data[]byte)(err error)  {

	//send len of data
	var pkgLen uint32
	pkgLen=uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4],pkgLen)

	//写入数据总长
	n,err:=this.Conn.Write(this.Buf[:4])
	if n!=4||err!=nil{
		fmt.Println("conn.Write err=",err)
		return
	}

	//写入所有数据
	n,err=this.Conn.Write(data)
	if n!=int(pkgLen)||err!=nil{
		fmt.Println("conn.Write err=",err)
		return
	}
	return
}