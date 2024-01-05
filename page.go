/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Text []byte

type Page []Text

type IO interface {
	/*
	 */
	Read(os.File) (Page, error)
	/*
	 */
	Write(os.File) (error)
}

type Coder interface {
	/*
	 */
	Encode() (Page)
	/*
	 */
	Decode() (Page)
}

func (this Page) Read(src *os.File) (that Page, er error) {
	var fi os.FileInfo
	fi, er = src.Stat()
	if nil != er {
		return nil, fmt.Errorf("Read error file not found: %w",er)
	} else {
		var reader *bufio.Reader = bufio.NewReader(src)
		var inl []byte
		var isp bool
		inl, isp, er = reader.ReadLine()
		for true {
			if nil != er {
				if io.EOF == er {
					return this, nil
				} else {
					return nil, fmt.Errorf("Read error from '%s': %w",fi.Name(),er)
				}
			} else if isp {
				return nil, fmt.Errorf("Read error from '%s'.",fi.Name())
			} else {
				var line Text = inl
				this = append(this,line)
				inl, isp, er = reader.ReadLine()
			}
		}
		return nil, nil
	}
}

func (this Page) Write(tgt *os.File) (error) {
	var writer *bufio.Writer = bufio.NewWriter(tgt)

	for _, line := range this {

		writer.Write(line)
		writer.WriteByte('\n')
	}

	writer.Flush()

	return nil
}
var encodehead Page = Page{ Text{ 0x3C, 0x3F, 0x78, 0x6D, 0x6C, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6F, 0x6E, 0x3D, 0x22, 0x31, 0x2E, 0x30, 0x22, 0x20, 0x73, 0x74, 0x61, 0x6E, 0x64, 0x61, 0x6C, 0x6F, 0x6E, 0x65, 0x3D, 0x22, 0x6E, 0x6F, 0x22, 0x3F, 0x3E }, Text{ 0x3C, 0x3F, 0x78, 0x6D, 0x6C, 0x2D, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3D, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3A, 0x2F, 0x2F, 0x63, 0x64, 0x6E, 0x2E, 0x6A, 0x73, 0x64, 0x65, 0x6C, 0x69, 0x76, 0x72, 0x2E, 0x6E, 0x65, 0x74, 0x2F, 0x67, 0x68, 0x2F, 0x61, 0x61, 0x61, 0x61, 0x6B, 0x73, 0x68, 0x61, 0x74, 0x2F, 0x63, 0x6D, 0x2D, 0x77, 0x65, 0x62, 0x2D, 0x66, 0x6F, 0x6E, 0x74, 0x73, 0x40, 0x6C, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2F, 0x66, 0x6F, 0x6E, 0x74, 0x73, 0x2E, 0x63, 0x73, 0x73, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3D, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2F, 0x63, 0x73, 0x73, 0x22, 0x3F, 0x3E }, Text{ 0x3C, 0x21, 0x44, 0x4F, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x73, 0x76, 0x67, 0x20, 0x50, 0x55, 0x42, 0x4C, 0x49, 0x43, 0x20, 0x22, 0x2D, 0x2F, 0x2F, 0x57, 0x33, 0x43, 0x2F, 0x2F, 0x44, 0x54, 0x44, 0x20, 0x53, 0x56, 0x47, 0x20, 0x31, 0x2E, 0x31, 0x2F, 0x2F, 0x45, 0x4E, 0x22, 0x20, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3A, 0x2F, 0x2F, 0x77, 0x77, 0x77, 0x2E, 0x77, 0x33, 0x2E, 0x6F, 0x72, 0x67, 0x2F, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x2F, 0x53, 0x56, 0x47, 0x2F, 0x31, 0x2E, 0x31, 0x2F, 0x44, 0x54, 0x44, 0x2F, 0x73, 0x76, 0x67, 0x31, 0x31, 0x2E, 0x64, 0x74, 0x64, 0x22, 0x3E }, Text{ 0x3C, 0x73, 0x76, 0x67, 0x20, 0x78, 0x6D, 0x6C, 0x6E, 0x73, 0x3D, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3A, 0x2F, 0x2F, 0x77, 0x77, 0x77, 0x2E, 0x77, 0x33, 0x2E, 0x6F, 0x72, 0x67, 0x2F, 0x32, 0x30, 0x30, 0x30, 0x2F, 0x73, 0x76, 0x67, 0x22, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6F, 0x6E, 0x3D, 0x22, 0x31, 0x2E, 0x31, 0x22, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x76, 0x69, 0x65, 0x77, 0x42, 0x6F, 0x78, 0x3D, 0x22, 0x30, 0x20, 0x30, 0x20, 0x35, 0x30, 0x30, 0x20, 0x35, 0x30, 0x30, 0x22, 0x3E }, Text{ 0x20, 0x20, 0x3C, 0x64, 0x65, 0x66, 0x73, 0x3E }, Text{ 0x20, 0x20, 0x20, 0x20, 0x3C, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3D, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2F, 0x63, 0x73, 0x73, 0x22, 0x3E }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x7B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x66, 0x61, 0x6D, 0x69, 0x6C, 0x79, 0x3A, 0x20, 0x22, 0x43, 0x6F, 0x6D, 0x70, 0x75, 0x74, 0x65, 0x72, 0x20, 0x4D, 0x6F, 0x64, 0x65, 0x72, 0x6E, 0x20, 0x53, 0x65, 0x72, 0x69, 0x66, 0x22, 0x2C, 0x20, 0x73, 0x65, 0x72, 0x69, 0x66, 0x3B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20, 0x31, 0x31, 0x3B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x69, 0x6C, 0x6C, 0x3A, 0x20, 0x62, 0x6C, 0x61, 0x63, 0x6B, 0x3B }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7D }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2E, 0x74, 0x69, 0x74, 0x6C, 0x65, 0x20, 0x7B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20, 0x32, 0x32, 0x3B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3A, 0x20, 0x62, 0x6F, 0x6C, 0x64 }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7D }, Text{ 0x20, 0x20, 0x20, 0x20, 0x3C, 0x2F, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x3E }, Text{ 0x20, 0x20, 0x3C, 0x2F, 0x64, 0x65, 0x66, 0x73, 0x3E }, Text{ 0x20, 0x20, 0x3C, 0x72, 0x65, 0x63, 0x74, 0x20, 0x78, 0x3D, 0x22, 0x30, 0x22, 0x20, 0x79, 0x3D, 0x22, 0x30, 0x22, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x66, 0x69, 0x6C, 0x6C, 0x3D, 0x22, 0x23, 0x46, 0x46, 0x46, 0x22, 0x20, 0x73, 0x74, 0x72, 0x6F, 0x6B, 0x65, 0x3D, 0x22, 0x23, 0x30, 0x30, 0x30, 0x22, 0x20, 0x73, 0x74, 0x72, 0x6F, 0x6B, 0x65, 0x2D, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3D, 0x22, 0x33, 0x30, 0x22, 0x20, 0x2F, 0x3E }}

