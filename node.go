package infinite

import (
	"path"

	"github.com/blang/vfs"
)

const maxInt = int(^uint(0) >> 1)

// Node represents a node in the database.
//
// Each node contains a value and some child nodes. Nodes represent a directory
// in the filesystem, while its value is derived from the names of all the files
// within it. Any subdirectories would be child nodes.
//
// A node can sometimes be in an unloaded state, which means that its value and
// its children have not been read from the filesystem. This is desirable in
// circumstances where the amount of data contained in the child nodes is large
// and should only be loaded on-demand.
type Node struct {
	conn *Conn

	loaded   bool
	value    []byte
	children map[string]*Node
}

// Load loads a node at the given path from the OS filesystem.
//
// For more details, see LoadVirtual.
func Load(path string) (*Node, error) {
	return LoadVirtual(path, vfs.OS())
}

// LoadDepth loads a node at the given path from the OS filesystem, up to the
// given depth.
//
// For more details, see LoadVirtualDepth.
func LoadDepth(path string, depth int) (*Node, error) {
	return LoadVirtualDepth(path, vfs.OS(), depth)
}

// LoadVirtual loads a node at the given path from the given filesystem.
//
// All data contained in the node (including child nodes) is read from the
// filesystem in a recursive manner.
func LoadVirtual(path string, fs vfs.Filesystem) (*Node, error) {
	return LoadVirtualDepth(path, fs, maxInt)
}

// LoadVirtualDepth loads a node at the given path from the given filesystem, up
// to the given depth.
//
// Data contained in the node (including child nodes) is read from the
// filesystem in a recursive manner, up to the given depth.
func LoadVirtualDepth(path string, fs vfs.Filesystem, depth int) (*Node, error) {
	return load(path, fs, depth, 0)
}

// load loads a node at the given path from the given filesystem, up to the
// given depth, if the current depth is still valid.
func load(nodePath string, fs vfs.Filesystem, depth int, curDepth int) (*Node, error) {
	node := Node{
		conn: &Conn{
			path: nodePath,
			fs:   fs,
		},
	}

	// If the depth has been reached, then don't load anything else
	if curDepth == depth {
		return &node, nil
	}

	files, dirs, err := node.conn.Load()
	if err != nil {
		return nil, err
	}

	// Decode and set the value
	node.value, err = decodeValue(files)
	if err != nil {
		return nil, err
	}

	// Load and set each child
	node.children = make(map[string]*Node)
	for _, dir := range dirs {
		child, err := load(path.Join(nodePath, dir), fs, depth, curDepth+1)
		if err != nil {
			return nil, err
		}

		node.children[dir] = child
	}

	node.loaded = true

	return &node, nil
}

// Save saves the node to the OS filesystem.
//
// For more details, see SaveVirtual.
func (n *Node) Save() error {
	return n.SaveVirtual(vfs.OS())
}

// SaveVirtual saves the node to the given filesystem.
//
// All data contained in the node (including child nodes) is written into the
// filesystem in a recursive manner. Any data that is not defined in the node
// is removed from the filesystem.
func (n *Node) SaveVirtual(fs vfs.Filesystem) error {
	return nil
}

// Value returns the value of the node.
//
// Will fail if the node has not been loaded.
func (n *Node) Value() ([]byte, error) {
	if !n.loaded {
		return nil, ErrNotLoaded
	}
	return n.value, nil
}

// Child returns the child node with the corresponding key.
//
// Will fail if the node has not been loaded or if the child cannot be found.
func (n *Node) Child(key string) (*Node, error) {
	if !n.loaded {
		return nil, ErrNotLoaded
	}
	c, ok := n.children[key]
	if !ok {
		return nil, ErrNotFound
	}
	return c, nil
}
