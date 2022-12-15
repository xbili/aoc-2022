package day7

import "testing"

func TestParseFileNodes(t *testing.T) {
	assertParseNode(t, "123 file", "file", 123, File)
	assertParseNode(t, "123 f.txt", "f.txt", 123, File)
	assertParseNode(t, "3 f", "f", 3, File)
}

func TestParseDirectoryNodes(t *testing.T) {
	assertParseNode(t, "dir dir", "dir", 0, Directory)
	assertParseNode(t, "dir d", "d", 0, Directory)
}

func assertParseNode(t *testing.T, line, name string, size int, nodeType NodeType) {
	parsedName, parsedSize, parsedNodeType := ParseNode(line)

	if parsedName != name {
		t.Fatalf("Expected name to be %s, but got %s", name, parsedName)
	}

	if parsedSize != size {
		t.Fatalf("Expected size to be %d, but got %d", size, parsedSize)
	}

	if parsedNodeType != nodeType {
		t.Fatalf("Expected nodeType to be %d, but got %d", nodeType, parsedNodeType)
	}
}
