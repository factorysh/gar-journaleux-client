package main

import (
	"github.com/urfave/cli"
	//"sort"

	"gitlab.bearstech.com/factory/journaleux/client/command"

	"os"
)

var (
	git_version = ""
)

func main() {

	app := cli.NewApp()
	app.Name = "Journaleux"
	app.Version = git_version
	app.EnableBashCompletion = true
	app.Usage = "Client CLI application for journaleuxd"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "domain, d",
			Value: "rpc.example.com",
			Usage: "Target RPC server",
		},
		cli.StringFlag{
			Name:  "server, s",
			Value: "servers.example",
			Usage: "Target RPC server from .journaleux.toml",
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
			Name:    "environments",
			Aliases: []string{"e"},
			Usage:   "Get your environments for a project",
			Action:  cmd.Environments,
		},
		{
			Name:    "journal",
			Aliases: []string{"j"},
			Usage:   "Read journal `NAME`",
			Action:  cmd.Journal,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "lines, l",
					Value: -10,
					Usage: "Number of lines",
				},
				cli.BoolFlag{
					Name:  "follow, f",
					Usage: "Follow logs",
				},
				cli.StringFlag{
					Name:  "regexp, r",
					Usage: "Filter by regexp, using RE2 syntax https://github.com/google/re2/wiki/Syntax",
				},
			},
		},
	}

	//sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
