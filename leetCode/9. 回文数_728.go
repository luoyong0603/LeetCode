package main

import (
	"fmt"
	"strconv"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func main() {

	a :=1233212

	fmt.Println(isPalindrome(a))

}
func isPalindrome(x int) bool {
	if x < 0{
		return false
	}
	y := strconv.Itoa(x)
	for i := 0; i < len(y)/2; i++ {
		fmt.Printf("str[%d]=%c\n", i, y[i])
		fmt.Printf("str[%d]=%c\n", len(y)-i-1, y[len(y)-i-1])
		if y[i] !=y[len(y)-i-1]{
			return false
		}

	}
	return true
}


func isReturn(x,y int)int {
	var v int
	var b,k int
	if x%v > v{
		b= v%10
		k++
		return isReturn(x,b)
	}




}