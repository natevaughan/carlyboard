package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DbHost   string `yaml:"dbhost"`
	DbName   string `yaml:"name"`
	HttpPort string `yaml:"port"`
}

func (c *Config) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, c)

	if err != nil {
		return err
	}
	println("dbHost: " + c.DbHost)
	println("dbName: " + c.DbName)
	println("httpPort: " + c.HttpPort)
	return nil
}
