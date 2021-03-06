package commands

import (
	"go-discord-bot/log"
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var KickCommand = utils.ICommand{Name: "kick", Description: "Kick a user", Command: kickCommand}

func kickCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate) {
	if len(args) != 2 {
		log.Error("Please supply exactly one argument for this command", s, message)
		return
	}
	if len(message.Mentions) != 1 {
		log.Error("Please mention a user to kick", s, message)
		return
	}

	target := message.Mentions[0]
	err := s.GuildMemberDelete(message.GuildID, target.ID)

	if err != nil {
		log.Error("Could not kick user "+target.Username, s, message)
		return
	}

	log.Success("Kicked user "+target.Username, s, message)
}
