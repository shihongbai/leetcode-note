package hot_one_hundred

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	mProfit := 0

	for i := 1; i < len(prices); i++ {
		// 如果当前价格比最小价格低，更新最小价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// 更新最大利润
			mProfit = max(mProfit, prices[i]-minPrice)
		}
	}

	return mProfit
}

// 55. 跳跃游戏
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}

	// 最远能达到的距离
	maxReach := 0

	// 遍历
	for i := 0; i < n; i++ {
		// 如果当前位置超过了最远距离，则返回false
		if i > maxReach {
			return false
		}

		// 更新最远距离
		maxReach = max(maxReach, i+nums[i])
		// 如果最远距离超过了数组长度，则返回true
		if maxReach >= n-1 {
			return true
		}
	}

	return maxReach >= n-1
}

// 46. 全排列
// 时间复杂度：O(n!)
// 空间复杂度：O(n)
func permute(nums []int) [][]int {
	var res [][]int

	// 定义回溯函数, index 表示当前递归到哪个位置
	var backtrack func(int)
	backtrack = func(index int) {
		// 递归终止条件
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}

		// 遍历当前位置只有所有的元素
		for i := index; i < len(nums); i++ {
			// 交换当前元素与第i个元素
			nums[index], nums[i] = nums[i], nums[index]
			// 递归处理
			backtrack(index + 1)
			// 回溯
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	backtrack(0)
	return res
}

// 208. 实现 Trie (前缀树)
type Trie struct {
	children [26]*Trie // 子节点数组，索引对应小写字母a-z
	isEnd    bool      // 是否是单词结尾
}

func ConstructorTrie() Trie {
	return Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
}

// 时间复杂度：O(n)
// 空间复杂度：O(n)
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		// 计算索引
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		// 移动到子节点
		node = node.children[idx]
	}
	// 标记单词结束
	node.isEnd = true
}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		// 计算索引
		idx := ch - 'a'
		if node.children[idx] == nil {
			// 路径不存在， 返回失败
			return false
		}
		// 移动到子节点
		node = node.children[idx]
	}
	// 检查是否单词结尾
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		// 计算索引
		idx := ch - 'a'
		if node.children[idx] == nil {
			// 路径不存在， 返回失败
			return false
		}
		// 移动到子节点
		node = node.children[idx]
	}
	return true
}

// 78. 子集
// 回溯
func subsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(int, []int) // 定义回溯函数

	backtrack = func(start int, current []int) {
		// 每次进入递归，将当前子集加入结果集
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i]) // 选择当前元素
			backtrack(i+1, current)            // 递归处理
			current = current[:len(current)-1] // 回溯
		}
	}

	backtrack(0, []int{})
	return result
}
