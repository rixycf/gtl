package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rixycf/gtd/todo"
	"github.com/urfave/cli"
)

var Add = cli.Command{
	Name:   "add",
	Usage:  "add todo",
	Action: add,
}

func add(c *cli.Context) error {
	t := todo.List{}
	t.Todo = "hello"
	fmt.Println(t.Todo)

	jsonBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Json Marshal error: ", err)
	}

	fmt.Printf("%v\n", jsonBytes)
	//TODO Create file and write json to file
	//TODO get filepath
	path, err := filepath.Abs(".")
	if err != nil {
		fmt.Println(path)
	}

	file, err := os.Create(path + "/test.json")
	if err != nil {
		fmt.Println("file create error: ", err)
	}
	defer file.Close()
	file.Write(jsonBytes)

	fmt.Println(string(jsonBytes))

	return nil
}
