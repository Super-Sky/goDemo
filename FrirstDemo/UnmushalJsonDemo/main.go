// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/14 4:14 下午
// @Update: xxx 2020/10/14 4:14 下午

package main

import (
	"encoding/json"
	"log"
)

type PosInfo struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

// 队列消息分组包头
type CommonGroupQueueMsg struct {
	GroupID  int32       `json:"group_id"`   // 分组ID，指定分组进行路由，第一优先级。
	UserID   int64       `json:"user_id"`    // 用户ID，用于路由时分发到对就的集群。 如果有则按用户ID 进行路由， 第二优先级
	SrcSvrID int32       `json:"src_svr_id"` // 发送此消息的，服务ID
	JsonData interface{} `json:"json_data"`  // 如果要传输的数据是 json 格式， 则将内容放置在 json_data 字段中。
	PbData   string      `json:"pb_data"`    // base64 string 如果要传输的数据是PB格式，则将PB序列化后，在进行base64，放置在 pb_data 字段中
}

// 自由出航战斗火苗打点
type FreeSailHotSpotNotify struct {
	AttackId       uint64  `json:"attack_id"`        // 攻击方
	DefendId       uint64  `json:"defend_id"`        // 被攻击方
	Count          int32   `json:"count"`            // 战斗次数
	MapId          uint64  `json:"map_id"`           // 地图 id
	AttackLeagueId uint64  `json:"attack_league_id"` // 攻击方联盟 id
	DefendLeagueId uint64  `json:"defend_league_id"` // 被攻击方联盟 id
	Pos            PosInfo `json:"pos"`              // 位置信息
	OpTime         int64   `json:"op_time"`          // 战斗时间
}

func main() {
	bingo := &FreeSailHotSpotNotify{
		AttackId:   9.00720000000096e+15,
		DefendId:    9.007200000000898e+15 ,
		Count: 1,
		MapId:19,
		AttackLeagueId: 6.882989145459786e+18,
		DefendLeagueId: 6.882989145459786e+18,
		Pos:PosInfo{
			X:5.546533e+06,
			Y:1.1771557e+07,
		},
		OpTime:1.602662336518e+12,
	}
	data, _ := json.Marshal(bingo)
	com := &CommonGroupQueueMsg {
		JsonData:data,
	}
	comdata, err := json.Marshal(com)
	if err != nil {
		panic(err)
	}
	var JsonData interface{}
	err = json.Unmarshal(comdata, &JsonData)
	if err != nil{
		log.Println("Unmarshal error", err)
	}
	mapData, ok := JsonData.(*FreeSailHotSpotNotify)
	if !ok {
		log.Println(ok)
		return
	}
	log.Println(mapData)
	//jsonStr, err := json.Marshal(mapData)
	//if err != nil {
	//	log.Println("Marshal json Err", err)
	//}
	//var msg *FreeSailHotSpotNotify
	//json.Unmarshal(jsonStr, &msg)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("msg",msg.AttackId)
}