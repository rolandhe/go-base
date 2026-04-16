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

	ret := pq.PopToSlice()

	for _, v := range ret {
		fmt.Println(v)
	}
}

type Integer struct {
	v int
}

func TestStPQ(t *testing.T) {
	k := 5

	// 小顶堆：最小的在顶
	pq := NewLimitedPriorityQueue[*Integer](k, true, func(a, b *Integer) bool {
		return a.v < b.v
	})

	nums := []int{5, 2, 55, 9, 1, 32, 7, 16, 3, 10, 8}

	for _, num := range nums {
		pq.Push(&Integer{
			v: num,
		})
	}

	ret := pq.OnceToSlice()

	for _, v := range ret {
		fmt.Println(v.v)
	}
}

func TestIndexSort(t *testing.T) {
	k := 5
	// 小顶堆：最小的在顶
	pq := NewLimitedPriorityQueue[*Integer](k, true, func(a, b *Integer) bool {
		return a.v < b.v
	})

	nums := []int{5, 2, 55, 9, 1, 32, 7, 16, 3, 10, 8}

	for _, num := range nums {
		pq.Push(&Integer{
			v: num,
		})
	}

	ret := pq.OnceToSlice()
	var ids []int
	for _, v := range ret {
		ids = append(ids, v.v)
	}
	ret[0], ret[1] = ret[1], ret[0]

	err := SortByIndex(ids, ret, func(v *Integer) int {
		return v.v
	})
	if err != nil {
		t.Error(err)
	}
	for _, v := range ret {
		fmt.Println(v)
	}
	//log.Println(ret)
}
