package RedisSubScribeServer

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
	"time"
)

var Count int64
var once  sync.Once
var startTime int64
var endTime int64

type SubScribeInfo struct {
	network,address string
	conn    redis.Conn
	psc     redis.PubSubConn
}

func NewSubScribeInfo(network,address string) (*SubScribeInfo) {
	subScribeInfo := &SubScribeInfo{network: network,address: address}
	return subScribeInfo
}

func (si *SubScribeInfo) SubScribeInit() (error) {
	c, err := redis.Dial(si.network, si.address)
	if err != nil {
		return err
	}
	si.conn = c
	psc := redis.PubSubConn{Conn: c}
	si.psc = psc
	return nil
}

func (si *SubScribeInfo) ReceiveData(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		switch res := si.psc.Receive().(type) {
		case redis.Message:
			once.Do(func() {
				startTime = time.Now().Unix()
			})
			//channel := res.Channel
			//message := res.Data
			Count++
			//fmt.Printf("channel: %v, message: %v\n", channel, string(message))
		case error:
			fmt.Println("Receive Err", res.Error())
			continue
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)

		}
		if Count == 30000000 {
			endTime = time.Now().Unix()
			diffTime := endTime - startTime
			QPS := 30000000/diffTime
			log.Println("SubScribe QPS: ", QPS)
			break
		}
	}
}

func (si *SubScribeInfo) StartSubScribe(channel interface{}) error {
	err := si.psc.Subscribe(channel)
	if err != nil {
		return err
	}
	return nil
}

//func main() {
//	var wg sync.WaitGroup
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	panic(err)
	//}
	//defer c.Close()
	//psc := redis.PubSubConn{Conn: c}
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for {
	//		switch res := psc.Receive().(type){
	//		case redis.Message:
	//			channel := res.Channel
	//			message := res.Data
	//			fmt.Printf("channel: %v, message: %v\n", channel, string(message))
	//		case error:
	//			fmt.Println("Receive Err",res.Error())
	//			continue
	//		case redis.Subscription:
	//			fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
	//		}
	//	}
	//}()
	//err = psc.Subscribe("c1")
	//if err != nil {
	//	panic(err)
	//}
	//defer psc.Unsubscribe()
	//wg.Wait()
//}
