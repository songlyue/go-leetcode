package main

/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val和next。val是当前节点的值，next是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性prev以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第index个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为val的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第index个节点之前添加值为val 的节点。如果index等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引index 有效，则删除链表中的第index 个节点。

*/

func main() {

}

type node struct {
	val        int
	next, prev *node
}

type MyLinkedList struct {
	head, tail *node
	size       int
}

func Constructor() MyLinkedList {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{head, tail, 0}
}
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	curr := l.head
	for i := 0; i <= index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}
	var pred, succ *node
	pred = l.head
	for i := 0; i < index; i++ {
		pred = pred.next
	}
	succ = pred.next
	l.size++
	toAdd := &node{val, succ, pred}
	pred.next = toAdd
	succ.prev = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	var pred, succ *node
	pred = l.head
	for i := 0; i < index; i++ {
		pred = pred.next
	}
	succ = pred.next.next
	l.size--
	pred.next = succ
	succ.prev = pred

}
