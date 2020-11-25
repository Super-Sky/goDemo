//server.go

package main

import (
	"SecondDemo/echoDemo/gen-go/my/test/demo"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

const (
	NetworkAddr = "0.0.0.0:10086"
)

type ClassMemberImpl struct {
}

func (c *ClassMemberImpl) Add(ctx context.Context, s *demo.Student) (err error) {
	fmt.Println(s)
	students[s.Sid] = s
	return nil
}

func (c *ClassMemberImpl) List(ctx context.Context, callTime int64) (r []*demo.Student, err error) {
	for _, s := range students {
		r = append(r, s)
	}
	return r, nil
}

func (c *ClassMemberImpl) IsNameExist(ctx context.Context, callTime int64, name string) (r bool, err error) {
	for _, s := range students {
		if s.Sname == name {
			return true, nil
		}
	}
	return false, nil
}

var students map[int32]*demo.Student

func main() {

	students = make(map[int32]*demo.Student, 5)

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &ClassMemberImpl{}
	processor := demo.NewClassMemberProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}