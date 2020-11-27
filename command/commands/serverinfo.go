package commands

import (
	"go-discord-bot/utils"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var InfoCommand = utils.ICommand{Name: "serverinfo", Description: "Return server information", Command: infoCommand}

func infoCommand(s *discordgo.Session, message *discordgo.MessageCreate) {
	s.RequestGuildMembers(message.GuildID, "", 0, true)

	embed := &discordgo.MessageEmbed{Title: "Server info", Color: 0x34b4eb}
	server, _ := s.Guild(message.GuildID)
	owner, _ := s.User(server.OwnerID)

	embed.Fields = append(
		embed.Fields,
		&discordgo.MessageEmbedField{Name: "Server name", Value: server.Name},
		&discordgo.MessageEmbedField{Name: "Owner", Value: owner.Username},
		&discordgo.MessageEmbedField{Name: "Region", Value: server.Region},
		&discordgo.MessageEmbedField{Name: "Total members", Value: strconv.FormatInt(int64(server.ApproximateMemberCount), 10)},
		&discordgo.MessageEmbedField{Name: "Online", Value: strconv.FormatInt(int64(server.ApproximatePresenceCount), 10)},
	)

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}
