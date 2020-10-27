package main

import "fmt"

func main() {
	var target int = 9
	fmt.Print(findContinuousSequence(target))
}

func findContinuousSequence(target int) [][]int {
	var i = 1  // 滑动窗口的左边界
	var j = 1  // 滑动窗口的右边界
	var sum = 0 // 滑动窗口中数字的和

	var l [][]int
	for i <= (target+1)/2 {
		// 如果总和小于目标值
		if sum < target {
			// 右边界向右移动
			sum += j
			j++
		} else {
			// 左边界向右移动
			sum = sum - i
			i++
		}
		// 记录结果
		if sum == target {
			var tmp []int
			for k := i; k < j; k++ {
				tmp = append(tmp, k)
			}
			l = append(l, tmp)
			// 左边界向右移动
			sum = sum - i
			i++
		}
	}
	return l
}
