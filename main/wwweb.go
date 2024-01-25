/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"fmt"
	"os"
	wwweb "github.com/syntelos/go-wwweb"
)

func usage(){
	fmt.Println(`
Synopsis

    wwweb [no|re] src ...     -- List input of ...

    wwweb [no|re] tgt ...     -- List output of ...

    wwweb [no|re] enc <t> <s> -- Produce TGT from SRC.

    wwweb [no|re] upd <tgt>   -- Produce TGT from DATA.

    wwweb [no|re] con <tgt>   -- Produce TGT from DATA.

    wwweb [no|re] tab <tgt>   -- Produce TGT from DATA.

Description

    The "wwweb" application has "notes" and "recent"
    transformation contexts for producing "wwweb notes"
    online content.  The "notes" applet produces SVG
    textboxes from TXT preformatted text.  The "recent"
    applet produces JSON indeces from a Google Drive file
    listing of PDF files.

    The WWWeb directory structure is

      <tgt>/<YYYY>/<MM>

    having index JSON

      <tgt>/<YYYY>/<MM>/<YYYY><MM><DD>.json

    and embed targets

      <tgt>/<YYYY>/<MM>/<tablename>-<YYYY><MM><DD>.svg
      <tgt>/<YYYY>/<MM>/<tablename>-<YYYY><MM><DD>.png

  Operators

      The operational symbols are recognized by short and
      long character symbols.

        "no" "not" "notes"
        "re" "rec" "recent"
        "src" "source"
        "tgt" "target"
        "enc" "encode"
        "upd" "update"
        "con" "contents"
        "tab" "tabulate"

`)
	os.Exit(1)
}


func main(){

	if wwweb.Configure(os.Args[1:]) {

		switch wwweb.ConfigurationOperation() {

		case wwweb.ClassSource:
			for _, src := range wwweb.SourceList(wwweb.ConfigurationSource()) {
				fmt.Println(src)
			}
			os.Exit(0)
		case wwweb.ClassTarget:
			for _, tgt := range wwweb.TargetList(wwweb.ConfigurationTarget()) {
				fmt.Println(tgt)
			}
			os.Exit(0)

		default:
			if wwweb.DataTransform() {

				os.Exit(0)
			}			
		}

		os.Exit(1)
	} else {
		usage()
	}
}
