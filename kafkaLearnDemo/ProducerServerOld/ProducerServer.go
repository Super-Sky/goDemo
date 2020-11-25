package ProducerServerOld

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
	"time"
)

var Count int32 = 0

type ProducerInfo struct {
	Wg *sync.WaitGroup
	Producer sarama.SyncProducer
	Brokers []string
	Count   int32
	MsgTopic,MessageKey,MessageValue string
}

func configInit() (sarama.SyncProducer,error) {
	config := sarama.NewConfig()
	//  config.ProducerServer.RequiredAcks = sarama.WaitForAll
	//  config.ProducerServer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(strings.Split("192.168.100.190:9092", ","), config)
	return p,err
}

func StartProducer(wg *sync.WaitGroup) {
	defer wg.Done()
	producer,err := configInit()
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage {
		//Topic: "test",//包含了消息的主题
		Partition: int32(0),//
		Key:        sarama.StringEncoder("keysd"),//
	}

	//var value string
	var msgType string
	var tempvalue string
	for i:=0;i<10000;i++ {
		tempvalue += "m"
	}
	tempvalue += "_%d"
	msgType = "pressure1"
	//fmt.Println("msgType =",msgType,",value =",tempvalue)
	msg.Topic = msgType
	for j:=1;j<=10000000;j++{

		tempvalue := fmt.Sprintf(tempvalue,j)
		//将字符串转换为字节数组
		msg.Value = sarama.ByteEncoder(tempvalue)
		//_, err := fmt.Scanf("%s", &value)
		//if err != nil {
		//	break
		//}
		//fmt.Scanf("%s",&msgType)
		//fmt.Println(value)
		//SendMessage：该方法是生产者生产给定的消息
		//生产成功的时候返回该消息的分区和所在的偏移量
		//生产失败的时候返回error
		_, _, err := producer.SendMessage(msg)
		Count ++
		if err != nil {
			fmt.Println(err)
			fmt.Println("Send message Fail")
		}

		//_, _, err := producer.SendMessage(msg)

		//fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}