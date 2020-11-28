package commands

import (
	"fmt"
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var RoleCommand = utils.ICommand{Name: "role", Description: "Toggle a role for a user", Command: roleCommand}

func roleCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate) {
	if len(args) != 3 {
		s.ChannelMessageSend(message.ChannelID, "**Error:** Please supply exactly 2 arguments for this command!")
		return
	}
	if len(message.Mentions) > 1 {
		s.ChannelMessageSend(message.ChannelID, "**Error:** Please only mention one user!")
		return
	}

	guildRoles, _ := s.GuildRoles(message.GuildID)

	for _, rol := range guildRoles {
		fmt.Println(rol)
		if rol.Name == args[2] {
			err := s.GuildMemberRoleAdd(message.GuildID, message.Mentions[0].ID, rol.ID)

			if err != nil {
				s.ChannelMessageSend(message.ChannelID, "**Error:** Failed to add role "+rol.Name+" for "+message.Mentions[0].Username+": "+err.Error())
			} else {
				s.ChannelMessageSend(message.ChannelID, "Modified roles for "+args[1])
			}
		}
	}
}
