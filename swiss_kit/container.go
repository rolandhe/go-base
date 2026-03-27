package swiss_kit

import (
	"slices"
	"sync"
)

func NewSafeMap[K comparable, T any]() *SafeMap[K, T] {
	return &SafeMap[K, T]{
		m: make(map[K]T),
	}
}

type SafeMap[K comparable, T any] struct {
	sync.RWMutex
	m map[K]T
}

func (s *SafeMap[K, T]) Set(key K, value T) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *SafeMap[K, T]) Get(key K) (bool, T) {
	s.RLock()
	defer s.RUnlock()
	v, ok := s.m[key]
	return ok, v
}

func (s *SafeMap[K, T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[K]T)
}

func (s *SafeMap[K, T]) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.m)
}

func NewSafeSet[T comparable]() *SafeSet[T] {
	return &SafeSet[T]{
		guard: make(map[T]struct{}),
	}
}

type SafeSet[T comparable] struct {
	sync.RWMutex
	guard map[T]struct{}
	l     []T
}

func (s *SafeSet[T]) Set(key T) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.guard[key]; ok {
		return
	}
	s.guard[key] = struct{}{}
	s.l = append(s.l, key)
}

func (s *SafeSet[T]) Copy() []T {
	s.RLock()
	defer s.RUnlock()
	return slices.Clone(s.l)
}

func (s *SafeSet[T]) Len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.l)
}

func (s *SafeSet[T]) Exists(key T) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.guard[key]
	return ok
}

func (s *SafeSet[T]) WalkRead(wrk func(e T) bool) {
	s.RLock()
	defer s.RUnlock()
	for _, v := range s.l {
		if wrk(v) {
			break
		}
	}
}
