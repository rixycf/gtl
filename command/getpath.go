package cmd

import (
	"os"
	"path/filepath"
)

func getPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "gtl_todo.json")
	return path, nil
}
