package swiss_kit

import (
	"slices"
	"sync"
)

func NewSafeList[T any]() *SafeList[T] {
	return &SafeList[T]{}
}

type SafeList[T any] struct {
	sync.RWMutex
	items []T
}

func (s *SafeList[T]) Append(item T) {
	s.Lock()
	defer s.Unlock()
	s.items = append(s.items, item)
}

func (s *SafeList[T]) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.items)
}

func (s *SafeList[T]) Items() []T {
	s.RLock()
	defer s.RUnlock()
	return slices.Clone(s.items)
}

func (s *SafeList[T]) Walk(fn func(T) bool) {
	s.RLock()
	defer s.RUnlock()
	for _, v := range s.items {
		if fn(v) {
			break
		}
	}
}
