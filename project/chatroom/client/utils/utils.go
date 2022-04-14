package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/Learn-Go/project/chatroom/common/message"
	"net"
)

//将这些方法关联到结构体中
type Transfer struct {
	//分析应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时的缓冲

}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err=", err)
		//err = errors.New("read pkg header error") //自定义error
		return
	}
	//根据读到的buf[:4]长度 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据 pkgLen 读取内容
	n, err := this.Conn.Read(this.Buf[:pkgLen]) //把从conn独到的[:pkgLen]这个长度放到buf中，并检验读取的长度
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn.Read fail err=", err)
		//err = errors.New("read pkg body error")
		return
	}

	//把pkgLen反序列化成 → message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //&mes
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail", err)
		return
	}

	//发送 data 本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return

}
