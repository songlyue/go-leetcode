package main

import "fmt"

/**
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有push to top,peek/pop from top,size, 和is empty操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


*/

func main() {
	ints := []int{1, 2, 3}
	fmt.Printf("%v", ints[1:])
}

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}
func (q *MyQueue) Pop() int {
	for len(q.outStack) == 0 {
		q.in2out()
	}
	result := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return result
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	result := q.outStack[len(q.outStack)-1]
	return result
}

func (q *MyQueue) Empty() bool {
	return len(q.outStack) == 0 && len(q.inStack) == 0

}
