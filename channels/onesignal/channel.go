package onesignal

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bregydoc/dmt"
)

type Channel struct {
	AppID string
	APIKey string
	works []dmt.Work
	subscribers []func([]dmt.Work)
}

const ChannelName = "one-signal"

func (c *Channel) propagateWorksChangeState() {
	for _, s := range c.subscribers {
		s(c.works)
	}
}

func (c *Channel) onNewWork(index int) {
	c.propagateWorksChangeState()

	if err := c.works[index].ExecuteTask(); err != nil {
		fmt.Println("[ERROR]", err.Error())
		return
	}

	c.propagateWorksChangeState()

}

func (c *Channel) Name() dmt.ChannelName {
	return ChannelName
}

func (c *Channel) AddTask(task dmt.Task) error {
	if task.Channel != c.Name() {
		return errors.New("invalid task for this channel")
	}

	if c.works == nil {
		c.works = []dmt.Work{}
	}

	if task.Type == "push-notification" {
		d, err := json.Marshal(task.Params["contents"])
		if err != nil {
			return err
		}

		contents := map[Language]string{}
		if err = json.Unmarshal(d, &contents); err != nil {
			return err
		}

		push := &PushNotification{
			appID:    c.AppID,
			apiKey:   c.APIKey,
			done:     false,
			Contents: contents,
		}
		c.works = append(c.works, push)
		c.onNewWork(len(c.works) - 1)
		return nil
	} else if task.Type == "email" {

	}

	return errors.New("invalid task type")
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


