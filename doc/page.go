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

func bar(this notes.Page) (bar int) {
	bar = -1

	for tix, txt := range this {

		if txt.IsText() {

			bar = tix
			break
		}
	}
	return bar
}

func Enumerate(this notes.Page) {

	var cc int = 0

	switch operand() {

	case "head":
		var bar int = bar(this)
		fmt.Print("var encodehead Page = Page{")
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
		fmt.Print("var encodetail Page = Page{")
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
		fmt.Print("var page Page = Page{")
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

func List(this notes.Page) {

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
	fi, er = os.Open("page.svg")
	if nil != er {
		fi, er = os.Open("doc/page.svg")
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

  page list                   -- List (head|tail|all).

  page enumerate              -- Enumerate (head|tail).

`)
	os.Exit(1)
}

func main(){

	var file *os.File = open()
	if nil != file {
		defer file.Close()

		var page notes.Page
		var er error

		page, er = page.Read(file)
		if nil != er {
			log.Fatal(er)
		} else {
			switch operator() {

			case "enumerate":
				Enumerate(page)

			case "list":
				List(page)

			default:
				usage()
			}
		}
	} else {
		log.Fatalf("Missing source 'page.svg'.",)
	}
}
