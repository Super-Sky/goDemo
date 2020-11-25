package RedisProducerServer

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
	"time"
)

var Count int64 = 0

type redisInfo struct {
	m            sync.Mutex
	network      string
	address      string
	Pool         *redis.Pool
	TimeOUt      int32
}

func NewProducerInfo() (producerInfo *redisInfo) {
	producerInfo = &redisInfo{}
	return
}

func (ri *redisInfo) RedisConfigInit (network, address string){
	ri.network = network
	ri.address = address
}

func (ri *redisInfo) NewProducerPool () {
	pool := &redis.Pool{
		MaxIdle: 20,
		IdleTimeout: 240 * time.Second,
		//IdleTimeout: 0,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial(ri.network, ri.address) },
	}
	ri.Pool = pool
}

func (ri *redisInfo) Push (key, valuebase string, wg *sync.WaitGroup) {
	defer wg.Done()

	//ri.m.Lock()
	coon := ri.Pool.Get()
	//err := coon.Send("LPUSH", key, value)
	//if err != nil {
	//	log.Println("Redis Send Error ：", err)
	//}
	//err = coon.Flush()
	//if err != nil {
	//	log.Println("Redis Flush Error ：", err)
	//}
	//index, err := redis.Int64(coon.Receive())
	for i:=0;i<1000000;i++ {
		value := fmt.Sprintf(valuebase,i)
		_, err := redis.Int64(coon.Do("LPUSH", key, value))
		Count ++
		if err != nil {
			log.Println("Redis Receice Error ：", err)
		}
	}
	coon.Close()
	//ri.m.Unlock()

	//else {
	//	log.Printf("Receive Key %s Value %v Index %d", key, value, index)
	//}
}