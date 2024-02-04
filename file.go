/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type IO interface {
	Read() []byte
	Write([]byte)
}

type FileLocation struct {
	/*
	 * Type class of location
	 */
	typeclass FileTypeClass
	/*
	 * Source of derivation
	 */
	source string
	/*
	 * Location & components
	 */
	dirname, basename, location string

	tablename TableName

	datetime DateTime
}

type FileIx string

type FileCollectionList map[FileIx]FileLocation

type FileId string

type FileLocationList map[FileId]FileLocation

type FileTypeClass uint16

const (
	FileClassAbstract FileTypeClass = 0b1000000000000000
	FileClassIndex    FileTypeClass = 0b0100000000000000
	FileClassTable    FileTypeClass = 0b0010000000000000

	FileTypeTXT       FileTypeClass = 0b0000000000100000
	FileTypeJSN       FileTypeClass = 0b0000000000010000
	FileTypeHTL       FileTypeClass = 0b0000000000001000
	FileTypeSVG       FileTypeClass = 0b0000000000000100
	FileTypePNG       FileTypeClass = 0b0000000000000010
	FileTypeJPG       FileTypeClass = 0b0000000000000001
)

const (
	FileClass FileTypeClass = (FileClassAbstract | FileClassIndex | FileClassTable)
	FileType  FileTypeClass = (FileTypeTXT | FileTypeJSN | FileTypeHTL | FileTypeSVG | FileTypePNG | FileTypeJPG)
)

var FileTypeClassList []FileTypeClass = []FileTypeClass{FileClassIndex, FileClassTable, FileTypeTXT, FileTypeJSN, FileTypeHTL, FileTypeSVG, FileTypePNG, FileTypeJPG}

type CC uint8

const (
	CCNone   CC = 0
	CCLetter CC = 0b01000000
	CCNumber CC = 0b00100000
	CCOther  CC = 0b00010000
)

