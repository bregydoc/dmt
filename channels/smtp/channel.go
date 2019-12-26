package smtp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bregydoc/dmt"
	"gopkg.in/gomail.v2"
)

type Channel struct {
	dialer *gomail.Dialer
	works       []dmt.Work
	subscribers []func([]dmt.Work)
}

const ChannelName = "smtp"


func (c *Channel) propagateWorksChangeState() {
	for _, s := range c.subscribers {
		s(c.works)
	}
}

func (c *Channel) executeTask(index int) {
	if err := c.works[index].ExecuteTask(); err != nil {
		fmt.Println("[ERROR]", err.Error())
	}
	c.propagateWorksChangeState()
}

func (c *Channel) onNewWork(index int) {
	c.propagateWorksChangeState()
	go c.executeTask(index)
}


func (c *Channel) Name() dmt.ChannelName {
	return ChannelName
}

func (c *Channel) AddTask(task dmt.Task) error {
	if task.Channel != c.Name() {
		return errors.New("invalid task for this channel")
	}

	if task.Type == "send-email" {
		d, err := json.Marshal(task.Params)
		if err != nil {
			return err
		}
		sendEmail := &SendEmail{}
		if err = json.Unmarshal(d, sendEmail); err != nil {
			return err
		}

		sendEmail.dialer = c.dialer

		if c.works == nil {
			c.works = []dmt.Work{}
		}

		c.works = append(c.works, sendEmail)
		c.onNewWork(len(c.works) - 1)
	} else {
		return errors.New("invalid task type")
	}
	return nil
}

func (c *Channel) FlushAll() error {
	c.works = []dmt.Work{}
	return nil
}

func (c *Channel) Observe(hook func([]dmt.Work)) error {
	if c.subscribers == nil {
		c.subscribers = []func([]dmt.Work){}
	}

	c.subscribers = append(c.subscribers, hook)

	return nil
}



