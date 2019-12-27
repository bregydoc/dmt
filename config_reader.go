package dmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func extractConfigFromYAML(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	if err = yaml.Unmarshal(data, config); err != nil {
		log.Error(err)
		return nil, err
	}

	return config, nil
}

func extractConfigFromJSON(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

func extractConfigFromYAMLOrJSON(filename string) (*Config, error) {
	chunks := strings.Split(filename, ".")
	firstYaml := true
	if len(chunks) > 1 {
		fileType := chunks[1]
		if fileType == "json" || fileType == "js" {
			firstYaml = false
		}
	}

	var conf *Config
	var err error

	fmt.Println(firstYaml)
	if firstYaml {
		conf, err = extractConfigFromYAML(filename)
	} else {
		conf, err = extractConfigFromJSON(filename)
	}

	if err != nil {
		if firstYaml {
			conf, err = extractConfigFromJSON(filename)
		} else {
			conf, err = extractConfigFromYAML(filename)
		}
		if err != nil {
			return nil, err
		}
		return conf, nil
	}

	return conf, err
}
