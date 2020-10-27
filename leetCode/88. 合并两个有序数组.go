package main

import (
	"fmt"
	"sort"
)

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/

func main() {

	var nums1 = []int{6, 7, 8, 0, 0, 0}
	var m = 3
	var n = 3
	var nums2 = []int{2, 4, 6}
	//nums1 = append(nums1[:m], nums2...)
	//nums1 = append(nums1[:m], append(nums2[:n], nums1[m:]...)...)

	l := merges(nums1, m, nums2, n)
	l1 := merge4(nums1, m, nums2, n)
	fmt.Print(l)
	fmt.Print(l1)
}

func merges(nums1 []int, m int, nums2 []int, n int) []int {
	k := m + n
	//nums1从后往前放元素
	for i := k - 1; i >= 0; i-- {
		//n为零时说明nums2中所有的元素已经放到nums1里面了
		if n == 0 {
			break
		}
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
	return nums1
}

func merge4(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
	return nums1
}

func merge3(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m <= 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}


func merge5(nums1 []int, m int, nums2 []int, n int){
	k:=m+n
	for i:=k-1;i>0;i--{
		//n为零时说明nums2中所有的元素已经放到nums1里面了
		if n == 0 {
			break
		}
		if m>0 && nums1[m-1] >nums2[n-1]{
			nums1[i] = nums1[m-1]
			m--
		}else{
			nums1[i] = nums2[n-1]
			n--
		}
	}


}