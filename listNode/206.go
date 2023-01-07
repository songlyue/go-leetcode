package main

/*
*
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/
func main() {

}

type ListNode struct {
	val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	prev := &ListNode{}
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
