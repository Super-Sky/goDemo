// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/19 11:40 上午
// @Update: xxx 2020/11/19 11:40 上午

package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(map[string]int)
	go func() {//开一个协程写map
		for j := 0; j < 1000000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() {    //开一个协程读map
		for j := 0; j < 1000000; j++ {
			fmt.Println(c[fmt.Sprintf("%d",j)])
		}
	}()

	time.Sleep(time.Second*20)

}



