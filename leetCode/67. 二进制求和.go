package main

import (
	"fmt"
	"strconv"
)

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"

示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
func main() {
	fmt.Print(addBinary("100", "110010"))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		//始终让len(a)>len(b)
		a, b = b, a
	}
	var str = ""
	//需要进位 则变为 1
	var flag = 0
	var lenAB = len(a) - len(b)
	//补零操作
	if len(a) != len(b) {
		for i := 0; i < lenAB; i++ {
			b = "0" + b
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'

		sum := int(aNum+bNum) + flag
		if sum == 2 {
			str = "0" + str
			flag = 1
		} else if sum == 3 {
			str = "1" + str
			flag = 1
		} else {
			str = strconv.Itoa(sum) + str
			flag = 0
		}
	}
	if flag == 1 {
		str = "1" + str
	}
	return str
}
