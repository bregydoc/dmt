package main

import (
	"errors"

	"github.com/bregydoc/dmt"
	"github.com/bregydoc/dmt/channels/onesignal"
	"github.com/bregydoc/dmt/channels/smtp"
	"github.com/bregydoc/dmt/channels/twilio"
)

func inflateChannelFromConfig(config dmt.ChannelConfig) (dmt.Channel, error) {
	switch config.Name {
	case onesignal.ChannelName:
		return onesignal.NewChannelFromMap(config.Config)
	case smtp.ChannelName:
		return smtp.NewChannelFromMap(config.Config)
	case twilio.ChannelName:
		return twilio.NewChannelFromMap(config.Config)
	default:
		return nil, errors.New("invalid channel name")
	}
}

func main() {
	engine, err := dmt.NewEngine("./config.yaml")
	if err != nil {
		panic(err)
	}

	for _, ch := range engine.Config.Channels {
		channel, err := inflateChannelFromConfig(ch)
		if err != nil {
			panic(err)
		}

		if err = engine.AddNewChannel(channel); err != nil {
			panic(err)
		}
	}

	if err = engine.Run(); err != nil {
		panic(err)
	}

}
