package twilio

import (
	"github.com/bregydoc/dmt"
	"github.com/sfreiberg/gotwilio"
)

func NewChannel(accountSID, authToken string) (*Channel, error) {
	return &Channel{
		client:      gotwilio.NewTwilioClient(accountSID, authToken),
		works:       []dmt.Work{},
		subscribers: []func([]dmt.Work){},
	}, nil
}