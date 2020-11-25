package ProducerServer

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type ProducerInfo struct {
	Pool         *redis.Pool
}

type producerMessage struct {
	Topic         string    `json:"topic"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
}

func NewProducerInfo() (producerInfo *ProducerInfo) {
	producerInfo = &ProducerInfo{}
	return
}

func (pi *ProducerInfo) NewProducerPool(network, address string) {
	pool := &redis.Pool{
		MaxIdle: 20,
		IdleTimeout: 240 * time.Second,
		//IdleTimeout: 0,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial(network, address) },
	}
	pi.Pool = pool
}

func (pi *ProducerInfo) Push (Topic, key, value string) {
	msg := &producerMessage{
		Topic: Topic,
		Key: key,
		Value: value,
	}
	coon := pi.Pool.Get()
	defer coon.Close()
	msgJson, err := json.Marshal(msg)
	if err != nil {
		log.Println("Msg Marshal Err: ",err)
	}
	_, err = redis.Int64(coon.Do("LPUSH", Topic, msgJson))
	if err != nil {
		log.Println("Redis Push Error ：", err)
	}
}