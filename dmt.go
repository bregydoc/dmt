package dmt

import "time"

type WorkType string

type WorkState int

const (
	WorkCreated WorkState = iota
	WorkSent
	WorkPending
	WorkRetrying
	WorkDone
)

type Work interface {
	Type() WorkType
	State() WorkState
	ExecuteTask() error
	IsDone() bool
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

