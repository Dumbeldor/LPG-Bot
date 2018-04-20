package bot

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"LPG-Bot/LPGBot/config"
	"strings"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate)  {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotId {
			return
		}

		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}

		if m.Content == "!joke" {
			err := sendJoke()
			if err != nil {
				fmt.Println(err)
			}
		}

		if m.Content == "!test" {
			err := testJoke()
			if err != nil {
				fmt.Println(err)
			}
		}


	}

}


