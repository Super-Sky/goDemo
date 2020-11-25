package RedisPublishServer

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

var Count int64


type PublishInfo struct {
	Pool       *redis.Pool
}

func (pbl *PublishInfo)GetPool(network, address string)  {
	pool := &redis.Pool{
		MaxIdle: 20,
		IdleTimeout: 240 * time.Second,
		//IdleTimeout: 0,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial(network, address) },
	}
	pbl.Pool = pool
}

func (pbl *PublishInfo)Publish(chanel, msg string, wg *sync.WaitGroup)  {
	defer wg.Done()
	for i:=0;i<1000000;i++ {
		conn := pbl.Pool.Get()
		_, err := conn.Do("PUBLISH", chanel, msg)
		conn.Close()
		Count++
		if err != nil{
			panic(err)
		}
		//fmt.Println("reply1",reply)
	}
}
//func main() {
//	c, err := redis.Dial("tcp", "127.0.0.1:6379")
//
//	defer c.Close()
//
//	reply, err := c.Do("PUBLISH", "c1", "hello")
//	if err != nil{
//		panic(err)
//	}
//	fmt.Println("reply1",reply)
//	reply, err = c.Do("PUBLISH", "c1", "world")
//	if err != nil{
//		panic(err)
//	}
//	fmt.Println("reply1",reply)
//	reply, err = c.Do("PUBLISH", "c1", "goodbye")
//	if err != nil{
//		panic(err)
//	}
//	fmt.Println("reply1",reply)
//
//}