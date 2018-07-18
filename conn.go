package infinite

// Conn represents the underlying connection to the file system.
//
// Since Infinite does not contain an actual connection to anything, this simply
// keeps track of a path within the filesystem and also provides checks for
// corruption.
type Conn struct {
	Path string
}
