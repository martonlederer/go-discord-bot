package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"regexp"
	"io/ioutil"
	"log"
)

func main () {
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
	bot.AddHandler(commandHandler)

	// don't shut down
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}

func ready (s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("ready")
}

func commandHandler (s *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == s.State.User.ID {
		return
	}

	// does it start with ! (prefix)
	matched, _ := regexp.Match(`^!`, []byte(message.Content))

	if !matched {
		return
	}

	s.ChannelMessageSend(message.ChannelID, "You said " + message.Content)

}

func getToken () string {
	content, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}