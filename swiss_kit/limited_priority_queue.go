package swiss_kit

import (
	"container/heap"
	"slices"
)

type LimitedPriorityQueue[T any] struct {
	st    *storeCore[T]
	limit int
}

func NewLimitedPriorityQueue[T any](limit int, top bool, cmp func(a, b T) bool) *LimitedPriorityQueue[T] {
	return &LimitedPriorityQueue[T]{
		st: &storeCore[T]{
			cmp: func(a T, b T) bool {
				ret := cmp(a, b)
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

func (lpq *LimitedPriorityQueue[T]) CloneToSlice() []T {
	if lpq.st.Len() == 0 {
		return nil
	}
	clonePq := &LimitedPriorityQueue[T]{
		st: &storeCore[T]{
			cmp:   lpq.st.cmp,
			items: slices.Clone(lpq.st.items),
		},
	}
	return clonePq.ToSlice()
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

type storeCore[T any] struct {
	items []T
	cmp   func(T, T) bool
}

func (s *storeCore[T]) Len() int {
	return len(s.items)
}
func (s *storeCore[T]) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}
func (s *storeCore[T]) Less(i, j int) bool {
	return s.cmp(s.items[i], s.items[j])
}
func (s *storeCore[T]) Push(v any) {
	s.items = append(s.items, v.(T))
}
func (s *storeCore[T]) Pop() any {
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v
}
