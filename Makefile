env_dir := $(shell go env GOPATH)

gob_dir := $(env_dir)/bin

main_target := $(gob_dir)/wwweb

sources := $(shell ls *.go | egrep -v '_test.go')


$(main_target): main/wwweb.go $(sources)
	go build -o $@ $<

table.go: doc/table_generate.go doc/source_table.txt
	go run $<

page.go: doc/page_generate.go doc/page.svg doc/catalog.svg 
	go run $<

clean:
	$(RM) $(main_target)
