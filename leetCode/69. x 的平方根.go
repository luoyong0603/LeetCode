package main

import "fmt"

func main() {
	fmt.Print(mySqrt1(102221212))
}

//暴力破解
func mySqrt(x int) int {
	var num int = 0
	for num*num <= x {
		if (num+1)*(num+1) > x {
			return num
		}
		num++
	}
	return num
}

//二分法
/*
由于 x平方根的整数部分是满足 k^2≤x 的最大 k值 因此我们可以对 k 进行二分查找，从而得到答案
*/

func mySqrt1(x int) int {
	left, right := 0, x
	ans := -1
	for left < right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
