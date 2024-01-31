/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"bytes"
)

func (this FileLocation) NotesEncode() { // [TODO]
	var tgt FileLocation = this
	var src FileLocation = this.Source(ConfigurationSource())
	if tgt.IsValid() && src.IsValid() {
		var source Note
		if source.Read(src) {
			var wb bytes.Buffer

			wb.Write(page_head.Encode())


			wb.Write(page_tail.Encode())

			tgt.Write(wb.Bytes())
		}
	}
}

func (this FileLocation) NotesUpdate() { // [TODO]
}

func (this FileLocation) NotesContents() {
}

func (this FileLocation) NotesTabulate() {
}
/*
 * Note text is preformatted plain text.  Lines are wrapped
 * to presentation width by line terminals (i.e. '\n',
 * 0x0A).  Lines may be associated to link URLs with a
 * terminal tab (i.e. '\t' 0x09) separator.  The first line
 * is employed as a title.  Paragraphs are separated by
 * empty lines.
 */
type NoteText struct {
	text []byte
	link []byte
}

func (this NoteText) IsText() bool {
	return 0 != len(this.text)
}

func (this NoteText) IsLink() bool {
	return 0 != len(this.text) && 0 != len(this.link)
}

type Note []NoteText


func (this Note) Read(file FileLocation) bool {
	var source []byte = file.Read()
	var z int = len(source)
	if 0 < z {
		var x, begin, end, tab int = 0, 0, -1, -1
		var ch byte
		var line, link []byte

		for ; x < z; x++ {
			ch = source[ch]
			switch ch {

			case '\t':
				tab = x
				if nil == line || 0 == len(line) {
					line = source[begin:tab]
				}

			case '\r', '\n':
				end = x
				if -1 != tab {
					begin = (tab+1)
					if nil == link || 0 == len(link) {

						if nil != line && 0 != len(line) {

							link = source[begin:end]

							var text NoteText = NoteText{line,link}

							this = append(this,text)

							line = nil
							link = nil
						}
						tab = -1
					}
				} else if begin < end {
					line = source[begin:end]

					var text NoteText = NoteText{line,nil}

					this = append(this,text)

					line = nil
				}
				begin = (end+1)
			}
		}
	}
	return false
}
