client: bin vendor
	go build -ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/journaleux github.com/factorysh/gar-journaleux-client

vendor:
	dep ensure

bin:
	mkdir bin

bin/darwin:
	mkdir -p bin/darwin

bin/linux_amd64:
	mkdir -p bin/linux_amd64

clean:
	rm -rf bin
	rm -rf vendor
	rm -rf gopath

pull:
	docker pull bearstech/golang-dep
	docker pull bearstech/upx

protoc:
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	protoc journaleux.proto --go_out=plugins=grpc:rpc_journal

client-all: client-linux client-darwin

client-linux: bin/linux_amd64 vendor
	GOOS=linux GOARCH=amd64 go build \
		-ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/linux_amd64/journaleux \
		github.com/factorysh/gar-journaleux-client

client-darwin: bin/darwin vendor
	GOOS=darwin GOARCH=amd64 go build \
		-ldflags "-X main.gitVersion=`git rev-parse HEAD`" \
		-o bin/darwin/journaleux \
		github.com/factorysh/gar-journaleux-client

upx-client-linux: client-linux
	upx bin/linux_amd64/journaleux

upx-client-darwin: client-darwin
	upx bin/darwin/journaleux

upx-client: client
	upx bin/journaleux

docker-client-linux:
	docker run -t --rm \
		-u `id -u` \
		-v `pwd`:/go/src/github.com/factorysh/gar-journaleux-client/ \
		-v `pwd`/.cache:/.cache \
		-w /go/src/github.com/factorysh/gar-journaleux-client/ \
		bearstech/golang-dep \
		make client-linux
	docker run -t --rm \
		-u `id -u` \
		-v `pwd`:/upx \
		bearstech/upx \
		upx bin/linux_amd64/journaleux

docker-client-darwin:
	docker run -t --rm \
		-u `id -u` \
		-v `pwd`:/go/src/github.com/factorysh/gar-journaleux-client/ \
		-v `pwd`/.cache:/.cache \
		-w /go/src/github.com/factorysh/gar-journaleux-client/ \
		bearstech/golang-dep \
		make client-darwin
	docker run -t --rm \
		-u `id -u` \
		-v `pwd`:/upx \
		bearstech/upx \
		upx bin/darwin/journaleux

docker-clients: docker-client-linux docker-client-darwin