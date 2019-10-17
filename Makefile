# CHANGED BY CHRISTIAN ALEXANDER BAHRDT
# This file is derivative from the following file 
# https://github.com/edgexfoundry/device-sdk-go/blob/edinburgh/Makefile

.PHONY: run test clean
.PHONY: build build-engine build-service
.PHONY: docker docker-device-bitflow docker-engine

GO=CGO_ENABLED=0 GO111MODULE=on go

SERVICE=cmd/device-bitflow/device-bitflow
ENGINE=cmd/engine/engine
.PHONY: $(SERVICE)
.PHONY: $(ENGINE)

VERSION=$(shell cat ./VERSION)

GOFLAGS=-ldflags "-X github.com/datenente/device-bitflow.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(SERVICE) $(ENGINE)
	$(GO) install -tags=safe

build-engine: $(ENGINE)
	$(GO) install -tags=safe

build-service: $(SERVICE)
	$(GO) install -tags=safe

cmd/device-bitflow/device-bitflow:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/device-bitflow

cmd/engine/engine:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/engine

docker: docker-device-bitflow docker-engine
	echo "docker..."

docker-device-bitflow:
	docker build \
		-f ./Dockerfile_Device_Service \
		--label "git_sha=$(GIT_SHA)" \
		-t datenente/device-bitflow:$(GIT_SHA) \
		-t datenente/device-bitflow:$(VERSION) \
		.

docker-engine:
	docker build \
		-f ./Dockerfile_Script_Execution_Engine \
		--label "git_sha=$(GIT_SHA)" \
		-t datenente/device-bitflow-engine:$(GIT_SHA) \
		-t datenente/device-bitflow-engine:$(VERSION) \
        .

run:
	cd cmd/device-bitflow ; ./device-bitflow --registry=consul://localhost:8500 --confdir=./res/

test:
	$(GO) vet ./...
	gofmt -l .
	$(GO) test -tags unit -coverprofile=coverage.out ./...
	$(GO) test -tags package -coverprofile=coverage.out ./...

clean:
	rm -f $(SERVICE) $(SERVICE).log
	rm -f $(ENGINE)
