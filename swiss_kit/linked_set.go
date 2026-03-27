package swiss_kit

import (
	"container/list"
	"iter"
)

type LinkedSet[K comparable] struct {
	m    map[K]*list.Element
	list *list.List
}

func NewLinkedSet[K comparable]() *LinkedSet[K] {
	return &LinkedSet[K]{
		m:    make(map[K]*list.Element),
		list: list.New(),
	}
}

func NewLinkedSetFrom[K comparable](items []K) *LinkedSet[K] {
	s := &LinkedSet[K]{
		m:    make(map[K]*list.Element, len(items)),
		list: list.New(),
	}
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			el := s.list.PushBack(item)
			s.m[item] = el
		}
	}
	return s
}

func (s *LinkedSet[K]) Add(key K) {
	if _, ok := s.m[key]; ok {
		return
	}
	el := s.list.PushBack(key)
	s.m[key] = el
}

func (s *LinkedSet[K]) Remove(key K) {
	if el, ok := s.m[key]; ok {
		s.list.Remove(el)
		delete(s.m, key)
	}
}

func (s *LinkedSet[K]) Has(key K) bool {
	_, ok := s.m[key]
	return ok
}

func (s *LinkedSet[K]) Len() int {
	return len(s.m)
}

func (s *LinkedSet[K]) ToSlice() []K {
	result := make([]K, 0, len(s.m))
	for el := s.list.Front(); el != nil; el = el.Next() {
		result = append(result, el.Value.(K))
	}
	return result
}

func (s *LinkedSet[K]) Range() iter.Seq[K] {
	return func(yield func(K) bool) {
		for el := s.list.Front(); el != nil; el = el.Next() {
			if !yield(el.Value.(K)) {
				return
			}
		}
	}
}
