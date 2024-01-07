/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"fmt"
	"os"
	notes "github.com/syntelos/go-notes"
)

func usage(){
	fmt.Println(`
Synopsis

    notes encode <txt>        -- Produce SVG from TXT.

    notes update              -- Index content with JSON.

Description

    Update WWWeb Notes directory files.  If the local
    directory "notes" is found, it is employed as the WWWeb
    Notes directory structure.

  Encode

      Given one of more directories of TXT files, or
      individual TXT files, produce corresponding SVG files
      in place or in structure.  Sources are filtered by
      filename (*.txt) as well as tablename
      (e.g. "politics-*.txt", or "sociology-*.txt").
      Targets are WWWeb Notes SVG text boxes.

  Update

      Update WWWeb Notes JSON indeces.  Note that existing
      JSON index files are not overwritten.

`)
	os.Exit(1)
}

func haveOperator() bool {
	return (1 < len(os.Args))
}

func operator() (arg string) {
	if 1 < len(os.Args) {
		return os.Args[1]
	} else {
		return arg
	}
}

func haveOperands() bool {

	return (2 < len(os.Args))
}

func operands() []string {

	if 2 < len(os.Args) {
		return os.Args[2:]
	} else {
		return make([]string,0)
	}
}

func main(){

	if haveOperator() {

		switch operator() {
		case "encode":
			notes.Init()
			if haveOperands() {
				var target notes.FileName
				for _, opd := range operands() {

					for _, target = range notes.ListTextFiles(opd) {

						target.CodeWrite()
					}
				}
			} else {
				usage()
			}
		case "update":
			if notes.Init() {
				var target notes.IndexTarget
				for _, target = range notes.ListIndexFiles() {
					
					target.IndexWrite()
				}
			} else {
				usage()
			}
		default:
			usage()
		}
	} else {
		usage()
	}
}
