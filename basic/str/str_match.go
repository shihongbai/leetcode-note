package str

func computeNext(str string) []int {
	if len(str) == 0 {
		return []int{}
	}

	m := len(str)
	next := make([]int, m)
	j := 0

	for i := 1; i < m; i++ {
		// 回溯j直到找到匹配的前后缀
		for j > 0 && str[i] != str[j] {
			j = next[j-1]
		}

		if str[i] == str[j] {
			j++ // 前后缀长度增加
		}
		next[i] = j
	}
	return next
}

// kmp算法实现字符串匹配
func kmp(s string, pattern string) []int {
	n, m := len(s), len(pattern)
	next := computeNext(pattern)
	j := 0            // 匹配子串 p 的指针
	result := []int{} // 记录匹配位置

	for i := 0; i < n; i++ {
		// 如果不匹配，按照next数组回溯j
		for j > 0 && s[i] != pattern[j] {
			j = next[j-1]
		}
		if s[i] == pattern[j] {
			j++
		}

		// 完整匹配 pattern, 记录位置
		if j == m {
			result = append(result, i-m+1)
			j = next[j-1] // 继续匹配
		}
	}

	return result
}
