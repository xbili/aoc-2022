package day7

import (
	"bufio"
	"os"
	"strings"
)

// TerminalReader reads the terminal output from a file.
type TerminalReader struct {
	scanner *bufio.Scanner
	command *Command // internal state to keep track of the current command
}

func NewTerminalReader(f *os.File) *TerminalReader {
	tr := new(TerminalReader)
	tr.scanner = bufio.NewScanner(f)

	return tr
}

// Reads the next command and output pair
func (tr *TerminalReader) Next() (command *Command, ok bool) {
	current := tr.command

	for {
		tr.scanner.Scan()
		line := tr.scanner.Text()

		// End of file - return the last read command
		if line == "" {
			return current, false
		}

		if strings.HasPrefix(line, "$") {
			// Command
			if current == nil {
				// First command
				current = ParseCommand(line)
			} else {
				// Reached the next command
				tr.command = ParseCommand(line)
				return current, true
			}
		} else {
			// Output for the command
			current.AddOutput(line)
		}
	}
}
