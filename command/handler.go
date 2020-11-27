package command

import (
	"regexp"
	"strings"

	"go-discord-bot/command/commands"

	"github.com/bwmarrin/discordgo"
)

var prefix string = "!"

func CommandHandler(s *discordgo.Session, message *discordgo.MessageCreate) {

	/* Commands */
	allCommands := make(map[string]func(*discordgo.Session, *discordgo.MessageCreate))
	allCommands["help"] = commands.HelpCommand

	// don't reply to ourselves
	if message.Author.ID == s.State.User.ID {
		return
	}

	// does it start with ! (prefix)
	matched, _ := regexp.Match(`^`+prefix, []byte(message.Content))

	if !matched {
		return
	}

	// parse command
	withoutPrefix := strings.Replace(message.Content, "!", "", 1)
	commandArgs := strings.Split(withoutPrefix, " ")
	cmd := allCommands[commandArgs[0]]

	if cmd == nil {
		s.ChannelMessageSend(message.ChannelID, "Could not find command")
		return
	}

	cmd(s, message)
}
