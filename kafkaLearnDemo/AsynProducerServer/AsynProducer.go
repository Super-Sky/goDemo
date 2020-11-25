package AsynProducerServer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

type Producer struct {
	config   *sarama.Config
	Producer sarama.AsyncProducer
	Topic    string
	Value    string
	brokers  []string
}

var Count int32 = 0
var m sync.RWMutex

func (pr *Producer)ProducerConfigInit(brokers []string)  {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1
	pr.config = config
	pr.brokers = brokers
	fmt.Println("start make producer")
}

func (pr *Producer)NewProducer() (err error) {
	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(pr.brokers, pr.config)
	if e != nil {
		fmt.Println(e)
		return err
	}
	pr.Producer = producer
	return nil
}

func (pr *Producer)ReciveSucesses(wg *sync.WaitGroup)  {
	defer wg.Done()
	//循环判断哪个通道发送过来数据.
	for{
		select {
		case <-pr.Producer.Successes():
			//value,_ := suc.Value.Encode()
			//fmt.Println("offset: ", suc.Offset,"Topic",suc.Topic ,"Value: ", string(value))
		case fail := <-pr.Producer.Errors():
			fmt.Println("err: ", fail.Err)
		}
	}
}

func (pr *Producer)Send(value,Topic string,wg *sync.WaitGroup) {
	defer wg.Done()
	// 发送的消息,主题。
	// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
	msg := &sarama.ProducerMessage{
		Topic: Topic,
	}
	//将字符串转化为字节数组
	msg.Value = sarama.ByteEncoder(value)
	//fmt.Println(value)

	//使用通道发送
	for i:=0;i<10000000;i++ {
		Count++
		pr.Producer.Input() <- msg
	}
	//fmt.Println("Count=======",Count)
}

//func SaramaProducer()  {


	//defer producer.AsyncClose()

	//循环判断哪个通道发送过来数据.
	//fmt.Println("start goroutine")
	//go func(p sarama.AsyncProducer) {
	//	for{
	//		select {
	//		case  <-p.Successes():
	//			//fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
	//		case fail := <-p.Errors():
	//			fmt.Println("err: ", fail.Err)
	//		}
	//	}
	//}(producer)

	//var value string
	//for i:=0;;i++ {
	//
	//}
//}
