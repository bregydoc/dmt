package smtp

import (
	"errors"

	"github.com/bregydoc/dmt"
	"gopkg.in/gomail.v2"
)

type SendEmail struct {
	dialer      *gomail.Dialer
	done        bool
	From        Email   `json:"from"`
	To          []Email `json:"to"`
	ContentType string  `json:"content_type"`
	Subject     string  `json:"subject"`
	Body        []byte  `json:"body"`
}

func (s *SendEmail) Type() dmt.WorkType {
	return SendEmailTask
}

func (s *SendEmail) State() dmt.WorkState {
	if s.done {
		return dmt.WorkDone
	}

	return dmt.WorkPending
}

func (s *SendEmail) ExecuteTask() error {
	if s.dialer == nil {
		return errors.New("uninitialized work")
	}
	s.done = false
	m := gomail.NewMessage()

	name, email, err := extractNameAndEmailFromString(string(s.From))
	if err != nil {
		return err
	}

	if name != "" {
		m.SetAddressHeader("From", email, name)
	} else {
		m.SetHeader("From", email)
	}

	tos := make([]string, 0)
	for _, to := range s.To {
		tos = append(tos, string(to))
	}
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", s.Subject)
	m.SetBody(s.ContentType, string(s.Body))

	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}

	s.done = true
	return nil
}

func (s *SendEmail) IsDone() bool {
	return s.done
}
