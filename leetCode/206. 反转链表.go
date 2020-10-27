package main

func main() {
	//var str = "12345"
}

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

 type ListNode struct {
	     Val int
	     Next *ListNode
 }

 //利用 Go 的多重赋值语法糖，一步到位
func reverseList(head *ListNode) *ListNode {
	var prev  *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

 //非递归：
func reverseList1(head *ListNode) *ListNode {
	var pre  *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
