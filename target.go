/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"os"
)

var NotesTarget FileName

func HaveTarget() bool {

	return 0 != len(NotesTarget)
}

func InitTarget(tgt string) bool {
	var target FileName = FileName(tgt)

	var fo *os.File
	var er error

	fo, er = os.Open(string(target))
	if nil == er {
		defer fo.Close()

		var fi os.FileInfo
		fi, er = fo.Stat()
		if nil == er {

			if fi.IsDir() {

				NotesTarget = target
				return true
			}
		}
	}
	return false
}
