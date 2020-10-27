package main

/*
翻转一棵二叉树。

示例：

输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/

//递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//root.Left, root.Right = root.Right, root.Left
	//invertTree(root.Left)
	//invertTree(root.Right)

	//等价于下边这种写法
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

//前序遍历
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node:=&TreeNode{Val: root.Val}
	node.Left = invertTree1(root.Right)
	node.Right = invertTree1(root.Left)
	return node
}