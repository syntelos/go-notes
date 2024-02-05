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
	TargetOperationNotesCollection TargetOperationClass = 0b10000000
	TargetOperationNotesPeer       TargetOperationClass = 0b01000000
	TargetOperationRecentFetch     TargetOperationClass = 0b00100000
)

func TargetOperation() TargetOperationClass {
	switch ConfigurationContext() {

	case ClassNotes:
		switch ConfigurationTransform() {

		case ClassEncode:
			return TargetOperationNotesPeer

		case ClassUpdate, ClassContents, ClassTabulate:
			return TargetOperationNotesCollection

		default:
			return 0
		}

	case ClassRecent:
		switch ConfigurationTransform() {

		case ClassFetch:
			return TargetOperationRecentFetch

		default:
			return 0
		}

	default:
		return 0
	}
}

func TargetDefine() bool {
	/*
	 * Initialize
	 */
	var typeclass_tgt FileTypeClass = ConfigurationTarget()
	var typeclass_src FileTypeClass = ConfigurationSource()

	var list FileLocationList = targets[typeclass_tgt]
	if 0 == len(list) {
		list = make(FileLocationList)
		targets[typeclass_tgt] = list
	}
	/*
	 * Collect projection
	 */
	switch TargetOperation() {

	case TargetOperationNotesCollection:
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
		/*
		 * Map into targets set.
		 */
		for _, to := range unique {

			list = targets[typeclass_tgt]

			list[to.FileIdentifier()] = to

			targets[typeclass_tgt] = list
		}
		return true

	case TargetOperationNotesPeer:
		/*
		 * Map into targets set.
		 */
		for _, from := range SourceList(typeclass_src) {
			var to FileLocation = from.Target(typeclass_tgt)

			list = targets[typeclass_tgt]

			list[to.FileIdentifier()] = to

			targets[typeclass_tgt] = list
		}
		return true

	case TargetOperationRecentFetch:
		var to FileLocation = RecentFetchTarget()
		if typeclass_tgt == to.typeclass {

			list = targets[typeclass_tgt]

			list[to.FileIdentifier()] = to

			targets[typeclass_tgt] = list

			return true
		} else {
			return false
		}
	default:
		return false
	}
}
