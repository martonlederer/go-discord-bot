package command

import (
	"errors"
	"regexp"
	"strings"

	"go-discord-bot/command/commands"
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var prefix string = "!"
var allCommands = []utils.ICommand{commands.HelpCommand}

func CommandHandler(s *discordgo.Session, message *discordgo.MessageCreate) {
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
	command, err := findCommand(commandArgs[0])

	if err != nil {
		s.ChannelMessageSend(message.ChannelID, err.Error())
		return
	}

	command.Command(s, message)
}

func findCommand(cmd string) (command utils.ICommand, err error) {
	for _, c := range allCommands {
		if c.Name == cmd {
			return c, nil
		}
	}
	return command, errors.New("Could not find command")
}
