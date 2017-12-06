package cmd

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

// Done define done subcommand
var Done = cli.Command{
	Name:   "done",
	Usage:  "check off todo",
	Action: done,
}

// done command check todo list
func done(c *cli.Context) error {

	// get id to check TODOList
	ids := c.Args()
	path, err := getPath()
	if err != nil {
		return err
	}
	// read todolist
	todos := readList(path)
	//fmt.Println(todos)

	// check todoList
	for _, id := range ids {

		id, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if id > len(todos) {
			return nil
		}

		if todos[id-1].Done == false {
			todos[id-1].Done = true
			fmt.Printf(" Done %v\n", todos[id-1].Todo)
		} else {
			todos[id-1].Done = false
			fmt.Printf(" UnDone %v\n", todos[id-1].Todo)
		}
	}

	writeList(path, todos)

	return nil
}
