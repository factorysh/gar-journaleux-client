client: bin vendor
	go build -ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/journaleux github.com/factorysh/gar-journaleux-client

vendor:
	dep ensure

bin:
	mkdir bin

clean:
	rm -rf bin
	rm -rf vendor
	rm -rf gopath

protoc:
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	protoc journaleux.proto --go_out=plugins=grpc:rpc_journal

client-all: client-linux client-darwin

client-linux: bin vendor
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/journaleux_linux_amd64 github.com/factorysh/gar-journaleux-client

client-darwin: bin vendor
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/journaleux_darwin_amd64 github.com/factorysh/gar-journaleux-client