/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

func TestNotesSource(t *testing.T) {

	if Configure([]string{"notes", "encode", "tst/notes", "tst/txt"}) {

		fmt.Printf("[TestNotesSource] (%s)\n", Operand(1))

		var src uint32 = 0
		for _, file := range SourceList(ConfigurationSource()) {
			src += 1

			fmt.Printf("[TestNotesSource] %s\n", file)
		}

		if 4 != src {
			t.Fatalf("[TestNotesSource] Count TXT %d expected 4.", src)
		}
	} else {
		t.Fatal("[TestNotesSource] Failed to configure.")
	}
}

func TestNotesTarget(t *testing.T) {

	if Configure([]string{"notes", "encode", "tst/notes", "tst/txt"}) {

		fmt.Printf("[TestNotesTarget] (%s)\n", Operand(1))

		var tgt uint32 = 0
		for _, file := range TargetList(ConfigurationTarget()) {
			tgt += 1

			fmt.Printf("[TestNotesTarget] %s\n", file)
		}

		if 4 != tgt {
			t.Fatalf("[TestNotesTarget] Count TXT %d expected 4.", tgt)
		}
	} else {
		t.Fatal("[TestNotesTarget] Failed to configure.")
	}
}
