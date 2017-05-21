package main

import (
	"log"
	"os"

	"gitlab.bearstech.com/bearstech/journaleux/rpc"
	"golang.org/x/net/context"

	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

var (
	git_version = ""
)

func main() {
	go internalServer.ListenAndServe()

	domain := os.Args[1]

	conf := &Conf{domain}
	err := conf.makeDomainHome()
	if err != nil {
		log.Fatalf("Can't build local conf folder: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(domain+port,
		grpc.WithPerRPCCredentials(&DummyAuth{}),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpc.NewJournaleuxClient(conn)
	h := rpc.NewHelloServiceClient(conn)

	ctx := context.Background()

	hello, err := h.SayHello(ctx, &rpc.HelloRequest{os.Args[2]})
	if err != nil {
		log.Fatalf("Can't hello: %v", err)
	}
	log.Println(hello)

	tail, err := c.Tail(ctx, &rpc.Predicate{
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
