package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {

	var list = []int{2, 7, 11, 15}
	var target = 17
	ints := searchOther(list, target)
	fmt.Println(ints)
}

func twoSum(nums []int, target int) []int {
	var list []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return append(list, i, j)
			}
		}

	}
	return nil
}

//寻找法
func searchOther(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		find := target - v
		if _, e := m[find]; e {
			return []int{m[find], i}
		}
		m[v] = i
	}
	return []int{}

}
