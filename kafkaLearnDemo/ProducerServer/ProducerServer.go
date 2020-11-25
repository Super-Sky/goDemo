package ProducerServer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

//var COUNT int32 = 0
//var m sync.Mutex

type ProducerInfo struct {
	Wg *sync.WaitGroup
	Producer sarama.AsyncProducer
	Brokers []string
	Count   int32
	MsgTopic,MessageKey,MessageValue string
}

func (pr *ProducerInfo) NewProducer() (err error){
	//设置配置
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_11_0_0

	//使用配置,新建一个异步生产者
	producers, err := sarama.NewAsyncProducer(pr.Brokers, config)
	if err != nil {
		return nil
	}
	pr.Producer = producers
	return nil
}

func (pr *ProducerInfo) Send(msgValue string){
	pr.MessageValue = msgValue
	defer pr.Wg.Done()
	//构建发送的消息，
	msg := &sarama.ProducerMessage {
		//Topic: "test",//包含了消息的主题
		Partition: int32(0),//
		Key:       sarama.StringEncoder(pr.MessageKey),
		Value:     sarama.ByteEncoder(pr.MessageValue),
		Topic:     pr.MsgTopic,
	}
	fmt.Printf("Send Msg Partition:%d Key:%s Topic:%s \n", msg.Partition, msg.Key, msg.Topic)
	//_, _, err := pr.ProducerProcess.SendMessage(msg)
	pr.Producer.Input() <- msg
	//m.Lock()
	//COUNT++
	//m.Unlock()
	return
}



