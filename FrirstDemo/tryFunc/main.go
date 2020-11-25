// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/14 10:50 上午
// @Update: xxx 2020/10/14 10:50 上午

package main

import (
	"FrirstDemo/"
)

func main() {
	// 自定义AddressBook内容
	book := &FrirstDemo.AddressBook{
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
}