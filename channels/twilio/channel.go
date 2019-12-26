package twilio

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bregydoc/dmt"
	"github.com/sfreiberg/gotwilio"

)

type Channel struct {
	client *gotwilio.Twilio
	works       []dmt.Work
	subscribers []func([]dmt.Work)
	defaultFromNumber string
}

const ChannelName = "twilio"

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

	if task.Type == SendSMSTask {
		sendSMS := &SendSMS{}

		data, err := json.Marshal(task.Params)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(data, sendSMS); err != nil {
			return err
		}

		sendSMS.defaultFrom = c.defaultFromNumber
		sendSMS.client = c.client

		if c.works == nil {
			c.works = []dmt.Work{}
		}

		c.works = append(c.works, sendSMS)
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


