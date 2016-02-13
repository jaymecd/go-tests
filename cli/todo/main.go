package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.HideHelp = true
	app.EnableBashCompletion = true

	app.Author = "jaymecd"
	app.Email = "www@qqq.ee"
	app.Usage = "Make an explosive entrance"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
