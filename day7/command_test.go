package day7

import "testing"

func TestParseCommand(t *testing.T) {
	assertParseCommand(t, "$ cd /", Cd, []string{"/"})
	assertParseCommand(t, "$ cd ..", Cd, []string{".."})
	assertParseCommand(t, "$ cd 1 2 3 4", Cd, []string{"1", "2", "3", "4"}) // invalid, but we should include these in parameters
	assertParseCommand(t, "$ ls", Ls, []string{})
	assertParseCommand(t, "$ cd", Cd, []string{})
}

func assertParseCommand(t *testing.T, line string, commandType CommandType, parameters []string) {
	command := ParseCommand(line)

	if command.commandType != commandType {
		t.Fatalf("Expected commandType %d, but got %d", commandType, command.commandType)
	}

	if len(command.parameters) != len(parameters) {
		t.Fatalf("Expected length of parameters to be %d, but got %d", len(parameters), len(command.parameters))
	}

	for i, param := range command.parameters {
		if parameters[i] != param {
			t.Fatalf("Expected parameter at %d to be %s, but got %s", i, parameters[i], param)
		}
	}
}
