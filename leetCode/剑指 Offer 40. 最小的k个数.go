package main

import (
	"fmt"
	"sort"
)

/*
输入整数数组 arr ，找出其中最小的 k 个数。
例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]

*/

func main() {
	var l = []int{3, 2, 1}
	var k = 2
	fmt.Print(getLeastNumbers(l, k))
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[0:k]
}
