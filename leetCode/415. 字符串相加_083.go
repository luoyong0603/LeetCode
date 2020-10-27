package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/
/*

思路: 对应取出每一位数进行相加，最后返回拼接好的数值即可
	1、对齐两字符串，短的一方前面补0
	2、任意取一个字符串进行遍历，对应每一位数组下标相加，逢十进一，
	3、定义一个变量flag记录两数大于10，
	4、在判断是否相加大于10之前，每次都加一次flag，保证逢十进一
	5、处理边界问题，最后一次遍历可能直接跳出，有可能会相加大于10，所以，只要flag不等于0.即返回，flag+returnStr


*/


func main() {

	var str = "1"
	var str1 = "9"
	s := addStrings(str, str1)
	fmt.Println(s)
}

func addStrings(num1 string, num2 string) string {
	var len1 = len(num1)
	var len2 = len(num2)
	var length = len2
	var returnStr = ""

	//先补零
	if len1 > len2 {
		length = len1
		num2 = strings.Repeat("0", len1-len2) + num2
	} else {
		num1 = strings.Repeat("0", len2-len1) + num1
	}
	//用于记录进位，逢10进1
	flag := 0
	//从个位开始相加
	for i := length - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		y := int(num2[i] - '0')
		sum := x + y + flag
		if sum >= 10 {
			flag = 1
			sum = sum % 10
		} else {
			flag = 0
		}
		returnStr = strconv.Itoa(sum) + returnStr
	}
	//处理最后两位数相加大于10 ，直接跳出循环，则应把flog 加上
	if flag != 0 {
		returnStr = strconv.Itoa(flag) + returnStr
	}

	return returnStr
}
