package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."
*/
func main() {
	var str = "We are happy."
	fmt.Println(replaceSpace("We are happy."))
	fmt.Println(replaceSpace1("We are happy."))
	fmt.Println(replaceSpace2("We are happy."))
	fmt.Println(replaceSpace3(str))
}

//遍历匹配
func replaceSpace(s string) string {
	var str string
	for _, v := range s {
		if v == ' ' {  //或者使用这种方式匹配空格  unicode.IsSpace(v)
			str += "%20"
			continue
		}
		str += string(v)

	}
	return str
}

//使用Unicode 匹配空格
func replaceSpace1(s string) string {
	var str string
	for _, v := range s {
		if unicode.IsSpace(v) {
			str += "%20"
			continue
		}
		str += string(v)
	}
	return str
}

//
func replaceSpace2(s string) string {
	var str strings.Builder
	for _, v := range s {
		if unicode.IsSpace(v) {
			str.WriteString("%20")
			continue
		}
		str.WriteByte(byte(v))
	}
	return str.String()
}

//直接使用ReplaceAll 快速替换替换
func replaceSpace3(s string) string {
	return strings.ReplaceAll(s," ","%20")
}
