package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{1, nil, nil}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}

	fmt.Print(getMinimumDifference(&root))
}

func getMinimumDifference(root *TreeNode) int {
	tmp, ans := -1, math.MaxInt64

	var dfs func(p *TreeNode)
	dfs = func(p *TreeNode) {
		if p.Left != nil {
			dfs(p.Left)
		}
		if tmp >= 0 && p.Val-tmp < ans {
			ans = p.Val - tmp
		}
		tmp = p.Val
		if p.Right != nil {
			dfs(p.Right)
		}
	}
	dfs(root)
	return ans
}
