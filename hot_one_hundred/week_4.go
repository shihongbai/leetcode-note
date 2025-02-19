package hot_one_hundred

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	// 二分法
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (right-left)>>1 + left
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 1. 二分法找到第一个值等于给定值的位置
// 2. 遍历数组，找到最后一个值等于给定值的位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstTarget := bSearchFirstTarget(nums, target)

	if firstTarget != -1 {
		r := firstTarget
		for r < len(nums) && nums[r] == target {
			r++
		}

		return []int{firstTarget, r - 1}
	}

	return []int{-1, -1}
}

func bSearchFirstTarget(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
