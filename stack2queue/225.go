package main

func main() {

}

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	m.queue2 = append(m.queue2, x)
	for len(m.queue1) > 0 {
		m.queue2 = append(m.queue2, m.queue1[0])
		m.queue1 = m.queue1[1:]
	}
	m.queue2, m.queue1 = m.queue1, m.queue2
}

func (m *MyStack) Pop() int {
	v := m.queue1[0]
	m.queue1 = m.queue1[1:]
	return v
}

func (m *MyStack) Top() int {
	return m.queue1[0]
}

func (m *MyStack) Empty() bool {
	return len(m.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
