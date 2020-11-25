package main

import (
	"fmt"
	"kafkaDemo/ConsumerServer"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	groupID := "group-1"
	fmt.Println("EndBigStr")
	topicName := "oneTopicName"
	brokers := []string{"192.168.100.190:9092"}
	//创建一个消费者
	consumerInfo := ConsumerServer.ConsumerInfo{Brokers: brokers, Topic: topicName, GroupID: groupID, Wg: &wg, Count: 0}
	err := consumerInfo.NewConsumer()
	if err != nil {
		fmt.Println("NewConsumer Err")
		panic(err)
	}
	//消费者接收消息
	err = consumerInfo.ReceiveNews()
	if err != nil {
		fmt.Println("ReceiveNews Err")
		panic(err)
	}
}
