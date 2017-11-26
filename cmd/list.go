package cmd

import (
	"fmt"
	_ "os"
	_ "path/filepath"

	_ "github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

var List = cli.Command{
	Name:   "list",
	Usage:  "show todo list",
	Action: list,
}

func list(c *cli.Context) error {
	fmt.Println("list is working")

	//TODO read file and marshal jsonfile
	//TODO write stdout todo list
	// readfile .testjson

	// todos is todo.List slice
	todos := readList(`./test.json`)

	for num, todo := range todos {
		fmt.Println(num, todo.Todo, todo.Done)
	}
	fmt.Println(todos)
	return nil
}
