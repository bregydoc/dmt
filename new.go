package dmt

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewEngine(configFilename string) (*Engine, error) {
	conf, err := extractConfigFromYAMLOrJSON(configFilename)
	if err != nil {
		return nil, err
	}

	engine := new(Engine)
	engine.StartedAt = time.Now()
	engine.Config = conf
	engine.channels = []Channel{}
	engine.rest = &API{
		dmt: engine,
		engine: gin.Default(),
		host: "0.0.0.0",
		port: 8080,
	}


	// for _, ch := range conf.Channels {
	// 	channel, err := inflateChannelFromConfig(ch)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	engine.channels = append(engine.channels, channel)
	// }

	return engine, nil
}
