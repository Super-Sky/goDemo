// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/10 11:36 上午
// @Update: xxx 2020/11/10 11:36 上午

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	GolangLabSvr "pdDemo/pb"
	"strconv"
	"time"
)

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

type StreamService struct {}

func (s *StreamService) ListValue(req *GolangLabSvr.TestRQ, srv GolangLabSvr.StreamServer_ListValueServer) error {
	for n :=0;n<50000;n++{
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err !=nil {
			return err
		}
		time.Sleep(20*time.Millisecond)
	}
	return nil
}

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil{
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})
	err = grpcServer.Serve(listener)
	if err != nil{
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}