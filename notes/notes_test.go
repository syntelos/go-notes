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
	if DefineObjectiveDirectory(ObjectiveKeyTargetWeb,"tst/notes") {
		DefineTextFiles([]string{"tst/txt"})
		var source IndexFile
		for _, source = range ListIndexSource(IndexFileTypeTXT) {

			fmt.Printf("[TestEncode] (CodeWrite) %s\n",source)

			source.CodeWrite()
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}

func TestUpdate(t *testing.T){
	if DefineObjectiveDirectory(ObjectiveKeyTargetWeb,"tst/notes") {
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
	if DefineObjectiveDirectory(ObjectiveKeyTargetWeb,"tst/notes") {
		var target IndexTarget
		for _, target = range ListIndexFiles() {

			fmt.Printf("[TestContent] (CatalogRead) %s\n",target.path)

			var text Catalog = target.CatalogEncode()

			for _, line := range text {

				fmt.Println(string(line))
			}
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}

func TestTabulate(t *testing.T){
	if DefineObjectiveDirectory(ObjectiveKeyTargetWeb,"tst/notes") {
		var target IndexTarget
		for _, target = range ListIndexFiles() {

			fmt.Printf("[TestTabulate] (CatalogRead) %s\n",target.path)

			var text Tabulation = target.TabulateEncode()

			for _, line := range text {

				fmt.Println(string(line))
			}
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}
