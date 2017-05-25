package command

import (
	_command "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
)

type Client struct {
	_command.Client
}

func NewClient() *Client {
	return &Client{}
}
