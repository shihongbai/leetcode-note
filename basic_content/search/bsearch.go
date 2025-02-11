package search

import "math"

// 二分查找的递归实现
func bSearchRecursion(arr []int, target int) int {
	return bsearchInternally(arr, 0, len(arr)-1, target)
}

func bsearchInternally(arr []int, low int, high int, target int) int {
	if low > high {
		return -1
	}

	mid := low + ((high - low) >> 1)
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return bsearchInternally(arr, low, mid-1, target)
	} else {
		return bsearchInternally(arr, mid+1, high, target)
	}
}

// 二分查找的非递归实现
func bSearch(arr []int, target int) bool {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}

// 通过二分查找实现求a的平方根，保留小数点后六位
// 解题思路：通过二分查找分别确认整数位、小数第一位、小数第二位...
func squareRoot(a int) float64 {
	// 对于 x 为负数时，平方根是无定义的
	if a < 0 {
		return math.NaN()
	}

	// 对于 0 和 1，直接返回
	if a == 0 || a == 1 {
		return float64(a)
	}

	result := float64(0)
	l, r := 1, a
	var mid = 0
	// 求解整数位
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == a {
			return result
		} else if mid*mid > a {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 求解小数第一位
	result = result + float64(mid)
	l = 0
	r = 9
	for l <= r {
		mid = l + (r-l)/2
		temp := float64(mid)/10 + result
		if temp == float64(a) {
			result = temp
		} else if temp > float64(a) {
			r = int(mid) - 1
		} else {
			l = int(mid) + 1
		}
	}

	for i := 1; i <= 6; i++ {
		value := findTargetPlaceValue(a, i, result)
		result = result + float64(value)/math.Pow(10, float64(i))
	}

	return result
}

func findTargetPlaceValue(target, place int, prefix float64) int {

	l, r := 0, 9
	mid := 0
	for r-l > 1 {
		mid = l + (r-l)/2
		temp := float64(mid)/math.Pow(10, float64(place)) + prefix
		if temp*temp > float64(target) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return mid
}

// 方法二：精度确认法
// sqrtBinarySearch 使用二分法计算平方根，保留到小数点后第六位
func sqrtBinarySearch(x float64) float64 {
	// 设置精度，保留到小数点第六位
	const epsilon = 1e-7

	// 对于 x 为负数时，平方根是无定义的
	if x < 0 {
		return math.NaN()
	}

	// 对于 0 和 1，直接返回
	if x == 0 || x == 1 {
		return x
	}

	// 初始化左边界和右边界
	left, right := 0.0, x

	// 进行二分法查找平方根
	for right-left > epsilon {
		// 计算中间值
		mid := (left + right) / 2
		square := mid * mid

		// 判断中间值的平方和目标值的大小
		if square == x {
			return mid
		} else if square < x {
			left = mid // mid 可能是平方根
		} else {
			right = mid // mid 可能大于平方根
		}
	}

	// 返回结果，保留到小数点第六位
	return math.Round((left+right)/2*1e6) / 1e6
}
