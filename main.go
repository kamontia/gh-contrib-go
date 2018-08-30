package main

import (
	"fmt"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"regexp"
)

var GlobalFlags = []cli.Flag{}
var Commands = []cli.Command{}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
func existUser(u string) bool {
	url := "http://github.com/" + u

	response, err := http.Get(url)
	if err != nil {
		os.Exit(2)
	}

	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func validate(u string, s string, e string) {
	r := regexp.MustCompile(`^(\d{6})$`)

	if len(u) == 0 || len(s) == 0 || len(e) == 0 {
		fmt.Printf("-u, -s, -e is required.]\n")
		os.Exit(1)
	} else if !r.MatchString(s) || !r.MatchString(e) {
		fmt.Printf("Invalid arguments --start %s --end %s\n", s, e)
		os.Exit(1)
	} else if !existUser(u) {
		fmt.Printf("User %s is not found\n", u)
		os.Exit(1)
	} else {
		fmt.Printf("Validation is OK!\n")
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

		/* search api */

    /* parse json */

    /* put std.out */

    return nil
	}
	app.Run(os.Args)
}
