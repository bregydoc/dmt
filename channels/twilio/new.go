package twilio

import (
	"errors"

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

func NewChannelFromMap(params map[string]interface{}) (*Channel, error) {
	accountSID, ok := params["accountSID"].(string)
	if !ok {
		return nil, errors.New("accountSID param not found")
	}

	authToken, ok := params["authToken"].(string)
	if !ok {
		return nil, errors.New("authToken param not found")
	}

	defaultFromNumber, ok := params["defaultFromNumber"].(string)
	if !ok {
		defaultFromNumber = ""
	}

	if defaultFromNumber == "" {
		return NewChannel(accountSID, authToken)
	}

	return NewChannel(accountSID, authToken, defaultFromNumber)
}
