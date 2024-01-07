env_dir := $(shell go env GOPATH)
gob_dir := $(env_dir)/bin
sources := $(shell ls *.go) 

$(gob_dir)/notes: cmd/notes/main.go $(sources)
	go build -o $@ $<

clean:
	$(RM) $(shell find tst -type f -name '*.svg')
	$(RM) $(shell find tst -type f -name '*.json')

test: clean
	go test
