package swiss_kit

import (
	"errors"
	"fmt"
	"sort"
)

func SortByIndex[K comparable, V any](ids []K, list []V, getIdFunc func(v V) K) error {
	return sortByIndexCore(ids, false, list, getIdFunc)
}
func SortByIndexDesc[K comparable, V any](ids []K, list []V, getIdFunc func(v V) K) error {
	return sortByIndexCore(ids, true, list, getIdFunc)
}

func sortByIndexCore[K comparable, V any](ids []K, desc bool, list []V, getIdFunc func(v V) K) error {
	if len(list) != len(ids) {
		return errors.New("list must have the same length")
	}
	if len(list) == 0 {
		return nil
	}
	idMap := make(map[K]int, len(ids))
	for i, id := range ids {
		var v int
		if desc {
			v = len(list) - i - 1
		} else {
			v = i
		}
		idMap[id] = v
	}
	var err error
	acceptExistFunc := func(id K) {
		if err != nil {
			return
		}
		err = fmt.Errorf("%v not found", id)
	}
	sort.Slice(list, func(i, j int) bool {
		if err != nil {
			return false
		}
		iId := getIdFunc(list[i])
		jId := getIdFunc(list[j])
		iIndex, exist := idMap[iId]
		if !exist {
			acceptExistFunc(iId)
			return false
		}
		jIndex, exist := idMap[jId]
		if !exist {
			acceptExistFunc(jId)
			return false
		}
		return iIndex < jIndex
	})
	return err
}
