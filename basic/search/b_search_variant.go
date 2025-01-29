package search

// 二分查找的变体问题

// 变体一：查找第一个值等于给定值的元素
// arr：有序数组，但是存在重复数字，请找到第一个目标数值的下标，如果没有返回-1
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
