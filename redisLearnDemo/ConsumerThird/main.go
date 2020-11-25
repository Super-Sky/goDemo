package main

import (
	"RedisDemo/RedisConsumerServer"
)

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
	consumerInfo.Receive(key)
}
