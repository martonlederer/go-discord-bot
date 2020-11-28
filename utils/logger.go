package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// error
func LogError(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0xff0000}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
	// also print to console
	fmt.Println("Error: " + log)
}

// warning
func LogWarning(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0xff7300}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}

// success
func LogSuccess(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0x00ff22}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}

// error something
func LogInfo(log string, s *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Title: log, Color: 0x00d5ff}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}
