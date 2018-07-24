package infinite

import "fmt"

// Common errors.
var (
	ErrNotFound      = fmt.Errorf("infinite: node not found")
	ErrNotLoaded     = fmt.Errorf("infinite: node not loaded")
	ErrAlreadyExists = fmt.Errorf("infinite: node already exists")
	ErrInvalidValue  = fmt.Errorf("infinite: value is invalid")
)
