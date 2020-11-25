package main

import (
	"fmt"
	"kafkaDemo/AsynProducerServer"
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
				DiffProducerCount := AsynProducerServer.Count - OldProducerCount
				//log.Println("DiffProducerCount",DiffProducerCount)
				//log.Println("NewTime",NewTime)
				//log.Println("OldTime",OldTime)
				if NewTime == OldTime {
					continue
				}
				OriginalQPS := float64(DiffProducerCount)/float64(NewTime -OldTime)
				//log.Println("OriginalQPS",OriginalQPS)
				QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
				if QPS == "+Inf" || QPS == "Nan" {
					log.Println("DiffProducerCount",DiffProducerCount)
					log.Println("NewTime",NewTime)
					log.Println("OldTime",OldTime)
					panic("Err")
				}
				log.Println("单位时间发送数据(SendCount/s) = ", QPS)
				//log.Println("Count----------",AsynProducerServer.Count)
				//runnerGo := runtime.NumGoroutine()
				//fmt.Println("runnerGo",runnerGo)
				QPSNum,err := strconv.ParseFloat(QPS,64)
				if QPS == "+Inf" || err != nil {
					QPSNum = 0
				}
				OldProducerCount = AsynProducerServer.Count
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
		fmt.Println( "MaxCount ", MaxCount, "MinCount", MinCount)
	}()
	var wg sync.WaitGroup
	var producer AsynProducerServer.Producer
	topic := "AsyncFirst1"
	basevalue := "mxt2_%d"
	brokers := []string{"192.168.100.190:9092"}
	producer.ProducerConfigInit(brokers)
	countTimer()
	err := producer.NewProducer()
	if err != nil {
		panic(err)
	}
	for i:=0;i<5000000;i++ {
		value := fmt.Sprintf(basevalue,i)
		wg.Add(1)
		producer.Send(value,topic,&wg)
	}
	defer producer.Producer.AsyncClose()
	wg.Wait()
}
