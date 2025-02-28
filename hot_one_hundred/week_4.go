package hot_one_hundred

import (
	"container/heap"
	"math/rand"
	"sort"
	"time"
)

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

// 215. 数组中的第K个最大元素
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// findKthLargest 寻找数组中第k大的元素。
// nums: 输入的整数数组。
// k: 指定的第k大元素。
// 返回值: 数组中第k大的元素。
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	// 使用quickSelect算法找出第k大的元素，转换为寻找第n-k小的元素。
	return quickSelect(nums, 0, n-1, n-k)
}

// quickSelect 快速选择算法，用于寻找数组中第k小的元素。
// nums: 输入的整数数组。
// l: 数组左边界。
// r: 数组右边界。
// k: 指定的第k小元素。
// 返回值: 数组中第k小的元素。
func quickSelect(nums []int, l, r, k int) int {
	// 当左右边界相等时，说明找到了目标元素。
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	// 重新调整数组，使得所有小于partition的元素在左边，所有大于partition的元素在右边。
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		// 如果i<j，交换位置。
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 根据k的位置，决定是在左边还是右边继续寻找。
	if k <= j {
		return quickSelect(nums, l, j, k)
	} else {
		return quickSelect(nums, j+1, r, k)
	}
}

// 堆实现
// findKthLargestV2 寻找数组中的第K个最大元素。
// 通过构建最大堆的方式，将数组的前K个元素构建为最大堆，然后通过交换堆顶元素到数组末尾并调整堆的结构，
// 最终找到第K大的元素。
// 参数:
//
//	nums []int: 输入的整数数组。
//	k int: 要寻找的第K大元素的索引。
//
// 返回值:
//
//	int: 数组中第K大的元素。
func findKthLargestV2(nums []int, k int) int {
	heapSize := len(nums)
	// 构建最大堆
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 交换堆顶元素到数组末尾
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		// 调整堆的结构，保持最大堆的性质
		maxHeapify(nums, 0, heapSize)
	}
	// 返回第K大的元素
	return nums[0]
}

// buildMaxHeap 构建最大堆。
// 从数组的中间位置开始，逐个调整堆的结构，确保每个节点的值大于其子节点的值。
// 参数:
//
//	a []int: 输入的整数数组。
//	heapSize int: 堆的大小。
func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

// maxHeapify 最大堆化。
// 给定一个节点索引i，调整堆的结构，确保节点i的值大于其子节点的值，并递归地调整其子树以保持最大堆的性质。
// 参数:
//
//	a []int: 输入的整数数组。
//	i int: 要调整的节点索引。
//	heapSize int: 堆的大小。
func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	// 找到左子节点中最大的值
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	// 找到右子节点中最大的值
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	// 如果最大值不是当前节点，交换值并递归调整子树
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

// 347. 前 K 个高频元素
// topKFrequent 返回数组中出现频率最高的 k 个元素。
// nums: 输入的整数数组
// k: 需要返回的频率最高元素的数量
func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	// 获取每个数字出现次数
	for _, num := range nums {
		occurrences[num]++
	}
	values := [][]int{}
	// 将map中的键值对作为数组的元素，方便后续排序
	for key, value := range occurrences {
		values = append(values, []int{key, value})
	}
	ret := make([]int, k)
	// 使用快速排序算法对values数组进行排序，以找出出现频率最高的k个元素
	qsort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

// qsort 是一个辅助函数，用于对values数组进行快速排序，以便找出出现频率最高的前k个元素。
// values: 需要排序的数组，每个元素是一个包含数字和其出现次数的二元数组
// start, end: 定义了需要排序的数组部分的起始和结束索引
// ret: 用于存储结果的数组
// retIndex: 当前在ret数组中的填充位置
// k: 需要找出的频率最高元素的数量
func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	// 随机选择一个元素作为基准值，以提高排序效率
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	// 使用双指针快速排序算法，将不小于基准值的元素放到左边，小于基准值的元素放到右边
	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]

	// 根据k值和当前的index判断前k大的值在左侧的子数组里还是在右侧的子数组里
	if k <= index-start {
		// 前 k 大的值在左侧的子数组里
		qsort(values, start, index-1, ret, retIndex, k)
	} else {
		// 前 k 大的值等于左侧的子数组全部元素加上右侧子数组中前 k - (index - start + 1) 大的值
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}

// 295. 数据流的中位数
// MedianFinder 是一个用于查找中位数的数据结构，包含两个堆：queMin 和 queMax。
// queMin 是一个小顶堆，queMax 是一个大顶堆。
type MedianFinder struct {
	queMin, queMax hp
}

// ConstructorMedianFinder 是构造函数，返回一个 MedianFinder 实例。
func ConstructorMedianFinder() MedianFinder {
	return MedianFinder{}
}

// AddNum 向数据结构中添加一个数字。
// 如果 num 小于或等于小顶堆的最大值，则将其添加到小顶堆；
// 否则将其添加到大顶堆。这样可以确保小顶堆中的元素都小于等于大顶堆中的元素。
func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

// FindMedian 计算并返回当前所有已添加数字的中位数。
// 如果总数字个数为奇数，则小顶堆的堆顶即为中位数；
// 如果为偶数，则小顶堆和大顶堆堆顶的平均值即为中位数。
func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

// hp 是一个实现了 heap.Interface 的结构体，用于作为小顶堆或大顶堆。
type hp struct{ sort.IntSlice }

// Push 向堆中插入一个元素。
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }

// Pop 从堆中移除并返回最后一个元素。
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 73. 矩阵置零
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 标记第一行和第一列是否需要置零
	rowZero, colZero := false, false

	// 检查第一行是否需要置零
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
			break
		}
	}

	// 检查第一列是否需要置零
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}

	// 从第二行和第二列开始，检查每个元素是否需要置零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 从第二行和第二列开始，将需要置零的元素置零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 将第一行和第一列标记为0
	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if rowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
}
