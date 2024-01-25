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

func TestTarget(t *testing.T){
	if Configure([]string{"notes","encode","tst/notes","tst/txt"}) { // [TODO] (review)
		SourceDefine(Operand(1))
		TargetDefine()

		fmt.Printf("[TestTarget] (%s)\n",Operand(1))

		var tgt uint32 = 0
		for _, file := range TargetList(TestTargetClassSVG) {
			tgt += 1

			fmt.Printf("[TestTarget] %s\n",file)
		}

		if 4 != tgt {
			t.Fatalf("[TestTarget] Count SVG %d expected 4.",tgt)
		}
	} else {
		t.Fatal("[TestTarget] Failed to configure.");
	}
}
