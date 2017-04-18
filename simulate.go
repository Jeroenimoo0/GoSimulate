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

type Simulation interface {
	Run()
	Stop()

	Add(actor Actor, delay time.Duration)

	GetRunTime() time.Duration
}

type SimulationInstant struct {
	l sync.Mutex
	waiting []waitingActor
	lastTime time.Duration

	running bool
	finished bool
	asleep bool
	awake chan bool

	period time.Duration

	totalTime time.Duration
	totalTimeAsleep time.Duration
}

func NewSimulationInstant(period time.Duration) *SimulationInstant {
	return &SimulationInstant{
		sync.Mutex{},
		[]waitingActor{},
		0,
		false,
		false,
		false,
		make(chan bool),
		period,
		0,
		0,
	}
}

func (s *SimulationInstant) Run() {
	if s.running {
		panic("Already running")
	}

	if s.finished {
		panic("Already finished")
	}

	s.running = true

	loop:
	for {
		s.l.Lock()

		length := len(s.waiting)
		if length == 0 {
			s.asleep = true
			s.l.Unlock()

			now := time.Now()
			select {
			case <-s.awake:
				s.totalTimeAsleep += time.Now().Sub(now)
				continue
			case <-time.After(Timeout):
				break loop
			}
		}

		entry := s.waiting[0]
		s.waiting = s.waiting[1:]

		wait := entry.date - s.lastTime
		s.totalTime += wait

		if s.totalTime > s.period {
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

func (s *SimulationInstant) Stop() {
	panic("Not supported")
}

func (s *SimulationInstant) Add(actor Actor, delay time.Duration) {
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

func (s *SimulationInstant) GetRunTime() time.Duration {
	s.l.Lock()
	timeRunning := s.totalTime
	s.l.Unlock()

	return timeRunning
}

type SimulationRealtime struct {
	l sync.Mutex
	initActors []waitingActor

	running bool
	finished bool

	TimeMultiplier float32

	startTime time.Time
}

func NewSimulationRealtime() *SimulationRealtime {
	return &SimulationRealtime{
		sync.Mutex{},
		[]waitingActor{},
		false,
		false,
		1,
		time.Now(),
	}
}

func (s *SimulationRealtime) Run() {
	if s.running {
		panic("Already running")
	}

	if s.finished {
		panic("Already finished")
	}

	s.running = true
	s.startTime = time.Now()

	for _, waiting := range s.initActors {
		go waiting.actor.Run()
	}
}

func (s *SimulationRealtime) Stop() {
	s.l.Lock()

	if !s.running {
		panic("Not running")
	}

	s.running = false
	s.finished = true
	s.l.Unlock()
}

func (s *SimulationRealtime) Add(actor Actor, delay time.Duration) {
	if s.running {
		go func() {
			time.Sleep(time.Duration(float32(delay) * s.TimeMultiplier))
			actor.Run()
		}()
	} else if !s.finished {
		s.initActors = append(s.initActors, waitingActor{
			actor,
			delay,
		})
	}
}

func (s *SimulationRealtime) GetRunTime() time.Duration {
	return time.Now().Sub(s.startTime)
}