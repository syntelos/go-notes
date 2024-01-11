/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"testing"
)

func TestEncode(t *testing.T){
	if Init("tst/notes") {
		var source FileName
		for _, source = range ListTextFiles("tst/txt") {
			
			source.CodeWrite()
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}

func TestUpdate(t *testing.T){
	if Init("tst/notes") {
		var target IndexTarget
		for _, target = range ListIndexFiles() {
			
			target.IndexWrite()
		}		
	} else {
		t.Fatal("Missing 'tst/notes'.")
	}
}
