package main

import (
	"RedisDemo/RedisSubScribeServer"
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

func countTimer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <- ticker.C:
				NewTime = time.Now().Unix()
				NewCount = RedisSubScribeServer.Count
				DiffTime = NewTime - OldTime
				DiffCount = NewCount - OldCount
				OriginalQPS := float64(DiffCount)/float64(DiffTime)
				QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
				log.Println("SubScribe QPS = ", QPS)
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
	subScribeInfo := RedisSubScribeServer.NewSubScribeInfo(network,address)
	err := subScribeInfo.SubScribeInit()
	if err != nil {
		panic(err)
	}
	countTimer()
	wg.Add(1)
	go subScribeInfo.ReceiveData(&wg)
	err = subScribeInfo.StartSubScribe("publish")
	if err != nil {
		panic(err)
	}
	wg.Wait()
}
