package main

import (
	"fmt"

	"github.com/ravernkoh/infinite"
	"github.com/ravernkoh/kubo"
)

func init() {
	children = &kubo.Command{
		Name:        "children",
		Aliases:     []string{"c"},
		Description: "prints the list of children",
		Arguments: []kubo.Argument{
			{Name: "path"},
		},
		Flags: []kubo.Flag{
			{
				Name:        "new",
				Aliases:     []string{"n"},
				Description: "creates a new child with the given key",
			},
		},
		Run: func(ctx *kubo.Context) error {
			path, err := ctx.Argument("path")
			if err != nil {
				return err
			}

			node, err := infinite.Load(path)
			if err != nil {
				return err
			}

			children, err := node.Children()
			if err != nil {
				return err
			}

			for key := range children {
				fmt.Fprintln(ctx.Stdout(), key)
			}

			newKey, err := ctx.Flag("new")
			if err != nil {
				return nil
			}

			if _, err := node.NewChild(newKey); err != nil {
				return err
			}

			return node.Save()
		},
	}

	children.Add(children.Help())
}
