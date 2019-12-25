package dmt

import (
	"errors"
	"time"
)

type Engine struct {
	StartedAt time.Time
	works []Work
}

func (e *Engine) AddNewWork(w Work) error {
	if w.IsDone() {
		return errors.New("work already done")
	}
	if e.works == nil {
		e.works = []Work{}
	}
	e.works = append(e.works, w)
	return nil
}
