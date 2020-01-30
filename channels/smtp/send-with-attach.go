package smtp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bregydoc/dmt"
	"gopkg.in/gomail.v2"
)

type SendEmailWithAttach struct {
	dialer      *gomail.Dialer
	done        bool
	From        Email             `json:"from"`
	To          []Email           `json:"to"`
	ContentType string            `json:"content_type"`
	Subject     string            `json:"subject"`
	Body        []byte            `json:"body"`
	Attachments map[string][]byte `json:"attachments"`
}

const SendEmailWithAttachTask = "send-email-with-attachments"

func (s *SendEmailWithAttach) Type() dmt.WorkType {
	return SendEmailWithAttachTask
}

func (s *SendEmailWithAttach) State() dmt.WorkState {
	if s.done {
		return dmt.WorkDone
	}

	return dmt.WorkPending
}

func (s *SendEmailWithAttach) ExecuteTask() error {
	if s.dialer == nil {
		return errors.New("uninitialized work")
	}
	s.done = false
	m := gomail.NewMessage()

	name, email, err := extractNameAndEmailFromString(string(s.From))
	if err != nil {
		return err
	}

	fmt.Println(name, email)

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

	for name, data := range s.Attachments {
		f, err := ioutil.TempFile(os.TempDir(), "dmtjs")
		if err != nil {
			return err
		}

		if _, err := f.Write(data); err != nil {
			return err
		}

		fmt.Println(f.Name())
		m.Attach(f.Name(), gomail.Rename(name))

		// if err = f.Close(); err != nil {
		// 	return err
		// }
	}
	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}

	s.done = true
	return nil
}

func (s *SendEmailWithAttach) IsDone() bool {
	return s.done
}
