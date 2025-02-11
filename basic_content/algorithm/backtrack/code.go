package backtrack

/**
回溯算法的思想非常简单，大部分情况下，都是用来解决广义的搜索问题，
也就是，从一组可能的解中，选择出一个满足要求的解。回溯算法非常适合用递归来实现，
在实现的过程中，剪枝操作是提高回溯效率的一种技巧。利用剪枝，我们并不需要穷举搜索所有的情况，从而提高搜索效率。
*/

// 八皇后问题
func solveNQueens(n int) [][]string {
	return nil
}

type packageProblem struct {
	maxWeight int
	capacity  int
	weight    []int
	mem       [][]bool
	res       int
}

// 0-1背包问题
// 输入：
// 1. 背包容量
// 2. 背包内物品的重量
// 3. 物品重量
// 输出：
// 1. 物品最大重量
// i: 决策当前放入那个物品，cw 以及当前重量值
func (p *packageProblem) packageBacktrack(i int, cw int) {
	if len(p.weight) == 0 || p.maxWeight <= 0 || p.capacity <= 0 {
		return
	}

	if i == p.capacity || cw == p.maxWeight {
		// 装满或者装完
		if cw >= p.res {
			p.res = cw
		}
		return
	}

	if p.mem[i][cw] {
		// 已经计算
		return
	}

	// 决策不放入
	p.packageBacktrack(i+1, cw)
	p.mem[i+1][cw] = true
	// 决策放入
	if i < len(p.weight) && cw+p.weight[i] <= p.maxWeight {
		p.packageBacktrack(i+1, cw+p.weight[i])
	}
}

// 正则表达式匹配

//
