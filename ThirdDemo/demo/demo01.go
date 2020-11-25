// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/24 5:31 下午
// @Update: xxx 2020/10/24 5:31 下午

package main---

import (
	"log"
)

var CityHouseBeingInfo = make(map[uint64]struct{})

func setCityHouseBeingInfo(userID uint64)  {
	CityHouseBeingInfo[userID] = struct{}{}
}

func getCityHouseBeingInfo(userID uint64) bool {
	if _, ok := CityHouseBeingInfo[userID]; ok {
		return true
	}
	return false
}

func main() {
	var userID uint64 = 123456
	first := getCityHouseBeingInfo(userID)
	if first {
		log.Println("first")
	}
	setCityHouseBeingInfo(userID)
	second := getCityHouseBeingInfo(userID)
	if second {
		log.Println("second")
	}
	var userID1 uint64 = 1234561
	var userID2 uint64 = 1234562
	var userID3 uint64 = 1234563
	setCityHouseBeingInfo(userID1)
	setCityHouseBeingInfo(userID2)
	setCityHouseBeingInfo(userID3)
	log.Println("CityHouseBeingInfo", CityHouseBeingInfo)
}
