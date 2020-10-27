package main

import (
	"fmt"
	"strings"
)

/*
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
示例:
输入: "Hello World"
输出: 5
*/
func main() {

	//var str = "a "
	//fmt.Println(lengthOfLastWord(str))
	//fmt.Println(lengthOfLastWord1(str))


	str := "空格, 有空格, haha haha!"
	strList := strings.Fields(str)
	fmt.Printf("%q\n", strList) //["Hello," "世界!" "Hello!"]
}

func lengthOfLastWord(s string) int {
	list := strings.Fields(s)
	if len(list) == 0 {
		return 0
	}
	s2 := list[len(list)-1]
	return len(s2)
}

func lengthOfLastWord1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var l = len(s)
	var flog = 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if flog > 0 {
				break
			}
		} else {
			flog ++
		}
	}
	return flog
}
