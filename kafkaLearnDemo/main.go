package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	//2020/09/19 15:57:18 DiffProducerCount 37500
	//2020/09/19 15:57:18 NewTime 1600502238
	//2020/09/19 15:57:18 OldTime 1600502238
	DiffProducerCount := 37500
	NewTime := 1600502238
	OldTime := 1600502238
	OriginalQPS := float64(DiffProducerCount)/float64(NewTime -OldTime)
	fmt.Println(float64(NewTime -OldTime))
	QPS := strconv.FormatFloat(OriginalQPS, 'f', 1, 64)
	log.Println("Consumer QPS = ", QPS)
}