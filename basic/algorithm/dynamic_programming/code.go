package dynamic_programming

// 0-1背包问题
// weights 物品重量
// capacity 物品容量
// maxWeight 背包最大重量
// 对于一组不同重量、不可分割的物品，我们需要选择一些装入背包，在满足背包最大重量限制的前提下，背包中物品总重量的最大值是多少呢？
// v1: 版本，空间复杂度较高
func packageProblemV1(weights []int, capacity int, maxWeight int) int {
	if len(weights) == 0 || maxWeight <= 0 || capacity <= 0 {
		return 0
	}

	// step1 初始化动态转移数组与边界值，dp[i][j]表示装到容量为i的背包中，重量不超过j的物品，此时最大重量为j
	dp := make([][]bool, capacity, maxWeight+1)
	dp[0][0] = true
	if weights[0] <= maxWeight {
		dp[0][weights[0]] = true
	}

	// step2 状态转移方程：dp[i][j] = dp[i-1][j] || dp[i-1][j-weights[i]]
	for i := 1; i < capacity; i++ {
		// 不放入第i个物品
		for j := 0; j <= maxWeight; j++ {
			if dp[i-1][j] {
				dp[i][j] = dp[i-1][j]
			}
		}

		// 放入第i个物品
		for j := 1; j <= maxWeight-weights[i]; j++ {
			if dp[i-1][j] {
				dp[i][j+weights[i]] = true
			}
		}
	}

	// step 遍历状态数组，找到最大重量
	for i := maxWeight; i >= 0; i-- {
		if dp[capacity-1][i] {
			return i
		}
	}

	return 0
}

// v2: 版本，空间复杂度较低
func packageProblemV2(weights []int, capacity int, maxWeight int) int {
	if len(weights) == 0 || maxWeight <= 0 || capacity <= 0 {
		return 0
	}

	// step1 初始化动态转移数组
	dp := make([]bool, maxWeight+1)
	// 初始化边界值
	dp[0] = true
	if weights[0] <= maxWeight {
		dp[weights[0]] = true
	}

	for i := 0; i <= capacity; i++ {
		// 决策是否放入第i个物品
		for j := maxWeight - weights[i]; j >= 0; j-- {
			if dp[j] {
				dp[j+weights[i]] = true
			}
		}
	}

	for i := maxWeight; i >= 0; i-- {
		if dp[i] {
			return i
		}
	}

	return 0
}

// v3: 版本，求最大价值
// 对于一组不同重量、不可分割的物品, 不同价值，我们需要选择一些装入背包，在满足背包最大重量限制的前提下，背包中物品最大价值是多少呢？
func packageProblemV3(weights []int, values []int, maxWeight int) int {
	if len(weights) == 0 || len(values) == 0 || len(weights) != len(values) || maxWeight <= 0 {
		return 0
	}

	// step1 初始化动态转移数组
	n := len(weights)
	dp := make([][]int, n, maxWeight+1)
	dp[0][0] = 0
	if weights[0] <= maxWeight {
		dp[0][weights[0]] = values[0]
	}

	// step2 状态转移方程：dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i]] + values[i])
	for i := 1; i < n; i++ {
		// 决策是否放入第i个物品
		for j := maxWeight - weights[i]; j >= 0; j-- {
			if dp[i-1][j] > dp[i-1][j-weights[i]]+values[i] {
				// 不放入第i个物品
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-weights[i]] + values[i]
			}
		}
	}

	// step3 遍历状态数组，找到最大重量
	result := 0
	for i := maxWeight; i >= 0; i++ {
		if result < dp[n-1][i] {
			result = dp[n-1][i]
		}
	}

	return result
}
