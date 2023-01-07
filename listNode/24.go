package main

import "fmt"

/*
*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
type ListNode struct {
	val  int
	Next *ListNode
}

func main() {
	node2 := &ListNode{3, nil}
	node1 := &ListNode{2, node2}
	head := &ListNode{1, node1}
	swapPairs(head)
	fmt.Printf("%v", head)
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
