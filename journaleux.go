package main

import (
	"github.com/urfave/cli"

	"fmt"
	"gitlab.bearstech.com/bearstech/journaleux/client/command"

	"gitlab.bearstech.com/bearstech/journaleux/rpc"
	gl_rpc "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/client"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {
	var domain string

	app := cli.NewApp()
	app.Name = "Journaleux"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "domain, d",
			Value:       "rpc.example.com",
			Usage:       "Target RPC server",
			Destination: &domain,
		},
	}

	conn, err := client.NewConn(domain)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cmd := command.NewClient(conn)
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
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Version release",
			Action:  cmd.Version,
		},
		{
			Name:    "journal",
			Aliases: []string{"j"},
			Usage:   "Journal",
			Action:  cmd.Journal,
		},
	}

	app.Run(os.Args)

}
