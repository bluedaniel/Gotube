package main

import (
	"os"

	"github.com/bluedaniel/gotube/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tube"
	app.Version = "0.1.0"
	app.Author = "bluedaniel"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	if len(os.Args) == 1 {
		app.Action = command.CmdStatus
	} else {
		app.Action = command.CmdStation
	}

	app.Run(os.Args)
}
