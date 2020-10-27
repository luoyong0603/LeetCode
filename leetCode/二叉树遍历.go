package main

import "fmt"

//创建一个结构体
type treeNode struct {
	No          int
	value       string
	left, right *treeNode
}

func main() {
	//创建一棵树
	root := treeNode{1, "A", nil, nil}
	root.left = &treeNode{value: "B", No: 2}
	root.right = &treeNode{value: "C", No: 3}
	root.left.left = &treeNode{value: "D", No: 4}
	root.left.right = &treeNode{value: "F", No: 5}
	root.left.right.left = &treeNode{value: "E", No: 8}
	//root.left.right.left = new(treeNode)
	//root.left.right.left.value = "E"
	//root.left.right.No = 8
	root.right.left = &treeNode{value: "G", No: 6}
	root.right.left.right = &treeNode{value: "H", No: 9}
	root.right.right = &treeNode{value: "I", No: 7}
	fmt.Println("先序遍历：")
	traverse(&root)
	fmt.Println("")
	fmt.Println("中序遍历：")
	traverse1(&root)
	fmt.Println("")
	fmt.Println("后序遍历：")
	traverse2(&root)

}

//先序遍历
func traverse(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Printf("no=%d name=%s\n", node.No, node.value)
	traverse(node.left)
	traverse(node.right)
}

//中序遍历
func traverse1(node *treeNode) {
	if node == nil {
		return
	}
	traverse1(node.left)
	fmt.Printf("no=%d name=%s\n", node.No, node.value)
	traverse1(node.right)
}

//后序遍历
func traverse2(node *treeNode) {

	if node == nil {
		return
	}
	traverse2(node.left)
	traverse2(node.right)
	fmt.Printf("no=%d name=%s\n", node.No, node.value)
}
