package main

import (
	"RedisDemo/RedisProducerServer"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var NewTime int64
var NewCount int64
var OldTime int64
var OldCount int64
var DiffTime int64
var DiffCount int64
var MaxCount float64
var MinCount float64
var Count int64

func countTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <- ticker.C:
				NewTime = time.Now().UnixNano()
				NewCount = Count
				DiffTime = NewTime - OldTime
				DiffCount = NewCount - OldCount
				OriginalQPS := float64(DiffCount)/float64(DiffTime)
				QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
				log.Println("ProducerServer QPS = ", QPS)
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
	producerInfo := RedisProducerServer.NewProducerInfo()
	producerInfo.RedisConfigInit(network,address)
	producerInfo.NewProducerPool()
	countTimer()
	defer producerInfo.Pool.Close()
	key := "ProducerDemo"
	valuebase := "mxt1_%d"
	for i :=0;i<1000000;i++ {
		value := fmt.Sprintf(valuebase,i)
		Count++
		wg.Add(1)
		producerInfo.Push(key, value, &wg)
	}
	wg.Wait()
}
