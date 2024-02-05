/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"testing"
)

const TestRecentGdrFileSource string = "tst/gdr/gdr_files_any-20231115_094508.json"
const TestRecentGdrFileTarget string = "tst/notes/gdr_files_any-20231115_094508.json"

func TestRecentSource(t *testing.T) {

	if Configure([]string{"recent", "get", TestRecentGdrFileTarget, TestRecentGdrFileSource}) {

		fmt.Printf("[TestRecentSource] %s\n", Operand(1))

		var src uint32 = 0
		for _, file := range SourceList(ConfigurationSource()) {
			src += 1

			fmt.Printf("[TestRecentSource] %s\n", file)
		}

		if 1 != src {
			t.Fatalf("[TestRecentSource] Count %d expected 1.", src)
		}

	} else {
		t.Fatal("[TestRecentSource] Failed to configure.")
	}
}

func TestRecentTarget(t *testing.T) {

	if Configure([]string{"recent", "get", TestRecentGdrFileTarget, TestRecentGdrFileSource}) {

		fmt.Printf("[TestRecentTarget] %s\n", Operand(0))

		var tgt uint32 = 0
		for _, file := range TargetList(ConfigurationTarget()) {
			tgt += 1

			fmt.Printf("[TestRecentTarget] %s\n", file)
		}

		if 1 != tgt {
			t.Fatalf("[TestRecentTarget] Count %d expected 1.", tgt)
		}

	} else {
		t.Fatal("[TestRecentTarget] Failed to configure.")
	}
}
