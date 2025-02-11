package skip_list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewSkipList(t *testing.T) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	sl := NewSkipList()

	// 插入元素
	sl.Insert(10)
	sl.Insert(20)
	sl.Insert(15)
	sl.Insert(30)
	sl.Insert(5)
	sl.Insert(1)
	sl.Insert(3)
	sl.Insert(2)
	sl.Insert(8)
	sl.Insert(17)
	sl.Insert(13)
	sl.Insert(9)

	// 跳表结构
	sl.Print()

	// 查找元素
	fmt.Println("查找 15:", sl.Search(15))   // true
	fmt.Println("查找 100:", sl.Search(100)) // false

	// 删除元素
	sl.Delete(15)
	fmt.Println("删除 15 之后：")
	sl.Print()
}
