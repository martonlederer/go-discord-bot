package command

import (
	"errors"
	"regexp"
	"strings"

	"go-discord-bot/command/commands"
	"go-discord-bot/log"
	"go-discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

var allCommands = []utils.ICommand{
	commands.InfoCommand,
	commands.RoleCommand,
	commands.KickCommand,
}

func CommandHandler(s *discordgo.Session, message *discordgo.MessageCreate, config *utils.Config) {
	// don't reply to ourselves
	if message.Author.ID == s.State.User.ID {
		return
	}

	// does it start with ! (prefix)
	matched, _ := regexp.Match(`^`+config.Prefix, []byte(message.Content))

	if !matched {
		return
	}

	// parse command
	withoutPrefix := strings.Replace(message.Content, config.Prefix, "", 1)
	commandArgs := strings.Split(withoutPrefix, " ")

	// the help command is not
	// added to the commands list
	// i have to figure this out
	// but if I add it, it throws
	// an initialization loop error
	if commandArgs[0] == "help" {
		helpCommand(commandArgs, s, message, config)
		return
	}

	command, err := findCommand(commandArgs[0])

	if err != nil {
		log.Error(err.Error(), s, message)
		return
	}

	// execute command
	command.Command(commandArgs, s, message)
}

// find a command from the command list
func findCommand(cmd string) (command utils.ICommand, err error) {
	for _, c := range allCommands {
		if c.Name == cmd {
			return c, nil
		}
	}
	return command, errors.New("Could not find command")
}

// list all commands
func helpCommand(args []string, s *discordgo.Session, message *discordgo.MessageCreate, config *utils.Config) {
	embed := &discordgo.MessageEmbed{Title: "Help", Color: 0x4ceb34}

	for _, c := range allCommands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: config.Prefix + c.Name, Value: c.Description})
	}

	s.ChannelMessageSendEmbed(message.ChannelID, embed)
}
