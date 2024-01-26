/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"io/fs"
	"os"
)

var sources map[FileTypeClass]FileLocationList = make(map[FileTypeClass]FileLocationList)

func SourceList(typeclass FileTypeClass) FileLocationList {

	return sources[typeclass]
}

func sourceDefineWalker(path string, d fs.DirEntry, er error) error {

	if nil != d && !d.IsDir() {

		var ixfil FileIndex = FileClassify(path)
		if ixfil.IsValid() {

			var lofil FileLocation = ixfil.Condense()
			if lofil.IsValid() {

				var locationList FileLocationList = sources[ixfil.typeclass]
				if 0 == len(locationList) {
					locationList = make(FileLocationList)
				}

				locationList[lofil.FileIdentifier()] = lofil

				sources[ixfil.typeclass] = locationList
			}
		}
	}
	return nil
}

func SourceDefine() bool {

	if HaveOperand(1) {

		var src string = Operand(1)

		var dir fs.FS = os.DirFS(".")

		fs.WalkDir(dir,src,sourceDefineWalker)

		return true

	} else if HaveOperand(0) {

		var tgt string = Operand(0)

		var dir fs.FS = os.DirFS(".")

		fs.WalkDir(dir,tgt,sourceDefineWalker)

		return true

	} else {
		return false
	}
}

func SourceClassCount() uint8 {

	return uint8(len(sources))
}

func SourceClassList() (list []FileTypeClass) {

	for cl, _ := range sources {

		list = append(list,cl)
	}
	return list
}
