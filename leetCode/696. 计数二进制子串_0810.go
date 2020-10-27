package main

import "fmt"

/*
给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。
示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
请注意，一些重复出现的子串要计算它们出现的次数。
另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
*/

//题解：

func main() {
	var str = "00000"
	i := countBinarySubstrings(str)
	fmt.Println("i=", i)
}

func countBinarySubstrings(s string) int {
	var list []int
	for i := 0; i < len(s); {
		count := 0
		v := s[i]
		for i < len(s) && s[i] == v {
			count++
			i++
		}
		list = append(list, count)
	}
	var count int
	for j := 0; j < len(list); j++ {
		if j+1 == len(list) {
			break
		}
		v := returnMinV(list[j], list[j+1])
		count += v
	}
	return count
}

func returnMinV(x, y int) int {
	if x > y {
		return y
	}
	return x

}
