package main

/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
*/

func preorderTraversal(root *TreeNode) []int {

	var nodeList []int
	if root == nil {
		return nil
	}
	nodeList = append(nodeList, root.Val)
	nodeList = append(nodeList, preorderTraversal(root.Left)...)
	nodeList = append(nodeList, preorderTraversal(root.Right)...)
	return nodeList
}
