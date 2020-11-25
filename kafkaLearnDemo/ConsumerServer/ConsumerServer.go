package ConsumerServer

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"os"
	"strings"
	"sync"
	"time"
)

var tempMap sync.Map
var readData sync.RWMutex
//var once sync.Once
var m sync.Mutex
var Count int32

type ConsumerInfo struct {
	Wg            *sync.WaitGroup
	Consumer      *cluster.Consumer
	Brokers       []string
	Topic         string
	GroupID       string
	partitionList []int32
	Count         int32
}

func (consumerInfo *ConsumerInfo) NewConsumer() (err error) {
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	saramaConsumer, err := cluster.NewConsumer(consumerInfo.Brokers, consumerInfo.GroupID, strings.Split(consumerInfo.Topic, ","), config)
	if err != nil {
		return
	}
	consumerInfo.Consumer = saramaConsumer
	return nil
}

func (consumerInfo *ConsumerInfo) ReceiveNews() (err error) {
	defer consumerInfo.Consumer.Close()
	//for partition := range consumerInfo.partitionList {
		//ConsumePartition根据主题,分区和给定的偏移量创建了相应分区的消费者
		//如果该分区消费者已经消费了将返回err
		//sarama.OffsetNewest 表示最新消息
		//partitionConsumer, err := consumerInfo.ConsumerProcess.ConsumePartition("oneTopicName", int32(partition), sarama.OffsetNewest)
		//if err != nil {
		//	return nil
		//}
	consumerInfo.Wg.Add(1)
	go consumerInfo.ReceiveNew()
	for msg := range consumerInfo.Consumer.Messages(){
		m.Lock()
		Count++
		m.Unlock()
		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		consumerInfo.Consumer.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset

	}
	return nil
}

func (consumerInfo ConsumerInfo) ReceiveNew() {
	defer consumerInfo.Wg.Done()
	errors := consumerInfo.Consumer.Errors()
	noti := consumerInfo.Consumer.Notifications()
	for {
		select {
		//接收消息通道和错误通道的内容.
		//case msg := <-partitionConsumer.Messages():
		//case <-partitionConsumer.Messages():
			case <- noti:
			case err := <- errors:
				fmt.Println(err)
		}
	}
}

