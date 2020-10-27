package main

import (
	"fmt"
)

/*
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:
输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true

*/

/*
	解题思路:  类似插入，不能改变s1 s2 原字符串的先后顺序
		1、
		2、
		3、
*/

func main() {
	var s1 = "aabcc"
	var s2 = "dbbca"
	var s3 = "aadbbacbcc"

	bool := isInterleave(s1, s2, s3)
	fmt.Println(bool)
}

func isInterleave(s1 string, s2 string, s3 string) bool {

	return dfs(s1, s2, s3, 0, 0, 0)
}

func dfs(s1, s2, s3 string, i, j, k int) bool {
	if k == len(s3) && i == len(s1) && j == len(s2) {
		return true
	}
	if k >= len(s3) {
		return false
	}
	// bound check
	if i < len(s1) {
		if s1[i] == s3[k] {

			if dfs(s1, s2, s3, i+1, j, k+1) {
				return true
			}
		}
	}
	if j < len(s2) {
		if s2[j] == s3[k] {
			if dfs(s1, s2, s3, i, j+1, k+1) {
				return true
			}
		}
	}
	return false
	//类似插入，不能改变s1 s2 原字符串的顺序

}
