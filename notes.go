/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"bytes"
	"fmt"
)

func (this FileLocation) NotesEncode() {
	var tgt FileLocation = this
	var src FileLocation = this.Source(ConfigurationSource())
	if tgt.IsValid() && src.IsValid() {
		var txt Note
		if txt.Read(src) {
			var svg bytes.Buffer
			/*
			 * Head
			 */
			svg.Write(page_head.Encode())
			/*
			 * Body
			 */
			var title bool = true
			const bhi int = 18
			var px, py int = 30, 50

			for _, line := range txt.hyperlines {

				if line.IsText() {

					if title {
						svg.Write(line.Encode("title",px,py))
						svg.WriteByte('\n')

						title = false
					} else {
						svg.Write(line.Encode("text",px,py))
						svg.WriteByte('\n')
					}
				}
				py += bhi
			}
			/*
			 * Tail
			 */
			svg.Write(page_tail.Encode())

			tgt.Write(svg.Bytes())
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
	return 0 != len(this.link)
}

func (this NoteText) Encode(c string, x int, y int) (empty []byte) {
	if this.IsText() {
		if this.IsLink() {

			var str string = fmt.Sprintf("  <a href=\"%s\"><text class=\"%s\" x=\"%d\" y=\"%d\">%s</text></a>",string(this.link),c,x,y,string(this.text))

			return []byte(str)

		} else {

			var str string = fmt.Sprintf("  <text class=\"%s\" x=\"%d\" y=\"%d\">%s</text>",c,x,y,string(this.text))

			return []byte(str)
		}
	} else {
		return empty
	}
}
/*
 * Note text documents employ tabs and lines to encode
 * a trivial hypertext format.
 */
type Note struct {
	location FileLocation
	hyperlines []NoteText
}

func (this *Note) Read(file FileLocation) bool {
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
				if nil == line || 0 == len(line) && begin < tab {

					line = source[begin:tab]
				}

			case '\r', '\n':
				end = x
				if -1 != tab {
					begin = (tab+1)
					if (nil == link || 0 == len(link)) &&
					   (nil != line && 0 != len(line)) {

						link = source[begin:end]

						var text NoteText = NoteText{line,link}

						this.hyperlines = append(this.hyperlines,text)

						line = nil
						link = nil
					}
					tab = -1

				} else if begin < end {
					line = source[begin:end]

					var text NoteText = NoteText{line,nil}

					this.hyperlines = append(this.hyperlines,text)

					line = nil
				}
				begin = (end+1)
			}
		}

		if 0 != len(this.hyperlines) {
			this.location = file
			return true
		}
	}
	return false
}
