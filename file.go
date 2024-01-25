/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"strings"
)

type FileTypeClass uint8

const (
	FileClassIndex FileTypeClass = 0b10000000
	FileClassTable FileTypeClass = 0b01000000

	FileTypeTXT    FileTypeClass = 0b00100000
	FileTypeJSN    FileTypeClass = 0b00010000
	FileTypeHTL    FileTypeClass = 0b00001000
	FileTypeSVG    FileTypeClass = 0b00000100
	FileTypePNG    FileTypeClass = 0b00000010
	FileTypeJPG    FileTypeClass = 0b00000001
)

const (
	FileClass FileTypeClass = (FileClassIndex|FileClassTable)
	FileType  FileTypeClass = (FileTypeTXT|FileTypeJSN|FileTypeHTL|FileTypeSVG|FileTypePNG|FileTypeJPG)
)

var FileTypeClassList []FileTypeClass = []FileTypeClass{FileClassIndex,FileClassTable,FileTypeTXT,FileTypeJSN,FileTypeHTL,FileTypeSVG,FileTypePNG,FileTypeJPG}

type CC uint8

const (
	CCNone   CC = 0
	CCLetter CC = 0b01000000
	CCNumber CC = 0b00100000
	CCOther  CC = 0b00010000
)

type FileLocation struct {
	typeclass FileTypeClass
	tablename string
	datetime string
	location string
}

type FileLocationList []FileLocation

type FileIndex struct {
	location []byte
	ix_head, ix_date, ix_fext int
	cc_head, cc_date, cc_fext CC
	typeclass FileTypeClass
}

func (this FileTypeClass) BitString() string {
	switch this {
	case FileClassIndex:
		return "index"
	case FileClassTable:
		return "table"
	case FileTypeTXT:
		return "text"
	case FileTypeJSN:
		return "json"
	case FileTypeHTL:
		return "html"
	case FileTypeSVG:
		return "svg"
	case FileTypePNG:
		return "png"
	case FileTypeJPG:
		return "jpg"
	default:
		return ""
	}
}

func (this FileTypeClass) ByteString() string {
	var str strings.Builder
	for _, bit := range FileTypeClassList {
		if bit == (this & bit) {
			if 0 != str.Len() {
				str.WriteByte(' ')
			}
			str.WriteString(bit.BitString())
		}
	}
	return str.String()
}

func FileClassify(location string) (index FileIndex) {
	{
		index.location = []byte(location)
		index.ix_head = -1
		index.cc_head = CCNone
		index.ix_date = -1
		index.cc_date = CCNone
		index.ix_fext = -1
		index.cc_fext = CCNone
	}

	var z int = len(location)
	if 0 < z {
		var y int = (z-1)
		if 0 < y {
			var x int = (y-1)
			if 0 < x {
				var cc CC = 0
				var ch byte

				scan:for ; 0 < x; x-- {
					ch = index.location[x]
					switch ch {
					case '/':
						index.ix_head = x
						index.cc_head = cc
						cc = 0
						break scan
					case '-':
						index.ix_date = x
						index.cc_date = cc
						cc = 0
					case '.':
						index.ix_fext = x
						index.cc_fext = cc
						cc = 0
					default:
						if ('a' <= ch && 'z' >= ch) || ('A' <= ch && 'Z' >= ch) {

							cc |= CCLetter

						} else if '0' <= ch && '9' >= ch {

							cc |= CCNumber

						} else if '_' != ch {

							cc |= CCOther
						}
					}
				}

				switch index.cc_head {
				case CCLetter, CCNumber:

					if -1 != index.ix_fext && CCLetter == index.cc_fext {
						var begin int = (index.ix_fext+1)
						var fext string = location[begin:]
						switch fext {
						case "txt":
							index.typeclass |= FileTypeTXT
						case "json":
							index.typeclass |= FileTypeJSN
						case "html":
							index.typeclass |= FileTypeHTL
						case "svg":
							index.typeclass |= FileTypeSVG
						case "png":
							index.typeclass |= FileTypePNG
						case "jpg", "jpeg":
							index.typeclass |= FileTypeJPG
						}
					}

					switch index.cc_head {
					case CCLetter:
						if -1 != index.ix_date && CCNumber == index.cc_date {

							index.typeclass |= FileClassTable
						}
					case CCNumber:
						if -1 == index.ix_date && CCNone == index.cc_date {

							index.typeclass |= FileClassIndex
						}		
					}
				}
			}
		}
	}
	return index
}

func (this FileIndex) IsValid() bool {
	return 0 != this.typeclass
}

func (this FileIndex) Condense() (that FileLocation) {
	if this.IsValid() {
		that.typeclass = this.typeclass
		that.location = string(this.location)

		switch (this.typeclass & FileClass) {
		case FileClassIndex:
			if 0 < this.ix_head && 0 < this.ix_fext {
				var begin, end int = this.ix_head+1, this.ix_fext
				that.datetime = string(this.location[begin:end])
			}
		case FileClassTable:
			if 0 < this.ix_head && 0 < this.ix_date {
				var begin, end int = this.ix_head+1, this.ix_date
				that.tablename = string(this.location[begin:end])

				if 0 < this.ix_date && 0 < this.ix_fext {
					var begin, end int = this.ix_date+1, this.ix_fext
					that.datetime = string(this.location[begin:end])
				}
			}
		}
	}
	return that
}

func (this FileIndex) String() string {
	return string(this.location)
}

func (this FileLocation) IsValid() bool {
	return 0 != this.typeclass
}

func (this FileLocation) String() string {
	return this.location
}

func (this FileLocation) YYYYMM() string {

	if 6 <= len(this.datetime) {
		return this.datetime[0:6]
	} else {
		return ""
	}
}

func (this FileLocation) YYYYMMDD() string {

	if 8 <= len(this.datetime) {
		return this.datetime[0:8]
	} else {
		return ""
	}
}

func (this FileLocation) YYYYMMDD_HHMMSS() string {

	if 15 <= len(this.datetime) && '_' == this.datetime[8] {
		return this.datetime[0:15]
	} else {
		return ""
	}
}

func (this FileLocation) HHMMSS() string {

	if 15 <= len(this.datetime) && '_' == this.datetime[8] {
		return this.datetime[10:15]
	} else {
		return ""
	}
}

func (this FileLocation) Transform(to FileTypeClass) (empty FileLocation) {
	var from FileTypeClass = (this.typeclass & FileType)
	to = (to & FileType)
	if from == to {

		return this

	} else {
		var fext int = strings.LastIndexByte(this.location,'.')
		var to_string string
		switch (to) {
		case FileTypeTXT:
			to_string = this.location[0:fext]+".txt"
		case FileTypeJSN:
			to_string = this.location[0:fext]+".json"
		case FileTypeHTL:
			to_string = this.location[0:fext]+".html"
		case FileTypeSVG:
			to_string = this.location[0:fext]+".svg"
		case FileTypePNG:
			to_string = this.location[0:fext]+".png"
		case FileTypeJPG:
			to_string = this.location[0:fext]+".jpeg"
		default:
			return empty
		}
		return FileClassify(to_string).Condense()
	}
}
