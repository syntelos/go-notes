/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"bytes"
	"fmt"
)
/*
 * Encode SRC TXT to TGT SVG as the "notes textbox" found in
 * PAGE.
 */
func (this FileLocation) NotesEncode() {
	var tgt FileLocation = this
	var src FileLocation = this.Source(ConfigurationSource())
	if tgt.IsValid() && src.IsValid() {
		var txt Note = Note{this, []NoteText{}}
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
						svg.Write(line.Encode("title", px, py))

						title = false

					} else if line.IsLink() {

						svg.Write(line.Encode("link", px, py))
					} else {
						svg.Write(line.Encode("text", px, py))
					}
					svg.WriteByte('\n')
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
/*
 * Write TGT JSN catalog index for SRC SVG filtered and
 * ordered.
 */
func (this FileLocation) NotesUpdate() {
	var tgt FileLocation = this
	/*
	 * Do not overwrite existing target
	 */
	if tgt.IsValid() && tgt.NotExists() {
		var membership FileIx = tgt.FileIndex()
		/*
		 * Include source list as ordered members of {FileIx} `dirname`
		 */
		var src []FileLocation
		{
			for _, rev := range SourceList(ConfigurationSource()) {
				var rel FileIx = rev.FileIndex()
				if rel == membership {

					src = append(src, rev)
				}
			}
			src = FileSort(src)
		}

		if 0 < len(src) {

			var cat Index = CatalogIndex(src)

			tgt.Write(cat.Encode())
		}
	}
}

func (this FileLocation) NotesContents() {
}

func (this FileLocation) NotesTabulate() {
}

func (this FileLocation) NotesFetch() {
}

/*
 * Note text is preformatted plain text.  Lines are wrapped
 * to presentation width by unix line terminals (i.e. '\n',
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

			var str string = fmt.Sprintf("  <a href=\"%s\"><text class=\"%s\" x=\"%d\" y=\"%d\">%s</text></a>", string(this.link), c, x, y, string(this.text))

			return []byte(str)

		} else {

			var str string = fmt.Sprintf("  <text class=\"%s\" x=\"%d\" y=\"%d\">%s</text>", c, x, y, string(this.text))

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
	location   FileLocation
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
			ch = source[x]
			switch ch {

			case ' ':
				if x == begin {
					/*
					 * Truncate line prefix whitespace.
					 */
					begin += 1
				}

			case '\t':
				if x == begin {
					/*
					 * Truncate line prefix whitespace.
					 */
					begin += 1

				} else {
					tab = x
					if nil == line || 0 == len(line) && begin < tab {

						line = source[begin:tab]
						link = nil
					}
				}

			case '\r':
				return false // (format error)

			case '\n':
				end = x
				if -1 != tab {
					begin = (tab + 1)
					if (nil == link || 0 == len(link)) &&
						(nil != line && 0 != len(line)) {

						link = source[begin:end]

						var hypertext NoteText = NoteText{line, link}

						this.hyperlines = append(this.hyperlines, hypertext)

						line = nil
						link = nil
					}
					tab = -1

				} else if begin < end {
					line = source[begin:end]

					var plaintext NoteText = NoteText{line, nil}

					this.hyperlines = append(this.hyperlines, plaintext)

					line = nil
					link = nil
				} else {
					var newline NoteText = NoteText{[]byte{}, []byte{}}

					this.hyperlines = append(this.hyperlines, newline)

					line = nil
					link = nil
				}
				begin = (end + 1)
			}
		}

		return 0 != len(this.hyperlines)
	} else {
		return false
	}
}
