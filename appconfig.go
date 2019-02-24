package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Dbconn string `yml:"dbconn"`
	Port string `yml:"port"`
}

func (c *Config) Load(filename string) (error) {
	file, err := ioutil.ReadFile(filename) 

	if err != nil {
		return err
	}
	
	err = yaml.Unmarshal(file, c)

	if err != nil {
		return err
	}
	return nil
}
