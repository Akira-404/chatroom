package main

import(
	"redigo/redis"
	"time"
)

var pool *redis.Pool

//初始化redis 连接池
func initPool(address string,maxIdle,maxActive int,idleTimeout time.Duration)  {
	
	pool=&redis.Pool{
		MaxIdle:maxIdle,
		MaxActive :maxActive,
		IdleTimeout:idleTimeout,

		//func Dial(network, address string, options ...DialOption) (Conn, error)
		/*usage:
		Dial connects to the Redis server at the given network and address using the specified options.
		*/
	
		Dial:func()(redis.Conn,error){
			return redis.Dial("tcp",address)
		},
	}	
}