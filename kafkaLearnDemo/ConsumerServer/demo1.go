package ConsumerServer

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"strings"
	"time"
)

func main() {

	groupID := "group-1"
	topic := "test_topic"
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	c, err := cluster.NewConsumer([]string{"192.168.100.190:9092"}, groupID, strings.Split(topic, ","), config)
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()
	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				panic(err)
			case <-noti:
			}
		}
	}(c)
	fmt.Println("c.Messages()",c.Messages(),len(c.Messages()))
	for msg := range c.Messages() {
		fmt.Println("msg offset:", msg.Offset, " partition:", msg.Partition, " timestrap:", msg.Timestamp.Format("2006-Jan-02 15:04"), " value:", msg.Value, " key:", msg.Key)
		//fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}