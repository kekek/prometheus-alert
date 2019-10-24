## If the first argument is "run"...
# ifeq (run,$(firstword $(MAKECMDGOALS)))
##   use the rest as arguments for "run"
#  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
#  # ...and turn them into do-nothing targets
#  $(eval $(RUN_ARGS):;@:)
# endif


## color output
NO_COLOR=\x1b[0m
OK_COLOR=\x1b[32;01m
ERROR_COLOR=\x1b[31;01m
WARN_COLOR=\x1b[33;01m

# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

# build version info
GIT_VERSION=$(shell git rev-parse --short HEAD)
DATE_VERSION=$(shell date +%Y%m%d%H%M)

PACKAGE=$(shell head -n 1 go.mod | cut -f 2 -d " ")

# Binary names
PROJECT = $(shell pwd | xargs basename)

BINARY_FILE_NAME=$(PROJECT)
#BINARY_UNIX_NAME=$(PROJECT)_linux
#VERSION_BINARY_UNIX_NAME=$(BINARY_UNIX)-$(GIT_VERSION)-$(DATE_VERSION)

BIN_DIR=bin

### build编译时传入version参数
IDFLAG_DSC=-ldflags "-X main.gitVersion=${GIT_VERSION} -X main.dateVersion=${DATE_VERSION} "
#IDFLAG_DSC=-ldflags "-X ${PACKAGE}/main.gitVersion=${GIT_VERSION} -X  ${PACKAGE}/main.dateVersion=${DATE_VERSION} "
#$package/config.GitLog=$GITLOGVERSION  -X $package/config.DateVersion=$NOW

.PHONY : build clean build-exec run run-exec

build: clean build-exec

run: run-exec

build-exec:
	if [ ! -d $(BIN_DIR) ]; then mkdir $(BIN_DIR) ; fi
	$(GOBUILD) ${IDFLAG_DSC} -v -o $(BIN_DIR)/$(BINARY_FILE_NAME) main.go


run-exec:
	$(BIN_DIR)/$(BINARY_FILE_NAME)


clean:
	rm -rf $(BIN_DIR)/

color:

	@echo "${OK_COLOR} ok ${NO_COLOR}"
	@echo "${ERROR_COLOR} error ${NO_COLOR}"
	@echo "${WARN_COLOR} warn ${NO_COLOR}"



