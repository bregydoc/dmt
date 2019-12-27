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
}


