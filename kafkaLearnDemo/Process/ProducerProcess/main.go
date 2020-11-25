package main

import (
	"fmt"
	"kafkaDemo/ProducerServer"
	"sync"
)

var m sync.RWMutex
var Count int32 = 0

func tempfunc(p ProducerServer.ProducerInfo) {
	//defer p.Wg.Done()
	for {
		select {
		case suc := <-p.Producer.Successes():
			if suc != nil {
				fmt.Printf("succeed, offset=%d, timestamp=%s, partitions=%d\n", suc.Offset, suc.Timestamp.String(), suc.Partition)
				//fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			}
		case fail := <-p.Producer.Errors():
			if fail != nil {
				fmt.Printf("error= %v\n", fail.Err)
			}
		}
	}
}

func GetBigStr() string{
	baseStr := "m"
	bigStr := ""
	for i :=0;i<1000;i++ {
		bigStr += baseStr
	}
	return bigStr
}

func main() {
	var wg sync.WaitGroup
	bigStr := GetBigStr()
	msgKey := "oneMsgKey"
	baseMsgValue := bigStr+"oneMsgValue__%d"
	fmt.Println("EndBigStr")
	topicName := "oneTopicName"
	brokers := []string{"192.168.100.190:9092"}
	//创建一个生产者
	producerInfo := ProducerServer.ProducerInfo{Brokers: brokers, MsgTopic: topicName, MessageKey: msgKey, Wg: &wg, Count: 0}
	err := producerInfo.NewProducer()
	if err != nil {
		fmt.Println("NewProducer Err")
		panic(err)
	}
	producerInfo.Wg.Add(1)
	defer producerInfo.Producer.AsyncClose()
	//go tempfunc(producerInfo)
	for i :=0;i<1000000;i++ {
		msgValue := fmt.Sprintf(baseMsgValue,i)
		wg.Add(1)
		// 生产者发送消息
		m.Lock()
		Count++
		m.Unlock()
		go producerInfo.Send(msgValue)
	}
	tempfunc(producerInfo)
	wg.Wait()
}
