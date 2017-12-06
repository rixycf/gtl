package cmd

import (
	"fmt"
	"strconv"

	"github.com/rixycf/gtl/todo"
	"github.com/urfave/cli"
)

// Remove define "remove" subcommand
var Remove = cli.Command{
	Name:   "remove",
	Usage:  "remove todo from List",
	Action: remove,
}

func remove(c *cli.Context) error {

	// get id to remove todo
	ids := c.Args()
	path, err := getPath()
	if err != nil {
		return err
	}
	// read todoList
	todos := readList(path)

	// remove todo
	for i, id := range ids {
		id, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if id < 1 || id > len(todos)-i {
			fmt.Println(" wrong id ")
			return nil
		}

		fmt.Println("remove task: ", todos[id-1].Todo)
		todos = rmSlice(id-i-1, todos)
	}

	for i := 0; i < len(todos); i++ {
		todos[i].Id = i + 1
	}
	//write todoList
	writeList(path, todos)
	return nil
}

// rmSlice remove slice
func rmSlice(id int, td []todo.List) []todo.List {
	td = append(td[:id], td[id+1:]...)

	newtd := make([]todo.List, len(td))
	copy(newtd, td)
	return newtd
}
