package smtp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bregydoc/dmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type Channel struct {
	dialer      *gomail.Dialer
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

	log.Info(task)

	if task.Type == SendEmailTask {
		d, err := json.Marshal(task.Params)
		if err != nil {
			return err
		}

		log.Info(string(d))
		sendEmail := &SendEmail{}

		if err = json.Unmarshal(d, sendEmail); err != nil {
			log.Error(err)
			return err
		}

		sendEmail.dialer = c.dialer

		if sendEmail.ContentType == "" {
			sendEmail.ContentType = "text/plain"
		}

		spew.Dump(sendEmail)

		if c.works == nil {
			c.works = []dmt.Work{}
		}

		c.works = append(c.works, sendEmail)
		c.onNewWork(len(c.works) - 1)
	} else if task.Type == SendEmailWithAttachTask {
		d, err := json.Marshal(task.Params)
		if err != nil {
			return err
		}

		sendEmail := &SendEmailWithAttach{}
		if err = json.Unmarshal(d, sendEmail); err != nil {
			return err
		}

		sendEmail.dialer = c.dialer

		if sendEmail.ContentType == "" {
			sendEmail.ContentType = "text/plain"
		}

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
