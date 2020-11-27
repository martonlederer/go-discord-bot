package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	"go-discord-bot/command"
	"go-discord-bot/utils"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	if os.Getenv("TOKEN") == "" {
		log.Fatal("No token provided. Exiting...")
	}

	bot, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	bot.AddHandler(ready)

	err = bot.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}

	// config
	configFile, err := os.Open("config.json")

	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	byteValue, _ := ioutil.ReadAll(configFile)

	var config *utils.Config
	json.Unmarshal(byteValue, &config)

	fmt.Println(config)

	// commands
	bot.AddHandler(func(s *discordgo.Session, message *discordgo.MessageCreate) {
		command.CommandHandler(s, message, config)
	})

	fmt.Println("Running bot. Ctrl + C to quit...")

	// don't shut down
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is online...")
}
