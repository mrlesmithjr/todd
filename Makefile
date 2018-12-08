SHELL=/bin/bash

all: compile

clean:
	rm -f $(GOPATH)/bin/todd-server
	rm -f $(GOPATH)/bin/todd
	rm -f $(GOPATH)/bin/todd-agent

build:
	rm -f api/v1/generated/*
	docker build -t toddproject/todd -f Dockerfile .

compile:

	# Generate API code
	rm -rf api/exp/generated/ && mkdir -p api/exp/generated/ && protoc -I api/exp/definitions/ api/exp/definitions/* --go_out=plugins=grpc:api/exp/generated/

	# Installing testlets
	./scripts/gettestlets.sh

	# Installing ToDD
	go install ./cmd/...

fmt:
	go fmt github.com/toddproject/todd/...

test: 
	go test ./... -cover
	scripts/start-containers.sh integration

lint:
	scripts/lint.sh

update_deps:
	go get -u github.com/tools/godep
	godep save ./...

update_assets:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -o assets/assets_unpack.go -pkg="assets" -prefix="agent" agent/testing/bashtestlets/... agent/facts/collectors/...

start: compile

	# This mode is just to get a demo of ToDD running within the VM quickly.
	# It made sense to re-use the configurations for integration testing, so
	# that's why "server-int.cfg" and "agent-int.cfg" are being used here.
	start-containers.sh 6 /etc/todd/server-int.cfg /etc/todd/agent-int.cfg

install:

	# Set capabilities on testlets
	./scripts/set-testlet-capabilities.sh

	# Copy configs if etc and /etc/todd aren't linked
	if ! [ "etc" -ef "/etc/todd" ]; then mkdir -p /etc/todd && cp -f ./etc/{agent,server}.cfg /etc/todd/; fi
	mkdir -p /opt/todd/{agent,server}/assets/{factcollectors,testlets}
	chmod -R 777 /opt/todd
