package hot_one_hundred

import "math"

// 543. 二叉树的直径
var ans = 0

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	depth(root)
	return ans - 1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := depth(node.Left)
	r := depth(node.Right)
	ans = max(l+r+1, ans) // 计算d_node即L+R+1 并更新ans
	return max(l, r) + 1
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	result := make([][]int, 0)
	for len(queue) > 0 {
		// 获取当前队列长度
		layer := make([]int, 0)
		// 遍历当前队列
		for i := 0; i < len(queue); i++ {
			layer = append(layer, queue[i].Val)
		}
		result = append(result, layer)

		// 更新队列
		newQ := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQ = append(newQ, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQ = append(newQ, queue[i].Right)
			}
		}
		queue = newQ
	}

	return result
}

// 108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	// 递归：数组二分
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, l int, r int) bool {
	if node == nil {
		return true
	}

	if node.Val <= l || node.Val >= r {
		return false
	}

	return helper(node.Left, l, node.Val) && helper(node.Right, node.Val, r)
}

// 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}

// 199. 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := make([]int, 0)
	for len(queue) > 0 {
		// 获取当前队列长度
		layer := make([]int, 0)
		// 遍历当前队列
		for i := 0; i < len(queue); i++ {
			layer = append(layer, queue[i].Val)
		}
		result = append(result, layer[len(layer)-1])

		// 更新队列
		newQ := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQ = append(newQ, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQ = append(newQ, queue[i].Right)
			}
		}
		queue = newQ
	}

	return result
}

// 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	helper2(root)
}

func helper2(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	left := helper2(node.Left)
	right := helper2(node.Right)

	node.Left = nil
	node.Right = left
	p := node
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return node
}

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
