// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/26 10:16 上午
// @Update: xxx 2020/10/26 10:16 上午

package main

import (
	"log"
	"sync"
)

var gCityHouseInfo *CityHouseInfo


type CityHouseInfo1 struct {
	CityHouseBeingInfos sync.Map
	expiredTime        int64
}

func (c *CityHouseInfo)setCityHouseBeingInfos(userID uint64)  {
	c.CityHouseBeingInfos.Store(userID, struct{}{})
}

func (c *CityHouseInfo)getCityHouseBeingInfos(userID uint64) bool {
	if _, ok := c.CityHouseBeingInfos.Load(userID); ok {
		return true
	}
	return false
}

func newCityHouseInfo() *CityHouseInfo {
	return &CityHouseInfo{}
}

func GetCityHouseInfo() *CityHouseInfo {
	if gCityHouseInfo == nil {
		log.Println("11")
		gCityHouseInfo = newCityHouseInfo()
	}
	return gCityHouseInfo
}

func main() {
	var userID uint64 = 123456
	cityHouseInfo := GetCityHouseInfo()
	first := cityHouseInfo.getCityHouseBeingInfos(userID)
	if first {
		log.Println("first")
	}
	cityHouseInfo.setCityHouseBeingInfos(userID)
	second := cityHouseInfo.getCityHouseBeingInfos(userID)
	if second {
		log.Println("second")
	}
	cityHouseInfos := GetCityHouseInfo()
	var userID1 uint64 = 1234561
	var userID2 uint64 = 1234562
	var userID3 uint64 = 1234563
	cityHouseInfos.setCityHouseBeingInfos(userID1)
	cityHouseInfos.setCityHouseBeingInfos(userID2)
	cityHouseInfos.setCityHouseBeingInfos(userID3)
	cityHouseInfos.CityHouseBeingInfos.Range(func(k, v interface{}) bool {
		log.Println("iterate:", k, v)
		return true
	})
}

