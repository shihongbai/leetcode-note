package graph

type Queue struct {
	data []int
}

// 入队
func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

// 出队
func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		panic("Queue is empty")
	}
	val := q.data[0]
	q.data = q.data[1:] // 删除队首元素
	return val
}

// 获取队列长度
func (q *Queue) Size() int {
	return len(q.data)
}

// 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
