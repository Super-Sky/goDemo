// @Desc: \\todo
// @Author: MaXiaoTian 2020/10/20 3:56 下午
// @Update: xxx 2020/10/20 3:56 下午
package main

import (
	"log"
)

func GetGCD(x, y int64) int64 { // 获取两数最大公约数
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func GetGCD4(x, y int64) int64 { // 获取两数最大公约数
	x, y = y, x%y
	if y == 0 {
		return x
	}
	return GetGCD4(x , y)
}

func GetGCD2(x, y int64) int64 {

	if x == 0 {
		return y
	}else if y == 0 {
		return x
	}
	if x&1 !=1 && y&1 !=1 { //都是偶数
		return GetGCD2(x >> 1, y >> 1) << 1
	}else if x&1 !=1 && y&1 ==1 { // x 偶数 y奇数
		return GetGCD2( x>> 1, y)
	}else if x&1 ==1 && y&1 !=1 { // x 奇数 y 奇数
		return GetGCD2(x, y >> 1)
	}
	if x >y {
		return GetGCD2((x - y) >> 1, y)
	}
	return GetGCD2((y - x) >> 1, x)

}

func GetGCD3(num1, num2 int64) int64 {
	var factor int64= 1
	if num1 < num2 {
		return GetGCD3(num2, num1)
	}
	for num1 != num2{
		if num1&1 !=1 && num2&1 !=1 { // 均为偶数
			num1 = num1 >> 1
			num2 = num2 >> 2
			factor *= 2
		}else if num1&1 !=1 && num2&1 ==1{
			num1 = num1 >> 1
		}else if num1&1 ==1 && num2&1 !=1 {
			num2 = num2 >> 1
		}else {
			if num1 > num2 {
				num1 = num1 - num2
			} else {
				num2 = num2 - num1
			}

		}

	}
	return factor*num1
}

func main(){
	//var b int64= 999999342353200
	//var a int64= 777774234
	//startTime1 := time.Now().UnixNano()
	//for i := 0;i<1000;i++{
	//	_ = GetGCD4(a, b)
	//}
	//c := GetGCD4(a, b)
	//startTime2 := time.Now().UnixNano()
	//for i := 0;i<1000;i++{
	//	_ = GetGCD(a, b)
	//}
	//d := GetGCD(a, b)
	//endTime3 := time.Now().UnixNano()
	//log.Println("c",c, "Time", startTime2-startTime1)
	//log.Println("d",d, "Time", endTime3 - startTime2)
	subGemItemLists := make([]string, 4)
	//subGemItemList := []string{}
	//var subGemItemList []string
	//subGemItemList = append(subGemItemList, "1234")
	subGemItemLists[0] = "1"
	subGemItemLists[1] = "2"
	subGemItemLists[2] = "3"
	subGemItemLists[3] = "4"
	subGemItemList := make([]string, len(subGemItemLists))
	for i, v := range subGemItemLists {
		key := v+v
		subGemItemList[i] = key
	}
	log.Println(subGemItemList)
}


//func main() {
//	a := errors.New("添加物品达到单笔上限")
//	log.Println(a)
//	b := errors.New("添加物品达到单笔上限")
//	log.Println(b)
//	if a.Error() == b.Error() {
//		log.Println("111")
//	}
//}
