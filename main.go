package main

import (
	"os"

	"github.com/rixycf/gtd/command"
	"github.com/urfave/cli"
)

const (
	Version = "1.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "gtd"
	app.Version = Version
	app.Usage = "make todo list"
	app.Author = "riki kaida"
	app.Email = "kasnake1013@gmail.com"

	app.Commands = []cli.Command{
		cmd.Add,
		cmd.List,
		cmd.Done,
		cmd.Remove,
	}

	app.Run(os.Args)
}
