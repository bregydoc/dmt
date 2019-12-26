package main

import (
	"fmt"
	"io/ioutil"
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
		"2d0ec0212c07016a6ba717caa212c57f",
		"eaa24c87ca5bf24d1debd0ae7c648c63",
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

	invoiceData, err := ioutil.ReadFile("/Users/bregy/Documents/COMPROBANTE.pdf")
	if err != nil {
		panic(err)
	}

	if err = smtpChannel.AddTask(dmt.Task{
		Channel: smtp.ChannelName,
		Type:    smtp.SendEmailWithAttachTask,
		Params:  map[string]interface{}{
			"from": "<Bregy Malpartida>bregy.malpartida@utec.edu.pe",
			"to": []string{
				"bregy.malpartida@utec.edu.pe",
				"bregymr@gmail.com",
			},
			"subject": "DMT Test with attach",
			"content_type": "text/html",
			"body": []byte("<h2> Hello World with attachment</h2>"),
			"attachments": map[string][]byte{
				"invoice.pdf": invoiceData,
			},
		},
	}); err != nil {
		panic(err)
	}

	<- a
}
