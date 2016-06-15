package main

import (
	"fmt"
	"os"

	"github.com/bluedaniel/gotube/command"
	"github.com/urfave/cli"
)

// GlobalFlags is an array of flags
var GlobalFlags = []cli.Flag{}

// Commands are all the apps commands
var Commands = []cli.Command{
	{
		Name:   "status",
		Usage:  "",
		Action: command.CmdStatus,
		Flags:  []cli.Flag{},
	}, {
		Name:   "station",
		Usage:  "",
		Action: command.CmdStation,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound prints usage info
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
