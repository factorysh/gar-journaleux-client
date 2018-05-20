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

client-linux: bin vendor
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/journaleux_linux_amd64 gitlab.bearstech.com/factory/journaleux/client

client-darwin: bin vendor
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/journaleux_darwin_amd64 gitlab.bearstech.com/factory/journaleux/client
