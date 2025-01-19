package sort

// 快速排序的递归公式
// 递归公式：quickSort(l, r) = quickSort(l...pivot) + quick(pivot+1...r)
// 推出条件：pivot >= r

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(arr, l, r)
	quickSort(arr, l, pivot-1)
	quickSort(arr, pivot+1, r)
}

// 生成分区点
func partition(arr []int, l int, r int) int {
	if l >= r {
		return l
	}

	pivot := arr[r]
	i := l

	for j := l; j <= r-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[r] = arr[r], arr[i]
	return i
}

// 使用快排的思想， 求解一个无序数组中，第k大的元素下标
// 时间复杂度是多少？
