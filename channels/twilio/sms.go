package twilio

import (
	"errors"

	"github.com/bregydoc/dmt"
	"github.com/sfreiberg/gotwilio"
	"github.com/sirupsen/logrus"
)

type SendSMS struct {
	done bool
	retryingCount uint8
	client *gotwilio.Twilio
	defaultFrom string
	From string `json:"from"`
	To []string `json:"to"`
	Content string `json:"content"`
}


func (s *SendSMS) Type() dmt.WorkType {
	return SendSMSTask
}

func (s *SendSMS) State() dmt.WorkState {
	if s.done {
		return dmt.WorkDone
	}

	if s.retryingCount > 0 {
		return dmt.WorkRetrying
	}

	return dmt.WorkPending
}

func (s *SendSMS) ExecuteTask() error {
	if s.done {
		return errors.New("your task is already done")
	}

	from := s.defaultFrom
	if s.From != "" {
		from = s.From
	}

	statues := make([]string, 0)
	for _, to := range s.To {
		res, exp, err := s.client.SendSMS(from, to, s.Content, "", "")
		if err != nil {
			return err
		}

		if exp != nil {
			if exp.Code == gotwilio.ErrorQueueAlreadyExists {
				s.done = false
				return nil
			}
			logrus.Info("exception status: ", exp.Status)
			s.retryingCount += 1
			break
		}

		price := "0.0"
		if res.Price != nil {
			price = *res.Price
		}

		statues = append(statues, res.Status)

		logrus.Infof("sms cost: %s, sms status: %s", price, res.Status)
	}

	s.done = true
	for _, status := range statues {
		if status != "queued" {
			s.done = false
		}
	}

	return nil
}

func (s *SendSMS) IsDone() bool {
 	return s.done
}
