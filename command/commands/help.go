package commands

import "github.com/bwmarrin/discordgo"

func HelpCommand(s *discordgo.Session, message *discordgo.MessageCreate) {
	s.ChannelMessageSend(message.ChannelID, "*The dark side never helps, it laughs*")
}
