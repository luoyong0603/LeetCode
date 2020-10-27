package main

import "fmt"

func main() {
	var nums = []int{0,0,1,1,1,2,2,3,3,4}
	i:= removeDuplicates(nums)
	//for _,v:=range i {
	//	fmt.Print(v)
	//}
	fmt.Print(i)

}

//双指针法

/*
数组完成排序后，我们可以放置两个指针 ii 和 jj，其中 ii 是慢指针，而 jj 是快指针。
只要 nums[i] = nums[j]nums[i]=nums[j]，我们就增加 jj 以跳过重复项。
当我们遇到 nums[j] \neq nums[i]nums[j]=nums[i] 时，跳过重复项的运行已经结束，
因此我们必须把它（nums[j]nums[j]）的值复制到 nums[i + 1]nums[i+1]。
然后递增 ii，接着我们将再次重复相同的过程，直到 jj 到达数组的末尾为止。
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	var i = 1
	for j:=1;j<len(nums);j++{
		if nums[j] !=nums[j-1]{
			nums[i] = nums[j]
			i++
		}
	}
	return i
}