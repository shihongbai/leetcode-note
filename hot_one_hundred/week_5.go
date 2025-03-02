package hot_one_hundred

import "sort"

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	res := []int{}
	// 定义边界
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	for left <= right && top <= bottom {
		// 顺时针打印
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}

// 48. 旋转图像
func rotateGraph(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 74. 搜索二维矩阵
// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
func searchMatrixV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1 // 从右上角开始

	// 逐步向左或向下移动
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++ // 当前元素小于目标，向下移动
		} else {
			j-- // 当前元素大于目标，向左移动
		}
	}

	return false // 未找到目标
}

// 200. 岛屿数量
// dfs实现
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	isLandCount := 0

	// 定义四个方向
	// 0 上 1 下 2 左 3 右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 越界或者当前不是陆地
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}

		// 当前是陆地，标记为已访问
		grid[i][j] = '0'

		// 递归遍历四个方向
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}
	}

	// 遍历整个矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// 找到一个陆地，进行深度优先搜索
				dfs(i, j)
				isLandCount++
			}
		}
	}

	return isLandCount
}

// 994. 腐烂的橘子
// bfs实现
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{}
	freshCount := 0

	// 方向数组：上下左右
	directions := [][]int{
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
		{0, 1},  // 右
	}

	// 初始化队列，添加所有腐烂的橙子，并统计新鲜橙子的数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j}) // 将腐烂的橙子加入队列
			} else if grid[i][j] == 1 {
				freshCount++ // 统计新鲜橙子的数量
			}
		}
	}

	// 如果没有新鲜橙子，直接返回 0
	if freshCount == 0 {
		return 0
	}

	minutes := 0
	// BFS 扩散过程
	for len(queue) > 0 {
		minutes++
		// 当前层级的腐烂橙子的数量
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:] // 弹出队列中的元素

			// 扩散到四个方向
			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]
				// 判断新的坐标是否有效且是新鲜橙子
				if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 {
					grid[newX][newY] = 2                     // 使新鲜橙子腐烂
					freshCount--                             // 新鲜橙子减少
					queue = append(queue, []int{newX, newY}) // 将腐烂的橙子加入队列
				}
			}
		}

		// 如果在这一轮中没有新鲜橙子被腐烂，说明腐烂过程完成
		if freshCount == 0 {
			return minutes
		}
	}

	// 如果结束时还有新鲜橙子没有腐烂，返回 -1
	return -1
}

