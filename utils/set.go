package utils

import "sync"

type SetItem map[interface{}]struct{}

type Set struct {
	SetItem
	lock *sync.RWMutex
}

func NewSet() *Set {
	s := make(SetItem, 5)
	return &Set{
		SetItem: s,
		lock:    new(sync.RWMutex),
	}
}

func (s *Set) Add(key interface{}) {
	if !s.Has(key) {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.SetItem[key] = struct{}{}
	}
}

func (s *Set) Has(key interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.SetItem[key]
	return ok
}

func (s *Set) ToArray() (res []interface{}) {
	res = make([]interface{}, 0, s.Len())
	for key := range s.SetItem {
		res = append(res, key)
	}
	return res
}

func (s *Set) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.SetItem)
}

func (s *Set) Remove(key interface{}) bool {
	if !s.Has(key) {
		return false
	} else {
		s.lock.Lock()
		defer s.lock.Unlock()
		delete(s.SetItem, key)
	}
	return true
}
