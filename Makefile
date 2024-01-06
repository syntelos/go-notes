godir := $(shell go env GOPATH)
tgtdir := $(godir)/bin
sources := $(shell ls *.go) 

all: $(tgtdir)/encode $(tgtdir)/update

$(tgtdir)/encode: cmd/encode/main.go $(sources)
	go build -o $@ $<

$(tgtdir)/update: cmd/update/main.go $(sources)
	go build -o $@ $<

