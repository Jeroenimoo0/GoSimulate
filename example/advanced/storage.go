package main

import (
	"sync"
)

type Storage struct {
	sL      sync.RWMutex
	storage []Resource

	cL          sync.Mutex
	addCallback map[Resource][]chan bool
}

func NewStorage(size int) *Storage {
	return &Storage{
		sync.RWMutex{},
		make([]Resource, size),
		sync.Mutex{},
		make(map[Resource][]chan bool),
	}
}

type TransferResult int

const (
	Success          TransferResult = iota
	NoSpace
	ResourceNotFound
)

func (s *Storage) Transfer(other *Storage, resource Resource) (bool, TransferResult) {
	s.sL.Lock()
	other.sL.Lock()
	defer s.sL.Unlock()
	defer other.sL.Unlock()

	slot := -1

	for i, res := range s.storage {
		if res == resource {
			slot = i
			break
		}
	}

	if slot < 0 {
		return false, ResourceNotFound
	}

	for i, res := range other.storage {
		if res == None {
			s.storage[slot] = None
			other.storage[i] = resource

			return true, Success
		}
	}

	return false, NoSpace
}

func (s *Storage) TransferOrWait(other *Storage, resource Resource) {
	_, result := s.Transfer(other, resource)

	if result == Success {
		return
	}

	callback := make(chan bool)
	// It is no guaranteed that when the resource was added we will be
	// able to get it, therefore keep on trying until we actually get a
	// successful transfer
	for {
		s.WaitFor(resource, callback)
		<-callback

		_, result = s.Transfer(other, resource)
		if result == Success {
			return
		}
	}
}

func (s *Storage) Remove(resource Resource) bool {
	s.sL.RLock()

	for i, res := range s.storage {
		if res == resource {
			s.sL.RUnlock()

			s.sL.Lock()
			s.storage[i] = None
			s.sL.Unlock()

			return true
		}
	}

	s.sL.RUnlock()
	return false
}

func (s *Storage) Add(resource Resource) bool {
	s.sL.RLock()

	for i, res := range s.storage {
		if res == None {
			s.sL.RUnlock()

			s.sL.Lock()
			s.storage[i] = resource
			s.sL.Unlock()

			s.cL.Lock()

			callbacks := s.addCallback[resource]
			if len(callbacks) > 0 {
				callback := callbacks[0]
				callbacks = callbacks[1:]
				s.addCallback[resource] = callbacks

				s.cL.Unlock()

				go func() {
					callback <- true
				}()
			} else {
				s.cL.Unlock()
			}

			return true
		}
	}

	s.sL.RUnlock()
	return false
}

func (s *Storage) WaitFor(resource Resource, callback chan bool) {
	s.cL.Lock()

	callbacks := s.addCallback[resource]

	callbacks = append(callbacks, callback)
	s.addCallback[resource] = callbacks

	s.cL.Unlock()
}

func (s *Storage) GetResources() []Resource {
	//Avoid slow defer unlock
	s.sL.RLock()

	cpy := make([]Resource, len(s.storage))
	copy(cpy, s.storage)

	s.sL.RUnlock()

	return cpy
}

func (s *Storage) Full() bool {
	s.sL.RLock()
	cpy := s.storage
	s.sL.RUnlock()

	for _, i := range cpy {
		if i == None {
			return false
		}
	}

	return true
}

func (s *Storage) full() bool {
	for _, i := range s.storage {
		if i == None {
			return false
		}
	}

	return true
}

func (s *Storage) Contains(resource Resource) bool {
	for _, i := range s.storage {
		if i == resource {
			return true
		}
	}

	return false
}

func Equal(first []Resource, second []Resource) bool {
	fLen := len(first)

	if fLen != len(second) {
		return false
	}

	for i := 0; i < fLen; i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}
