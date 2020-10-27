package main

import (
	"fmt"
	"strconv"
)

/*
给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

注意：整数序列中的每一项将表示为一个字符串。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
*/
func main() {

	str := countAndSay(4)
	fmt.Print(str)

}

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}
	if n == 2{
		return "11"
	}
	var str,s string
	var count = 1
	str = countAndSay(n - 1)
	if n > 2 {
		for i := 1; i <= len(str)-1; i++ {
			if str[i-1] != str[i] {
				s += strconv.Itoa(count) + string(str[i-1])
				count = 1
			}else{
				count++
			}
			//处理边界
			if i+1 == len(str){
				s += strconv.Itoa(count) +string(str[i])
				return s
			}
		}
	}
	return s
}
