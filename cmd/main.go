package main

import (
	"fmt"

	"github.com/bregydoc/dmt"
)

func main() {
	engine, err := dmt.NewEngine("./config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(engine.StartedAt)

	if err = engine.Run(); err != nil {
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
}
