package main

import (
	"fmt"
	"os"

	"github.com/ravernkoh/kubo"
)

var (
	root     *kubo.Command
	value    *kubo.Command
	children *kubo.Command
)

func init() {
	root = &kubo.Command{
		Name:        "infinite",
		Description: "the database that can store the internet",
	}

	root.Add(root.Help())
}

func main() {
	root.Add(value)
	root.Add(children)

	if err := kubo.NewApp(root).Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
