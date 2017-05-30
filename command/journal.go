package command

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/urfave/cli"
	gl_rpc "gitlab.bearstech.com/factory/gitlab-authenticated-rpc/rpc"
	"gitlab.bearstech.com/factory/journaleux/rpc"
	"io"
	"log"
	"regexp"
	"time"
)

func (c *Client) Journal(_cli *cli.Context) error {
	if _cli.Int("lines") > 0 && _cli.Bool("follow") {
		return errors.New("You can't both reads n lines, and follows")
	}
	if _cli.String("regexp") != "" {
		_, err := regexp.Compile(_cli.String("regexp"))
		if err != nil {
			return err
		}
	}
	err := c.SetDomain(_cli.GlobalString("domain"))
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
		Lines:   int32(_cli.Int("lines")),
		Follow:  _cli.Bool("follow"),
		Regexp:  _cli.String("regexp"),
	})
	if err != nil {
		return err
	}

	for {

		event, err := tail.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Fatalf("Trouble with this event: %v", err)
		}

		m := event.GetRealtime()
		s := m / 1000000
		t := time.Unix(int64(s), 0)
		fmt.Printf("%s !%d %s ", t.Format(time.RFC3339),
			event.GetPriority(),
			event.Fields["CONTAINER_NAME"])
		fmt.Println(" ", event.GetMessage())

	}

	return nil

}
