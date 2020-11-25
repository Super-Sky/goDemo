package main

import (
	"fmt"
	"kafkaDemo/AsynConsumerServer"
	"log"
	"strconv"
	"sync"
	"time"
)

var OldProducerCount int32
var OldTime          int64
var NewTime          int64
var MaxCount         float64
var MinCount         float64
var TempFloat        float64

func countTimer()  {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				NewTime = time.Now().Unix()
				DiffConsumerCount := AsynConsumerServer.Count - OldProducerCount
				//log.Println("DiffConsumerCount",DiffConsumerCount)
				//log.Println("NewTime",NewTime)
				//log.Println("OldTime",OldTime)
				OriginalQPS := float64(DiffConsumerCount)/float64(NewTime -OldTime)
				QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
				log.Println("Consumer QPS = ", QPS)
				//log.Println("Count----------",AsynConsumerServer.Count)
				//runnerGo := runtime.NumGoroutine()
				//fmt.Println("runnerGo",runnerGo)
				QPSNum,err := strconv.ParseFloat(QPS,64)
				if QPS == "+Inf" || err != nil {
					QPSNum = 0
				}
				OldProducerCount = AsynConsumerServer.Count
				OldTime = NewTime
				if QPSNum >= MaxCount {
					MaxCount = QPSNum
				}
				if QPSNum <= MinCount && QPSNum != TempFloat {
					MinCount = QPSNum
				}
			}
		}
	}()
}


func main() {
	defer func() {
		fmt.Printf(" %d messagesCount %v MaxCount %v MinCount \n",
			AsynConsumerServer.Count, MaxCount, MinCount)
	}()
	var wg sync.WaitGroup
	var consumer AsynConsumerServer.Consumer
	groupId := "Group-1"
	topics := []string{"AsyncFirst2"}
	brokers := []string{"192.168.100.190:9092"}
	consumer.ConsumerConfigInit(groupId,topics,brokers)
	countTimer()
	err := consumer.NewConsumer()
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go consumer.Receive(&wg)
	wg.Wait()
}