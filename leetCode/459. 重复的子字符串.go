package main

import "fmt"

/*
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。
给定的字符串只含有小写英文字母，并且长度不超过10000。
示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
*/

func main() {
	var str = "abcabc"
	fmt.Print(repeatedSubstringPattern(str))
}

func repeatedSubstringPattern(s string) bool {
	//子串
	var str1 string
	//子串长度
	var strLen int
	//全串长度
	var l = len(s)
	for i := 1; i <= l; i++ {
		str1 = s[0:i]
		strLen = len(str1)
		//若子串跟全串一样，直接返回false
		if len(str1) == len(s) {
			return false
		}
		//根据子串去匹配全串是否全部由子串组成
		for j := 0; j <= l-strLen; j+=strLen{
			//按子串str1长度从全串中取得str2
			str2 := s[j : j+strLen]
			//若str2不匹配，直接跳出匹配循环
			if str1 != str2 {
				break
			}
			//表示子串全部匹配成功，则一定会执行返回true
			if j == l-strLen{
				return true
			}
		}
	}
	return false
}
