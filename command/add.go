package cmd

import (
	"fmt"

	"github.com/rixycf/gtl/todo"
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
		cli.IntFlag{
			Name:  "i, id",
			Usage: "this option specify ID to add processes",
		},
	},
}

func add(c *cli.Context) error {

	path, err := getPath()
	fmt.Println(path)
	if err != nil {
		return err
	}

	// get task from commanline argment
	tasks := c.Args()
	for _, task := range tasks {
		fmt.Println(task)
	}

	todos := readList(path)

	var label = len(todos)

	// add task to todo.List
	for i, task := range tasks {
		var t todo.List
		t.Todo = task
		t.Id = label + i + 1
		todos = append(todos, t)

		fmt.Printf(" Add task: %v\n", task)
	}

	// option "-i"
	var id = c.Int("i")
	if id > 0 {

		if len(todos) < id {
			fmt.Println("wrong id")
			return nil
		}

		for _, p := range c.StringSlice("p") {
			todos[id-1].Proc = append(todos[id-1].Proc, p)
		}
	}

	writeList(path, todos)
	return nil
}
