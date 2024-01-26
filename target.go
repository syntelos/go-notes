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

func TargetDefine() bool {
	var typeclass_tgt FileTypeClass = ConfigurationTarget()
	var typeclass_src FileTypeClass = ConfigurationSource()

	switch TargetOperation() {

	case TargetOperationClassMonthly:
		var unique FileCollectionList = make(FileCollectionList)
		for _, fil := range SourceList(typeclass_tgt) {

			var fil_id FileId = fil.FileIdentifier()
			var fil_ix FileIx = fil.FileIndex()
			if fil_id.IsValid() && fil_ix.IsValid() {

				var inf FileLocation = unique[fil_ix]
				if inf.IsValid() {

					var inf_id FileId = inf.FileIdentifier()
					var inf_ix FileIx = inf.FileIndex()
					if inf_id.IsValid() && inf_ix.IsValid() {
						if inf_id > fil_id {

							unique[fil_ix] = inf
						}
					}
				} else {
					unique[fil_ix] = fil
				}
			}
		}

		var list FileLocationList
		for _, file := range unique {
			list = targets[file.typeclass]

			if 0 == len(list) {
				list = make(FileLocationList)
			}

			list[file.FileIdentifier()] = file

			targets[file.typeclass] = list
		}
		return true

	case TargetOperationClassPeer:
		var list FileLocationList
		for _, from := range SourceList(typeclass_src) {
			var to FileLocation = from.Transform(typeclass_tgt)

			list = targets[to.typeclass]

			if 0 == len(list) {
				list = make(FileLocationList)
			}

			list[to.FileIdentifier()] = to

			targets[to.typeclass] = list
		}
		return true
	default:
		return false
	}
}
