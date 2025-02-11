package hot_one_hundred

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	// 构建索引
	index := make(map[int]int, len(nums))

	for i, _ := range nums {
		index[nums[i]] = i
	}

	for i, _ := range nums {
		if j, ok := index[target-nums[i]]; ok && j != i {
			return []int{i, j}
		}
	}

	return nil
}

// 739. 每日温度: 单调栈
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{} // 单调递减栈，存索引

	for i := 0; i < n; i++ {
		// 维护单调栈：栈顶元素比当前元素小，则可以找到答案
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
			res[prevIndex] = i - prevIndex
		}
		// 当前索引入栈
		stack = append(stack, i)
	}

	return res
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p1, p2 := headA, headB

	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		}

		if p2 == nil {
			p2 = headA
		}

		if p2 == p1 {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
