package commands

import (
	"go-discord-bot/log"
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var UnbanCommand = utils.ICommand{Name: "unban", Description: "Unban a user", Command: unbanCommand}

func unbanCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate) {
	if len(args) != 2 {
		log.Error("Please supply exactly one argument for this command", s, message)
		return
	}

	target := args[1]
	bans, err := s.GuildBans(message.GuildID)

	if err != nil {
		log.Error("Error regesting banlist", s, message)
		return
	}

	for _, ban := range bans {
		if ban.User.Username == target {
			err := s.GuildBanDelete(message.GuildID, ban.User.ID)

			if err != nil {
				log.Error("Error removing ban for user", s, message)
			} else {
				log.Success("Unbanned user "+target, s, message)
			}
		}
	}
}
