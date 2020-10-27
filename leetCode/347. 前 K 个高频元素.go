package main

import (
	"fmt"
)

func main() {

	var l = []int{1, 1, 1, 2, 2, 3}
	var k = 2
	frequent := topKFrequent(l, k)
	fmt.Print(frequent)

}

func topKFrequent(nums []int, k int) []int {

	var mapV = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := mapV[nums[i]]
		if ok {
			mapV[nums[i]] ++
			continue
		} else {
			mapV[nums[i]] = 1
		}
	}
	var l []int
	for k1, _ := range mapV {
		//if len(mapV) <= k {
		//	l = append(l, k1)
		//	continue
		//}else{
		//
		//
		//
		//}
		l = append(l, k1)


		//if v >= k {
		//	l = append(l, k1)
		//}
	}
	//从小到大排序
	for i:=0;i<len(l);i++{
		for j:=1;i<len(l);j++{
			if mapV[i] >mapV[j]{
				mapV[i],mapV[j] = mapV[j],mapV[i]
			}
		}

	}
	if len(l) <= k{
		return l
	}
	//取k个数
	return l[len(l)-k-1:len(l)-1]
}
