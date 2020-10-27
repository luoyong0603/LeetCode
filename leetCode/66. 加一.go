package main

import (
	"fmt"
)

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

func main() {
	var li = []int{9, 9, 9, 9}
	one := plusOne1(li)
	for _, v := range one {
		fmt.Print(v)
	}

}

//递归
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}
	var l = []int{1}
	var flag = 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] ++
			return digits
		}
		b := isAdd(digits[i] + flag)
		//说明需要进位
		if b {
			flag = 1
			digits[i] = 0
			if i-1 < 0 {
				digits = append(digits[:0], append(l, digits[0:]...)...)
				break
			}
		} else {
			digits[i] = digits[i] + flag
			flag = 0
		}
	}
	return digits
}

//校验是否需要+1
func isAdd(num int) bool {
	if num == 10 {
		return true
	}
	return false
}


//优化
func plusOne1(digits []int) []int {
	l := len(digits)
	var one = []int{1}
	if l == 0 {
		return []int{1}
	}
	for i := l-1; i>= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append(one, digits...)
}