package log

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// error
func Error(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0xff0000}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
	// also print to console
	fmt.Println("Error: " + log)
}

// warning
func Warning(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0xff7300}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}

// success
func Success(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0x00ff22}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}

// error something
func Info(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0x00d5ff}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}
