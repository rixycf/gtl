package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/rixycf/gtl/todo"
	"github.com/urfave/cli"
)

// List define list subcommand
var List = cli.Command{
	Name:   "list",
	Usage:  "show todo list",
	Action: list,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "detail, d",
			Usage: "detail todo list",
		},
	},
}

func list(c *cli.Context) error {

	path, err := getPath()
	if err != nil {
		return err
	}

	todos := readList(path)

	detail := c.Bool("d")

	if detail == true {
		detailTodo(todos)
	} else {
		makeTable(todos)
	}

	return nil
}

func detailTodo(todos []todo.List) {

	for _, todo := range todos {
		var check string
		if todo.Done == true {
			check = "[ok]"
		} else {
			check = "[  ]"
		}
		color.Cyan(" %v %v -> %v", check, todo.Id, todo.Todo)
		if len(todo.Proc) != 0 {
			for _, proc := range todo.Proc {
				fmt.Printf("\t-> %v\n", proc)
			}
		}
	}
}

func makeTable(todos []todo.List) {

	data := make([][]string, 0)

	for _, todo := range todos {
		var check string
		if todo.Done == true {
			check = "  ok  "
		} else {
			check = "  -  "
		}

		//var row []string = {todo.check,todo.Id,todo.Todo}
		row := make([]string, 3)
		row[0] = check
		row[1] = strconv.Itoa(todo.Id)
		row[2] = todo.Todo
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Done", "  ID  ", "  Todo  "})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
