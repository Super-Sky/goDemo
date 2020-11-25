// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/2 4:40 下午
// @Update: xxx 2020/11/2 4:40 下午

package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)



var (
	CityHouseDuration int64 = 180
	GCityHouseInfo *CityHouseInfo
)

type CityHouseInfo struct {
	CityHouseBeingInfo sync.Map
}
func init()  {
	GCityHouseInfo = &CityHouseInfo{}
	nowTime := time.Now().Unix()
	// 缓存城主府信息
	expireTime := nowTime + CityHouseDuration
	for i:=0;i<100;i++ {
		GCityHouseInfo.CityHouseBeingInfo.Store(uint64(i), expireTime)
	}
}

func main() {
	ticker := time.NewTicker(1*time.Second)
	go func() {
		for {
			select {
			case now:= <-ticker.C:
				nowTime := now.Unix()
				log.Println(nowTime)
				GCityHouseInfo.CityHouseBeingInfo.Range(func(k, v interface{}) bool {
					log.Println(k)
					userId,_ := k.(uint64)
					fmt.Println("del key"/* 语句中的注释 */,userId)
					GCityHouseInfo.CityHouseBeingInfo.Delete(uint64(userId))
					return true
				})
			}
		}
	}()
	time.Sleep(1000*time.Second)
	//fun2()
}
