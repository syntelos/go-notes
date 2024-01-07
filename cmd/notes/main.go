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

Encode

    Given one of more directories of TXT files, or
    individual TXT files, produce corresponding SVG files in
    place.  

    Sources are filtered by filename (*.txt) as well as
    tablename (e.g. "politics-*.txt", or "sociology-*.txt").

    Targets are WWWeb Notes SVG text boxes.

    If the local directory "notes" is found, it is employed
    as the WWWeb Notes directory target structure.

Update

    Update WWWeb Notes directory structure with JSON
    indeces.  

    Note that existing JSON index files are not overwritten.

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

func haveOperand(idx int) bool {
	idx += 2

	return (idx < len(os.Args))
}

func operand(idx int) (arg string) {
	idx += 2

	if idx < len(os.Args) {
		return os.Args[idx]
	} else {
		return arg
	}
}

func main(){

	if haveOperator() {

		switch operator() {
		case "encode":
			notes.Init()
			var target notes.FileName
			for _, opd := range os.Args[1:] {

				for _, target = range notes.ListTextFiles(opd) {

					target.CodeWrite()
				}
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
