package sort

// 冒泡排序、插入排序、选择排序

// 最好O(n)
// 最坏O(n^2)
// 平均O(n^2)
func bubbleSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	for i := 0; i < n; i++ {
		flag := false // 提前退出循环的标志
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	return arr
}

func insertSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		// 选择插入位置
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}

	return arr
}

// 最好、最坏都是O(n^2)
func selectSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

func shellSort(arr []int) {
	n := len(arr)

	if n <= 1 {
		return
	}

	// 进行gap拆分
	for gap := n / 2; gap > 0; gap /= 2 {
		// gap内部进行插入排序
		for i := gap; i < n; i += 1 {
			// 从gap开始，对每个分组进行插入排序
			value := arr[i]
			j := i
			for j >= gap && arr[j-gap] > value {
				arr[j] = arr[j-gap] // 将前面元素后移
				j -= gap
			}

			arr[j] = value // 将元素插入正确的位置
		}
	}

	return
}
