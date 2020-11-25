package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)


var (
	newHotSpotsInfo = &HotSpotsInfo{}
	oldHotSpotsInfo = &HotSpotsInfo{Spots: make(map[string]*Spot)}
	ch chan string
)

type HotSpotsInfo struct {
	Spots map[string]*Spot // posKey: *Spot
	sync.RWMutex
}

type Spot struct {
	LeagueIds map[uint64]struct{}
}

func init()  {
	newHotSpotsInfo = &HotSpotsInfo{
		Spots: make(map[string]*Spot),
	}
}

func (r *HotSpotsInfo) getSpotStr() (ret string) {
	r.RLock()
	defer r.RUnlock()
	for k, v := range r.Spots {
		ret += fmt.Sprintf("%s:%v ", k, *v)
	}
	//r.RUnlock()
	return
}

func readMap1()  {
	for {
		newHotSpotsInfo.RLock()
		//log.Println("newHotSpotsInfo.Spots",newHotSpotsInfo.Spots)
		for newK,newV :=range newHotSpotsInfo.Spots {
			//log.Println("k: ", k,"v: ",v)
			//log.Println("newK",newK,"newV",newV)
			oldHotSpotsInfo.Spots[newK] =newV
		}
		newHotSpotsInfo.RUnlock()
		readMap2()
	}

}

func readMap2()  {
	for {
		newHotSpotsInfo.getSpotStr()
		oldHotSpotsInfo.getSpotStr()
		log.Println("newHotSpotsInfo", newHotSpotsInfo.getSpotStr(), "oldHotSpotsInfo",oldHotSpotsInfo.getSpotStr())
	}

}

func writeMap1()  {
	for i:=0;i<1000000;i++{
		//log.Println("WriteMap1,",i)
		spot := &Spot{
			LeagueIds: make(map[uint64]struct{}),
		}
		spot.LeagueIds[uint64(i)] = struct{}{}
		newHotSpotsInfo.Lock()
		str := "WriteMap1" + strconv.FormatInt(int64(i),10)
		newHotSpotsInfo.Spots[str] = spot
		newHotSpotsInfo.Unlock()

	}
}

func WriteMap2()  {
	for i:=0;i<1000000;i++{
		//log.Println("WriteMap2,",i)
		spot := &Spot{
			LeagueIds: make(map[uint64]struct{}),
		}
		spot.LeagueIds[uint64(i)] = struct{}{}
		newHotSpotsInfo.Lock()
		str := "WriteMap2" + strconv.FormatInt(int64(i),10)
		newHotSpotsInfo.Spots[str] = spot
		newHotSpotsInfo.Unlock()
	}
}

func main() {
	ch = make(chan string)
	go readMap1()
	go readMap1()
	go readMap1()
	for i:=0;i<10000;i++{
		go writeMap1()
		go WriteMap2()
	}
	time.Sleep(time.Second*6000)
}

