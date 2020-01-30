package onesignal

import (
	"errors"

	"github.com/bregydoc/dmt"
)

func NewChannel(appID, apiKey string) (*Channel, error) {
	return &Channel{
		AppID:       appID,
		APIKey:      apiKey,
		works:       []dmt.Work{},
		subscribers: []func([]dmt.Work){},
	}, nil
}

func NewChannelFromMap(params map[string]interface{}) (*Channel, error) {
	appID, ok := params["appID"].(string)
	if !ok {
		return nil, errors.New("appID param not found")
	}

	apiKey, ok := params["apiKey"].(string)
	if !ok {
		return nil, errors.New("apiKey param not found")
	}

	return NewChannel(appID, apiKey)
}
