package hot_one_hundred

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0 // 进位
	// 遍历链表
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		// 两数相加
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// 19. 删除链表的倒数第 N 个结点
// 进阶: 只扫描一次就删除倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first, second := head, dummy

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for ; first != nil; first = first.Next {
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

// 146. LRU 缓存机制
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	// var inorder func(node *TreeNode)
	// inorder = func(node *TreeNode) {
	//     if node == nil {
	//         return
	//     }

	//     inorder(node.Left)
	//     res = append(res, node.Val)
	//     inorder(node.Right)
	// }

	// inorder(root)
	// return

	// 递归实现
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummy.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

// 138. 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	// 递归加回溯实现
	//cachedNode = map[*Node]*Node{}
	//return deepCopy(head)

	// 迭代 + 节点拆分实现
	if head == nil {
		return nil
	}

	// 链表复制
	for curr := head; curr != nil; curr = curr.Next.Next {
		curr.Next = &Node{Val: curr.Val, Next: curr.Next}
	}

	// 随机节点复制
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
	}

	// 链表拆分
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}

	return headNew
}

// 148. 排序链表
// 自底向上归并实现
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

/*
sortList 使用归并排序算法对链表进行排序
参数:
  - head: *ListNode 链表头节点指针

返回值:
  - *ListNode 排序后的链表头节点指针
*/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 计算链表总长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	// 自底向上的归并排序实现
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		// 遍历整个链表进行分组归并
		for cur != nil {
			// 获取第一个子链表头节点
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			// 获取第二个子链表头节点并切断第一个子链表
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			// 保存后续节点指针并切断第二个子链表
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			// 遍历并切断第二个子链表
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			// 合并两个子链表并连接到结果链表
			prev.Next = merge(head1, head2)

			// 移动prev到已排序部分的末尾
			if prev != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	// 递归实现
	//if root == nil {
	//	return true
	//}
	//
	//return check(root.Left, root.Right)

	// 迭代实现: 队列实现
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		u, v := queue[0], queue[1]
		queue = queue[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}

		queue = append(queue, u.Left)
		queue = append(queue, v.Right)

		queue = append(queue, u.Right)
		queue = append(queue, v.Left)
	}
	return true
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}
