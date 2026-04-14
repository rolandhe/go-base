package swiss_kit

import (
	"container/heap"
)

type LimitedPriorityQueue[T any] struct {
	st    *storeCore
	limit int
}

func NewLimitedPriorityQueue[T any](limit int, top bool, cmp func(a, b T) bool) *LimitedPriorityQueue[T] {
	return &LimitedPriorityQueue[T]{
		st: &storeCore{
			cmp: func(a any, b any) bool {
				ret := cmp(a.(T), b.(T))
				if top {
					return ret
				}
				return !ret
			},
		},
		limit: limit,
	}
}

func (lpq *LimitedPriorityQueue[T]) Len() int { return lpq.st.Len() }

func (lpq *LimitedPriorityQueue[T]) Push(v T) {
	if lpq.st.Len() < lpq.limit {
		heap.Push(lpq.st, v)
		return
	}
	top := lpq.st.items[0]
	if lpq.st.cmp(top, v) {
		lpq.st.items[0] = v
		heap.Fix(lpq.st, 0)
	}
}

func (lpq *LimitedPriorityQueue[T]) Pop() T {
	v := heap.Pop(lpq.st)
	return v.(T)
}

func (lpq *LimitedPriorityQueue[T]) ToSlice() []T {
	if lpq.st.Len() == 0 {
		return nil
	}
	ret := make([]T, lpq.st.Len())
	indexFunc := nextIndex(lpq.Len())

	for lpq.Len() > 0 {
		v := lpq.Pop()
		i := indexFunc()
		ret[i] = v
	}
	return ret
}

func nextIndex(originalLen int) func() int {
	index := -1
	return func() int {
		if index == -1 {
			index = originalLen - 1
			return index
		}
		index--
		return index
	}
}

type storeCore struct {
	items []any
	cmp   func(any, any) bool
}

func (s *storeCore) Len() int {
	return len(s.items)
}
func (s *storeCore) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}
func (s *storeCore) Less(i, j int) bool {
	return s.cmp(s.items[i], s.items[j])
}
func (s *storeCore) Push(v any) {
	s.items = append(s.items, v)
}
func (s *storeCore) Pop() any {
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v
}
