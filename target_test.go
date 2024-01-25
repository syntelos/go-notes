/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

const TestTargetClassSVG FileTypeClass = (FileClassTable|FileTypeSVG)
const TestTargetClassTXT FileTypeClass = (FileClassTable|FileTypeTXT)

func TestTarget(t *testing.T){
	Configure([]string{"notes","encode","tst/notes","tst/txt"}) // [TODO] (review)
	SourceDefine(Operand(1))
	TargetDefine()

	fmt.Println("[TestTarget]")

	var svg uint32 = 0
	for _, file := range TargetList(TestTargetClassSVG) {
		svg += 1

		fmt.Printf("[TestTarget] %s\n",file)
	}

	if 2 == svg {

		var txt uint32 = 0
		for _, file := range SourceList(TestTargetClassTXT) {
			txt += 1

			fmt.Printf("[TestTarget] %s\n",file)
		}

		if 4 != txt {
			t.Fatalf("[TestTarget] Count TXT %d expected 4.  Classes %d.",svg,SourceClassCount())
		}
	} else {
		t.Fatalf("[TestTarget] Count SVG %d expected 2.  Classes %d.",svg,SourceClassCount())
	}
}
