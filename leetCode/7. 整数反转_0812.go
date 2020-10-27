package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
*/

/*
	题解：转成字符串，再反转输出
	注意边界：假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。
	请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

func main() {
	var i = 1534236469
	i2 := reverse(i)
	fmt.Println(i2)
}
func reverse(x int) int {
	var b = true
	if x == 0 {
		return 0
	}
	if x < 0 {
		b = false
		x = -x
	}
	//s := fmt.Sprintf("%v", x)
	s := strconv.Itoa(x)
	str := returnStr(s)
	atoi, _ := strconv.Atoi(str)
	if atoi > math.MaxInt32 {
		return 0
	}
	if atoi < math.MinInt32 {
		return 0
	}
	if !b {
		atoi = -atoi
	}
	return atoi
}

//遍历字符反转
func returnStr(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		temp := r[i]
		r[i] = r[len(r)-i-1]
		r[len(r)-i-1] = temp
	}
	return string(r)
}
