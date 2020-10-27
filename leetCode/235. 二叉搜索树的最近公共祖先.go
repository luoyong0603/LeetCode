package main

import "fmt"

type TreeNode3 struct {
	Val         int
	Left, Right *TreeNode3
}

//深度尽可能大
func main() {
	root := &TreeNode3{6, nil, nil}
	root.Right = &TreeNode3{Val: 8}
	root.Right.Left = &TreeNode3{Val: 7}
	root.Right.Right = &TreeNode3{Val: 9}
	root.Left = &TreeNode3{Val: 2}
	root.Left.Left = &TreeNode3{Val: 0}
	root.Left.Right = &TreeNode3{Val: 4}
	root.Left.Right.Left = &TreeNode3{Val: 3}
	root.Left.Right.Right = &TreeNode3{Val: 5}
	p := &TreeNode3{2, nil, nil}
	q := &TreeNode3{4, nil, nil}
	ancestor := lowestCommonAncestor(root, p, q)
	fmt.Print(ancestor)
}

func lowestCommonAncestor(root, p, q *TreeNode3) (ancestor *TreeNode3) {
	//1、找出两数的路径存list1，list2
	//2、从左开始遍历，因为要保证x尽可能的最大，

	//var tv *TreeNode3
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	//x的深度尽可能大
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return
}

/*

我们从根节点开始遍历；

如果当前节点就是 p，那么成功地找到了节点；
如果当前节点的值大于 p 的值，说明 p 应该在当前节点的左子树，因此将当前节点移动到它的左子节点；
如果当前节点的值小于 p 的值，说明 p 应该在当前节点的右子树，因此将当前节点移动到它的右子节点。
*/
func getPath(root, target *TreeNode3) (path []*TreeNode3) {
	node := root
	for node != target && node != nil {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}
