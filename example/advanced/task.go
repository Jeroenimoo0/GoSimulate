package main

import (
	"time"
)

type Task interface {
	GetTimeToComplete(worker *Worker) time.Duration
	Start(worker *Worker) bool
	Complete(worker *Worker) bool
	GetName() string
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
	t.from.TransferOrWait(&worker.Storage, t.resource)

	return true
}

func (t *TaskMoving) Complete(worker *Worker) bool {
	worker.TransferOrWait(t.to, t.resource)

	return true
}

func (t *TaskMoving) GetName() string {
	return t.name
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