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

    notes list [enc txt|upd]  -- List input.

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

func operator() (arg string) {
	if 1 < len(os.Args) {
		return os.Args[1]
	} else {
		return arg
	}
}

func haveOperand(first int) bool {
	first += 2

	return (first < len(os.Args))
}

func getOperand(first int) string {
	first += 2

	if first < len(os.Args) {
		return os.Args[first]
	} else {
		return ""
	}
}
func listOperands(first int) []string {
	first += 2

	if first < len(os.Args) {
		return os.Args[first:]
	} else {
		return make([]string,0)
	}
}

func main(){

	switch operator() {
	case "list":
		if haveOperand(0) {
			switch getOperand(0) {
			case "enc", "encode":
				notes.Init()
				if haveOperand(1) {
					var target notes.FileName
					for _, opd := range listOperands(1) {

						for _, target = range notes.ListTextFiles(opd) {

							fmt.Println(target)
						}
					}
					os.Exit(0)
				} else {
					usage()
				}
			case "upd", "update":
				if notes.Init() {
					var target notes.IndexTarget
					for _, target = range notes.ListIndexFiles() {
						
						fmt.Println(target)
					}
					os.Exit(0)
				} else {
					usage()
				}
			default:
				usage()
			}
		} else {
			usage()
		}
	case "encode":
		notes.Init()
		if haveOperand(0) {
			var target notes.FileName
			for _, opd := range listOperands(0) {

				for _, target = range notes.ListTextFiles(opd) {

					target.CodeWrite()
				}
			}
			os.Exit(0)
		} else {
			usage()
		}
	case "update":
		if notes.Init() {
			var target notes.IndexTarget
			for _, target = range notes.ListIndexFiles() {
				
				target.IndexWrite()
			}
			os.Exit(0)
		} else {
			usage()
		}
	default:
		usage()
	}
}
