package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis" //引入redis包
)

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //和数据库最大的连接数 0 表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接
			return redis.Dial(
				"tcp",
				"localhost:6379",
				redis.DialDatabase(1),
				redis.DialPassword("145622"))
		},
	}
}

func main() {
	//从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "汤姆")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	//取出
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

	//如果要用pool连接，一定要保证连接池是没关闭的
	pool.Close()
	conn2 := pool.Get()

	_, err = conn2.Do("set", "name2", "汤姆猫2")
	if err != nil {
		fmt.Println("conn.Do err1=", err)
		return
	}
}
