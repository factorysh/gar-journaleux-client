package main

import (
	_cli "github.com/factorysh/gitlab-authenticated-rpc-client/cli"

	"os"

	_command "github.com/factorysh/gitlab-authenticated-rpc-client/command"

	"github.com/factorysh/gar-journaleux-client/command"
)

var (
	gitVersion = ""
)

func main() {

	app := _cli.NewApp()
	app.Name = "Journaleux"
	app.Version = gitVersion
	app.Usage = "Client CLI application for journaleuxd"

	cmd := _cli.NewClient()
	jj, err := command.NewJournalClient(cmd)
	if err != nil {
		panic(err)
	}
	_cli.Register(app,
		_command.NewAuthClient(cmd),
		_command.NewGitlabClient(cmd),
		jj,
	)

	//sort.Sort(cli.FlagsByName(app.Flags))
	//sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}
