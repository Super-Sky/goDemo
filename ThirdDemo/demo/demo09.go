// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/11 10:46 下午
// @Update: xxx 2020/11/11 10:46 下午

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Set(firstSet map[string]struct{}, store map[string]struct{}) error   {
	s := make(map[string]map[string]struct{})
	for dmn, _ := range store {
		for cId, _ := range firstSet {
			if _, ok := s[dmn]; !ok {
				s[dmn] = make(map[string]struct{})
			}
			fmt.Println(cId)
		}
	}
	return nil
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func main() {
	for i :=0;i<100;i++{
		randNum := RandInt64(60,180)
		log.Println(randNum)
	}
}
