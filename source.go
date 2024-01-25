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

				locationList = append(locationList,lofil)

				sources[ixfil.typeclass] = locationList
			}
		}
	}
	return nil
}

func SourceDefine(tgt string) {

	var dir fs.FS = os.DirFS(".")

	fs.WalkDir(dir,tgt,sourceDefineWalker)
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
