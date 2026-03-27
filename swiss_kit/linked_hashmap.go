package swiss_kit

import (
	"container/list"
	"iter"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LinkedHashMap[K comparable, V any] struct {
	m    map[K]*list.Element
	list *list.List
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		m:    make(map[K]*list.Element),
		list: list.New(),
	}
}

func (lm *LinkedHashMap[K, V]) Put(key K, value V) {
	if el, ok := lm.m[key]; ok {
		el.Value.(*entry[K, V]).value = value
		return
	}
	el := lm.list.PushBack(&entry[K, V]{key: key, value: value})
	lm.m[key] = el
}

func (lm *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	if el, ok := lm.m[key]; ok {
		return el.Value.(*entry[K, V]).value, true
	}
	var zero V
	return zero, false
}

func (lm *LinkedHashMap[K, V]) Delete(key K) {
	if el, ok := lm.m[key]; ok {
		lm.list.Remove(el)
		delete(lm.m, key)
	}
}

func (lm *LinkedHashMap[K, V]) Has(key K) bool {
	_, ok := lm.m[key]
	return ok
}

func (lm *LinkedHashMap[K, V]) Len() int {
	return len(lm.m)
}

func (lm *LinkedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, lm.Len())
	for el := lm.list.Front(); el != nil; el = el.Next() {
		keys = append(keys, el.Value.(*entry[K, V]).key)
	}
	return keys
}

func (lm *LinkedHashMap[K, V]) Values() []V {
	values := make([]V, 0, lm.Len())
	for el := lm.list.Front(); el != nil; el = el.Next() {
		values = append(values, el.Value.(*entry[K, V]).value)
	}
	return values
}

func (lm *LinkedHashMap[K, V]) Walk(fn func(K, V) bool) {
	for el := lm.list.Front(); el != nil; el = el.Next() {
		e := el.Value.(*entry[K, V])
		if !fn(e.key, e.value) {
			return
		}
	}
}

func (lm *LinkedHashMap[K, V]) Range() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for el := lm.list.Front(); el != nil; el = el.Next() {
			e := el.Value.(*entry[K, V])
			if !yield(e.key, e.value) {
				return
			}
		}
	}
}
