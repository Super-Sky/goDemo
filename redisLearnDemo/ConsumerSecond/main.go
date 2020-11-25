package main

import (
	"RedisDemo/RedisConsumerServer"
	"log"
	"strconv"
	"time"
)

var NewTime int64
var NewCount int64
var OldTime int64
var OldCount int64
var DiffTime int64
var DiffCount int64
var MaxCount float64
var MinCount float64

func countTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <- ticker.C:
				NewTime = time.Now().Unix()
				NewCount = RedisConsumerServer.Count
				DiffTime = NewTime - OldTime
				DiffCount = NewCount - OldCount
				OriginalQPS := float64(DiffCount)/float64(DiffTime)
				QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
				log.Println("Consumer QPS = ", QPS)
				QPSNum,err := strconv.ParseFloat(QPS,64)
				if QPS == "+Inf" || err != nil {
					QPSNum = 0
				}
				OldCount = NewCount
				OldTime = NewTime
				if QPSNum >= MaxCount {
					MaxCount = QPSNum
				}
				if QPSNum <= MinCount && QPSNum != float64(0) {
					MinCount = QPSNum
				}
			}
		}
	}()
}

func main() {
	network := "tcp"
	address := "127.0.0.1:6379"
	timeout := 0
	consumerInfo := RedisConsumerServer.NewConsumerInfo()
	consumerInfo.RedisConfigInit(network,address, int32(timeout))
	err := consumerInfo.NewConsumer()
	if err != nil {
		panic(err)
	}
	key := "ProducerDemo"
	countTimer()
	consumerInfo.Receive(key)
	log.Println("Max QPS : ",MaxCount)
	log.Println("Min QPS: ",MinCount)
}
