package command

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
	"gitlab.bearstech.com/bearstech/journaleux/rpc"
	gl_rpc "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
	"log"
	"time"
)

func (c *Client) Journal(_cli *cli.Context) error {
	err := c.SetDomain(_cli.String("domain"))
	if err != nil {
		return err
	}
	j := rpc.NewJournaleuxClient(c.Conn)
	gl := gl_rpc.NewGitlabClient(c.Conn)
	_, err = gl.Ping(c.Ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	tail, err := j.Tail(c.Ctx, &rpc.Predicate{
		Project: _cli.Args().First(),
	})
	if err != nil {
		return err
	}

	for {

		event, err := tail.Recv()
		if err != nil {
			log.Fatalf("Trouble with this event: %v", err)
		}

		m := event.GetRealtime()
		s := m / 1000000
		t := time.Unix(int64(s), 0)
		fmt.Printf("%s !%d %s ", t.Format(time.RFC3339),
			event.GetPriority(),
			event.GetContainerName())
		fmt.Println(" ", event.GetMessage())

	}

	return nil

}
