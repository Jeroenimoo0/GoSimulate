package simulate

import (
	"fmt"
	"time"
)

type Task interface {
	GetTimeToComplete(worker *Worker) time.Duration
	Start(worker *Worker) bool
	Complete(worker *Worker) bool
	GetName() string
	CanBeCompleted(worker *Worker) bool
}

type TaskMoving struct {
	name string
	timeToComplete time.Duration
	resource Resource
	from *Storage
	to *Storage
}

func NewTaskMoving(name string, duration time.Duration, resource Resource, from *Storage, to *Storage) *TaskMoving {
	return &TaskMoving{
		name,
		duration,
		resource,
		from,
		to,
	}
}

func (t *TaskMoving) GetTimeToComplete(worker *Worker) time.Duration {
	return t.timeToComplete
}

func (t *TaskMoving) Start(worker *Worker) bool {
	_, result := t.from.Transfer(&worker.Storage, t.resource)

	switch result {
	case NoSpace:
		fmt.Println("Could not move, NoSpace")
		return false
	case ResourceNotFound:
		fmt.Println("Could not move, ResourceNotFound")
		return false
	}

	return true
}

func (t *TaskMoving) Complete(worker *Worker) bool {
	_, result := worker.Transfer(t.to, t.resource)

	switch result {
	case NoSpace:
		fmt.Println("Could not move, NoSpace")
		return false
	case ResourceNotFound:
		fmt.Println("Could not move, ResourceNotFound")
		return false
	}

	return true
}

func (t *TaskMoving) GetName() string {
	return t.name
}

func (t *TaskMoving) CanBeCompleted(worker *Worker) bool {
	return t.from.Contains(t.resource) && !t.to.Full() && !worker.Full()
}

type TaskResource struct {
	name string
	storage *Storage
	timeToComplete time.Duration
	input []Resource
	output []Resource
}

func NewTaskResource(name string, storage *Storage, duration time.Duration, input []Resource, output []Resource) *TaskResource {
	return &TaskResource{
		name,
		storage,
		duration,
		input,
		output,
	}
}

func (t *TaskResource) GetTimeToComplete(worker *Worker) time.Duration {
	return t.timeToComplete
}

func (t *TaskResource) Start(worker *Worker) bool {
	success := true

	for _, resource := range t.input {
		if !t.storage.Remove(resource) {
			success = false
		}
	}

	return success
}

func (t *TaskResource) Complete(worker *Worker) bool {
	success := true

	for _, resource := range t.output {
		if !t.storage.Add(resource) {
			success = false
		}
	}

	return success
}

func (t *TaskResource) GetName() string {
	return t.name
}


func (t *TaskResource) CanBeCompleted(worker *Worker) bool {
	inputLoop:
	for _, inputRes := range t.input {
		for _, res := range t.storage.GetResources() {
			if res == inputRes {
				continue inputLoop
			}
		}

		return false
	}

	return true
}
