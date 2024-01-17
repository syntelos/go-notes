/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"os"
)

type ObjectiveKey uint8
const (
	ObjectiveKeyUnknown    ObjectiveKey = 0
	ObjectiveKeySourceText ObjectiveKey = 1
	ObjectiveKeyTargetWeb  ObjectiveKey = 2
)

var objective map[ObjectiveKey]FileName = make(map[ObjectiveKey]FileName)

func ObjectiveDirectory(key ObjectiveKey) FileName {

	return objective[key]
}

func HaveObjective(key ObjectiveKey) bool {

	var target FileName = objective[key]

	return 0 != len(target)
}

func DefineObjectiveDirectory(key ObjectiveKey, tgt string) bool {
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

				objective[key] = target
				return true
			}
		}
	}
	return false
}
