package commands

import (
	"go-discord-bot/log"
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var BanCommand = utils.ICommand{Name: "ban", Description: "Ban a user", Command: banCommand}

func banCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate) {
	if len(args) != 2 {
		log.Error("Please supply exactly one argument for this command", s, message)
		return
	}
	if len(message.Mentions) != 1 {
		log.Error("Please mention a user to ban", s, message)
		return
	}

	target := message.Mentions[0]
	err := s.GuildBanCreate(message.GuildID, target.ID, 0)

	if err != nil {
		log.Error("Could not ban user "+target.Username, s, message)
		return
	}

	log.Success("Banned user "+target.Username, s, message)
}
