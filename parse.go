package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func parseFile(filename string) (interface{}, error) {

	var data interface{}
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, &data); err != nil {
		return nil, err
	}
	transformData(&data)
	return data, nil
}
