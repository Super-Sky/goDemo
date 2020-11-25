package AsynConsumerServer

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var Count int32 = 0
var once sync.Once
var StartTime int64
var EndTIme int64

type Consumer struct {
	config     *cluster.Config
	consumer   *cluster.Consumer
	brokers    []string
	topics     []string
	groupId    string
}

func (cos *Consumer)ConsumerConfigInit(groupId string,topics []string,brokers []string)  {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	cos.config = config
	cos.groupId = groupId
	cos.topics = topics
	cos.brokers = brokers
}

func (cos *Consumer)NewConsumer() (err error)  {
	// init consumer
	consumer, err := cluster.NewConsumer(cos.brokers, cos.groupId, cos.topics, cos.config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", cos.groupId, err)
		return
	}
	cos.consumer = consumer
	return nil
}

func (cos *Consumer)Receive(wg *sync.WaitGroup)  {
	defer wg.Done()
	defer cos.consumer.Close()
	// trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	// consume errors
	go func() {
		for err := range cos.consumer.Errors() {
			log.Printf("%s:Error: %s\n", cos.groupId, err.Error())
		}
	}()
	// consume notifications
	go func() {
		for ntf := range cos.consumer.Notifications() {
			log.Printf("%s:Rebalanced: %+v \n", cos.groupId, ntf)
		}
	}()

	// consume messages, watch signals
	var successes int
Loop:
	for {
		select {
		case msg, ok := <-cos.consumer.Messages():
			if ok {
				once.Do(func() {
					StartTime = time.Now().Unix()
				})
				//fmt.Fprintf(os.Stdout, "%s:%s/%d/%d\t%s\t%s\n", cos.groupId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				cos.consumer.MarkOffset(msg, "")  // mark message as processed
				successes++
				Count ++
				//fmt.Fprintf(os.Stdout, "%s consume %d messages \n", cos.groupId, successes)
			}
		case <-signals:
			break Loop
		}
		if Count == 300000000 {
			EndTIme = time.Now().Unix()
		}
	}
	SpendTime := EndTIme -StartTime
	QPS := float64(Count)/float64(SpendTime)
	fmt.Println("三个Producer情况下总--QPS = ", QPS)
	fmt.Fprintf(os.Stdout, "%s consume %d messages \n", cos.groupId, successes)
}

// 支持brokers cluster的消费者
//func clusterConsumer(wg *sync.WaitGroup,brokers, topics []string, groupId string)  {
//	defer wg.Done()
//
//
//
//	defer consumer.Close()
//
//	// trap SIGINT to trigger a shutdown
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, os.Interrupt)
//
//	// consume errors
//	go func() {
//		for err := range consumer.Errors() {
//			log.Printf("%s:Error: %s\n", groupId, err.Error())
//		}
//	}()
//
//	// consume notifications
//	go func() {
//		for ntf := range consumer.Notifications() {
//			log.Printf("%s:Rebalanced: %+v \n", groupId, ntf)
//		}
//	}()
//
//	// consume messages, watch signals
//	var successes int
//Loop:
//	for {
//		select {
//		case msg, ok := <-consumer.Messages():
//			if ok {
//				fmt.Fprintf(os.Stdout, "%s:%s/%d/%d\t%s\t%s\n", groupId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
//				consumer.MarkOffset(msg, "")  // mark message as processed
//				successes++
//			}
//		case <-signals:
//			break Loop
//		}
//	}
//	fmt.Fprintf(os.Stdout, "%s consume %d messages \n", groupId, successes)
//}
