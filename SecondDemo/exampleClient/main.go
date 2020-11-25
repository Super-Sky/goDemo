// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/19 11:22 上午
// @Update: xxx 2020/10/19 11:22 上午

package main

import (
	"SecondDemo/example"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
)

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := example.NewFormatDataClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST + ":" + PORT)
	}
	defer transport.Close()
	data := example.Data{Text:"hellO,world!"}
	ctx := context.Background()
	d, err := client.DoFormat(ctx , &data)
	fmt.Println(d)
}
