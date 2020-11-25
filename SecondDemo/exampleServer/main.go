// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/19 11:04 上午
// @Update: xxx 2020/10/19 11:04 上午

package main

import (
	"SecondDemo/example"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"strings"
)

type FormatDataImpl struct {}

func (fdi *FormatDataImpl) DoFormat(ctx context.Context, data *example.Data) (r *example.Data, err error) {
	var rData example.Data
	rData.Text = strings.ToUpper(data.Text)
	return &rData, nil
}

const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {
	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	serverTranSport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Println("Err: ", err)
	}
	//                        NewTBufferedTransportFactory
	tansportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//tansportFactory := thrift.NewTBufferedTransportFactory(256 * 1024)
	protocolFactoy := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTranSport, tansportFactory, protocolFactoy)
	fmt.Println("Running at:", HOST + ":" + PORT)
	server.Serve()

}