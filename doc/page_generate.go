/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type Text []byte
type Page []Text

func (this Text) IsText() (bool) {
	var found bool = false
	var first, last int = 0, 0

	for ix, ch := range this {

		switch ch {
		case '<':
			if 0 == first {
				first = ix
			}
		case '>':
			if 0 == last {
				last = ix

				break
			}
		default:
			if 0 < first {
				if ix == (first+1) {

					found = ('t' == ch)

				} else if ix == (first+2) {

					if found {

						found = ('e' == ch)
					}
				} else if ix == (first+3) {

					if found {

						found = ('x' == ch)
					}
				} else if ix == (first+4) {

					if found {

						found = ('t' == ch)
					}
				} else if ix == (first+5) {

					if found {

						found = ((' ' == ch)||('>' == ch))
					}
				}
			}
		}
	}

	if 0 < first && first < last {

		return found
	} else {
		return false
	}
}

func open(base string) (fi *os.File) {
	var er error
	fi, er = os.Open(base)
	if nil != er {
		fi, er = os.Open(path.Join("doc",base))
		if nil != er {
			return nil
		}
	}
	return fi
}

func create() (fi *os.File) {
	var er error
	fi, er = os.Create("page.go")
	if nil != er {
		return nil
	} else {
		return fi
	}
}

func PageRead() (this Page) {
	var src *os.File = open("page.svg")
	if nil != src {
		var reader *bufio.Reader = bufio.NewReader(src)
		var inl []byte
		var isp bool
		var er error
		inl, isp, er = reader.ReadLine()
		for true {
			if nil != er {
				if io.EOF == er {
					return this
				} else {
					return Page{}
				}
			} else if isp {
				return Page{}
			} else {
				var line Text = inl
				this = append(this,line)
				inl, isp, er = reader.ReadLine()
			}
		}
		return this
	} else {
		return Page{}
	}
}

func (this Page) Bar() (bar int) {
	bar = -1

	for tix, txt := range this {

		if txt.IsText() {

			bar = tix
			break
		}
	}
	return bar
}

func CatalogRead() (this Page) {
	var src *os.File = open("page.svg")
	if nil != src {
		var reader *bufio.Reader = bufio.NewReader(src)
		var inl []byte
		var isp bool
		var er error
		inl, isp, er = reader.ReadLine()
		for true {
			if nil != er {
				if io.EOF == er {
					return this
				} else {
					return Page{}
				}
			} else if isp {
				return Page{}
			} else {
				var line Text = inl
				this = append(this,line)
				inl, isp, er = reader.ReadLine()
			}
		}
		return this
	} else {
		return Page{}
	}
}

func main(){

	var page Page = PageRead()
	var page_bar = page.Bar()
	var catalog Page = CatalogRead()
	var catalog_bar = catalog.Bar()

	if 0 < page_bar && 0 < catalog_bar {
		var file *os.File = create()
		if nil != file {
			var cc int

			fmt.Fprintln(file,`/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"bufio"
	"bytes"
)

type Text []byte

type Page []Text

`)
			cc = 0
			fmt.Fprint(file,"var page_head Page = Page{")
			for tix, txt := range page {

				if tix < page_bar {
					if 0 != cc {
						fmt.Fprint(file,",")
					}
					fmt.Fprint(file," Text{")
					for cx, ch := range txt {
						if 0 != cx {
							fmt.Fprint(file,",")
						}
						fmt.Fprintf(file," 0x%02X",ch)
					}
					fmt.Fprint(file," }")
					cc += 1
				}
			}
			fmt.Fprintln(file,"}\n")

			cc = 0
			fmt.Fprint(file,"var page_tail Page = Page{")
			for tix, txt := range page {

				if page_bar < tix {
					if 0 != cc {
						fmt.Fprint(file,",")
					}
					fmt.Fprint(file," Text{")
					for cx, ch := range txt {
						if 0 != cx {
							fmt.Fprint(file,",")
						}
						fmt.Fprintf(file," 0x%02X",ch)
					}
					fmt.Fprint(file," }")
					cc += 1
				}
			}
			fmt.Fprintln(file,"}\n")

			cc = 0
			fmt.Fprint(file,"var catalog_head Page = Page{")
			for tix, txt := range catalog {

				if tix < catalog_bar {
					if 0 != cc {
						fmt.Fprint(file,",")
					}
					fmt.Fprint(file," Text{")
					for cx, ch := range txt {
						if 0 != cx {
							fmt.Fprint(file,",")
						}
						fmt.Fprintf(file," 0x%02X",ch)
					}
					fmt.Fprint(file," }")
					cc += 1
				}
			}
			fmt.Fprintln(file,"}\n")

			cc = 0
			fmt.Fprint(file,"var catalog_tail Page = Page{")
			for tix, txt := range catalog {
				
				if catalog_bar < tix {
					if 0 != cc {
						fmt.Fprint(file,",")
					}
					fmt.Fprint(file," Text{")
					for cx, ch := range txt {
						if 0 != cx {
							fmt.Fprint(file,",")
						}
						fmt.Fprintf(file," 0x%02X",ch)
					}
					fmt.Fprint(file," }")
					cc += 1
				}
			}

			fmt.Fprint(file,`}

func (this Page) Encode() []byte {
        var w bytes.Buffer
	for _, txt := range this {
		w.Write(txt)
		w.WriteByte('\n')
	}
	return w.Bytes()
}

func (this Page) Decode(src []byte) {

	var buffer *bytes.Buffer = bytes.NewBuffer(src)
	var reader *bufio.Reader = bufio.NewReaderSize(buffer,len(src))

	var inl []byte
	var isp bool
	var er error

	inl, isp, er = reader.ReadLine()
	for nil != er || isp {

		var line Text = inl
		this = append(this,line)
		inl, isp, er = reader.ReadLine()
	}
}
`)

			file.Close()
			os.Exit(0)
		} else {
			log.Fatal("Write failed.")
			os.Exit(1)
		}
	} else {
		log.Fatal("Read failed.")
		os.Exit(1)
	}
}