var encodetail Page = Page{ Text{ 0x3C, 0x2F, 0x73, 0x76, 0x67, 0x3E }}

func (this Page) Encode() (that Page) {
	/*
	 * Head
	 */
	for _, head := range encodehead {
		that = append(that,head)
	}
	/*
	 * Body
	 */
	var title bool = true
	var tem, thi int = 11, 23
	var bem, bhi int = 10, 12
	var px, py int = 40, 50
	var nti int = 0

	for _, body := range this {

		if 0 == len(body) {

			nti += 1

			if title {
				title = false
				px += (bem<<1)

			} else if 2 <= nti {
				title = true
				px = tem

			} else {
				py += bhi
			}

		} else if title {

			var str string = fmt.Sprintf("  <text class=\"title\" x=\"%d\" y=\"%d\">%s</text>",px,py,body)

			that = append(that,[]byte(str))

			if title {
				title = false
			}
			py += thi

		} else {
			nti = 0

			var str string = fmt.Sprintf("  <text x=\"%d\" y=\"%d\">%s</text>",px,py,body)

			that = append(that,[]byte(str))

			if title {
				title = false
			}
			py += bhi
		}
	}
	/*
	 * Tail
	 */
	for _, tail := range encodetail {
		that = append(that,tail)
	}
	return that
}

func (this Page) Decode() (that Page) {

	for _, txt := range this {

		if txt.IsText() {
			var begin = txt.first('>')+1
			var end = txt.last('<')

			var plain = txt[begin:end]

			that = append(that,plain)
		}
	}
	return that
}

func (this Text) IsText() (bool) {
	var found bool = false
	var first, last int = 0, 0

	for ix, ch := range this {

		switch ch {
		case '<':
			if 0 == first {
				first = ix
			}
		case '>':
			if 0 == last {
				last = ix

				break
			}
		default:
			if 0 < first {
				if ix == (first+1) {

					found = ('t' == ch)

				} else if ix == (first+2) {

					if found {

						found = ('e' == ch)
					}
				} else if ix == (first+3) {

					if found {

						found = ('x' == ch)
					}
				} else if ix == (first+4) {

					if found {

						found = ('t' == ch)
					}
				} else if ix == (first+5) {

					if found {

						found = ((' ' == ch)||('>' == ch))
					}
				}
			}
		}
	}

	if 0 < first && first < last {

		return found
	} else {
		return false
	}
}

func (this Text) first(ch byte) (spx int) {
	var ofs, len int = 0, len(this)
	
	spx = -1

	for ; ofs < len; ofs++ {

		if ch == this[ofs] {

			return ofs
		}
	}
	return spx
}

func (this Text) last(ch byte) (spx int) {
	var ofs, len int = 0, len(this)

	spx = -1

	for ofs = (len-1); 0 < ofs; ofs-- {

		if ch == this[ofs] {

			return ofs
		}
	}
	return spx
}
