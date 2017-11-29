package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rixycf/gtd/todo"
)

func readList(path string) []todo.List {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	// unmarshal json file
	var t []todo.List
	if err := json.Unmarshal(file, &t); err != nil {
		fmt.Println(err)
	}
	return t
}

func writeList(path string, list []todo.List) {
	// encode List slice
	td, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json marshal error ", err)
	}

	// write todo.List to file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("file crate error: ", err)
	}
	defer file.Close()

	file.Write(td)
}
