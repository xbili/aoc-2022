package day7

import "strings"

type CommandType int

const (
	InvalidCommand CommandType = iota
	Cd
	Ls
)

// A command represents a user input text based command that either creates
// new outputs, or modifies the current context.
type Command struct {
	commandType CommandType // type of command
	parameters  []string    // parameters passed into the command
	output      []string    // output of the command
}

func (command *Command) AddOutput(line string) {
	command.output = append(command.output, line)
}

func ParseCommand(command string) *Command {
	tokens := strings.Split(command, " ")

	// First token is $, we can safely ignore that
	commandType := getCommandType(tokens[1])

	return &Command{
		commandType: commandType,
		// Parameters goes from the second token onwards
		parameters: tokens[2:],
		output:     []string{},
	}
}

func getCommandType(token string) CommandType {
	switch token {
	case "cd":
		return Cd
	case "ls":
		return Ls
	default:
		return InvalidCommand
	}
}
