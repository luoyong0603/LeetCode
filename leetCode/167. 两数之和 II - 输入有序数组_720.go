package main

import (
	"fmt"
)

/*
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:
返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
*/

func main() {

	var list = []int{2, 3, 4}
	var target = 6
	output := SearchOther(list, target)
	for _, v := range output {
		fmt.Printf("%v,", v)
	}

}

//正常思路，循环找到每一个值
func TwoSum(numbers []int, target int) []int {
	var list []int
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return append(list, i+1, j+1)
			}
		}

	}
	return list
}

//寻找法
/*保存一个hsahmap,然后遍历数组的同时看看当前元素的补数(与target的差值)是否存在于hashmap中,
如果存在,就可以直接返回当前元素下标和hashmap中保存的下标,否则将当前元素的值和下标保存到hashmap中去
*/
//var list = []int{2, 7, 11, 15}
//	var target = 18
func SearchOther(numbers []int, target int) []int {
	m := map[int]int{} //map==> list.value:list.key
	for i, v := range numbers {
		find := target - v
		//根据value 判断是否存在对应的key
		if _, key := m[find]; key {
			return []int{m[find] + 1, i + 1}
		}
		m[v] = i
	}
	return []int{}
}
