// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/14 11:00 上午
// @Update: xxx 2020/10/14 11:00 上午

package main

import (
	pb "SecondDemo/addressbook"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	// 自定义AddressBook内容
	book := &pb.AddressBook{
		People: []*pb.Person {
			&pb.Person{
				Id: 1,
				Name: "zyq",
				Email: "77@qq.com",
				Phones: []*pb.Person_PhoneNumber{
					&pb.Person_PhoneNumber {
						Number: "11111",
						Type: pb.Person_MOBILE,
					},
					&pb.Person_PhoneNumber {
						Number: "22222",
						Type: pb.Person_HOME,
					},
				},
			},
		},
	}
	fmt.Println("Type: pb.Person_MOBILE",pb.Person_MOBILE)
	fmt.Println("book : ",book)
	fname := "address.dat"
	//将book进行序列化
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Marshal Err",err)
	}
	log.Println("Marshal out put ", out)
	//将序列化后的数据写入文件
	if err := ioutil.WriteFile(fname, out, 0644); err !=nil {
		log.Fatalln("out put write file Err", err)
	}
	//读取二进制数据
	in , err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("read file Err", err)
	}
	//定义一个空结构体
	UnmarshalBook := &pb.AddressBook{}
	//将文件从二进制进行反序列化
	if err := proto.Unmarshal(in, UnmarshalBook); err != nil {
		log.Fatalln("Unmarshal file data Err", err)
	}
	log.Println("Unmarshal file data: ")
	log.Println(UnmarshalBook)

}
