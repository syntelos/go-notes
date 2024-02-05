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
	var tgt, src string

	if HaveOperand(1) {
		tgt = Operand(0)
		src = Operand(1)
	} else if HaveOperand(0) {
		tgt = Operand(0)
	} else if HaveContext() {
		tgt = Context
	} else {
		return false
	}


	switch ConfigurationContext() {

	case ClassNotes:
		var walk []string

		if HaveOperand(1) {

			walk = FileList(src)

		} else if HaveOperand(0) {

			walk = FileList(tgt)

		} else if HaveContext() {

			walk = FileList(tgt)
		}

		if 0 != len(walk){

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

	case ClassRecent:

		switch ConfigurationTransform() {

		case ClassEncode:
			return false // [TODO]
		case ClassUpdate:
			return false // [TODO]

		case ClassFetch:
			var source FileLocation = RecentFetchSource(tgt,src)
			if source.IsValid() {

				var locationList FileLocationList = sources[source.typeclass]
				if 0 == len(locationList) {
					locationList = make(FileLocationList)
				}

				locationList[source.FileIdentifier()] = source

				sources[source.typeclass] = locationList
			}
			return true
		default:
			return false
		}
		return true

	default:
		return false
	}
}

func SourceClassCount() uint8 {

	return uint8(len(sources))
}

func SourceClassList() (list []FileTypeClass) {

	for cl, _ := range sources {

		list = append(list, cl)
	}
	return list
}
