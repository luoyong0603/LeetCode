package main



/*
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:      1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true

//思路：
深度优先搜索,如果两个二叉树都为空，则两个二叉树相同。
如果两个二叉树中有且只有一个为空，则两个二叉树一定不相同。
如果两个二叉树都不为空，那么首先判断它们的根节点的值是否相同，
若不相同则两个二叉树一定不同，若相同，再分别判断两个二叉树的左子树是否相同以及右子树是否相同。
这是一个递归的过程，因此可以使用深度优先搜索，递归地判断两个二叉树是否相同。

示例 1:
输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]
输出: true

示例 2:
输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]
输出: false
*/


func main()  {


}
type TreeNodes struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil{
		return true
	}
	if p == nil || q == nil{
		return false
	}
	if p.Val != q.Val{
		return false
	}
	return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}