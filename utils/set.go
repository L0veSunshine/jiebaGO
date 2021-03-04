package utils

import "sync"

type SetItem map[string]struct{}

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

func (s *Set) Add(str string) {
	if !s.Has(str) {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.SetItem[str] = struct{}{}
	}
}

func (s *Set) Has(str string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.SetItem[str]
	return ok
}

func (s *Set) ToArray() (res []string) {
	res = make([]string, 0, s.Len())
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

func (s *Set) Remove(str string) bool {
	if !s.Has(str) {
		return false
	} else {
		s.lock.Lock()
		defer s.lock.Unlock()
		delete(s.SetItem, str)
	}
	return true
}
