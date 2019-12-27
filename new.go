package dmt

import "github.com/davecgh/go-spew/spew"

func NewEngine(configFilename string) (*Engine, error) {
	conf, err := extractConfigFromYAMLOrJSON(configFilename)
	if err != nil {
		return nil, err
	}

	spew.Dump(conf)
	return nil, nil
}
