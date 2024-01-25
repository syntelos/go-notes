/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

var targets map[FileTypeClass]FileLocationList = make(map[FileTypeClass]FileLocationList)

func TargetList(typeclass FileTypeClass) FileLocationList {

	return targets[typeclass]
}

type TargetOperationClass uint8

const (
	TargetOperationClassMonthly TargetOperationClass = 0b10000000
	TargetOperationClassPeer    TargetOperationClass = 0b01000000
)

func TargetOperation() TargetOperationClass { // [TODO] (review)

	switch ConfigurationContext() {

	case ClassNotes:
		return TargetOperationClassMonthly

	case ClassRecent:
		return TargetOperationClassPeer

	default:
		return 0
	}

}

func TargetDefine() {
	var typeclass_tgt FileTypeClass = ConfigurationTarget()
	var typeclass_src FileTypeClass = ConfigurationSource()

	switch TargetOperation() {

	case TargetOperationClassMonthly:

		var unique map[string]FileLocation = make(map[string]FileLocation)
		for _, fil := range SourceList(typeclass_tgt) {

			var fil_yyyymm string = fil.YYYYMM()
			var fil_hhmmss string = fil.YYYYMMDD_HHMMSS()
			if 0 != len(fil_hhmmss) {

				var inf FileLocation = unique[fil_yyyymm]
				if inf.IsValid() {

					var inf_hhmmss string = inf.YYYYMMDD_HHMMSS()
					if 0 != len(fil_hhmmss) {
						if fil_hhmmss > inf_hhmmss {

							unique[fil_yyyymm] = fil
						}
					} else {
						unique[fil_yyyymm] = fil
					}
				} else {
					unique[fil_yyyymm] = fil
				}
			}
		}

		var list FileLocationList
		for _, file := range unique {
			list = targets[file.typeclass]

			list = append(list,file)

			targets[file.typeclass] = list
		}

	case TargetOperationClassPeer:
		var list FileLocationList
		for _, from := range SourceList(typeclass_src) {
			var to FileLocation = from.Transform(typeclass_tgt)

			list = targets[to.typeclass]

			list = append(list,to)

			targets[to.typeclass] = list
		}

	}
}
