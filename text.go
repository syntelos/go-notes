/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import "fmt"

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

func (this FileLocation) NoteRead() (that Note) {
	var source []byte = this.Read()
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
				return Note{} // (format error)

			case '\n':
				end = x
				if -1 != tab {
					begin = (tab + 1)
					if (nil == link || 0 == len(link)) &&
						(nil != line && 0 != len(line)) {

						link = source[begin:end]

						var hypertext NoteText = NoteText{line, link}

						that.hyperlines = append(that.hyperlines, hypertext)

						line = nil
						link = nil
					}
					tab = -1

				} else if begin < end {
					line = source[begin:end]

					var plaintext NoteText = NoteText{line, nil}

					that.hyperlines = append(that.hyperlines, plaintext)

					line = nil
					link = nil
				} else {
					var newline NoteText = NoteText{[]byte{}, []byte{}}

					that.hyperlines = append(that.hyperlines, newline)

					line = nil
					link = nil
				}
				begin = (end + 1)
			}
		}
	}
	return that
}

func (this Note) IsNotEmpty() bool {

	return 0 != len(this.hyperlines)
}
