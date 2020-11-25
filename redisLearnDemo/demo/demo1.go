package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
	}
	reply, err := redis.Values(conn.Do("HKEYS", "warnWordId:1"))
	if err !=nil{
		log.Println(err)
	}
	log.Println(reply)
	//for _, str :=range reply {
	//	log.Println(string(str.([]byte)))
	//}
}
