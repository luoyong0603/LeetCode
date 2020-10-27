package main

import (
	"fmt"
)

func main() {
	var list = []int{-2, 1}
	i := maxSubArray(list)
	fmt.Printf("i=%v", i)
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var max = nums[0]
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		//首先判断 当前值是否比max大，避免有可能max为负，但是当前值也为负但是比max大的情况
		if sum > max {
			max = sum
		}
		for j := i + 1; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

//优化  只需遍历一次
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
