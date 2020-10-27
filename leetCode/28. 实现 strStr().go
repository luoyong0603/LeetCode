package main

import (
	"fmt"
	"strings"
)

/*
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2
*/
func main() {
	var str1 = "hello"
	fmt.Print(str1[2:4])

	var str2 = "ll"
	index1 := strStr(str1, str2)
	index2 := strStr1(str1,str2)
	fmt.Print("index=",index1)
	fmt.Print("index=",index2)

}

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return 0
	}

	if len(needle) == 0 {
		return 0
	}
	var length = len(needle)
	for j := 0; j < len(haystack); j++ {

		if j+length <=len(haystack) && haystack[j:j+length] == needle {
			return j
		}
	}
	return -1
}

//无赖版
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack,needle)
}