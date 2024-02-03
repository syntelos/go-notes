/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

type Coder interface {
	Decode([]byte)
	Encode() []byte
}

func DataTransform() bool {

	switch ConfigurationContext() {

	case ClassNotes:
		switch ConfigurationTransform() {

		case ClassEncode:

			for _, file := range TargetList(ConfigurationTarget()) {
				file.NotesEncode()
			}
			return true

		case ClassUpdate:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.NotesUpdate()
			}
			return true

		case ClassContents:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.NotesContents()
			}
			return true

		case ClassTabulate:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.NotesTabulate()
			}
			return true

		case ClassFetch:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.NotesFetch()
			}
			return true

		default:
			return false
		}

	case ClassRecent:
		switch ConfigurationTransform() {

		case ClassEncode:

			for _, file := range TargetList(ConfigurationTarget()) {
				file.RecentEncode()
			}
			return true

		case ClassUpdate:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.RecentUpdate()
			}
			return true

		case ClassContents:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.RecentContents()
			}
			return true

		case ClassTabulate:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.RecentTabulate()
			}
			return true

		case ClassFetch:
			for _, file := range TargetList(ConfigurationTarget()) {
				file.RecentFetch()
			}
			return true

		default:
			return false
		}

	default:
		return false
	}
}
