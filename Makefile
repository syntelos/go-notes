env_dir := $(shell go env GOPATH)
gob_dir := $(env_dir)/bin
sources := $(shell find . -type f -name "*.go" | egrep -v '(main/wwweb|_test.go)')
targets := $(shell find tst -type f -name '*.svg' -o -name '*.json')
target := $(gob_dir)/wwweb

$(target): main/wwweb.go $(sources)
	go build -o $@ $<

table.go: doc/source_table.txt doc/table_generate.go
	go run $< $@ 

clean:
	$(RM) $(target)

test: clean
	$(RM) $(targets)
	go test
