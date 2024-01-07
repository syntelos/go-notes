gob_dir := $(shell go env GOPATH)
tgt_dir := $(gob_dir)/bin
sources := $(shell ls *.go) 

$(tgt_dir)/notes: cmd/notes/main.go $(sources)
	go build -o $@ $<

clean:
	$(RM) $(shell find tst -type f -name '*.svg')
	$(RM) $(shell find tst -type f -name '*.json')

test: clean
	go test
