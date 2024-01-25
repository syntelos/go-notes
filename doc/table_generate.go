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
	"strings"
)

/*
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

func create() (fi *os.File) {
	var er error
	fi, er = os.Create("table.go")
	if nil != er {
		return nil
	} else {
		return fi
	}
}

/*
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

/*
 */
type Record struct {
	name, path, link string
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

func main(){

	var file *os.File = open()
	if nil != file {

		var table Table
		var er error

		table, er = table.read(file)
		if nil != er {
			log.Fatal(er)
		} else {
			file.Close()

			file = create()

			fmt.Fprintln(file,`/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

type Table string
type TableName Table
type TablePath Table
type TableLink Table

const (
`)
			var cn string
			for _, re := range table {
				cn = camel(re.name)

				fmt.Fprintf(file,"\tTableName%s\tTableName\t= \"%s\"\n",cn,re.name)
			}
			fmt.Fprintln(file)

			for _, re := range table {
				cn = camel(re.name)

				fmt.Fprintf(file,"\tTablePath%s\tTablePath\t= \"%s\"\n",cn,re.path)
			}
			fmt.Fprintln(file)

			for _, re := range table {
				cn = camel(re.name)

				fmt.Fprintf(file,"\tTableLink%s\tTableLink\t= \"%s\"\n",cn,re.link)
			}

			fmt.Fprintln(file,`)

func IsTableName(name TableName) bool {
	if 0 != len(name) {
		var path TablePath = name.Path()
		return (0 != len(path))
	} else {
		return false
	}
}

func (this TableName) Path() TablePath {
	switch this {
`)
			for _, re := range table {
				cn = camel(re.name)

				fmt.Fprintf(file,"\tcase TableName%s:\n\t\treturn TablePath%s\n",cn,cn)

			}			
			fmt.Fprintln(file,`
	default:
		return ""
	}
}

func (this TableName) Link() TableLink {
	switch this {
`)
			for _, re := range table {
				cn = camel(re.name)

				fmt.Fprintf(file,"\tcase TableName%s:\n\t\treturn TableLink%s\n",cn,cn)
			}
			fmt.Fprintln(file,`
	default:
		return ""
	}
}
`)
			file.Close()
			os.Exit(0)
		}
	} else {
		log.Fatal("Missing source table.")
	}
}
