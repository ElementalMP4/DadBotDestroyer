package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type DadBotDestroyerConfig struct {
	Token  string
	Users  []string
	Guilds []string
}

var config DadBotDestroyerConfig

func main() {

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Couldn't find config!")
		return
	}
	config = DadBotDestroyerConfig{}
	err = json.Unmarshal([]byte(file), &config)

	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(reactionAdd)

	dg.Identify.Intents = discordgo.IntentsGuildMessageReactions

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Destroying Dad Bot as " + dg.State.User.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func reactionAdd(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if !contains(config.Guilds, m.GuildID) {
		return
	}

	msg, _ := s.ChannelMessage(m.ChannelID, m.MessageID)

	if !contains(config.Users, msg.Author.ID) {
		return
	}

	if m.Emoji.Name != "❌" {
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.MessageID)
}
