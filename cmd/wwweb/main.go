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

    wwweb source ...          -- List input of operation.

    wwweb target ...          -- List output of operation.

    wwweb encode <tgt> <src>  -- Produce SVG notes from TXT.

    wwweb update <tgt>        -- Index SVG notes with JSON.

    wwweb contents <tgt>      -- Write JSON index as SVG.

    wwweb tabulate <tgt>      -- Tabulate indeces.

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

  Contents

      Update WWWeb Notes graphical tables of contents.
      Existing tables of contents are overwritten.

  Source

      Enumerate inputs derived from operation "encode",
      "update", or "contents".

  Target

      Enumerate outputs implied by operation "encode",
      "update", or "contents".

  The principal operators, "encode", "update", and
  "contents" are recognized by their corresponding short and
  long character symbols.

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
					notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1))
					if haveOperand(2) {
						notes.DefineTextFiles(listOperands(2))

						var target notes.IndexFile
						for _, target = range notes.ListIndexSource(notes.IndexFileTypeTXT) {

							fmt.Println(target)
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
					if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1)) {
						var file notes.IndexFile
						for _, file = range notes.ListIndexSource(notes.IndexFileTypeSVG) {
							
							fmt.Println(file)
						}
						os.Exit(0)
					} else {
						usage()
					}
				} else {
					usage()
				}
			case "co", "con", "contents":
				if haveOperand(1) {
					if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1)) {
						var target notes.IndexTarget
						for _, target = range notes.ListIndexTarget(notes.IndexFileTypeJSN) {
							
							fmt.Println(target.Path())
						}
						os.Exit(0)
					} else {
						usage()
					}
				} else {
					usage()
				}
			case "ta", "tab", "tabulate":
				if haveOperand(1) {
					if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1)) {
						var target notes.IndexFile
						for _, target = range notes.ListIndexSource(notes.IndexFileTypeJSN) {
							
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
					notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1))
					if haveOperand(2) {
						notes.DefineTextFiles(listOperands(2))

						var target notes.IndexFile
						for _, target = range notes.ListIndexSource(notes.IndexFileTypeTXT) {

							fmt.Println(target.FileTarget("svg"))
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
					if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1)) {
						var target notes.IndexTarget
						for _, target = range notes.ListIndexFiles() {
							
							fmt.Println(target.Path())
						}
						os.Exit(0)
					} else {
						usage()
					}
				} else {
					usage()
				}
			case "co", "con", "contents":
				if haveOperand(1) {
					if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1)) {
						var target notes.IndexTarget
						for _, target = range notes.ListIndexTarget(notes.IndexFileTypeJSN) {
							
							fmt.Println(target.CatalogTarget())
						}
						os.Exit(0)
					} else {
						usage()
					}
				} else {
					usage()
				}
			case "ta", "tab", "tabulate":
				if haveOperand(1) {
					if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(1)) {
						var target notes.IndexTarget
						for _, target = range notes.ListIndexTarget(notes.IndexFileTypeJSN) {
							
							fmt.Println(target.TabulateTarget())
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
			notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(0))
			if haveOperand(1) {
				if haveOperand(2) {
					notes.DefineTextFiles(listOperands(2))

					var target notes.IndexFile
					for _, target = range notes.ListIndexSource(notes.IndexFileTypeTXT) {

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
			if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(0)) {
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
	case "co", "con", "contents":
		if haveOperand(0) {
			if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(0)) {
				var target notes.IndexTarget
				for _, target = range notes.ListIndexTarget(notes.IndexFileTypeJSN) {
					
					target.CatalogWrite()
				}
				os.Exit(0)
			} else {
				usage()
			}
		} else {
			usage()
		}
	case "ta", "tab", "tabulate":
		if haveOperand(0) {
			if notes.DefineObjectiveDirectory(notes.ObjectiveKeyTargetWeb,getOperand(0)) {
				var target notes.IndexTarget
				for _, target = range notes.ListIndexTarget(notes.IndexFileTypeJSN) {
					
					target.TabulateWrite()
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
