package cmd

import (
	"fmt"
	"strconv"

	"github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

// Remove define "remove" subcommand
var Remove = cli.Command{
	Name:   "remove",
	Usage:  "remove todo from List",
	Action: remove,
}

func remove(c *cli.Context) error {

	//TODO fix!
	//TODO remove todo function

	// get id to remove todo
	ids := c.Args()
	fmt.Printf("%+v", ids)

	// read todoList
	todos := readList("./test.json")

	// remove todo
	for i, id := range ids {
		id, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		todos = rmSlice(id-i-1, todos)
	}

	for i := 0; i < len(todos); i++ {
		todos[i].Id = i + 1
	}
	//write todoList
	writeList("./test.json", todos)
	return nil
}

// rmSlice remove slice
func rmSlice(id int, td []todo.List) []todo.List {
	td = append(td[:id], td[id+1:]...)

	newtd := make([]todo.List, len(td))
	copy(newtd, td)
	return newtd
}
