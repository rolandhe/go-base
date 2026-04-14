package swiss_kit

import (
	"fmt"
	"testing"
)

func TestPQ(t *testing.T) {
	k := 5

	// 小顶堆：最小的在顶
	pq := NewLimitedPriorityQueue[int](k, true, func(a, b int) bool {
		return a < b
	})

	nums := []int{5, 2, 55, 9, 1, 32, 7, 16, 3, 10, 8}

	for _, num := range nums {
		pq.Push(num)
	}

	ret := pq.ToSlice()

	for _, v := range ret {
		fmt.Println(v)
	}
}
