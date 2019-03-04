package config

import (
	"encoding/json"
	"io/ioutil"
)

//Cfg is config parameters
type Cfg struct {
	MC     mc     `json:"mc"`
	Server server `json:"server"`
}

type mc struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type server struct {
	Port int `json:"port"`
}

//LoadConfig load and parse config.json file
func LoadConfig() (*Cfg, error) {
	jsonString, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	c := new(Cfg)
	if err = json.Unmarshal(jsonString, c); err != nil {
		return nil, err
	}
	return c, nil
}
