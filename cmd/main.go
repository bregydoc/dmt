package main

import (
	"fmt"
	"time"

	"github.com/bregydoc/dmt"
	"github.com/bregydoc/dmt/channels/onesignal"
)

func main() {
	a := make(chan bool, 1)
	oneSignalChannel, err := onesignal.NewChannel(
		"09622104-abda-4456-bc7b-6d8495d8cd68",
		"NmNjMzQzMTAtMGYxYy00ZjBjLWExNmEtOWViNWQxNzEzNjMw",
	)
	if err != nil {
		panic(err)
	}

	_ = oneSignalChannel.Observe(func(works []dmt.Work) {
		fmt.Println(time.Now(), works)
	})

	if err = oneSignalChannel.AddTask(dmt.Task{
		Channel: onesignal.ChannelName,
		Type:    "push-notification",
		Params:  map[string]interface{}{
			"contents": map[string]string{
				"en": "Hello World",
				"es": "Hola Mundo",
			},
		},
	}); err != nil {
		panic(err)
	}
	<- a
}
