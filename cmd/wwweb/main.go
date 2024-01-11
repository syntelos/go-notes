/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"fmt"
	"os"
	notes "github.com/syntelos/go-wwweb"
)

func usage(){
	fmt.Println(`
Synopsis

    wwweb source [en|up] ...  -- List input.

    wwweb target [en|up] ...  -- List output.

    wwweb encode <tgt> <src>  -- Produce SVG from TXT.

    wwweb update <tgt>        -- Index content with JSON.

Description

    Update WWWeb Notes directory files.  The WWWeb Notes
    directory structure is

      <tgt>/<YYYY>/<MM>

    as for JSON index target

      <tgt>/<YYYY>/<MM>/<YYYY><MM><DD>.json

    or embed targets

      <tgt>/<YYYY>/<MM>/<tablename>-<YYYY><MM><DD>.svg
      <tgt>/<YYYY>/<MM>/<tablename>-<YYYY><MM><DD>.png

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

  Source

      Enumerate inputs derived from operation "encode" or
      "update".

  Target

      Enumerate outputs implied by operation "encode" or
      "update".

  The principal operators, "encode" and "update" are
  recognized by their corresponding two, three, and six
  character symbols.

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
	case "source":
		if haveOperand(0) {
			switch getOperand(0) {
			case "en", "enc", "encode":
				if haveOperand(1) {
					notes.Init(getOperand(1))
					if haveOperand(2) {
						var target notes.FileName
						for _, opd := range listOperands(2) {

							for _, target = range notes.ListTextFiles(opd) {

								fmt.Println(target)
							}
						}
						os.Exit(0)
					} else {
						usage()
					}
				} else {
					usage()
				}
			case "up", "upd", "update":
				if haveOperand(1) {
					if notes.Init(getOperand(1)) {
						var target notes.IndexTarget
						for _, target = range notes.ListIndexFiles() {
							
							fmt.Println(target)
						}
						os.Exit(0)
					} else {
						usage()
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
	case "target":
		if haveOperand(0) {
			switch getOperand(0) {
			case "en", "enc", "encode":
				if haveOperand(1) {
					notes.Init(getOperand(1))
					if haveOperand(2) {
						var target notes.FileName
						for _, opd := range listOperands(2) {

							for _, target = range notes.ListTextFiles(opd) {

								fmt.Println(target.Target())
							}
						}
						os.Exit(0)
					} else {
						usage()
					}
				} else {
					usage()
				}
			case "up", "upd", "update":
				if haveOperand(1) {
					if notes.Init(getOperand(1)) {
						var target notes.IndexTarget
						for _, target = range notes.ListIndexFiles() {
							
							fmt.Println(target.Target())
						}
						os.Exit(0)
					} else {
						usage()
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
	case "en", "enc", "encode":
		if haveOperand(0) {
			notes.Init(getOperand(0))
			if haveOperand(1) {
				var target notes.FileName
				for _, opd := range listOperands(1) {

					for _, target = range notes.ListTextFiles(opd) {

						target.CodeWrite()
					}
				}
				os.Exit(0)
			} else {
				usage()
			}
		} else {
			usage()
		}
	case "up", "upd", "update":
		if haveOperand(0) {
			if notes.Init(getOperand(0)) {
				var target notes.IndexTarget
				for _, target = range notes.ListIndexFiles() {
					
					target.IndexWrite()
				}
				os.Exit(0)
			} else {
				usage()
			}
		} else {
			usage()
		}
	default:
		usage()
	}
}
