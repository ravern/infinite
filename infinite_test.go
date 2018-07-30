package infinite_test

import (
	"fmt"

	"github.com/ravernkoh/infinite"
)

func Example() {
	// Load the root node
	node, err := infinite.Load("db")
	if err != nil {
		panic(err)
	}

	// Gets the child node
	child, err := node.Child("user")
	if err != nil {
		panic(err)
	}

	// Gets the value of the child node
	user, err := child.Value()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
