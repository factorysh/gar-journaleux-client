package command

import (
	_command "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/client/command"
	"google.golang.org/grpc"
)

type Client struct {
	_command.Client
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{*_command.NewClient(conn)}
}
