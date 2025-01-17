package queue

// 循环队列的实现

type MyCircularQueue struct {
	// 固定长度的数组
	items []int
	n     int
	head  int
	tail  int
}

// 入队列
func (q MyCircularQueue) enqueue(v int) bool {
	// 队列是否已经满
	if q.isFull() {
		return false
	}

	// 体会这段代码为什么能简化为下面这段代码
	//if q.tail == len(q.items)-1 {
	//	q.items[0] = v
	//	q.tail = 0
	//} else {
	//	q.items[q.tail] = v
	//	q.tail++
	//}

	// 简化后
	// 循环，周期性相关的题目，都要考虑到用取模操作
	q.items[q.tail] = v
	q.tail = (q.tail + 1) % q.n

	return true
}

// 出队列
func (q MyCircularQueue) dequeue() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}

	v := q.items[q.head]
	// 关键
	q.head = (q.head + 1) % q.n
	return v, true
}

func (q MyCircularQueue) isFull() bool {
	return (q.tail+1)%q.n == q.head
}

func (q MyCircularQueue) isEmpty() bool {
	return q.head == q.tail
}

func NewMyCircularQueue(n int) *MyCircularQueue {
	return &MyCircularQueue{
		items: make([]int, n),
		n:     n,
		head:  0,
		tail:  0,
	}
}
