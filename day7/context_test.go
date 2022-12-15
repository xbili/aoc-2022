package day7

import "testing"

func TestBuildContextRoot(t *testing.T) {
	context := NewContext()
	commands := []*Command{
		{commandType: Cd, parameters: []string{"/"}, output: []string{"/"}},
	}
	tr := NewTerminalReaderStub(commands)

	context.Build(tr)

	assertNode(t, context.root, "/")
	assertChildren(t, context.root, []*Node{})
}

func TestBuildContextLs(t *testing.T) {
	context := NewContext()
	commands := []*Command{
		{commandType: Cd, parameters: []string{"/"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"1 file1",
			"2 file2",
			"3 file3",
		}},
	}
	tr := NewTerminalReaderStub(commands)

	context.Build(tr)

	assertChildren(t, context.root, []*Node{
		{name: "file1"},
		{name: "file2"},
		{name: "file3"},
	})
}

func TestBuildContextCdRoot(t *testing.T) {
	context := NewContext()
	commands := []*Command{
		{commandType: Cd, parameters: []string{"/"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"dir dir1",
		}},
		{commandType: Cd, parameters: []string{"dir1"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"dir dir2",
		}},
		{commandType: Cd, parameters: []string{"dir2"}, output: []string{}},
		{commandType: Cd, parameters: []string{"/"}, output: []string{}},
	}
	tr := NewTerminalReaderStub(commands)

	context.Build(tr)

	assertCurrentDirectory(t, context, "/")
}

func TestBuildContextCdNamed(t *testing.T) {
	context := NewContext()
	commands := []*Command{
		{commandType: Cd, parameters: []string{"/"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"dir dir1",
		}},
		{commandType: Cd, parameters: []string{"dir1"}, output: []string{}},
	}
	tr := NewTerminalReaderStub(commands)

	context.Build(tr)

	assertCurrentDirectory(t, context, "dir1")
}

func TestBuildContextCdUp(t *testing.T) {
	context := NewContext()
	commands := []*Command{
		{commandType: Cd, parameters: []string{"/"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"dir dir1",
		}},
		{commandType: Cd, parameters: []string{"dir1"}, output: []string{}},
		{commandType: Cd, parameters: []string{".."}, output: []string{}},
	}
	tr := NewTerminalReaderStub(commands)

	context.Build(tr)

	assertCurrentDirectory(t, context, "/")
}

func TestComputeDirectoriesSizes(t *testing.T) {
	context := NewContext()
	commands := []*Command{
		{commandType: Cd, parameters: []string{"/"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"dir dir1",
			"30 file3",
		}},
		{commandType: Cd, parameters: []string{"dir1"}, output: []string{}},
		{commandType: Ls, parameters: []string{}, output: []string{
			"dir dir2",
			"10 file1",
			"20 file2",
		}},
	}
	tr := NewTerminalReaderStub(commands)

	context.Build(tr)
	sizes := context.ComputeDirectoriesSizes()

	// No way to perform reverse look up for name <> node, therefore we use switch
	// statement to assert sizes
	for dir, size := range sizes {
		switch dir.name {
		case "/":
			assertSize(t, size, 60)
		case "dir1":
			assertSize(t, size, 30)
		case "dir2":
			assertSize(t, size, 0)
		default:
			t.Fatal("Invalid directory found in result")
		}
	}
}

func assertNode(t *testing.T, node *Node, name string) {
	if node.name != name {
		t.Fatalf("Expected node name to be %s, but got %s instead", name, node.name)
	}
}

func assertChildren(t *testing.T, node *Node, children []*Node) {
	if len(node.children) != len(children) {
		t.Fatalf("Expected number of children to be %d, but got %d instead", len(children), len(node.children))
	}

	for i, child := range children {
		if node.children[i].name != child.name {
			t.Fatalf("Expected child at %d to be %s, but got %s instead", i, child.name, node.children[i].name)
		}
	}
}

func assertCurrentDirectory(t *testing.T, context *Context, cwd string) {
	if context.current.name != cwd {
		t.Fatalf("Expected cwd to be %s, but received %s instead", cwd, context.current.name)
	}
}

func assertSize(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Fatalf("Expected size to be %d, but received %d instead", expected, actual)
	}
}

// Faked terminal reader to stub out the input for building context
type TerminalReaderStub struct {
	commands []*Command
	index    int
}

func NewTerminalReaderStub(commands []*Command) *TerminalReaderStub {
	return &TerminalReaderStub{
		commands: commands,
		index:    0,
	}
}

func (tr *TerminalReaderStub) Next() (value *Command, ok bool) {
	value = tr.commands[tr.index]
	tr.index++

	return value, tr.index < len(tr.commands)
}
