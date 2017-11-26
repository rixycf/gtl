package cmd

import (
	"github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

// Add define add command
var Add = cli.Command{
	Name:   "add",
	Usage:  "add todo",
	Action: add,
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
		t.Id = label + i
		todos = append(todos, t)
	}

	writeList("./test.json", todos)
	return nil
}
