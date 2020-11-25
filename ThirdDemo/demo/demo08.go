// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/10 11:57 上午
// @Update: xxx 2020/11/10 11:57 上午

package main

import (
	"log"
	"time"
)

func main() {
	a :=time.Now().UnixNano()/ 1e6
	b := time.Now().Unix()
	log.Println(a)
	log.Println(b*1000)
}
