package sort

import (
	"fmt"
	"testing"
)

func Test_shellSort(t *testing.T) {
	arr := []int{45, 23, 89, 77, 12, 90, 2, 9, 54}
	shellSort(arr)
	fmt.Println(arr)
}

func Test_mergeSort(t *testing.T) {
	arr := []int{45, 23, 89, 77, 12, 90, 2, 9, 54}
	fmt.Println(mergeSort(arr))
}

func Test_quickSort(t *testing.T) {
	arr := []int{45, 23, 89, 77, 12, 90, 2, 9, 54}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
