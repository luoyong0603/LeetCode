package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
*/
/*
示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
*/
func main() {
	b := isPalindrome2("A man, a plan, a canal: Panama")
	fmt.Print(b)
}

//解题思路：1、只留字母跟数字 放进一个新数组中，2、然后按普通回文串判断即可
func isPalindrome1(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if returnB(s[i]) {
			sgood += string(s[i])
		}
	}
	//全部转换成小写
	sgood = strings.ToLower(sgood)
	for j := 0; j < len(sgood)/2; j++ {
		if sgood[j] != sgood[len(sgood)-1-j] {
			return false
		}
	}
	return true
}

//判断只含字母，跟数字
func returnB(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

/*
解题思路二：在原字符串上直接判断。
*/
func isPalindrome2(s string) bool {
	//转换成小写
	s = strings.ToLower(s)
	var left, right = 0, len(s) - 1
	for left < right {
		//只验证字母和数字字符，其余的字符 空格等直接跳过
		for left < right && !returnB(s[left]) {
			left++
		}
		//只验证字母和数字字符，其余的跳过
		for left < right && !returnB(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}

	}
	return true
}
