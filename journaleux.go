package main

import (
	_cli "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/cli"
	//"sort"

	"os"

	_command "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
	"gitlab.bearstech.com/factory/journaleux/client/command"
)

var (
	git_version = ""
)

func main() {

	app := _cli.NewApp()
	app.Name = "Journaleux"
	app.Version = git_version
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
