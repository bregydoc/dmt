package dmt

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type Engine struct {
	StartedAt time.Time
	channels  []Channel
	Config    *Config
	rest      *API
}

func (e *Engine) registerNewChannel(channel Channel) error {
	if e.channels == nil {
		e.channels = []Channel{}
	}

	log.Infof("channel '%s'registered", channel.Name())

	if err := channel.Observe(e.globalObserver); err != nil {
		return err
	}

	e.channels = append(e.channels, channel)

	return nil
}

func (e *Engine) registerNewTask(task Task) error {
	for _, ch := range e.channels {
		if ch.Name() == task.Channel {
			if err := ch.AddTask(task); err != nil {
				return err
			}
		}
	}

	return fmt.Errorf("channel '%s' not register on dmtjs", task.Channel)
}
