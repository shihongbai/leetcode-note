package sort

// 计数排序，arr是数组，n是数组大小。假设数组中存储的都是非负整数
func countSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	// 查找技术的数据范围
	var maxv = arr[0]
	for i := 1; i < n; i++ {
		if maxv < arr[i] {
			maxv = arr[i]
		}
	}

	// 申请一个计数数组
	c := make([]int, maxv+1)
	for i, _ := range c {
		c[i] = 0
	}

	// 依次累加
	for i := 1; i <= maxv; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 计数排序的关键逻辑
	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}

	// 将结果拷贝回arr
	for i := 0; i < n; i++ {
		arr[i] = r[i]
	}
	return arr
}
