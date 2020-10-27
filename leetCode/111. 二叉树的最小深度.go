package main

import "math"

type TreeNode1 struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minSum := math.MaxInt32
	if root.Left !=nil{
		minSum = min(minDepth(root.Left),minSum)
	}
	if root.Right !=nil{
		minSum = min(minDepth(root.Right),minSum)
	}
	return minSum+1
}


func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right != nil{
		ans = min(minDepth1(root.Left), minDepth1(root.Right)) + 1
	}else if root.Left != nil{
		ans = minDepth1(root.Left) + 1
	}else if root.Right != nil{
		ans = minDepth1(root.Right) + 1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
