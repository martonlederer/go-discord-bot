package commands

import (
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var InfoCommand = utils.ICommand{Name: "serverinfo", Description: "Return server information", Command: infoCommand}

func infoCommand(s *discordgo.Session, message *discordgo.MessageCreate) {
	s.ChannelMessageSend(message.ChannelID, "Info")
}
