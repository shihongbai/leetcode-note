package tree

// Heap 堆
type Heap struct {
	arr      []int
	count    int
	capacity int
}

func Sort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	// 堆化: 时间复杂度O(n)
	for i := n / 2; i >= 1; i-- {
		heapfiy(arr, n, i)
	}

	// 排序
	k := n
	for k > 1 {
		arr[1], arr[k] = arr[k], arr[1]
		k--
		heapfiy(arr, k, 1)
	}

	return arr
}

func heapfiy(arr []int, n int, i int) {
	for {
		maxPox := i
		if maxPox*2 < n && arr[maxPox*2] <= arr[maxPox] {
			arr[maxPox], arr[maxPox*2] = arr[maxPox*2], arr[maxPox]
			maxPox *= 2
		} else if maxPox*2+1 < n && arr[maxPox*2+1] <= arr[maxPox] {
			arr[maxPox], arr[maxPox*2+1] = arr[maxPox*2+1], arr[maxPox]
			maxPox = maxPox*2 + 1
		}

		if maxPox == i {
			break
		}

		i = maxPox
	}
}

// RemoveTop 移除堆顶元素
func (h *Heap) RemoveTop() bool {
	if h.count == 0 {
		return false
	}

	h.arr[1] = h.arr[h.count]
	h.heapfiy()
	h.count--
	return true
}

func (h *Heap) Insert(v int) bool {
	if h.count == h.capacity {
		// todo 动态扩容
		return false
	}

	// 插入并堆化
	h.count++
	h.arr[h.count] = v
	i := h.count
	for i/2 > 0 && h.arr[i] > h.arr[i/2] {
		h.arr[i], h.arr[i/2] = h.arr[i/2], h.arr[i]
		i /= 2
	}

	return true
}

// 自上往下堆化
func (h *Heap) heapfiy() {
	i := 1
	maxPos := i
	for {
		if i*2 < h.count && h.arr[i] > h.arr[i*2] {
			h.arr[i], h.arr[i*2] = h.arr[i*2], h.arr[i]
			maxPos = maxPos * 2
		} else if i*2+1 < h.count && h.arr[i] > h.arr[i*2+1] {
			h.arr[i], h.arr[i*2+1] = h.arr[i*2+1], h.arr[i]
			maxPos = maxPos*2 + 1
		}

		if maxPos == i {
			break
		}

		i = maxPos
	}
}

func NewHeap(capacity int) *Heap {
	if capacity <= 0 {
		return &Heap{
			arr:      make([]int, 10),
			count:    0,
			capacity: 10,
		}
	}

	return &Heap{
		arr:      make([]int, capacity),
		count:    0,
		capacity: capacity,
	}
}
