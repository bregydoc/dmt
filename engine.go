package dmt

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type Engine struct {
	StartedAt time.Time
	Channels []Channel
}

func (e *Engine) registerNewChannel(channel Channel) error {
	if e.Channels == nil {
		e.Channels = []Channel{}
	}

	log.Infof("channel '%s'registered", channel.Name())

	if err := channel.Observe(e.globalObserver); err != nil {
		return err
	}

	e.Channels = append(e.Channels, channel)

	return nil
}

func (e *Engine) registerNewTask(task Task) error {
	for _, ch := range e.Channels {
		if ch.Name() == task.Channel {
			if err := ch.AddTask(task); err != nil {
				return err
			}
		}
	}
	return fmt.Errorf("channel '%s' not register on dmt", task.Channel)
}

