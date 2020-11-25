// @Desc: \\todo
// @Author: MaXiaoTian 2020/11/7 11:28 上午
// @Update: xxx 2020/11/7 11:28 上午

package main

import (
	"encoding/json"
	"log"
)

type TaxInfo struct {
	LastTime   int32 `json:"last_time"`
	Remain     int64 `json:"remain"`
	SettleTime int32 `json:"settle_time"` //结算时间 税率改变就结算
	PeerTax    int64 `json:"peer_tax"`    //单位税率
}

type TaxLogInfo struct {
	LastTime   int32 `json:"last_time"`
	Remain     int64 `json:"remain"`
	SettleTime int32 `json:"settle_time"` //结算时间 税率改变就结算
}


func main() {
	info := &TaxInfo{
		LastTime: 1,
		Remain: 1,
		SettleTime: 1,
		PeerTax: 1,
	}
	tempData,_ := json.Marshal(info)
	info1 := &TaxInfo{}
	info2 := &TaxInfo{}
	_ = json.Unmarshal(tempData, &info1)
	_ = json.Unmarshal(tempData, info2)
	if *info1 == *info2 {
		log.Println("11111111111")
	}
}