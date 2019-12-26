package twilio

import (
	"github.com/bregydoc/dmt"
	"github.com/sfreiberg/gotwilio"

)

type Channel struct {
	client *gotwilio.Twilio
	works       []dmt.Work
	subscribers []func([]dmt.Work)
}

