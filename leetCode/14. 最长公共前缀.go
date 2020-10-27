package main

import (
	"fmt"
)

func main() {
	var str = []string{
		"foower", "flow", "fooght",
	}
	prefix := longestCommonPrefix(str)
	fmt.Print(prefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//初始化
	var str = strs[0]
	for i := 1; i < len(strs); i++ {
		//递归调用
		str = returnCommStr(str, strs[i])
	}
	return str
}

//找前缀公共元素
func returnCommStr(str1, str2 string) string {
	var str string
	//临时标记 目的是找出字符串较短的哪一个，防止遍历数组下标越界
	var flag = str1
	if len(str1) > len(str2) {
		flag = str2
	}
	for i := 0; i < len(flag); i++ {
		//前缀不相等，直接返回空
		if str1[0] != str2[0] {
			return ""
		}
		if str1[i] == str2[i] {
			str = str + (string(str1[i]))
		}
	}
	return str
}
