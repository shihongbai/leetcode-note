package hot_one_hundred

import (
	"slices"
)

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

// 49. 字母异位词分组
// 示例 1:
//
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 方案1：排序hash
// 方案2：计数hash
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	if len(strs) == 1 {
		return [][]string{strs}
	}

	// 构建一个hash索引
	sortedHash := make(map[string][]string)
	for _, s := range strs {
		bytes := []byte(s)
		slices.Sort(bytes)
		sortedStr := string(bytes)

		sortedHash[sortedStr] = append(sortedHash[sortedStr], s)
	}

	result := make([][]string, 0, len(sortedHash))
	for _, v := range sortedHash {
		result = append(result, v)
	}

	return result
}

// 128. 最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 示例 1：
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
func longestConsecutive(nums []int) int {
	// 创建一个哈希集合，存储数组中的所有数字
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0

	// 遍历每个数字
	for num := range numSet {
		// 只有当当前数字是一个序列的起始点时才进行计算
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// 向右扩展连续的序列
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// 更新最长序列的长度
			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}

	return longest
}

// 283. 移动零
// 示例 1:
//
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	// 指针 j 用于记录非零元素应该放置的位置
	j := 0

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 如果当前元素是非零元素，将其移动到前面的 j 位置
			nums[j], nums[i] = nums[i], nums[j]
			j++ // 移动 j 指针
		}
	}
}
