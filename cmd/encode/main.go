/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"fmt"
	"log"
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

    Targets are (*.svg) textboxes.

    If a local directory "notes" is found, it is employed as
    the WWWeb Notes directory target structure.

`)
	os.Exit(1)
}

func operator() (opr string) {
	if 1 < len(os.Args) {

		return os.Args[1]
	} else {
		return opr
	}
}

func operand(idx int) (opd string) {
	var argx int = (2+idx)
	if argx < len(os.Args) {

		return os.Args[argx]
	} else {
		return opd
	}
}

func main(){
	notes.Init()

	if 1 < len(os.Args) {

		for _, opd := range os.Args[1:] {
			
			var list notes.FileList = notes.ListTextFiles(opd)
			for _, fn := range list {
				var er error
				var tgt *os.File
				tgt, er = os.Create(string(fn.Target()))
				if nil != er {
					log.Fatalf("Error opening output '%s': %w",string(fn.Target()),er)
				} else {
					var src *os.File
					src, er = os.Open(string(fn.Source()))
					if nil != er {
						log.Fatalf("Error opening input '%s': %w",string(fn.Source()),er)
					} else {
						var txt, svg notes.Page
						txt, er = txt.Read(src)
						if nil != er {
							log.Fatal(er)
						} else {
							svg = txt.Encode()

							er = svg.Write(tgt)
							if nil != er {
								log.Fatal(er)
							} else {
								src.Close()
								tgt.Close()
							}
						}
					}
				}
			}
		}
	} else {
		usage()
	}
}
