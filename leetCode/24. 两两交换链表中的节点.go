package main

import "fmt"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

var head *ListNode  //头节点
var curr *ListNode  //当前节点

//创建头节点
func CreateHeadNode(data int) *ListNode{
	//创建Node头节点
	var node *ListNode = new(ListNode)
	//封装数据
	node.Val = data
	//指定下一个节点地址，因为还没添加所以是nil
	node.Next = nil
	//第一次创建头节点
	head = node
	curr = node
	return node
}

//向节点中添加数据
func AddNode (data int) *ListNode{
	var newNode *ListNode = new(ListNode)
	//添加信息
	newNode.Val = data
	newNode.Next = nil
	//指定刚添加的节点是当前节点
	curr.Next = newNode
	//新创建的节点在赋给当前节点
	curr = newNode
	//fmt.Println(curr)
	return newNode
}

func main() {
	//创建头节点
	CreateHeadNode(1)
	//向节点中添加数据
	AddNode(2)
	AddNode(3)
	AddNode(4)
	fmt.Print(swapPairs(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	//空直接返回
	if head == nil {
		return nil
	}
	//当个节点直接返回
	if head.Next == nil {
		return head
	}
	// 保存当前节点、当前节点的前置节点，最终返回的必然是第一次交换的head.Next（即：pre）

	curr, pre := head, head.Next
	// 当当前节点被交换到前置节点的位置时，当前节点的next应该是前置节点的Next
	curr.Next = swapPairs(pre.Next)
	// 前后对掉后，前置节点的next应该指向当前节点
	pre.Next = curr


	return pre
}
