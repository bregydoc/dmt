package onesignal

import "github.com/bregydoc/dmt"

func NewChannel(appID, aPIKey string) (*Channel, error) {
	return &Channel{
		AppID:       appID,
		APIKey:      aPIKey,
		works:       []dmt.Work{},
		subscribers: []func([]dmt.Work){},
	}, nil
}
