package main

/*
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [3,2,1]
*/

/*
  递归
*/
var list []int

func postorderTraversal(root *TreeNode) []int {
	defer func() {
		list = []int{}
	}()
	returnLists(root)
	return list
}

func returnLists(node *TreeNode) {
	if node == nil {
		return
	}
	returnLists(node.Left)
	returnLists(node.Right)
	list = append(list, node.Val)
}

/*
  递归  闭包方法
*/

func postorderTraversal1(root *TreeNode) []int {
	var list []int
	if root == nil {
		return list
	}

	var post func(node *TreeNode)
	post = func(node *TreeNode) {
		if node.Left != nil {
			post(node.Left)
		}
		if node.Right != nil {
			post(node.Right)
		}
		list = append(list, node.Val) //左右根
	}
	post(root)
	return list
}
