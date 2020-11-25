// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/26 11:12 上午
// @Update: xxx 2020/10/26 11:12 上午

package main

import "log"

type TaxInfo1 struct {
	LastTime   int32 `json:"last_time"`
	Remain     int64 `json:"remain"`
	SettleTime int32 `json:"settle_time"`
	PeerTax    int64 `json:"peer_tax"`
}

func setTaxInfo(taxInfo *TaxInfo) {
	taxInfo.LastTime = 123456
	taxInfo.Remain = 123456
	taxInfo.SettleTime = 123456
	taxInfo.PeerTax = 123456
}

func main() {
	taxInfo := &TaxInfo{}
	setTaxInfo(taxInfo)
	log.Println(taxInfo)
}
