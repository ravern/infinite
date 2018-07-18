package infinite

import "fmt"

// Common errors.
var (
	ErrNotFound     = fmt.Errorf("infinite: node not found")
	ErrNotLoaded    = fmt.Errorf("infinite: node not loaded")
	ErrInvalidValue = fmt.Errorf("infinite: value is invalid")
)
