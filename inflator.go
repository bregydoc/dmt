package dmt

// TODO:
// 	- Think in a new way to struct this code

// import (
// 	"errors"
//
// 	"github.com/bregydoc/dmt/channels/onesignal"
// 	"github.com/bregydoc/dmt/channels/smtp"
// 	"github.com/bregydoc/dmt/channels/twilio"
// )
//
// func inflateChannelFromConfig(config ChannelConfig) (Channel, error) {
// 	switch config.Name {
// 	case onesignal.ChannelName:
// 		return onesignal.NewChannelFromMap(config.Config)
// 	case smtp.ChannelName:
// 		return smtp.NewChannelFromMap(config.Config)
// 	case twilio.ChannelName:
// 		return twilio.NewChannelFromMap(config.Config)
// 	default:
// 		return nil, errors.New("invalid channel name")
// 	}
// }
