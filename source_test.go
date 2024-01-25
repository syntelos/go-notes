/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

const TestSourceClassSVG FileTypeClass = (FileClassTable|FileTypeSVG)
const TestSourceClassTXT FileTypeClass = (FileClassTable|FileTypeTXT)

func TestSource(t *testing.T){
	SourceDefine("tst")

	var svg uint32 = 0
	for _, file := range SourceList(TestSourceClassSVG) {
		svg += 1

		fmt.Println(file)
	}

	if 2 == svg {

		var txt uint32 = 0
		for _, file := range SourceList(TestSourceClassTXT) {
			txt += 1

			fmt.Println(file)
		}

		if 4 != txt {
			t.Fatalf("[TestSource] Count TXT %d expected 4.  Classes %d.",svg,SourceClassCount())
		}
	} else {
		t.Fatalf("[TestSource] Count SVG %d expected 2.  Classes %d.",svg,SourceClassCount())
	}
}
