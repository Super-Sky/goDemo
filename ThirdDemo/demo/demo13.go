// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/23 11:05 上午
// @Update: xxx 2020/11/23 11:05 上午

package main

import "log"

func main() {
	tempChan := make(chan int, 2)
	close(tempChan)
	for {
		select {
		case _, ok :=<-tempChan:
			if ok {
				log.Println("is ok")
			}else {
				log.Println("is ! ok",len(tempChan))
			}
		default:
			log.Println("default----------")
			return
		}
	}
}
