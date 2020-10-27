package main

import "fmt"

/*
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
假定 BST 有如下定义：
结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序
进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/
//创建一个结构体
type treeNode3 struct {
	value       int
	left, right *treeNode3
}

func main() {
	root := treeNode3{1, nil, nil}
	root.right = &treeNode3{value: 2}
	root.left = &treeNode3{value: 3}
	root.left.left = &treeNode3{value: 3}
	root.right.right = &treeNode3{value: 2}
	mode := findMode(&root)
	fmt.Print(mode)
}

func findMode(root *treeNode3) []int {
	var m = make(map[int]int)
	m = traverses(m, root)
	var list []int
	var list1 []int
	var max int = 1
	var index int
	for k, v := range m {
		if v > max {
			max = v
			index = k
			//清空与max相等追加进list1中的元素
			list1 = []int{}
		}
		//目的是为了防止出现同max的元素
		if v == max && v != 1 {
			list1 = append(list1, k)
		}
		list = append(list, k)
	}
	if max == 1 {
		return list
	}
	if len(list1) == 0 {
		list1 = []int{}
		list1 = append(list1, index)
	}
	return list1
}

//先序遍历
func traverses(m map[int]int, node *treeNode3) map[int]int {
	if node == nil {
		return map[int]int{}
	}
	//累积重复值
	m[node.value]++
	traverses(m, node.left)
	traverses(m, node.right)
	return m
}

