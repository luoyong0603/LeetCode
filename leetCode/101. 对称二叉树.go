package main

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
*/

func isSymmetric(root *TreeNode) bool {
	return isBool(root, root)
}

func isBool(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.Val == r2.Val && isBool(r1.Left, r2.Right) && isBool(r1.Right, r2.Left)
}
