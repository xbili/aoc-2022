package day7

import (
	"fmt"
	"xbili.com/aoc/utils"
)

// The context keeps track of the current state of the terminal based on
// commands that were ran.
type Context struct {
	root    *Node         // root directory
	current *Node         // current directory
	sizes   map[*Node]int // index of sizes for every file node
}

func NewContext() *Context {
	return &Context{
		root:    nil,
		current: nil,
		sizes:   map[*Node]int{},
	}
}

// Builds the context from the terminal output reader
func (context *Context) Build(tr utils.Iterator[*Command]) bool {
	for {
		command, ok := tr.Next()

		switch command.commandType {
		case Cd:
			// Find the node with the name under current parent, create one if it does
			// not exist
			context.changeDirectory(command)
		case Ls:
			// Create nodes if they do not exist under the current directory
			context.list(command)
		default:
			// ERROR: Invalid command
			fmt.Println("Invalid command type")
			return false
		}

		// Last command has been read
		if !ok {
			return true
		}
	}
}

// Returns the sizes of all directories as a list of tuples.
func (context *Context) ComputeDirectoriesSizes() map[*Node]int {
	directorySizes := map[*Node]int{}
	computeDirectorySize(context.root, context.sizes, directorySizes)

	return directorySizes
}

func (context *Context) changeDirectory(command *Command) bool {
	// Name of cd is the first parameter passed in, ignore the rest
	name := command.parameters[0]

	if name == "/" {
		// Create a new root directory
		if context.root == nil {
			context.root = &Node{
				name:     "/",
				nodeType: Directory,
				parent:   nil,
				children: []*Node{},
			}
		}

		// Go to the root directory
		context.current = context.root

		return true
	}

	// Go to the parent directory
	if name == ".." {
		if context.current.parent == nil {
			// Cannot go any higher, this should be an error
			fmt.Println("Invalid parent!")
			return false
		}

		context.current = context.current.parent
		return true
	}

	// Find the desired node (must be directory) in current list of children
	var target *Node
	for _, child := range context.current.children {
		if child.name == name && child.nodeType == Directory {
			target = child
		}
	}

	// Create a new directory in the tree if it does not exist
	// (I don't think this should happen though, but just in case)
	if target == nil {
		target = &Node{
			name:     name,
			nodeType: Directory,
			parent:   context.current,
			children: []*Node{},
		}

		if context.current != nil {
			context.current.AddChild(target)
		}
	}

	// Set the current directory to the target
	context.current = target

	return true
}

func (context *Context) list(command *Command) bool {
	for _, line := range command.output {
		name, size, nodeType := ParseNode(line)
		node := &Node{
			name:     name,
			nodeType: nodeType,
			parent:   context.current,
			children: []*Node{},
		}

		context.current.AddChild(node)

		if nodeType == File {
			context.sizes[node] = size
		}
	}

	return true
}

func computeDirectorySize(node *Node, fileSizes, directorySizes map[*Node]int) int {
	res := 0
	for _, child := range node.children {
		if child.nodeType == Directory {
			if cached, ok := directorySizes[child]; ok {
				res += cached
			} else {
				res += computeDirectorySize(child, fileSizes, directorySizes)
			}
		} else {
			res += fileSizes[child]
		}
	}

	// Cache this result for future uses
	directorySizes[node] = res

	return res
}
