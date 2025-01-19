package sort

// 归并排序
// mergeSort 实现归并排序
func mergeSort(arr []int) []int {
	n := len(arr)

	if n < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

// 考虑下使用哨兵如何简化边界值的判断
func merge(left []int, right []int) []int {
	i, j := 0, 0

	//temp := make([]int, len(left)+len(right)) 按照这种声明数组，最终结果有什么问题？
	temp := []int{}
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			// 保持稳定性，相等的优先选择左边
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	temp = append(temp, left[i:]...)
	temp = append(temp, right[j:]...)

	return temp
}
