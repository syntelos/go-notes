/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

func TestTarget(t *testing.T) {

	if Configure([]string{"notes", "encode", "tst/notes", "tst/txt"}) {

		fmt.Printf("[TestTarget] (%s)\n", Operand(1))

		var tgt uint32 = 0
		for _, file := range TargetList(ConfigurationTarget()) {
			tgt += 1

			fmt.Printf("[TestTarget] %s\n", file)
		}

		if 4 != tgt {
			t.Fatalf("[TestTarget] Count TXT %d expected 4.", tgt)
		}
	} else {
		t.Fatal("[TestTarget] Failed to configure.")
	}
}
