package commands

import (
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var HelpCommand = utils.ICommand{Name: "help", Description: "A help command", Command: helpCommand}

func helpCommand(s *discordgo.Session, message *discordgo.MessageCreate) {
	s.ChannelMessageSend(message.ChannelID, "*The dark side never helps, it laughs*")
}
