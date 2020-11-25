// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/10 11:36 上午
// @Update: xxx 2020/11/10 11:36 上午

package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"pdDemo/pb"
)

func listValue()  {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc",
	}
	// 调用服务中listvalue方法
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}
}
// Address 连接地址
const Address string = ":8000"

var grpcClient pb.StreamServerClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	// 建立gRPC连接
	grpcClient = pb.NewStreamServerClient(conn)
	listValue()
}