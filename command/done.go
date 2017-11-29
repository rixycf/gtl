package cmd

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

var Done = cli.Command{
	Name:   "done",
	Usage:  "check todo",
	Action: done,
}

// done command check todo list
func done(c *cli.Context) error {

	// get id to check TODOList
	ids := c.Args()

	// read todolist
	todos := readList("./test.json")
	//fmt.Println(todos)

	// check todoList
	for _, id := range ids {
		id, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if todos[id].Done == false {
			todos[id].Done = true
			fmt.Printf(" Done %v\n", todos[id].Todo)
		} else {
			todos[id].Done = false
			fmt.Printf(" UnDone %v\n", todos[id].Todo)
		}
	}

	writeList("./test.json", todos)

	return nil
}
