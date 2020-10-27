package main

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
输入: [1,3,5,6], 5
输出: 2
*/
func main() {

	var list = []int{1, 3, 5, 6}
	var target = 5
	index := guanfan1(list, target)
	fmt.Println(index)
}

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

//借助游标法
func guanfan(nums []int, target int) int {
	var left = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			left++
		} else {
			return left
		}
	}
	return left + 1
}

//官方  二分法 寻找
func guanfan1(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
