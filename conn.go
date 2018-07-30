package infinite

import "github.com/blang/vfs"

// conn represents the underlying connection to the file system.
//
// Usually, Conn will not be used directly but instead accessed by the its
// wrapper node, which provides higher level functionality and caching.
//
// Since Infinite does not contain an actual connection to anything, this simply
// keeps track of a path within the filesystem and also provides checks for
// corruption.
type conn struct {
	path string
	fs   vfs.Filesystem
}

// ReadDir loads data at the path from the filesystem.
//
// Returns the file names, directory names and a possible error.
func (c *conn) ReadDir() ([]string, []string, error) {
	infos, err := c.fs.ReadDir(c.path)
	if err != nil {
		return nil, nil, err
	}

	var (
		files []string
		dirs  []string
	)

	// Add the names to the appropriate slices
	for _, info := range infos {
		if info.IsDir() {
			dirs = append(dirs, info.Name())
		} else {
			files = append(files, info.Name())
		}
	}

	return files, dirs, nil
}

// Path returns the path of the connected directory.
func (c *conn) Path() string {
	return c.path
}
