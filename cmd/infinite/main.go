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
		node, err := infinite.Load("db")
		if err != nil {
			return err
		}

		value, err := node.Value()
		if err != nil {
			return err
		}

		fmt.Println(string(value))

		if err := node.SetValue([]byte("Boom!")); err != nil {
			return err
		}

		if err := node.Save(); err != nil {
			return err
		}

		return nil
	},
}

func main() {
	root.Add(root.Help())

	if err := kubo.NewApp(root).Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
