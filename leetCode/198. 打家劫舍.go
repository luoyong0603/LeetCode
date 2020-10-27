package main

import "fmt"

func main() {
	var l = []int{2,1,1,2}
	fmt.Print(rob(l))
}

//方法一：动态规划
//dp[i]=max(dp[i−2]+nums[i],dp[i−1])

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp:=make([]int,len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])
	for i:=2;i<len(nums);i++{
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
// 滚动数组
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i:=2;i<len(nums);i++{
		first, second = second, max(first + nums[i], second)
	}
	return second
}




func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}