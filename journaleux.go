package main

import (
	"fmt"
	"gitlab.bearstech.com/bearstech/journaleux/rpc"
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
	domain := os.Args[1]

	conn, err := client.NewConn(domain)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	j := rpc.NewJournaleuxClient(conn)

	ctx := context.Background()

	tail, err := j.Tail(ctx, &rpc.Predicate{
		Project: os.Args[2],
	})

	if err != nil {
		log.Fatalf("Can't tail: %v", err)
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
}
