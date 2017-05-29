package main

import (
	"github.com/urfave/cli"
	//"sort"

	"gitlab.bearstech.com/bearstech/journaleux/client/command"

	"os"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {

	app := cli.NewApp()
	app.Name = "Journaleux"
	app.Version = git_version
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Value: "rpc.example.com",
			Usage: "Target RPC server",
		},
	}

	cmd := command.NewClient()
	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Get yourself",
			Action:  cmd.User,
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "Get your projects",
			Action:  cmd.Projects,
		},
		{
			Name:    "journal",
			Aliases: []string{"j"},
			Usage:   "Read journal `NAME`",
			Action:  cmd.Journal,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "lines, l",
					Value: 0,
					Usage: "Number of lines",
				},
				cli.BoolFlag{
					Name:  "Follow, f",
					Usage: "Follow logs",
				},
			},
		},
	}

	//sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
