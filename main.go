package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var GlobalFlags = []cli.Flag{}
var Commands = []cli.Command{}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

func validate(u string, s string, e string) {
	if len(u) == 0 || len(s) == 0 || len(e) == 0 {
		fmt.Printf("Invalid arguments")
		os.Exit(1)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Takeshi Kondo"
	app.Email = "take.she12@gmail.com"
	app.Usage = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Usage: "Specify github user name",
		},
		cli.StringFlag{
			Name:  "start, s",
			Usage: "Specify start date. format is YYMMDD",
		},
		cli.StringFlag{
			Name:  "end, e",
			Usage: "Specify end date. format is YYMMDD",
		},
	}
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Action = func(c *cli.Context) error {

		/* parse arguments */
		user := c.String("user")
		startdate := c.String("start")
		enddate := c.String("end")
		validate(user, startdate, enddate)

		fmt.Printf("Hello, gh-contrib\n")
		return nil
	}
	app.Run(os.Args)
}
