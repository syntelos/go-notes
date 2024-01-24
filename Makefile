env_dir := $(shell go env GOPATH)
gob_dir := $(env_dir)/bin
sources := $(shell find . -type f -name "*.go")
targets := $(shell find tst -type f -name '*.svg' -o -name '*.json')

$(gob_dir)/wwweb: main/wwweb.go $(sources)
	go build -o $@ $<

clean:
	$(RM) $(targets)

test: clean
	go test
