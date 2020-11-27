package commands

import (
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var RoleCommand = utils.ICommand{Name: "role", Description: "Toggle a role for a user", Command: roleCommand}

func roleCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate) {
	if len(args) != 3 {
		s.ChannelMessageSend(message.ChannelID, "**Error:** Please supply exactly 2 arguments for this command!")
		return
	}
	// s.GuildMemberRoleAdd(message.GuildID, message.Member.User.ID, )
	// server, _ := s.Guild(message.GuildID)

	for _, u := range message.Mentions {
		for _, r := range message.MentionRoles {
			err := s.GuildMemberRoleAdd(message.GuildID, u.ID, r)

			if err != nil {
				s.ChannelMessageSend(message.ChannelID, "**Error:** Failed to add role "+r+" for "+u.Username+": "+err.Error())
			} else {
				s.ChannelMessageSend(message.ChannelID, "Modified roles for "+args[1])
			}
		}
	}
}
