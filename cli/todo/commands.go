package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jaymecd/go-test/cli/todo/command"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "ENV_DEBUG",
		Name:   "debug",
	},
}

var Commands = []cli.Command{
	{
		Name:   "template",
		Aliases:     []string{"tpl"},
		Usage:  "template related commands",
//		Action: command.CmdAdd,
		Flags:  []cli.Flag{},
		Subcommands: []cli.Command{
	    {
	      Name:  "add",
	      Usage: "add a new template",
				Flags:  []cli.Flag{},
	      Action: func(c *cli.Context) {
	        println("new task template: ", c.Args().First())
	      },
	    },
	    {
	      Name:  "remove",
	      Usage: "remove an existing template",
				Flags:  []cli.Flag{},
	      Action: func(c *cli.Context) {
	        println("removed task template: ", c.Args().First())
	      },
	    },
	  },
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "delete",
		Usage:  "",
		Action: command.CmdDelete,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
