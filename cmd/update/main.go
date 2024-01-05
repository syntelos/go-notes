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

`)
	os.Exit(1)
}

func main(){
	if notes.Init() {

		for _, ixfil := range notes.ListIndexFiles() {
			
			ixfil.IndexWrite()
		}
	} else {
		usage()
	}
}
