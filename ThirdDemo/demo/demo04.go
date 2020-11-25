// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/30 10:47 上午
// @Update: xxx 2020/10/30 10:47 上午

package main

import (
	"log"
	"time"
)

func sumStr(a, b string)  {
	log.Println(a+b)
}

func main() {
	a := "1234"
	b := "6789"
	go func(c,d string) {
		sumStr(c,d)
	}(a,b)
	time.Sleep(2*time.Second)
}
