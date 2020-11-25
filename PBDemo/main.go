// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/10 10:33 上午
// @Update: xxx 2020/11/10 10:33 上午

package main

import (
	proto "github.com/golang/protobuf/proto"
	"log"
	"pdDemo/pb"
)

func main() {
	s1 := &pb.Student{
		Name:    "李明",
		Age:     26,
		Address: "上",
		Cn:      pb.ClassName_class2,
	}
	ss := &pb.Students{
		School: "三小",
	}
	ss.Person = append(ss.Person, s1)
	s2 := &pb.Student{
		Name:    "张三",
		Age:     27,
		Address: "中",
		Cn:      pb.ClassName_class3,
	}
	ss.Person = append(ss.Person, s2)
	log.Println(ss)
	buffer, err :=proto.Marshal(ss)
	if err != nil {
		log.Println("marshal error ", err)
		return
	}
	log.Println("序列化后：", buffer)
	data := &pb.Students{}
	err = proto.Unmarshal(buffer, data)
	if err != nil {
		log.Println("unmarshal error", err)
	}
	log.Println("反序列化后：",data)
}