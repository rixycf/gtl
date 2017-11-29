package cmd

import (
	"fmt"

	"github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

// Add define add command
var Add = cli.Command{
	Name:   "add",
	Usage:  "add task",
	Action: add,
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "p, process",
			Usage: "this option can write processes for task",
		},
	},
}

func add(c *cli.Context) error {

	// TODO get command line argment
	// TODO assign todo.List

	// get todo from commanline argment
	tasks := c.Args()

	todos := readList("./test.json")

	var label int = len(todos)
	// add todo to todo.List
	for i, task := range tasks {
		var t todo.List
		t.Todo = task
		t.Id = label + i + 1
		todos = append(todos, t)

		fmt.Printf(" Add task: %v\n", task)
	}

	fmt.Println(c.StringSlice("p"))

	writeList("./test.json", todos)
	return nil
}
