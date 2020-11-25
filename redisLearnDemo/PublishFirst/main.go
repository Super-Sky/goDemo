package main

import (
	"RedisDemo/RedisPublishServer"
	"log"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var once sync.Once
var NewTime int64
var NewCount int64
var OldTime int64
var OldCount int64
var DiffTime int64
var DiffCount int64
var MaxCount float64
var MinCount float64
var startTime int64
var endTime   int64

func countTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <- ticker.C:
				NewTime = time.Now().Unix()
				NewCount = RedisPublishServer.Count
				DiffTime = NewTime - OldTime
				DiffCount = NewCount - OldCount
				OriginalQPS := float64(DiffCount)/float64(DiffTime)
				QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
				log.Println("Publish QPS = ", QPS)
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
	var publishInfo RedisPublishServer.PublishInfo
	network :="tcp"
	address :="127.0.0.1:6379"
	chanel :="publish"
	msg :="FirstMsg"
	publishInfo.GetPool(network, address)
	countTimer()
	for i:=0;i<15;i++ {
		once.Do(func() {
			startTime = time.Now().Unix()
		})
		wg.Add(1)
		go publishInfo.Publish(chanel, msg, &wg)
	}
	wg.Wait()
	defer func() {
		endTime = time.Now().Unix()
		diffTime := endTime - startTime
		QPS := 15000000/diffTime
		log.Println("Publish QPS ", QPS)
	}()
	defer publishInfo.Pool.Close()
}