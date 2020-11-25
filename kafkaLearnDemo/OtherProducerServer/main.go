package main

import (
"fmt"
"github.com/Shopify/sarama"
)

func main() {
	//创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"192.168.100.190:9092"},nil)
	if err != nil {
		panic(err)
	}
	//Partitions方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions("msgtype")
	if err != nil {
		panic(err)
	}
	fmt.Println("partitionList",partitionList)
	for partition := range partitionList {
		//ConsumePartition根据主题,分区和给定的偏移量创建了相应分区的消费者
		//如果该分区消费者已经消费了将返回err
		//sarama.OffsetNewest 表示最新消息
		partitionConsumer, err := consumer.ConsumePartition("msgtype", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		fmt.Println("start print")
		defer partitionConsumer.AsyncClose()
		for {
			select {
			//接收消息通道和错误通道的内容.
			case msg := <-partitionConsumer.Messages():
				fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
			case err := <-partitionConsumer.Errors():
				fmt.Println(err.Err)
			}
		}
	}
}
