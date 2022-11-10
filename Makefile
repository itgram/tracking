CUR_DIR = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ESDB_DIR := $(shell go list -m -f "{{.Dir}}" github.com/EventStore/EventStore-Client-Go/v3)

$(if $(shell go version),,$(error "warning: golang not exists in your path, this will probably not work"))

########################################################################################
# targets section

build: deps; @go run "$(CUR_DIR)Build.go" "$@" "$(target)" "$(os)" "$(arch)" "$(args)"

clean: deps; @go run "$(CUR_DIR)Build.go" "$@" "$(target)"

test: deps; @go run "$(CUR_DIR)Build.go" "$@" "$(module)"

package: deps
	@$(MAKE) build target="$(target)" os="linux" arch="amd64"
	@go run "$(CUR_DIR)Build.go" "$@" "$(target)"

deps:; -@cd src/ && go get -d -v

# check for stack/heap escape
# check: go build -gcflags="-m -l"
check:; go build -gcflags="-m=2 -l"

# fix the "proto: file "gossip.proto" is already registered" error
# using "go-proto-filename-prefixer"
# go install github.com/tuimeo/go-proto-filename-prefixer@latest
fix:; go-proto-filename-prefixer $(ESDB_DIR)/protos esdb.

help:; @printf "%s\n" \
	"application makefile help" \
	"" \
	"Usage: make [target]..." \
	"" \
	"targets:" \
	"  build         Compile the project" \
	"  clean         Delete temporary and output files" \
	"  test          Run the test for this project" \
	"  package       Creating Debian/Ubuntu .deb package" \
	"  deps          Download the dependencies" \
	"  help          Display this help and exit" \


.PHONY: build clean test package deps help


# # Makefile for go-template

# REPO:=gbaeke
# TAG:=latest
# IMAGE:=$(REPO)/go-template:$(TAG)


# test:
# 	go test -v -race ./...

# build:
# 	CGO_ENABLED=0 go build -installsuffix 'static' -o app cmd/app/*

# docker-build:
# 	docker build -t $(IMAGE) .

# docker-push:
# 	docker build -t $(IMAGE) .
# 	docker push $(IMAGE)

# swagger:
# 	cd pkg/api ; swag init -g server.go
