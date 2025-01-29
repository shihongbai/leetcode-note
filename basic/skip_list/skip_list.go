package skip_list

import (
	"fmt"
	"math/rand"
)

// golang实现跳表
// Java的实现参考：https://github.com/wangzheng0822/algo/blob/master/java/17_skiplist/SkipList.java

const maxLevel = 16 // 最高层数

// SkipListNode 代表跳表的一个节点
type SkipListNode struct {
	value int
	next  []*SkipListNode // 多层索引指针
}

// SkipList 跳表结构体
type SkipList struct {
	head  *SkipListNode // 头节点
	level int           // 当前最大层数
}

// NewSkipList 初始化跳表
func NewSkipList() *SkipList {
	head := &SkipListNode{
		value: -1,
		next:  make([]*SkipListNode, maxLevel),
	}
	return &SkipList{
		head:  head,
		level: 1,
	}
}

// 随机决定节点的层数
// 理论来讲，一级索引中元素个数应该占原始数据的 50%，二级索引中元素个数占 25%，三级索引12.5% ，一直到最顶层。
// 因为这里每一层的晋升概率是 50%。对于每一个新插入的节点，都需要调用 randomLevel 生成一个合理的层数。
// 该 randomLevel 方法会随机生成 1~MAX_LEVEL 之间的数，且 ：
//
//	  50%的概率返回 1
//	  25%的概率返回 2
//	12.5%的概率返回 3 ...
func (sl *SkipList) randomLevel() int {
	level := 1

	if rand.Float64() < 0.5 && sl.level < maxLevel {
		level++
	}
	return level
}

// Search 查找元素
func (sl *SkipList) Search(value int) bool {
	cur := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].value < value {
			cur = cur.next[i]
		}
	}
	cur = cur.next[0]

	return cur != nil && cur.value == value
}

// Insert 插入元素
func (sl *SkipList) Insert(value int) {
	update := make([]*SkipListNode, maxLevel)
	current := sl.head

	// 1. 查找插入的位置
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current // 记录每一层的前驱节点
	}

	// 2. 生成新节点层数
	level := sl.randomLevel()
	if level > sl.level {
		for i := 0; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	// 3. 每一层插入新节点
	newNode := &SkipListNode{
		value: value,
		next:  make([]*SkipListNode, maxLevel),
	}

	for i := 0; i < level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

// Print 打印跳表结构
func (sl *SkipList) Print() {
	for i := sl.level - 1; i >= 0; i-- {
		current := sl.head.next[i]
		fmt.Printf("Level %d: ", i+1)
		for current != nil {
			fmt.Printf("%d -> ", current.value)
			current = current.next[i]
		}
		fmt.Println("nil")
	}
}

// Delete 从跳表中删除元素
func (sl *SkipList) Delete(value int) {
	update := make([]*SkipListNode, maxLevel)
	current := sl.head

	// 1. 查找目标节点的前驱
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	// 2. 检查目标节点是否存在
	target := current.next[0]
	if target == nil || target.value != value {
		return // 不存在则直接返回
	}

	// 3. 逐层删除
	for i := 0; i < len(target.next); i++ {
		if update[i] != nil {
			update[i].next[i] = target.next[i]
		}
	}

	// 4. 更新跳表的最大层数
	for sl.level > 1 && sl.head.next[sl.level-1] == nil {
		sl.level--
	}
}
