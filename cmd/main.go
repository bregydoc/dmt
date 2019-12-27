package main

import "github.com/bregydoc/dmt"

func main() {
	a := make(chan bool, 1)
	// oneSignalChannel, err := onesignal.NewChannel(

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

	// smtpChannel, err := smtp.NewChannel(
	// 	"in-v3.mailjet.com",
	// 	587,

	// )
	// if err != nil {
	// 	panic(err)
	// }
	//
	// _ = smtpChannel.Observe(func(works []dmt.Work) {
	// 	fmt.Println(time.Now())
	// 	for i, w := range works {
	// 		fmt.Println(i, ":", w.IsDone())
	// 	}
	// })
	//
	// invoiceData, err := ioutil.ReadFile("/Users/bregy/Documents/COMPROBANTE.pdf")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// if err = smtpChannel.AddTask(dmt.Task{
	// 	Channel: smtp.ChannelName,
	// 	Type:    smtp.SendEmailWithAttachTask,
	// 	Params:  map[string]interface{}{
	// 		"from": "<Bregy Malpartida>bregy.malpartida@utec.edu.pe",
	// 		"to": []string{
	// 			"bregy.malpartida@utec.edu.pe",
	// 			"bregymr@gmail.com",
	// 		},
	// 		"subject": "DMT Test with attach",
	// 		"content_type": "text/html",
	// 		"body": []byte("<h2> Hello World with attachment</h2>"),
	// 		"attachments": map[string][]byte{
	// 			"invoice.pdf": invoiceData,
	// 		},
	// 	},
	// }); err != nil {
	// 	panic(err)
	// }

	// twChannel, err := twilio.NewChannel(

	// 	"+19382220921",
	// )
	// if err != nil {
	// 	panic(err)
	// }
	//
	//
	// _ = twChannel.Observe(func(works []dmt.Work) {
	// 	fmt.Println(time.Now())
	// 	for i, w := range works {
	// 		fmt.Println(i, ":", w.IsDone())
	// 	}
	// })
	//
	// if err := twChannel.AddTask(dmt.Task{
	// 	Channel: twilio.ChannelName,
	// 	Type:    twilio.SendSMSTask,
	// 	Params:  map[string]interface{}{
	// 		"to": []string{"+51957821858"},
	// 		"content": "Hello World from dmt",
	// 	},
	// }); err != nil {
	// 	panic(err)
	// }



	<-a
}
