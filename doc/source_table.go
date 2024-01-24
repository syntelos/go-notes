/*
 * Source table
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
/*
 *
 */
type Record struct {
	name, path, link string
}

func tab(src []byte, ofs, len int) (idx int) {
	for ofs < len {

		if '\t' == src[ofs] {
			return ofs
		} else {
			ofs += 1
		}
	}
	return len
}

func (this Record) read(line []byte) (Record) {
	var begin, end, len int = 0, 0, len(line)

	end = tab(line,begin,len)
	if 0 < end {

		this.name = string(line[begin:end])

		begin = end+1
		end = tab(line,begin,len)
		if 0 < end {

			this.path = string(line[begin:end])

			begin = end+1
			end = tab(line,begin,len)
			if begin < end && 0 < end {

				this.link = string(line[begin:end])
			} else {
				fmt.Fprintf(os.Stderr,"Skipping '%s'.\n",string(line))
			}
		}
	}
	return this
}
/*
 *
 */
type Table []Record

func (this Table) read(file *os.File) (that Table, e error) {
	var reader *bufio.Reader = bufio.NewReader(file)

	var inl []byte
	var isp bool = false
	var rec Record

	inl, isp, e = reader.ReadLine()

	for true {
		if nil != e {
			if io.EOF == e {

				return this, nil
			} else {
				return this, fmt.Errorf("Error reading '%s': %w",file.Name(),e)
			}
		} else if isp {
			return this, fmt.Errorf("Error reading '%s'.",file.Name())
		} else {

			this = append(this,rec.read(inl))

			inl, isp, e = reader.ReadLine()
		}
	}
	return this, nil
}

func camel(src_str string) (string) {
	var dst strings.Builder
	var src []byte = []byte(src_str)
	var src_len int = len(src)

	var w bool = true
	var x, z int = 0, src_len
	var y byte

	for ; x < z; x++ {
		y = src[x]
		if w {
			if 'a' <= y && 'z' >= y {
				y -= 'a'
				y += 'A'

				dst.WriteByte(y)

				w = false
			} else {
				dst.WriteByte(y)
			}
		} else if '-' == y || '_' == y {
			w = true
		} else if '0' <= y && '9' >= y {
			w = true
			dst.WriteByte(y)
		} else {
			dst.WriteByte(y)
		}
	}
	return dst.String()
}

func (this Table) List() {

	switch operand() {

	case "name":
		var cn string
		for _, re := range this {
			cn = camel(re.name)

			fmt.Printf("\tTableName%s\tTableName\t= \"%s\"\n",cn,re.name)
		}
	case "path":
		var cn string
		for _, re := range this {
			cn = camel(re.name)

			fmt.Printf("\tTablePath%s\tTablePath\t= \"%s\"\n",cn,re.path)
		}
	case "link":
		var cn string
		for _, re := range this {
			cn = camel(re.name)

			fmt.Printf("\tTableLink%s\tTableLink\t= \"%s\"\n",cn,re.link)
		}
	default:
		for ix, re := range this {

			fmt.Printf("%03o\t%s\t%s\t%s\n",(ix+1),re.name,re.path,re.link)
		}
	}
	os.Exit(0)
}

func (this Table) Enumerate() {

	switch operand() {

	case "path":
		var cn string
		for _, re := range this {
			cn = camel(re.name)

			fmt.Printf("\tcase TableName%s:\n\t\treturn TablePath%s\n",cn,cn)

		}
		os.Exit(0)
	case "link":
		var cn string
		for _, re := range this {
			cn = camel(re.name)

			fmt.Printf("\tcase TableName%s:\n\t\treturn TableLink%s\n",cn,cn)
		}
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
/*
 *
 */
func open() (fi *os.File) {
	var er error
	fi, er = os.Open("source_table.txt")
	if nil != er {
		fi, er = os.Open("doc/source_table.txt")
		if nil != er {
			return nil
		}
	}
	return fi
}

func operator() (opr string) {
	if 1 < len(os.Args) {
		return os.Args[1]
	} else {
		return opr
	}
}

func operand() (opd string) {
	if 2 < len(os.Args) {
		return os.Args[2]
	} else {
		return opd
	}
}

func usage() {
	fmt.Println(`
Synopsis

  source_table list           -- List (name|path|link).

  source_table enumerate      -- Enumerate (path|link).

`)
	os.Exit(1)
}

func main(){

	var file *os.File = open()
	if nil != file {
		defer file.Close()

		var table Table
		var er error

		table, er = table.read(file)
		if nil != er {
			log.Fatal(er)
		} else {
			switch operator() {

			case "list":
				table.List()

			case "enumerate":
				table.Enumerate()

			default:
				usage()
			}
		}
	} else {
		log.Fatal("Missing source table.")
	}
}
