package smtp

import (
	"crypto/tls"
	"errors"

	"gopkg.in/gomail.v2"
)

func NewChannel(host string, port int, username, password string) (*Channel, error) {
	ch := &Channel{
		dialer: gomail.NewDialer(host, port, username, password),
	}

	ch.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return ch, nil
}


func NewChannelFromMap(params map[string]interface{}) (*Channel, error) {
	host, ok := params["host"].(string)
	if !ok {
		return nil, errors.New("host param not found")
	}

	port, ok := params["port"].(int)
	if !ok {
		return nil, errors.New("port param not found")
	}

	username, ok := params["username"].(string)
	if !ok {
		return nil, errors.New("username param not found")
	}

	password, ok := params["password"].(string)
	if !ok {
		return nil, errors.New("password param not found")
	}

	return NewChannel(host, port, username, password)
}