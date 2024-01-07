gob_dir := $(shell go env GOPATH)
tgt_dir := $(gob_dir)/bin
sources := $(shell ls *.go) 

$(tgt_dir)/notes: cmd/notes/main.go $(sources)
	go build -o $@ $<

