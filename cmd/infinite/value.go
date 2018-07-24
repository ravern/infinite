package main

import (
	"fmt"

	"github.com/ravernkoh/infinite"
	"github.com/ravernkoh/kubo"
)

func init() {
	value = &kubo.Command{
		Name:        "value",
		Aliases:     []string{"v"},
		Description: "prints the value of the node",
		Arguments: []kubo.Argument{
			{Name: "path"},
		},
		Flags: []kubo.Flag{
			{
				Name:        "set",
				Aliases:     []string{"s"},
				Description: "sets the value to the given value",
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

			value, err := node.Value()
			if err != nil {
				return err
			}

			fmt.Fprintln(ctx.Stdout(), string(value))

			newValue, err := ctx.Flag("set")
			if err != nil {
				return nil
			}

			if err := node.SetValue([]byte(newValue)); err != nil {
				return err
			}

			return node.Save()
		},
	}

	value.Add(value.Help())
}
