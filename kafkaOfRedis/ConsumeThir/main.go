package main

import (
	"kafkaOfRedis/ConsumerServer"
)

func main() {
	consumerInfo := ConsumerServer.NewConsumerInfo()
	network := "tcp"
	address := "127.0.0.1:6379"
	Topic := "topic"
	GroupId := "group-2"
	consumerInfo.NewConsumer(network,address)
	consumerInfo.Receive(Topic, GroupId)
}
