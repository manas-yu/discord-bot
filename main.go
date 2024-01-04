package main

import (
	"fmt"

	"github.com/manas-yu/discord-ping/bot"
	"github.com/manas-yu/discord-ping/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}
