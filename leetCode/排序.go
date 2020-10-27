package main

import "fmt"

func main() {

	var l1 = []int{4, 2, 5, 3, 6, 9, 7, 1}
	var l2 = []int{4, 2, 5, 3, 6, 9, 7, 1}
	var l3 = []int{4, 2, 5, 3, 6, 9, 7, 1}
	var l4 = []int{4, 2, 5, 3, 6, 9, 7, 1}
	//排序默认：从小到大

	//选择排序
	fmt.Println(xuanze(l1))

	//冒泡排序
	fmt.Println(maopao(l2))

	//插入排序
	fmt.Println(charu(l3))

	//快速排序
	fmt.Println(kuaisu(l4))

}
/*
思路：
1、首先从原数组中取第一位作为基数，放首位
2、再从剩余未排序元素中继续寻找最小元素，然后放到已排序序列的后面
3、重复第二步、直到所有元素均排完序
*/
func xuanze(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		//用于标记j下标
		var min int = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				min = j
			}
		}
		//交换位置
		if min != i {
			nums[min],nums[i] = nums[i],nums[min]
		}

	}
	return nums
}
/*
思路：
1、从比较相邻两元素开始，往后遍历，有比自己更小的就交换位置，找最小的那个
2、重复1步骤，找第二小的值

*/
func maopao(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
/*
思路：
1、

*/
func charu(nums []int) []int {
	return nums
}


func kuaisu(nums []int) []int {
	return nums
}
