package util

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(createTree("1,null,2,3"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(s string) (tree *TreeNode) {
	if s == "" {
		return tree
	}
	str := strings.Split(s, ",")
	head, _ := strconv.Atoi(str[0])
	tree = &TreeNode{head, nil, nil}
	for i := 1; i < len(str)-1; i++ {
		if i == len(str)-1 {
			break
		}
		if str[i] != "null" && str[i+1] != "null" {
			strChangeInt1, _ := strconv.Atoi(str[i])
			strChangeInt2, _ := strconv.Atoi(str[i+1])
			tree.Left.Val = strChangeInt1
			tree.Right.Val = strChangeInt2
		}
		if str[i] != "null" && str[i+1] == "null" {
			strChangeInt1, _ := strconv.Atoi(str[i])
			strChangeInt2, _ := strconv.Atoi(str[i+1])
			tree.Left.Val = strChangeInt1
			tree.Right.Val = strChangeInt2
		}
		if str[i] != "null" && str[i+1] != "null" {
			strChangeInt1, _ := strconv.Atoi(str[i])
			strChangeInt2, _ := strconv.Atoi(str[i+1])
			tree.Left.Val = strChangeInt1
			tree.Right.Val = strChangeInt2
		}

	}

	return tree
}
