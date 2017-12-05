package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/rixycf/gtl/todo"
	"io/ioutil"
	"os"
)

// readList read json file and decode it
func readList(path string) []todo.List {

	c := checkFile(path)

	// if file already exist, read file
	if c == true {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Read File error: ", err)
		}

		// unmarshal json file
		var t []todo.List
		if err := json.Unmarshal(file, &t); err != nil {
			fmt.Println("json Unmarshal error: ", err)
		}
		return t

	} else {
		t := []todo.List{}
		return t
	}
}

// writeList encode struct and write to json file
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

func checkFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
