package main

import "fmt"



/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
func main() {

	var list = []int{0,1,1,2,2}
	i := removeElement1(list,1)
	fmt.Println()
	fmt.Println("i=:",i)
}


func removeElement(nums []int, val int) int {

	if len(nums) == 0{
		return 0
	}
	for i:=0;i<len(nums);i++{
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}


func removeElement1(nums []int, val int) int {

	var j =0
	for i:=0;i<len(nums);i++{
		if nums[i] != val{
			nums[j] = nums[i]
			j++
		}
	}
	for _,v:=range nums{
		fmt.Print(v)
	}
	return j
}
