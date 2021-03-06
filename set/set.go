package set

import "sync"

type inter interface {
}

type Set struct {
	m map[inter]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[inter]bool{},
	}
}

func (s *Set) Add(item inter) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) IsExist(item inter) bool {
	if s.m[item] {
		return true
	}

	return false
}