// 946. 验证栈序列
func validateStackSequences(pushed []int, popped []int) bool {
	st := []int{}
	j := 0
	for _, x := range pushed {
		st = append(st, x)
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}

// 70. 爬楼梯
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	// 滚动数组迭代
	// f(n) = f(n-1) + f(n-2)
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

// 118. 杨辉三角
// generate 生成杨辉三角的前 numRows 行。
// 杨辉三角是一个经典的数学问题，每一行的数字都是上一行相邻两数之和。
// 参数:
//
//	numRows - 需要生成的行数。
//
// 返回值:
//
//	一个二维切片，包含杨辉三角的前 numRows 行。
func generate(numRows int) [][]int {
	// 初始化一个二维切片 ans，用于存放杨辉三角的值。
	ans := make([][]int, numRows)
	for i := range ans {
		// 为每一行初始化一个切片，长度为 i+1，因为行数从0开始计数。
		ans[i] = make([]int, i+1)
		// 每一行的第一个和最后一个元素都是1，这是杨辉三角的特性。
		ans[i][0] = 1
		ans[i][i] = 1
		// 计算每一行中间的元素，中间元素的值是上一行相邻两元素之和。
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	// 返回生成的杨辉三角。
	return ans
}

// 198. 打家劫舍
// 动态规划 + 滚动数组
// rob 函数计算沿着数组掠夺能得到的最大金额。
// 参数 nums 是一个整数数组，代表每个房子中的金额。
// 返回值是能够掠夺到的最大金额。
func rob(nums []int) int {
	// 如果没有房子，掠夺者得不到任何金额。
	if len(nums) == 0 {
		return 0
	}

	// 如果只有一个房子，掠夺者掠夺该房子的金额。
	if len(nums) == 1 {
		return nums[0]
	}

	// first 存储前一个房子的最大掠夺金额。
	// second 存储当前房子的最大掠夺金额。
	first := nums[0]
	second := max(nums[0], nums[1])

	// 从第三个房子开始遍历，更新每个房子的最大掠夺金额。
	for i := 2; i < len(nums); i++ {
		// 更新 first 和 second 的值，以反映当前房子的最大掠夺金额。
		first, second = second, max(first+nums[i], second)
	}
	// 返回最后一个房子的最大掠夺金额。
	return second
}

// 279. 完全平方数
func numSquares(n int) int {
	dp := make([]int, n+1)

	// 初始化，dp[0] = 0，其余初始化为较大值
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	// 动态规划，计算 dp[i]
	for i := 1; i <= n; i++ {
		j := 1
		for j*j <= i {
			dp[i] = min(dp[i], dp[i-j*j]+1)
			j++
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	// 初始化 dp 数组，dp[i] 表示构成金额 i 所需的最小硬币数
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 初始化为一个较大的数
	}
	dp[0] = 0 // 0 元需要 0 个硬币

	// 动态规划计算每个金额所需的最小硬币数
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// 如果 dp[amount] 还是初始值，说明无法组成该金额
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

// 11. 盛最多水的容器
// 双指针实现
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	// 双指针遍历
	for left < right {
		// 计算当前容器高度较低的一个
		var h int
		if height[left] < height[right] {
			h = height[left]
		} else {
			h = height[right]
		}

		// 计算面积
		area := h * (right - left)
		if area > result {
			result = area
		}
		// 移动高度较小的一侧指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

// 15. 三数之和
// 双指针实现
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	// 数组排序
	sort.Ints(nums)

	// 双指针遍历
	for i := 0; i < n-2; i++ {
		// 固定第一个数, 因为数组有序，后面全部大于零，直接跳出循环
		if nums[i] > 0 {
			break
		}

		// 跳过重复数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针搜索
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// 找到满足条件的三元组
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return result
}

// 42. 接雨水
// 双指针实现
// 时间复杂度 O(n); 空间复杂度 O(1)
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	// 双指针移动左右两端
	if left < right {
		if height[left] < height[right] {
			// 左侧较小，决定了容器高度
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				// 累加水量
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

// 3. 无重复字符的最长子串
// 滑动窗口实现
func lengthOfLongestSubstring(s string) int {
	// 构建一个map， 用于记录字符上次出现的位置
	charIndexMap := make(map[rune]int)
	maxLength := 0
	left := 0

	// 遍历字符串，使用滑动窗口的方式
	for right, ch := range s {
		// 如果字符串出现过，并且出现的位置不在当前窗口的左边界
		if prevIndex, exists := charIndexMap[ch]; exists && prevIndex >= left {
			// 更新窗口的左边界
			left = prevIndex + 1
		}
		// 更新字符出现的位置
		charIndexMap[ch] = right
		// 计算窗口长度
		if currLen := right - left + 1; currLen > maxLength {
			maxLength = currLen
		}
	}

	return maxLength
}

// 438. 找到字符串中所有字母异位词
// 滑动窗口实现
func findAnagrams(s string, p string) []int {
	var res []int
	ns, np := len(s), len(p)
	if ns < np {
		return res
	}

	// 使用长度26的数组，统计p中每个字符出现的次数
	var pCount, sCount [26]int

	// 初始化
	for i := 0; i < np; i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	// 如果两个数组的类型和长度相同，可以直接使用 == 操作符进行比较。
	//这会逐个元素比较，即判断两个数组中每个位置的数字是否都相等。
	if pCount == sCount {
		res = append(res, 0)
	}

	// 滑动窗口
	for i := np; i < ns; i++ {
		// 加入新的字符串到窗口
		sCount[s[i]-'a']++
		// 移除窗口最左侧的字符
		sCount[s[i-np]-'a']--

		// 如果完全匹配，计入本次下标
		if pCount == sCount {
			res = append(res, i-np+1)
		}
	}
	return res
}
