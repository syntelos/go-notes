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

    encode <txt>              -- Produce SVG from TXT.

Description

    Given one of more directories of TXT files, or
    individual TXT files, produce corresponding SVG files in
    place.  

    Sources are filtered by filename (*.txt) as well as
    tablename (e.g. "politics-*.txt", or "sociology-*.txt").

    Targets are WWWeb Notes SVG text boxes.

    If the local directory "notes" is found, it is employed
    as the WWWeb Notes directory target structure.

`)
	os.Exit(1)
}

func main(){
	notes.Init()

	if 1 < len(os.Args) {
		var target notes.FileName
		for _, opd := range os.Args[1:] {

			for _, target = range notes.ListTextFiles(opd) {

				target.CodeWrite()
			}
		}
	} else {
		usage()
	}
}
