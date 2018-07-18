package infinite

import (
	"github.com/blang/vfs"
)

const maxInt = int(^uint(0) >> 1)

// Node represents a node in the database.
//
// Each node contains a value and some child nodes. Nodes represent a directory
// in the filesystem, while its value is derived from the names of all the files
// within it. Any subdirectories would be child nodes.
//
// A node can sometimes be in an unloaded state, which means that its children
// have not been read from the filesystem. This is desirable in circumstances
// where the amount of data contained in the child nodes is large and should
// only be loaded on-demand.
type Node struct {
	conn *Conn

	value    string
	children []*Node
}

// Load loads a node from the OS filesystem.
//
// For more details, see LoadVirtual.
func Load() (*Node, error) {
	return LoadVirtual(vfs.OS())
}

// LoadDepth loads a node from the OS filesystem, up to the given depth.
//
// For more details, see LoadDepthVirtual.
func LoadDepth(depth int) (*Node, error) {
	return LoadDepthVirtual(depth, vfs.OS())
}

// LoadVirtual loads a node from the given filesystem.
//
// All data contained in the node (including child nodes) is read from the
// filesystem in a recursive manner.
func LoadVirtual(fs vfs.Filesystem) (*Node, error) {
	return LoadDepthVirtual(maxInt, fs)
}

// LoadDepthVirtual loads a node from the given filesystem, up to the given
// depth.
//
// Data contained in the node (including child nodes) is read from the
// filesystem in a recursive manner, up to the given depth.
func LoadDepthVirtual(depth int, fs vfs.Filesystem) (*Node, error) {
	return nil, nil
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
