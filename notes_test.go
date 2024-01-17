/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T){
	if InitObjective("tst/notes") {
		var source FileName
		for _, source = range ListTextFiles("tst/txt") {

			fmt.Printf("[TestEncode] (CodeWrite) %s\n",source)

			source.CodeWrite()
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}

func TestUpdate(t *testing.T){
	if InitObjective("tst/notes") {
		var target IndexTarget
		for _, target = range ListIndexFiles() {

			fmt.Printf("[TestUpdate] (IndexWrite) %s\n",target.path)

			target.IndexWrite()
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}

func TestContent(t *testing.T){
	if InitObjective("tst/notes") {
		var target IndexTarget
		for _, target = range ListIndexFiles() {

			fmt.Printf("[TestContent] (IndexRead) %s\n",target.path)

			var text Catalog = target.Encode()

			for _, line := range text {

				fmt.Println(string(line))
			}
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}
