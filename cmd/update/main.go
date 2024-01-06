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

    update                    -- Index content with JSON.

Description

    Update WWWeb Notes directory structure with JSON
    indeces.  

    Note that existing JSON index files are not overwritten.

`)
	os.Exit(1)
}

func main(){
	if notes.Init() {
		var target IndexTarget
		for _, target = range notes.ListIndexFiles() {
			
			target.IndexWrite()
		}
	} else {
		usage()
	}
}
