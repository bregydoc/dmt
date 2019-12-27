package twilio

import (
	"github.com/bregydoc/dmt"
	"github.com/sfreiberg/gotwilio"
)

func NewChannel(accountSID, authToken string, defaultFromNumber ...string) (*Channel, error) {
	fromNumber := ""
	if len(defaultFromNumber) != 0 {
		fromNumber = defaultFromNumber[0]
	}

	return &Channel{
		client:            gotwilio.NewTwilioClient(accountSID, authToken),
		works:             []dmt.Work{},
		subscribers:       []func([]dmt.Work){},
		defaultFromNumber: fromNumber,
	}, nil
}
