package command

import (
	"errors"

	"github.com/urfave/cli"
	_cli "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/cli"
	"gitlab.bearstech.com/factory/journaleux/client/conf"
)

type JournalClient struct {
	Client *_cli.Client
	config *conf.ClientConfig
}

func NewJournalClient(client *_cli.Client) (*JournalClient, error) {

	j := &JournalClient{
		Client: client,
		config: new(conf.ClientConfig),
	}
	// read from config file
	err := conf.ReadClientConf(j.config, ".journaleux.toml")
	if err != nil {
		return nil, errors.New("Error reading client config file")
	}
	return j, nil
}

func (j *JournalClient) Register(app *cli.App) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    "journal",
		Aliases: []string{"j"},
		Usage:   "Read journal `NAME`",
		Action:  j.Journal,
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
	})
}
