package main

import (
	"kafkaDemo/AsynConsumerServer"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var consumer AsynConsumerServer.Consumer
	groupId := "Group-1"
	topics := []string{"AsyncFirst"}
	brokers := []string{"192.168.100.190:9092"}
	consumer.ConsumerConfigInit(groupId,topics,brokers)
	err := consumer.NewConsumer()
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go consumer.Receive(&wg)
	wg.Wait()
}