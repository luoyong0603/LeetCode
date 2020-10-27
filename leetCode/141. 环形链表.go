package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	//快慢两个指针
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		//慢指针每次走一步
		slow = slow.Next
		//快指针每次走两步
		fast = fast.Next.Next
		//如果相遇，说明有环，直接返回true
		if slow == fast {
			return true
		}
	}
	//否则就是没环
	return false
}
