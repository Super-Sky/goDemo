package main

import (
	"SecondDemo/echoDemo/gen-go/my/test/demo"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"net"
	"os"
	"time"
)

const (
	HOST = "127.0.0.1"
	PORT = "10086"
)

func main() {
	startTime := currentTimeMillis()
	ctx := context.Background()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, _ := transportFactory.GetTransport(transport)
	client := demo.NewClassMemberClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	var i int32
	for i = 0; i < 5; i++ {
		var s demo.Student
		s.Sid = i
		s.Sname = fmt.Sprintf("name_%d", i)
		s.Sage = int16(i + i)
		err := client.Add(ctx, &s)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
		fmt.Println("add", i, "student", s)
	}

	sList, err := client.List(ctx, currentTimeMillis())
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range sList {
		fmt.Println(s)
	}

	endTime := currentTimeMillis()
	fmt.Printf("calltime:%d-%d=%dms\n", endTime, startTime, (endTime - startTime))

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
