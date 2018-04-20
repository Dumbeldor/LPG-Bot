package main

import (
	"fmt"
)

import (
	"LPG-Bot/LPGBot/config"
	"LPG-Bot/LPGBot/bot"
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