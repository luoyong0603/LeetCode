package main

import (
	"fmt"
	"sort"
)

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func main() {
	var i = []int{8,8,7,7,7}

	fmt.Println(majorityElement(i))
	fmt.Print(majorityElement1(i))

}

func majorityElement(nums []int) int {
	//map[num:count]
	mapV := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//先判断是否存在这个key(num)
		_, ok := mapV[nums[i]]
		//存在即value+1
		if ok {
			mapV[nums[i]] += 1
			continue
		}
		//否则新增一个
		mapV[nums[i]] = 1
	}
	var max = 0
	var returnV int
	for k, v := range mapV {
		//根据value 找出最大的key
		if v > max {
			max = v
			returnV = k
		}
	}
	return returnV
}


//下标为 n/2的元素（下标从 0 开始）一定是众数
func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}