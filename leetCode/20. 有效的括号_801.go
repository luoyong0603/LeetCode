package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/


/*
解题思路
注意题目中有三种括号组合，先设置一个对应关系 Map
左括号类型入栈，右括号类型匹配栈顶元素，如果不是对应关系，则肯定 false，匹配上之后移除栈顶元素
最后看栈是否为空，空则是有效括号，非空则是无效的
同时注意题目中空字符串是返回 true
*/

func main() {
	var str = "{[(()])}"
	fmt.Println(isValid(str))

	//fmt.Print(str[:2])
}

func isValid(s string) bool {

	var length = len(s)
	if s == "" {
		return true
	}
	if length%2 == 1 {
		return false
	}

	var strck []byte
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			strck = append(strck, s[i])
		} else {
			if len(strck) == 0 {
				return false
			}
			if m[s[i]] != strck[len(strck)-1] {
				return false
			}
			strck = strck[:len(strck)-1] //匹配成功，即去掉strck最后一个
		}
	}
	return len(strck) == 0
}
