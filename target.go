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
		return TargetOperationClassPeer

	case ClassRecent:
		return TargetOperationClassMonthly

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
		for _, from := range SourceList(typeclass_src) {
			var to FileLocation = from.Target(typeclass_tgt)

			var to_id FileId = to.FileIdentifier()
			var to_ix FileIx = to.FileIndex()
			if to_id.IsValid() && to_ix.IsValid() {

				var ck FileLocation = unique[to_ix]
				if ck.IsValid() {

					var ck_id FileId = ck.FileIdentifier()
					var ck_ix FileIx = ck.FileIndex()
					if ck_id.IsValid() && ck_ix.IsValid() {
						if to_id > ck_id {

							unique[to_ix] = to
						}
					}
				} else {
					unique[to_ix] = to
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
			var to FileLocation = from.Target(typeclass_tgt)

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
