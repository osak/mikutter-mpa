GO_SRCS=$(wildcard src/**/*.go)

all: mpa

mpa: $(GO_SRCS)
	GOPATH=`pwd` go install mpa
