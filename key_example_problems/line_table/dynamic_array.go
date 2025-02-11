package line_table

import "errors"

// DynamicArray is a dynamic array.
// 实现一个支持动态扩容的数组
type DynamicArray struct {
	data []int
	size int
}

func (a *DynamicArray) Add(value int) {
	if a.size == cap(a.data) {
		// 动态扩容到原来的两倍
		newCap := max(cap(a.data)*2, 1)
		newData := make([]int, newCap)
		copy(newData, a.data) // 复制旧数据
		a.data = newData
	}

	a.data = append(a.data, value)
	a.size++
}

func (a *DynamicArray) Get(index int) (int, error) {
	if index < 0 || index > len(a.data) {
		return -1, errors.New("index is out of range")
	}

	return a.data[index], nil
}

// NewDynamicArray creates a new dynamic array.
func NewDynamicArray(capacity int) *DynamicArray {
	if capacity < 0 {
		return NewDynamicArray(1)
	}

	return &DynamicArray{
		data: make([]int, 0, capacity),
		size: 0,
	}
}
