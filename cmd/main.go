package main

import (
	"time"

	"github.com/bregydoc/dmt"
	"github.com/bregydoc/dmt/channels/onesignal"
)

func main() {
	os, err := onesignal.NewFabric(
		"",
		"",
	)
	if err != nil {
		panic(err)
	}

	e := dmt.Engine{
		StartedAt: time.Now(),
	}

	w, err := os.NewPushNotificationWork()
	if err != nil {
		panic(err)
	}

	err = w.ExecuteTask(onesignal.PushNotificationForAll{
		Contents: map[onesignal.Language]string{
			"es": "",
		},
	})

	if err = e.AddNewWork(w); err != nil {
		panic(err)
	}


}
