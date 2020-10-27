package main

import (
	"fmt"
)

/*
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
注意:
假设字符串的长度不会超过 1010。
示例 1:
输入:
"abccccdd"
输出:
7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7.
*/

func main() {
	var str = "abbbccccc"
	i := longestPalindrome(str)
	fmt.Println(i)
	fmt.Print(len(str))
}

func longestPalindrome(s string) int {
	var m = make(map[string]int)
	for i := 0; i < len(s); i++ {
		//统计所有字母出现次数
		m[string(s[i])]++
		//_, ok := m[string(s[i])]
		//if ok {
		//	m[string(s[i])]++
		//} else {
		//	m[string(s[i])] = 1
		//}
	}
	//回文长度
	var l int = 0
	//标记是否存在奇数对
	var flag = false
	//最长组合，即为，1：存在奇数串，l=所有偶数对长度+（所有奇数串长度-奇数串个数+1）2、不存在奇数串：l=所有偶数对长度
	for _, v := range m {
		if len(m) == 1 {
			return v
		}
		if v%2 == 0 {
			l += v
		} else {
			flag = true
			//奇数串长度变偶数长度
			l += v - 1
		}
	}
	//只要存在奇数串，则最大奇数串可以放最中间，则不需要做减1 处理
	if flag {
		return l + 1
	}
	return l
}

