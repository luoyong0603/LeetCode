package main

import (
	"fmt"
)

/*
例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

	解题思路：
	消层法：复制一个倒数一层数组dp[len-1]，然后取倒数第二层的每一个数都与倒数第一层的左右相邻节点求和，
	并用最小和替换掉数位置dp[len-1].直到第一层dp[0]，就是最小路径和
*/
func main() {
	//var sum int
	//给定三角形
	var isArr = [4][4]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	dp := isArr[len(isArr)-1]
	for i := len(isArr) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = returnMin(dp[j], dp[j+1]) + isArr[i][j]
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	fmt.Printf("自顶向下的最小路径和为：%v\n", dp[0])
}

func returnMin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
