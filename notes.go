/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import "bytes"

/*
 * Encode SRC TXT to TGT SVG as the "notes textbox" found in
 * PAGE.
 */
func (this FileLocation) NotesEncode() {
	var tgt FileLocation = this
	var src FileLocation = this.Source(ConfigurationSource())
	if tgt.IsValid() && src.IsValid() {
		var txt Note = src.NoteRead()
		if txt.IsNotEmpty() {
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

			var cat CatalogIndex = MakeCatalogIndex(src)

			tgt.CatalogIndexWrite(cat)
		}
	}
}

func (this FileLocation) NotesContents() {
}

func (this FileLocation) NotesTabulate() {
}

func (this FileLocation) NotesFetch() {
}
