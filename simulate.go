package simulate

import (
	"sort"
	"time"
	"sync"
)

type Actor interface {
	Run()
}

type waitingActor struct {
	actor Actor
	date time.Duration
}

var Timeout time.Duration = time.Millisecond * 1

type Simulation struct {
	l sync.Mutex
	waiting []waitingActor
	lastTime time.Duration

	running bool
	finished bool
	asleep bool
	awake chan bool

	TotalTime time.Duration
	TotalTimeAsleep time.Duration
}

func NewSimulation() *Simulation {
	return &Simulation{
		sync.Mutex{},
		[]waitingActor{},
		0,
		false,
		false,
		false,
		make(chan bool),
		0,
		0,
	}
}

func (s *Simulation) Run(period time.Duration) {
	if s.running {
		panic("Already running")
	}

	s.running = true

	simulate:
	for {
		s.l.Lock()

		length := len(s.waiting)
		if length == 0 {
			s.asleep = true
			s.l.Unlock()

			now := time.Now()
			select {
			case <-s.awake:
				s.TotalTimeAsleep += time.Now().Sub(now)
				continue
			case <-time.After(Timeout):
				break simulate;
			}
		}

		entry := s.waiting[0]
		s.waiting = s.waiting[1:]

		wait := entry.date - s.lastTime
		s.TotalTime += wait

		if s.TotalTime > period {
			s.l.Unlock()
			break
		}

		s.lastTime = entry.date

		go entry.actor.Run()

		s.l.Unlock()
	}

	s.l.Lock()
	s.running = false
	s.finished = true
	s.l.Unlock()
}

func (s *Simulation) Add(actor Actor, delay time.Duration) {
	s.l.Lock()

	if s.finished {
		s.l.Unlock()
		return
	}

	//TODO In realtime use current time instead of lastTime
	s.waiting = insertSorted(s.waiting, waitingActor{
		actor,
		s.lastTime + delay,
	})

	if s.asleep {
		s.asleep = false
		s.l.Unlock()

		s.awake <- true
		return
	}

	s.l.Unlock()
}

func insertSorted (slice []waitingActor, entry waitingActor) []waitingActor {
	length:=len(slice)
	if length == 0 {
		return []waitingActor{entry}
	}

	i := sort.Search(length, func(i int) bool {
		return slice[i].date >= entry.date
	})

	if i == 0 { // Not found, will be first value
		return append([]waitingActor{entry}, slice...)
	} else if i == length { // Not found, will be last value
		return append(slice, entry)
	}

	// Insert into array
	return append(slice[:i], append([]waitingActor{entry}, slice[i:]...)...)
}