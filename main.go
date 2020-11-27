package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"go-discord-bot/command"
)

func main() {
	token := getToken()
	bot, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	bot.AddHandler(ready)

	err = bot.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}

	// commands
	bot.AddHandler(command.CommandHandler)

	// don't shut down
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("ready")
}

func getToken() string {
	content, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
