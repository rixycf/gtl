package cmd

import (
	"strconv"

	"github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

var Remove = cli.Command{
	Name:   "remove",
	Usage:  "remove todo from List",
	Action: remove,
}

func remove(c *cli.Context) error {

	// get id to remove todo
	ids := c.Args()

	// read todoList
	todos := readList("./test.json")

	// remove todo
	// for _, id := range ids {
	// 	id, err := strconv.Atoi(id)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	todos = rmSlice(id, todos)
	// }

	for _, todo := range todos {
		for _, id := range ids {
			id, err := strconv.Atoi(id)
			if err != nil {
				return err
			}

			if todo.Id == id {
				todos = rmSlice(id, todos)
			}
		}
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
