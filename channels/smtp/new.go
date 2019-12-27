package smtp

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func NewChannel(host string, port int, username, password string) (*Channel, error) {
	ch := &Channel{
		dialer: gomail.NewDialer(host, port, username, password),
	}

	ch.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return ch, nil
}
