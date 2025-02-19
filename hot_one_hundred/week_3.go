package hot_one_hundred

import (
	"math"
	"sort"
)

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

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 找到中序列表的根节点
	root := &TreeNode{Val: preorder[0]}
	mid := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			mid = i
			break
		}
	}

	// 递归
	root.Left = buildTree(preorder[1:len(inorder[:mid])+1], inorder[:mid])
	root.Right = buildTree(preorder[len(inorder[:mid])+1:], inorder[mid+1:])
	return root
}

// 437. 路径总和 III
func pathSum(root *TreeNode, targetSum int) int {
	// 递归：前缀和: 由根结点到当前结点的路径上所有节点的和。
	// key: 前缀和，value: 出现次数
	prefix := map[int]int{0: 1}

	return dfs(root, prefix, 0, targetSum)
}

// dfs 深度优先搜索函数，用于计算满足条件的路径数量。
// node: 当前节点, prefix: 前缀和的记录, cur: 当前路径的前缀和, target: 目标和
// 返回值: 当前节点为起点的满足条件的路径数量
func dfs(node *TreeNode, prefix map[int]int, cur, target int) int {
	if node == nil {
		return 0
	}

	ret := 0
	cur += node.Val

	// 检查是否存在前缀和等于当前前缀和减去目标值的路径
	ret = prefix[cur-target]
	// 更新前缀和的记录
	prefix[cur] = prefix[cur] + 1
	// 递归搜索左子树和右子树
	ret += dfs(node.Left, prefix, cur, target)
	ret += dfs(node.Right, prefix, cur, target)
	// 回溯，恢复前缀和的记录
	prefix[cur] = prefix[cur] - 1

	return ret
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归实现
	if root == nil {
		return nil
	}

	// 遍历到目标子节点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 递归搜索左子树和右子树
	// 公共祖先不可能即在左边又在右边，满足条件的只能是根节点
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 检查左子树和右子树是否找到公共祖先
	if left == nil || right == nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	// 递归
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子树对总和的最大贡献值
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和为当前节点值加上左右子树最大贡献值
		priceNewPath := node.Val + leftGain + rightGain
		// 更新最大路径和
		maxSum = max(maxSum, priceNewPath)

		// 返回当前节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

// 53. 最大子数组和
// 设f(i)表示以nums[i]结尾的最大子数组和
// 动态转移方程为f(i) = max(f(i-1)+nums[i], nums[i])
func maxSubArray(nums []int) int {
	// 动态规划实现
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	return maxNum
}

// 56. 合并区间
// 如果当前区间的左端点在数组merged中最后一个区间的右端点之后，那么它们不会重合，我们可以直接将这个区间加入数组merged的末尾；
// 否则，它们重合，我们需要用当前区间的右端点更新数组merged中最后一个区间的右端点，将其置为二者的较大值。
func mergeFunc(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 基于左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		l := intervals[i][0]
		r := intervals[i][1]
		if len(merged) == 0 || merged[len(merged)-1][1] < l {
			merged = append(merged, []int{l, r})
		} else {
			// 更新merged中最后一个区间的右端点
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], r)
		}

	}
	return merged
}

// 189. 旋转数组
// 方法1: 创建一个新数组，将nums旋转k位，然后覆盖nums
// 方法2: 使用环状替换，将数组旋转k位，然后覆盖nums
// 方法3: 翻转数组，然后翻转前k位，再翻转剩余的元素
func rotate(nums []int, k int) {
	// 方法3
	k = k % len(nums)
	// 1. 反转整个数组
	reverse(nums)
	reverse(nums[:k])
	// 2. 反转后面 [k % n, n-1]
	reverse(nums[k:])

}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
