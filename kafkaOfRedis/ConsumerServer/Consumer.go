package ConsumerServer

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)
var once sync.Once

type consumerInfo struct {
	Conn         redis.Conn
	TimeOut      int32
	Offset       int64
	status       string
	markID       int64
	runMark      bool
	network      string
	address      string
}

func NewConsumerInfo() (consumer *consumerInfo)  {
	consumer = &consumerInfo{}
	return
}

func (ci *consumerInfo) NewConsumer(network, address string) {
	conn, err := redis.Dial(network, address)
	if err != nil {
		log.Println("Redis Dial Err: ",err)
	}
	ci.network = network
	ci.address = address
	ci.Conn = conn
	ci.status = "0"
	ci.markID = time.Now().UnixNano()
	ci.runMark = false
}

// GroupId:Topic:status     "0" 表示没有人在取  "1"表示已经有人在取
func (ci *consumerInfo) GetConsumerGroupStatus(Topic, GroupId string) {
	StatusKey := GroupId+":"+ Topic+ ":status"
	status , err := redis.String(ci.Conn.Do("GET", StatusKey))
	if err != nil && err != redis.ErrNil {
		log.Println("GET status Err ", err)
	}
	if err == redis.ErrNil {
		status = "0"
		ci.runMark = true
	}
	if status == "0" {
		ci.runMark = true
	}
	ci.status = status
}

func (ci *consumerInfo) SetConsumerGroupStatus(Topic, GroupId, status string) {
	StatusKey := GroupId+":"+ Topic+ ":status"
	_ , err := redis.String(ci.Conn.Do("SET", StatusKey, status))
	if err != nil {
		log.Println("SET status Err: ", err)
	}
	ci.status = status
}

// GroupId:Topic:markId  获取markId 查看是否是自己的markId如果是那就说明可以继续取，如果不是就说明已经有人在取
func (ci *consumerInfo) SetConsumerGroupFuncId(Topic, GroupId string) {
	MarkIdKey := GroupId+":"+ Topic+ ":markId"
	_ , err := redis.String(ci.Conn.Do("SET", MarkIdKey, ci.markID))
	if err != nil {
		log.Println("SET markId Err: ", err)
	}
}

func (ci *consumerInfo) GetConsumerGroupFuncId(Topic, GroupId string) {
	MarkIdKey := GroupId+":"+ Topic+ ":markId"
	MarkId , err := redis.Int64(ci.Conn.Do("GET", MarkIdKey))
	if err != nil && err != redis.ErrNil {
		log.Println("GET markId Err ", err)
	}
	if err == redis.ErrNil {
		MarkId = ci.markID
	}
	if ci.markID == MarkId {
		ci.runMark = true
	}else {
		ci.runMark = false
	}
}

func (ci *consumerInfo) DeleteConsumerGroupFuncId(Topic, GroupId string) {
	MarkIdKey := GroupId+":"+ Topic+ ":markId"
	_ , err := redis.Int64(ci.Conn.Do("DEL", MarkIdKey))
	if err != nil {
		log.Println("DEL markId Err ", err)
	}

}

func (ci *consumerInfo) GetOffset(Topic, GroupId string) {
	Offset, err := redis.Int64(ci.Conn.Do("hget", GroupId, Topic))
	if err != nil && err != redis.ErrNil {
		log.Println("HGET Err ", err)
	}
	if err == redis.ErrNil {
		//log.Println("First or Nil HGET ", GroupId, " ", Topic)
		Offset = 0
	}
	ci.Offset = Offset
}

func (ci *consumerInfo) MarkOffset(Topic, GroupId string, Offset int64) {
	_, err := redis.Int64(ci.Conn.Do("hset",GroupId,Topic, strconv.FormatInt(Offset, 10)) )
	if err != nil {
		log.Println("HSET Err: ",err)
	}
}

func (ci *consumerInfo) Receive(Topic, GroupId string)  {
	defer ci.Conn.Close()
	defer ci.SetConsumerGroupStatus(Topic, GroupId, "0")
	defer ci.DeleteConsumerGroupFuncId(Topic, GroupId)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for {
			select {
			case <-c:
				log.Println("Exit....")
				ci.NewConsumer(ci.network, ci.address)
				ci.SetConsumerGroupStatus(Topic, GroupId, "0")
				ci.DeleteConsumerGroupFuncId(Topic, GroupId)
				ci.Conn.Close()
				os.Exit(1)
			}
		}
	}()
	for {
		ci.GetConsumerGroupStatus(Topic, GroupId)
		if ci.status == "1" {
			ci.GetConsumerGroupFuncId(Topic, GroupId)
		}
		if !ci.runMark {
			continue
		}
		ci.GetOffset(Topic,GroupId)
		values, err := redis.Values(ci.Conn.Do("lrange", Topic, strconv.FormatInt(ci.Offset, 10), strconv.FormatInt(ci.Offset, 10)))
		if err != nil && err.Error() != "ERR value is not an integer or out of range" {
			log.Println(err)
		}
		if err != nil && err.Error() == "ERR value is not an integer or out of range" {
			continue
		}
		if len(values) == 0 {
			continue
		}
		once.Do(func() {
			ci.SetConsumerGroupStatus(Topic, GroupId, "1")
			ci.GetConsumerGroupFuncId(Topic, GroupId)
			if ci.runMark {
				ci.SetConsumerGroupFuncId(Topic, GroupId)
			}
		})
		if ci.runMark {
			log.Println(string(values[0].([]byte)))
			ci.MarkOffset(Topic, GroupId, ci.Offset-1)
		}
	}
}