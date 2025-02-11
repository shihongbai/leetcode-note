package link_list

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(4)

	cache.PutValue("1", 1)
	fmt.Println(cache.ToString())
	cache.PutValue("2", 2)
	fmt.Println(cache.ToString())
	cache.PutValue("3", 3)
	fmt.Println(cache.ToString())
	cache.PutValue("4", 4)
	fmt.Println(cache.ToString())
	cache.PutValue("5", 5)
	fmt.Println(cache.ToString())
	cache.GetValue("3")
	fmt.Println(cache.ToString())
	cache.PutValue("6", 6)
	fmt.Println(cache.ToString())
}
