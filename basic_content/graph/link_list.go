package graph

import "fmt"

type Node struct {
	value int
	next  *Node
}

// 定义链表结构
type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) Size() int {
	return l.size
}

// 添加节点（在链表末尾）
func (l *LinkedList) Add(value int) {
	newNode := &Node{value: value}
	l.size++
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// 删除节点（按值删除）
func (l *LinkedList) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		l.size--
		return
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
		l.size--
	}
}

func (l *LinkedList) Get(index int) int {
	if index < 0 || index >= l.size || l.head == nil {
		return -1 // 防止访问空链表
	}

	current := l.head
	for i := 0; i < index; i++ { // 遍历 index 次
		if current.next == nil { // 防止 nil 访问
			return -1
		}
		current = current.next
	}

	return current.value
}

// 打印链表
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
