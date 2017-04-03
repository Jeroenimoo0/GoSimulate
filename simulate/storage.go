package simulate

import "sync"

type Storage struct {
	sL sync.RWMutex
	storage []Resource
}

func NewStorage(size int) *Storage {
	return &Storage{
		sync.RWMutex{},
		make([]Resource, size),
	}
}

type TransferResult int
const (
	Success TransferResult = iota
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

			return true
		}
	}

	s.sL.RUnlock()
	return false
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
