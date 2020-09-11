//1.Redis中对string的操作
//2.Redis中Hash的使用
//3.Redis中List的使用
//4.Redis中Set的使用
//5.Redis中Hash的基本使用
//6.Go连接到Redis进行string和hash的操作
//7.Redis连接池的使用

package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis" //引入redis包
)

func main() {
	//通过go向redis 写入数据和读取数据
	//1. 链接到redis     //DialOption参数可以配置选择数据库、连接密码等
	//conn,err := redis.Dial("tcp","127.0.0.1:6379")
	conn, err := redis.Dial(
		"tcp",
		"127.0.0.1:6379",
		redis.DialDatabase(1),
		redis.DialPassword("145622"))

	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	fmt.Println("conn succ", conn)
	defer conn.Close() //关闭

	//2.通过go向redis写入数据 string [key-val]]
	_, err = conn.Do("Set", "name", "tom1")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	fmt.Println("操作ok")

	//3.通过go向redis读取数据 string [key-val]]
	//返回的r是interface{}
	//因为name对应的值为string 因此需要转换
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	fmt.Println("操作ok", r)
}
