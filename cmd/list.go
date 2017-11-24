package cmd

import (
	"fmt"
	_ "os"

	_ "github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

var List = cli.Command{
	Name:   "list",
	Usage:  "show todo list",
	Action: list,
}

func list(c *cli.Context) {
	//TODO read file and marchal jsonfile
	//TODO write stdout todo list

	fmt.Println("list is working")
}
