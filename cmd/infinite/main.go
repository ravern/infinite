package main

import (
	"fmt"
	"os"

	"github.com/ravernkoh/infinite"
	"github.com/ravernkoh/kubo"
)

var root = &kubo.Command{
	Name:        "infinite",
	Description: "the database that can store the internet",
	Run: func(ctx *kubo.Context) error {
		node, err := infinite.LoadDepth("db", 0)
		if err != nil {
			return err
		}

		node, err = node.Child("boom")
		if err != nil {
			return err
		}

		fmt.Println(node.Value())

		return nil
	},
}

func main() {
	root.Add(root.Help())

	if err := kubo.NewApp(root).Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
