package dmt

import "time"

type WorkType string

type Work interface {
	Type() WorkType
	ExecuteTask(TaskParams) error
	IsDone() bool
}

type TaskParams interface {
	Get(string) (interface{}, error)
}

type WorkerLife struct {
	OneTime bool
	PulseEach time.Duration
	StartAt time.Time
	EndAt time.Time
	Done bool
}

type Worker struct {
	Work Work
	life WorkerLife
}

