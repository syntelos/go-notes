/*
 * Source text
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"fmt"
	"log"
	"os"
	notes "github.com/syntelos/go-wwweb"
)

func bar(this notes.Catalog) (bar int) {
	bar = -1

	for tix, txt := range this {

		if txt.IsText() {

			bar = tix
			break
		}
	}
	return bar
}

func Enumerate(this notes.Catalog) {

	var cc int = 0

	switch operand() {

	case "head":
		var bar int = bar(this)
		fmt.Print("var catalog_head Catalog = Catalog{")
		for tix, txt := range this {

			if tix < bar {
				if 0 != cc {
					fmt.Print(",")
				}
				fmt.Print(" Text{")
				for cx, ch := range txt {
					if 0 != cx {
						fmt.Print(",")
					}
					fmt.Printf(" 0x%02X",ch)
				}
				fmt.Print(" }")
				cc += 1
			}
		}
		fmt.Println("}")

		os.Exit(0)

	case "tail":
		var bar int = bar(this)
		fmt.Print("var catalog_tail Catalog = Catalog{")
		for tix, txt := range this {

			if bar < tix {
				if 0 != cc {
					fmt.Print(",")
				}
				fmt.Print(" Text{")
				for cx, ch := range txt {
					if 0 != cx {
						fmt.Print(",")
					}
					fmt.Printf(" 0x%02X",ch)
				}
				fmt.Print(" }")
				cc += 1
			}
		}
		fmt.Println("}")

		os.Exit(0)

	case "all":
		fmt.Print("var catalog Catalog = Catalog{")
		for _, txt := range this {

			if 0 != cc {
				fmt.Print(",")
			}
			fmt.Print(" Text{")
			for cx, ch := range txt {
				if 0 != cx {
					fmt.Print(",")
				}
				fmt.Printf(" 0x%02X",ch)
			}
			fmt.Print(" }")
			cc += 1
		}
		fmt.Println("}")

		os.Exit(0)
	default:
		usage()
	}

}

func List(this notes.Catalog) {

	switch operand() {

	case "head":
		var bar int = bar(this)
		for tix, txt := range this {

			if tix < bar {

				fmt.Println(string(txt))
			}
		}

		os.Exit(0)

	case "tail":
		var bar int = bar(this)
		for tix, txt := range this {

			if bar < tix {

				fmt.Println(string(txt))
			}
		}

		os.Exit(0)

	case "all":
		for _, txt := range this {

			fmt.Println(string(txt))
		}

		os.Exit(0)
	default:
		usage()
	}

}

func open() (fi *os.File) {
	var er error
	fi, er = os.Open("catalog.svg")
	if nil != er {
		fi, er = os.Open("doc/catalog.svg")
		if nil != er {
			return nil
		}
	}
	return fi
}

func operator() (opr string) {
	if 1 < len(os.Args) {
		return os.Args[1]
	} else {
		return opr
	}
}

func operand() (opd string) {
	if 2 < len(os.Args) {
		return os.Args[2]
	} else {
		return opd
	}
}

func usage() {
	fmt.Println(`
Synopsis

  catalog list                   -- List (head|tail|all).

  catalog enumerate              -- Enumerate (head|tail).

`)
	os.Exit(1)
}

func main(){

	var file *os.File = open()
	if nil != file {
		defer file.Close()

		var catalog notes.Catalog
		var er error

		catalog, er = catalog.Read(file)
		if nil != er {
			log.Fatal(er)
		} else {
			switch operator() {

			case "enumerate":
				Enumerate(catalog)

			case "list":
				List(catalog)

			default:
				usage()
			}
		}
	} else {
		log.Fatalf("Missing source 'catalog.svg'.",)
	}
}
