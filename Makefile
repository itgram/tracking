CUR_DIR = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

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
check: go build -gcflags="-m=2 -l"

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
