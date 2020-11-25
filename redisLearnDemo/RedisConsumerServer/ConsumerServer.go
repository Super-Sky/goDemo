package RedisConsumerServer

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
	"time"
)

var Count int64 = 0
var once sync.Once

type redisInfo struct {
	network      string
	address      string
	Conn         redis.Conn
	TimeOUt      int32
}

func NewConsumerInfo() (consumerInfo *redisInfo) {
	consumerInfo = &redisInfo{}
	return
}

func (ri *redisInfo) RedisConfigInit (networt, address string, timeout int32){
	ri.network = networt
	ri.address = address
	ri.TimeOUt = timeout
}

func (ri *redisInfo) NewConsumer () (err error) {
	conn, err := redis.Dial(ri.network, ri.address)
	if err != nil {
		return err
	}
	ri.Conn = conn
	return nil
}

func (ri *redisInfo) Receive(key string)  {
	defer ri.Conn.Close()
	var startTime int64
	var endTime   int64
	for {
		_ , err := redis.Values(ri.Conn.Do("BLPOP", key, ri.TimeOUt))
		once.Do(func() {
			startTime = time.Now().Unix()
		})
		if err != nil {
			log.Println("BRPOP Error: ",err)
		} else {
			Count ++
			//log.Printf("receive Key %s Value %s", string(result[0].([]byte)), string(result[1].([]byte)))
		}
		if Count == 30000000 {
			endTime = time.Now().Unix()
			qps := float64(Count)/float64(endTime-startTime)
			log.Println("Avg QPS : ",qps)
			break
		}
	}
}
