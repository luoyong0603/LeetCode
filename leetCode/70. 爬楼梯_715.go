package main

import (
	"fmt"
)

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
 n=2  2
 n=3  3
 n=4  5

解题思路：发现规律 就是一个f(n) = f(n-2) + f(n-1)
*/
func main() {
	n := 2
	fmt.Println(sum(n))
}

func sum(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	x, y := 1, 2
	for i := 3; i < n; i++ {
		x, y = y, x+y
	}
	return x + y
}
