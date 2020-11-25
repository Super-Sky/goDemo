package main

import (
	"fmt"
	"kafkaDemo/ConsumerServer"
	"kafkaDemo/Process/ConsumerProcess"
	"kafkaDemo/Process/ProducerProcess"
	"runtime"
	"sync"
	"time"
)

var OldProducerCount int32 = 0
var OldConsumerCount int32 = 0

func countTimer()  {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				DiffProducerCount := ProducerProcess.Count - OldProducerCount
				DiffConsumerCount := ConsumerServer.Count - OldConsumerCount
				fmt.Println("DiffProducerCount",DiffProducerCount)
				fmt.Println("DiffConsumerCount",DiffConsumerCount)
				runnerGo := runtime.NumGoroutine()
				fmt.Println("runnerGo",runnerGo)
				OldProducerCount = ProducerProcess.Count
				OldConsumerCount = ConsumerServer.Count
			}
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	//countTimer()
	wg.Add(2)
	go func() {
		defer wg.Done()
		ProducerProcess.StartProducer()
	}()
	go func() {
		defer wg.Done()
		ConsumerProcess.StartConsumer()
	}()
	wg.Wait()
}
