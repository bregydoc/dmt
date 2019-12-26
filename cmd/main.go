package main

import (
	"fmt"
	"time"

	"github.com/bregydoc/dmt"
	"github.com/bregydoc/dmt/channels/smtp"
)

func main() {
	a := make(chan bool, 1)
	// oneSignalChannel, err := onesignal.NewChannel(
	// 	"09622104-abda-4456-bc7b-6d8495d8cd68",
	// 	"NmNjMzQzMTAtMGYxYy00ZjBjLWExNmEtOWViNWQxNzEzNjMw",
	// )
	// if err != nil {
	// 	panic(err)
	// }
	//
	// _ = oneSignalChannel.Observe(func(works []dmt.Work) {
	// 	fmt.Println(time.Now())
	// 	for i, w := range works {
	// 		fmt.Println(i, ":", w.IsDone())
	// 	}
	// })
	//
	// if err = oneSignalChannel.AddTask(dmt.Task{
	// 	Channel: onesignal.ChannelName,
	// 	Type:    "push-notification",
	// 	Params:  map[string]interface{}{
	// 		"contents": map[string]string{
	// 			"en": "Hello World",
	// 			"es": "Hola Mundo",
	// 		},
	// 	},
	// }); err != nil {
	// 	panic(err)
	// }

	smtpChannel, err := smtp.NewChannel(
		"in-v3.mailjet.com",
		587,
		"34fc23c6c6af1afa40fa61c9a4b2447f",
		"824d8412b0b4b41981c53009e2e03adf",
	)
	if err != nil {
		panic(err)
	}

	_ = smtpChannel.Observe(func(works []dmt.Work) {
		fmt.Println(time.Now())
		for i, w := range works {
			fmt.Println(i, ":", w.IsDone())
		}
	})

	if err = smtpChannel.AddTask(dmt.Task{
		Channel: smtp.ChannelName,
		Type:    "send-email",
		Params:  map[string]interface{}{
			"from": "dev@limacitypass.com",
			"to": []string{
				"bregy.malpartida@utec.edu.pe",
				"bregymr@gmail.com",
			},
			"subject": "DMT Test",
			"content_type": "text/html",
			"body": []byte("<h2> Hello World </h2>"),
		},
	}); err != nil {
		panic(err)
	}


	<- a
}
