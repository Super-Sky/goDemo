package main

import (
	"fmt"
	"kafkaOfRedis/ProducerServer"
	"strconv"
)

func main() {
	producerInfo := ProducerServer.NewProducerInfo()
	network:= "tcp"
	address:= "127.0.0.1:6379"
	Topic := "topic"
	key := "key1"
	basevalue := "value"
	producerInfo.NewProducerPool(network, address)
	for i :=0;i<100000;i++ {
		value := basevalue + strconv.FormatInt(int64(i),10)
		fmt.Println("value",value)
		producerInfo.Push(Topic, key, value)
	}
}
