package day7

import (
	"strconv"
	"strings"
)

type NodeType int

const (
	InvalidNodeType NodeType = iota
	File
	Directory
)

// A node represents an item in the file system
type Node struct {
	name     string   // name of node
	nodeType NodeType // type of node
	parent   *Node    // pointer to the parent of this node, nil if none
	children []*Node  // pointer to other nodes that belong to this node
}

// Adds a new node as the child of the current node
func (node *Node) AddChild(other *Node) {
	if node.nodeType != Directory {
		panic("Cannot add child to a non-directory node")
	}

	node.children = append(node.children, other)
}

// Returns node information from the terminal output
func ParseNode(line string) (name string, size int, nodeType NodeType) {
	tokens := strings.Split(line, " ")

	// Name is the second token
	name = tokens[1]

	// Node type depends on the first token
	if tokens[0] == "dir" {
		return name, 0, Directory
	} else {
		// Size of the file is the first token
		size, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}

		return name, size, File
	}
}
