package commands

import (
	"go-discord-bot/utils"
	"sort"

	"github.com/bwmarrin/discordgo"
)

var RoleCommand = utils.ICommand{Name: "role", Description: "Toggle a role for a user", Command: roleCommand}

func roleCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate) {
	if len(args) != 3 {
		utils.LogError("Please supply exactly 2 arguments for this command!", s, message)
		return
	}
	if len(message.Mentions) != 1 {
		utils.LogError("Please mention exactly one user!", s, message)
		return
	}

	guildRoles, _ := s.GuildRoles(message.GuildID)

	for _, rol := range guildRoles {
		if rol.Name == args[2] {
			member, err := s.GuildMember(message.GuildID, message.Mentions[0].ID)
			if err != nil {
				utils.LogError("Failed to find user "+message.Mentions[0].Username, s, message)
				return
			}

			// does the user have the role
			roleIndex := sort.SearchStrings(member.Roles, rol.ID)

			if roleIndex != len(member.Roles) {
				err := s.GuildMemberRoleRemove(message.GuildID, member.User.ID, rol.ID)

				if err != nil {
					utils.LogError("**Error:** Failed to modify roles ("+rol.Name+") for "+message.Mentions[0].Username, s, message)
				} else {
					utils.LogWarning("Removed role from "+message.Mentions[0].Username, s, message)
				}
			} else {
				err := s.GuildMemberRoleAdd(message.GuildID, member.User.ID, rol.ID)

				if err != nil {
					utils.LogError("**Error:** Failed to modify roles ("+rol.Name+") for "+message.Mentions[0].Username, s, message)
				} else {
					utils.LogSuccess("Added role for "+message.Mentions[0].Username, s, message)
				}
			}
		}
	}
}
