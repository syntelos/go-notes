/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

func TestSource(t *testing.T){

	if Configure([]string{"notes","encode","tst/notes","tst/txt"}) {

		fmt.Printf("[TestSource] (%s)\n",Operand(1))

		var src uint32 = 0
		for _, file := range SourceList(ConfigurationSource()) {
			src += 1

			fmt.Printf("[TestSource] %s\n",file)
		}

		if 4 != src {
			t.Fatalf("[TestSource] Count TXT %d expected 4.",src)
		}
	} else {
		t.Fatal("[TestSource] Failed to configure.")
	}
}
