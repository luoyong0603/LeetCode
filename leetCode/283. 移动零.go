package main

import (
	"fmt"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/

func main() {
	var list = []int{0, 0,1,3,0,9}
	fmt.Print(moveZeroes(list))
}

func moveZeroes(nums []int) []int {
	var j int
	for i := 0; i <= len(nums); i++ {
		//注意临界，跳出循环
		if i == len(nums) {
			return nums
		}
		if nums[i] != 0 {
			var temp =nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			j++
		}
	}
	return nums
}
