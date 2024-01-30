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

      wwweb <cx> <op> <tgt> [src]

      wwweb <cx> (src|tgt) <op> <tgt> [src]

Description

    The wwweb production manager has principal
    transformation contexts ("notes" and "recent") for
    producing "wwweb notes" and "recent documents" online.
    The "notes" applet produces SVG textboxes from TXT
    preformatted text.  The "recent" applet produces JSON
    indeces from a Google Drive file listing of PDF files.

    The WWWeb directory structure is

      <CX>/<YYYY>/<MM>

    having catalog indeces

      <CX>/<YYYY>/<MM>/<YYYY><MM><DD>.json

    and table records

      <CX>/<YYYY>/<MM>/<TABLENAME>-<YYYY><MM><DD>.svg
      <CX>/<YYYY>/<MM>/<TABLENAME>-<YYYY><MM><DD>.svg

    A catalog index relates to a table record by the tabular
    intersection of association in identity, location, and
    content.  The abstract catalog table includes the
    catalog index and table record.  The table record
    expresses the visual representation of the catalog table
    in the union of material information.

  Contexts

    The conventional symbols are recognized by short and
    long character symbols.

      "not", "notes"          -- TGT <SVG> SRC <TXT>
      "rec", "recent"         -- TGT <JSN> SRC <GDR>

  Operators

      "src", "source"         -- List sources of <op> ...
      "tgt", "target"         -- List targets of <op> ...
      "enc", "encode"         -- Encode <src> into <tgt>
      "upd", "update"         -- Update <tgt>
      "con", "contents"       -- Contents of <src> to <tgt>
      "tab", "tabulate"       -- Tabulation of <src> to <t>

  Operands

    The target operand is a production destination, and is
    always first.  The source operand is a content location,
    and is second when present.  Operands are directories
    and files.

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
			} else {

				os.Exit(1)
			}
		}
	} else {
		usage()
	}
}
