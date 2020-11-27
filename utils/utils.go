package utils

import "github.com/bwmarrin/discordgo"

type ICommand struct {
	Name, Description string
	Command           func(*discordgo.Session, *discordgo.MessageCreate)
}
