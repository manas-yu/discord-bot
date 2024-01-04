package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading Config file")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("error while reading config")
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("error while unmarshalling config")
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
