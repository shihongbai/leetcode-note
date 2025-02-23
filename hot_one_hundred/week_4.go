package hot_one_hundred

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	// 二分法
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (right-left)>>1 + left
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 1. 二分法找到第一个值等于给定值的位置
// 2. 遍历数组，找到最后一个值等于给定值的位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstTarget := bSearchFirstTarget(nums, target)

	if firstTarget != -1 {
		r := firstTarget
		for r < len(nums) && nums[r] == target {
			r++
		}

		return []int{firstTarget, r - 1}
	}

	return []int{-1, -1}
}

func bSearchFirstTarget(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// 25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 查看当前剩余是否满足翻转个数
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	// 翻转链表
	newHead := reverseK(head, k)

	head.Next = reverseKGroup(cur, k)
	return newHead
}

func reverseK(head *ListNode, k int) *ListNode {
	var pre *ListNode
	curr := head
	for i := 0; i < k; i++ {
		// 找到下一个节点
		next := curr.Next
		// 翻转
		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre
}

// 23. 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	// 分治实现
	// 时间复杂度：O(NlogK)
	// 空间复杂度：O(logK)
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := (l + r) >> 1
	return mergeTwoList(mergeList(lists, l, mid), mergeList(lists, mid+1, r))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l2 != nil {
		cur.Next = l2
	}

	if l1 != nil {
		cur.Next = l1
	}

	return dummy.Next
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1 // 如果数组为空，直接返回-1
	}

	// 如果数组长度为1，直接检查是否等于目标值
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	// 第一次二分查找，寻找旋转点
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[r] {
			// 旋转点在右半部分
			l = mid + 1
		} else if nums[mid] < nums[r] {
			// 旋转点在左半部分
			r = mid
		} else {
			// 如果 nums[mid] == nums[r]，我们无法确定旋转点的位置，缩小右边界
			r--
		}
	}

	// 旋转点的位置
	rotationPoint := l

	// 根据目标值与旋转点的位置判断目标值所在的部分
	if target >= nums[rotationPoint] && target <= nums[len(nums)-1] {
		l, r = rotationPoint, len(nums)-1
	} else {
		l, r = 0, rotationPoint-1
	}

	// 第二次二分查找，查找目标值
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1 // 如果数组为空，直接返回-1
	}

	// 如果数组长度为1，直接检查是否等于目标值
	if len(nums) == 1 {
		return nums[0]
	}

	// 第一次二分查找，寻找旋转点
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[r] {
			// 旋转点在右半部分
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 括号配对，出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

type MinStack struct {
	stack []int
	min   []int // 辅助栈，存放最小值
}

func ConstructorMinStack() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	// 检查辅助栈是否为空，如果为空，则将当前值作为最小值入栈
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if val == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// 260. 只出现一次的数字 III
func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res ^= num
	}

	return res
}

// 169. 多数元素
// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
