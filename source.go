/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

var sources map[FileTypeClass]FileLocationList = make(map[FileTypeClass]FileLocationList)

func SourceList(typeclass FileTypeClass) FileLocationList {

	return sources[typeclass]
}

func SourceDefine() bool {

	if HaveOperand(1) {

		var src string = Operand(1)

		var walk []string = FileList(src)

		for _, file := range walk {

			var ixfil FileIndex = FileClassify(file)
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
		return true

	} else if HaveOperand(0) {

		var tgt string = Operand(0)

		var walk []string = FileList(tgt)

		for _, file := range walk {

			var ixfil FileIndex = FileClassify(file)
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
		return true

	} else if HaveContext() {

		var tgt string = Context

		var walk []string = FileList(tgt)

		for _, file := range walk {

			var ixfil FileIndex = FileClassify(file)
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