type FileIndex struct {
	location                  []byte
	ix_head, ix_date, ix_fext int
	cc_head, cc_date, cc_fext CC
	typeclass                 FileTypeClass
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
		var y int = (z - 1)
		if 0 < y {
			var x int = (y - 1)
			if 0 < x {
				var cc CC = 0
				var ch byte
				/*
				 * Collection
				 */
			scan:
				for ; 0 < x; x-- {
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
				/*
				 * Analysis
				 */
				var fclass FileTypeClass
				switch index.cc_head {
				case CCLetter:
					if -1 != index.ix_date && CCNumber == index.cc_date {

						fclass = FileClassTable
					}
				case CCNumber:
					if -1 == index.ix_date && CCNone == index.cc_date {

						fclass = FileClassIndex
					}
				}
				var ftype FileTypeClass
				if -1 != index.ix_fext && CCLetter == index.cc_fext {
					var begin int = (index.ix_fext + 1)
					var fext string = location[begin:]
					switch fext {
					case "txt":
						ftype = FileTypeTXT
					case "json":
						ftype = FileTypeJSN
					case "html":
						ftype = FileTypeHTL
					case "svg":
						ftype = FileTypeSVG
					case "png":
						ftype = FileTypePNG
					case "jpg", "jpeg":
						ftype = FileTypeJPG
					}
				}
				var typeclass FileTypeClass = (fclass | ftype)
				index.typeclass = typeclass
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

		switch this.typeclass & FileClass {
		case FileClassIndex:
			if 0 < this.ix_head && 0 < this.ix_fext {
				var begin, end int = this.ix_head + 1, this.ix_fext
				that.datetime = DateTime(string(this.location[begin:end]))
				that.dirname = string(this.location[:begin])
				that.basename = string(this.location[begin:end])
			}
		case FileClassTable:
			if 0 < this.ix_head && 0 < this.ix_date {
				var begin, end int = this.ix_head + 1, this.ix_date
				that.tablename = TableName(string(this.location[begin:end]))
				that.dirname = string(this.location[:begin])
				if 0 < this.ix_fext {
					end = this.ix_fext
					that.basename = string(this.location[begin:end])
					if 0 < this.ix_date {
						var begin, end int = this.ix_date + 1, this.ix_fext
						that.datetime = DateTime(string(this.location[begin:end]))
					}
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
/*
 * Innumerable network data source.
 */
func (this FileLocation) IsAbstractClass() bool {
	return (FileClassAbstract == (this.typeclass & FileClassAbstract))
}
/*
 * Enumerable file data source.
 */
func (this FileLocation) IsNotAbstractClass() bool {
	return (FileClassAbstract != (this.typeclass & FileClassAbstract))
}

func (this FileLocation) IsIndexClass() bool {
	return (FileClassIndex == (this.typeclass & FileClassIndex))
}

func (this FileLocation) IsTableClass() bool {
	return (FileClassTable == (this.typeclass & FileClassTable))
}

func (this FileLocation) String() string {
	return this.location
}

func (this FileLocation) PathLocation() string {
	return this.location
}

func (this FileLocation) WebLocation() string {
	var str strings.Builder
	{
		str.WriteByte('/')
		str.WriteString(this.location)
	}
	return str.String()
}

func (this FileLocation) DirName() string {
	return this.dirname
}

func (this FileLocation) TableName() TableName {
	return TableName(this.tablename)
}

func (this FileLocation) BaseName() string {
	return this.basename
}

func (this FileLocation) HasDatetime() bool {

	return this.datetime.IsValid()
}

func (this FileLocation) YYYY() string {

	return this.datetime.YYYY()
}

func (this FileLocation) MM() string {

	return this.datetime.MM()
}

func (this FileLocation) YYYYMM() string {

	return this.datetime.YYYYMM()
}

func (this FileLocation) YYYYMMDD() string {

	return this.datetime.YYYYMMDD()
}

func (this FileLocation) YYYYMMDD_HHMMSS() string {

	return this.datetime.YYYYMMDD_HHMMSS()
}

func (this FileLocation) HHMMSS() string {

	return this.datetime.HHMMSS()
}

/*
 * <ID> := <YYYYMMDD_HHMMSS>
 */
func (this FileLocation) FileIdentifier() FileId {

	if this.datetime.IsLong() {

		return FileId(this.YYYYMMDD_HHMMSS())

	} else if this.datetime.IsShort() {

		return FileId(this.YYYYMM())
	} else {
		return FileId("")
	}
}

func (this FileId) IsValid() bool {
	switch len(this) {
	case 8, 15:
		return true
	default:
		return false
	}
}

/*
 * <IX> := <YYYYMM>
 */
func (this FileLocation) FileIndex() FileIx {

	if this.datetime.IsValid() {

		return FileIx(this.datetime.YYYYMM())

	} else {
		return FileIx("")
	}
}

func (this FileIx) IsValid() bool {
	return 6 == len(this)
}

/*
 * Structural analogue to Catalog#FileCatalog: <ID> =
 * <YYYYMMDD_HHMMSS>.
 */
func (this FileLocation) TableAnchor() string {
	if this.IsTableClass() {
		return string(this.tablename) + "/" + string(this.YYYYMMDD()) + "/" + this.YYYYMMDD_HHMMSS()
	} else {
		return ""
	}
}

func (this FileLocation) TableTabulate() string {

	if this.IsTableClass() {
		var anchor string = this.TableAnchor()
		var catalog Catalog = this.FileCatalog()

		return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s", anchor, catalog.id, catalog.icon, catalog.path, catalog.link, catalog.name, catalog.embed)
	} else {
		return ""
	}
}

/*
 * File type class derivation with target location.
 */
func (this FileLocation) Target(to FileTypeClass) (empty FileLocation) {
	var from FileTypeClass = (this.typeclass & FileType)
	to = (to & FileType)
	if from == to {

		return this

	} else if this.HasDatetime() {
		/*
		 * Express target state.
		 */
		var infix string = path.Join(this.YYYY(), this.MM())
		var prefix string = path.Join(OperandTarget(), infix, this.basename)
		var to_string string
		switch to {
		case FileTypeTXT:
			to_string = prefix + ".txt"
		case FileTypeJSN:
			to_string = prefix + ".json"
		case FileTypeHTL:
			to_string = prefix + ".html"
		case FileTypeSVG:
			to_string = prefix + ".svg"
		case FileTypePNG:
			to_string = prefix + ".png"
		case FileTypeJPG:
			to_string = prefix + ".jpeg"
		default:
			return empty
		}
		/*
		 * Classify and condense target location.
		 */
		var target FileLocation = FileClassify(to_string).Condense()
		/*
		 * Conserve target source.  The operand
		 * source location is conserved by copying
		 * this location to the target source field.
		 */
		target.source = this.location

		return target
	} else {
		return empty
	}
}

/*
 * File type class derivation within source location enabled
 * by the conservation of source operand location under
 * location transformations.
 */
func (this FileLocation) Source(to FileTypeClass) (empty FileLocation) {
	var tgt FileTypeClass = (this.typeclass & FileType)
	to = (to & FileType)
	if tgt == to {

		return this

	} else if 0 != len(this.source) {
		/*
		 * Classify and condense target location.
		 */
		var index FileIndex = FileClassify(this.source)

		var src FileTypeClass = (index.typeclass & FileType)
		if src == to {

			var source FileLocation = index.Condense()
			/*
			 * Source semantics conserve main
			 * operand source location.  In this
			 * case, a target is not a source as
			 * conservative of source semantics.
			 * The return value is a source as
			 * located in the operand source
			 * location, so the source field of
			 * the file location data structure
			 * is empty.
			 */
			return source
		} else {
			/*
			 * Express derivative state.
			 */
			var bax, fex int = index.ix_head, index.ix_fext
			if -1 < bax && -1 < fex {
				var prefix = this.location[0:fex]

				var to_string string
				switch to {
				case FileTypeTXT:
					to_string = prefix + ".txt"
				case FileTypeJSN:
					to_string = prefix + ".json"
				case FileTypeHTL:
					to_string = prefix + ".html"
				case FileTypeSVG:
					to_string = prefix + ".svg"
				case FileTypePNG:
					to_string = prefix + ".png"
				case FileTypeJPG:
					to_string = prefix + ".jpeg"
				default:
					return empty
				}
				/*
				 * Classify and condense derivative.
				 */
				var derivative FileLocation = FileClassify(to_string).Condense()
				/*
				 * Conserve derivative source.
				 */
				derivative.source = this.source

				return derivative
			}
		}
	}
	return empty
}

func (this FileLocation) Exists() bool {
	var finf os.FileInfo
	var er error

	finf, er = os.Stat(this.location)
	if nil == er {
		return 0 < finf.Size()
	} else {
		return false
	}
}

func (this FileLocation) NotExists() bool {
	var finf os.FileInfo
	var er error

	finf, er = os.Stat(this.location)
	if nil == er {
		return 1 > finf.Size()
	} else {
		return true
	}
}

func (this FileLocation) Read() []byte {
	var file *os.File
	var er error

	file, er = os.Open(this.location)
	if nil == er {

		var info os.FileInfo
		info, er = file.Stat()
		if nil == er {

			var size int64 = info.Size()
			if 0 < size && size < 0x7FFFFFFF {

				var z uint32 = uint32(size)
				var b []byte = make([]byte, z)
				var n int
				n, er = file.Read(b)
				if nil == er && n == int(z) {
					file.Close()

					return b

				} else {
					file.Close()
				}
			} else {
				file.Close()
			}
		}
	}
	return nil
}

func (this FileLocation) Write(content []byte) {
	if 0 != len(content) {
		var file *os.File
		var er error

		os.MkdirAll(this.dirname,0755) 

		file, er = os.Create(this.location)
		if nil == er {

			file.Write(content)

			file.Close()
		}
	}
}

func (this FileLocationList) List() (list []FileLocation) {
	for _, file := range this {

		list = append(list, file)
	}
	return list
}

/*
 * Index order is key descending sort product order.
 */
func FileSort(this []FileLocation) (array []FileLocation) {
	array = this

	var a, b int = 0, len(array)
	var x, y FileLocation
	/*
	 * Partition sort adapted from GOST
	 */
	for i := a + 1; i < b; i++ {

		for j := i; j > a; j-- {

			x = array[j]
			y = array[j-1]

			if x.FileIdentifier() > y.FileIdentifier() {

				array[j] = y
				array[j-1] = x
			} else {
				break
			}
		}
	}
	return array
}

func PathSplit(path string) (base, fext int) {
	{
		base = -1
		fext = -1
	}

	var z int = len(path)
	if 0 < z {
		var y int = (z - 1)
		if 0 < y {
			var x int = (y - 1)
			var ch byte

		scan:
			for ; 0 < x; x-- {
				ch = path[x]
				switch ch {
				case '/':
					base = x
					break scan
				case '.':
					fext = x
				}
			}
		}
	}
	return base, fext
}

/*
 * Given a file or directory, derive a list of files,
 * conserving relative and absolute path names.
 */
func FileList(src string) (list []string) {
	var finf os.FileInfo
	var dli []os.DirEntry
	var err error
	finf, err = os.Stat(src)
	if nil == err {

		if finf.IsDir() {
			dli, err = os.ReadDir(src)
			if nil == err {
				for _, dent := range dli {

					if dent.IsDir() {
						var name = path.Join(src, dent.Name())
						var slist []string = FileList(name)
						for _, sln := range slist {

							list = append(list, sln)
						}
					} else {
						var name = path.Join(src, dent.Name())

						list = append(list, name)
					}
				}
			}
		} else {
			list = append(list, src)
		}
	}
	return list
}
