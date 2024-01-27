/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

func TestCatalog(t *testing.T){

	if Configure([]string{"notes","encode","tst/notes","tst/txt"}) {

		fmt.Printf("[TestCatalog] (%s)\n",Operand(1))

		var tgt uint32 = 0
		var cat Catalog
		for _, file := range TargetList(ConfigurationTarget()) {
			tgt += 1

			cat = file.FileCatalog()

			fmt.Printf("[TestCatalog] %s\n",cat.LineString())
		}

		if 4 != tgt {

			t.Fatalf("[TestCatalog] Count TXT %d expected 4.",tgt)
		}
	} else {
		t.Fatal("[TestCatalog] Failed to configure.")
	}
}
