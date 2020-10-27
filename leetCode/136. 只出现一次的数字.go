package main

import "fmt"

func main() {
	var l = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber1(l))
}

//map法
func singleNumber(nums []int) int {

	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			//删除重复元素
			//delete(m, nums[i])
		} else {
			//存map 并value记录数组中的下标
			m[nums[i]] = i
		}
	}
	for k := range m {
		return k
	}
	return -1
}

//异或运算  相同数字异或为0
func singleNumber1(nums []int) int {
	res := 0
	for _, i := range nums {
		res = res ^ i
	}
	return res
}

//
