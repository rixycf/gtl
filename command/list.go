package cmd

import (
	"fmt"
	"os"
	_ "path/filepath"

	"github.com/olekukonko/tablewriter"
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

	// TODO output ASCII Table
	todos := readList(`./test.json`)

	for _, todo := range todos {
		var check string
		if todo.Done == true {
			check = "○"
		} else {
			check = "×"
		}
		fmt.Printf("%v  %v  -> %v\n", check, todo.Id, todo.Todo)
	}

	data := make([][]string, 0)

	for _, todo := range todos {
		var check string
		if todo.Done == true {
			check = "ok"
		} else {
			check = "no"
		}

		//var row []string = {todo.check,todo.Id,todo.Todo}
		row := make([]string, 3)
		row[0] = check
		row[1] = ""
		row[2] = todo.Todo
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"  Done  ", "  ID  ", "  Todo  "})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

	return nil
}
