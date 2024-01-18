/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"bufio"
	"io/fs"
	"os"
	sort "github.com/syntelos/go-sort"
)

type IndexFile FileName

const (
	IndexFextTXT IndexFile = "txt"
	IndexFextSVG IndexFile = "svg"
	IndexFextJSN IndexFile = "json"
)

type IndexList []IndexFile

type IndexTarget struct {
	dir IndexFile
	yyyymmdd_hhmmss IndexFile
	yyyymmdd IndexFile
	yyyymm IndexFile
	path IndexFile
	name TableName
}

type IndexTargetList []IndexTarget

type IndexCatalog struct {
	target IndexTarget
	id, icon, path, link, name, embed string
}

var sourceObjectiveList []IndexFile

var targetObjectiveList []IndexFile

var condensedObjectiveIndex map[IndexFile]IndexTarget = make(map[IndexFile]IndexTarget)

func indexListWalker(path string, d fs.DirEntry, er error) error {

	if nil != d && ! d.IsDir() {
		var ixfil IndexFile = IndexFile(path)

		if IndexFileTypeSVG == ixfil.FileType() {

			sourceObjectiveList = append(sourceObjectiveList,ixfil)

			var a IndexTarget = ixfil.Target()

			if a.IsValid() {

				var b IndexTarget = condensedObjectiveIndex[a.yyyymm]

				if b.IsInvalid() || a.yyyymmdd > b.yyyymmdd {

					condensedObjectiveIndex[a.yyyymm] = a
				}
			}
		} else if IndexFileTypeJSN == ixfil.FileType() {

			targetObjectiveList = append(targetObjectiveList,ixfil)
		}
	}
	return nil
}

func defineIndex(tgt string) {

	var dir fs.FS = os.DirFS(".")

	fs.WalkDir(dir,tgt,indexListWalker)
}

func ListIndexFiles() (list IndexTargetList) {

	for _, v := range condensedObjectiveIndex {

		list = append(list,v)
	}
	return list
}

func ListIndexSource() (list []IndexFile) {

	for _, v := range sourceObjectiveList {

		list = append(list,v)
	}
	return list
}

func ListIndexTarget() (list []IndexFile) {

	for _, v := range targetObjectiveList {

		list = append(list,v)
	}
	return list
}

type IndexFileType uint8

const (
	IndexFileTypeUNK IndexFileType = 0b00000000
	IndexFileTypeTXT IndexFileType = 0b00000001
	IndexFileTypeSVG IndexFileType = 0b00000010
	IndexFileTypeJSN IndexFileType = 0b00000100
)

func (this IndexFile) IsTXT() bool {
	return (IndexFileTypeTXT == this.FileType())
}
func (this IndexFile) IsSVG() bool {
	return (IndexFileTypeSVG == this.FileType())
}
func (this IndexFile) IsJSN() bool {
	return (IndexFileTypeJSN == this.FileType())
}

func (this IndexFile) FileType() IndexFileType {

	var first, end int = 0, len(this)

	first = (end-3)

	if 1 < first && '.' ==  this[first-1]{
		var fext IndexFile = this[first:end]

		switch fext {
		case IndexFextTXT:
			return IndexFileTypeTXT
		case IndexFextSVG:
			return IndexFileTypeSVG

		default:
			return IndexFileTypeUNK
		}

	} else {
		first = (end-4)

		if 1 < first && '.' ==  this[first-1] {
			var fext IndexFile = this[first:end]

			switch fext {
			case IndexFextJSN:
				return IndexFileTypeJSN

			default:
				return IndexFileTypeUNK
			}
		} else {
			return IndexFileTypeUNK
		}
	}
}

func (this IndexFile) LongKey() (that IndexFile) {

	switch this.FileType() {
	case IndexFileTypeTXT, IndexFileTypeSVG:
		var infix, postfix int = -1, -1

		var ofs, len int = 0, len(this)

		for ofs = (len-1); 0 < ofs; ofs-- {

			switch this[ofs] {

			case '-':
				if -1 == infix && 0 < postfix {
					infix = ofs

					return this[infix+1:postfix]
				}

			case '.':
				if -1 == postfix {
					postfix = ofs
				}
			}
		}
		return that

	default:
		return that
	}
}

func (this IndexFile) Target() (empty IndexTarget) {
	var ctor IndexTarget
	/*
	 * Parse filepath into directory, tablename,
	 * datetime, and filename extension.
	 */
	switch this.FileType() {
	case IndexFileTypeTXT, IndexFileTypeSVG:
		var prefix, infix, postfix, ppostfix int = -1, -1, -1, -1

		var ofs, len int = 0, len(this)

		for ofs = (len-2); 0 < ofs; ofs-- {

			switch this[ofs] {
			case '/':
				if -1 == prefix && -1 != infix {

					prefix = ofs
					ctor.dir = this[0:ofs]

					ctor.path = IndexFile(FileCat(string(ctor.dir),string(ctor.yyyymmdd))+"."+string(IndexFextJSN))
					ctor.name = TableName(this[prefix+1:infix])

					return ctor
				} else {
					return empty
				}
			case '-':
				if -1 == infix && 0 < postfix {
					infix = ofs
					ctor.yyyymmdd_hhmmss = this[infix+1:ppostfix]
					ctor.yyyymmdd = this[infix+1:postfix]
					ctor.yyyymm = ctor.yyyymmdd[0:6]
				}
			case '_':
				if -1 == postfix {
					postfix = ofs
				}
			case '.':
				if -1 == ppostfix {
					ppostfix = ofs
				}
			}
		}
	}
	return empty
}

func (this IndexTarget) IsInvalid() bool {

	return 0 == len(this.dir) || 0 == len(this.yyyymmdd_hhmmss) || 0 == len(this.yyyymmdd) || 0 == len(this.yyyymm) || 0 == len(this.path)
}

func (this IndexTarget) IsValid() bool {

	return 0 != len(this.dir) && 0 != len(this.yyyymmdd_hhmmss) && 0 != len(this.yyyymmdd) && 0 != len(this.yyyymm) && 0 != len(this.path)
}

func (this IndexTarget) Path() string {
	return string(this.path)
}

func (this IndexTarget) Name() string {
	return string(this.name)
}

func (this IndexTarget) Target() (empty FileName) {
	if this.IsValid() {
		var yyyy string = string(this.yyyymm[0:4])
		var mm string = string(this.yyyymm[4:6])

		if HaveObjective(ObjectiveKeyTargetWeb) {
			var root FileName = ObjectiveDirectory(ObjectiveKeyTargetWeb)

			return FileName(FileCat(FileCat(string(root),yyyy),mm))

		} else {
			return FileName(FileCat(FileCat("notes",yyyy),mm))
		}
	} else {
		return empty
	}
}
/*
 * An original source listing often differs from an
 * objective target listing because the file has been edited
 * manually since its initial production, as to include an
 * arbitrarily expanded date time range.
 */
func (this IndexTarget) IndexSourceList() (list []IndexCatalog) {

	var dl []os.DirEntry
	var er error

	dl, er = os.ReadDir(string(this.dir))
	if nil == er {
		/*
		 * Source list
		 */
		var ordering IndexList
		var directory map[IndexFile]IndexFile = make(map[IndexFile]IndexFile)
		{
			for _, de := range dl {

				var notes_svg IndexFile = IndexFile(FileCat(string(this.dir),de.Name()))

				if notes_svg.IsSVG(){

					var key IndexFile = notes_svg.LongKey()

					directory[key] = notes_svg

					ordering = append(ordering,key)
				}
			}

			sort.Descending(ordering)
		}
		/*
		 * Order list
		 */
		for _, key := range ordering {

			var notes_svg IndexFile = "/"+directory[key]
			var key IndexFile = notes_svg.LongKey()
			/*
			 * N.B. "this name" may be unrelated
			 * to "notes name", so it is derived
			 * from the source.
			 */
			var notes_idx IndexTarget = notes_svg.Target()

			var name TableName = notes_idx.name
			var path TablePath = name.Path()
			var link TableLink = name.Link()
			/*
			 * N.B. "this target" is
			 * objective, not
			 * figurative.
			 */
			var catalog IndexCatalog = IndexCatalog{this,string(key),"syntelos-catalog",string(path),string(link),string(name),string(notes_svg)}

			list = append(list,catalog)
		}
	}
	return list
}
/*
 * Write index source list to objective target file.
 */
func (this IndexTarget) IndexWrite() {
	/*
	 * Don't overwrite an existing target.
	 */
	var dir fs.FS = os.DirFS(".")
	var er error

	_, er = fs.Stat(dir,string(this.path))
	if nil != er {
	
		var tgt *os.File
		tgt, er = os.Create(string(this.path))
		if nil == er {
			var w *bufio.Writer = bufio.NewWriter(tgt)
					
			w.Write([]byte("[\n"))
			for x, record := range this.IndexSourceList() {

				if 0 != x {
					w.Write([]byte(",\n"))
				}
				w.Write([]byte(record.String()))
			}
			w.Write([]byte("\n]\n"))

			w.Flush()
			tgt.Close()
		}
	}
}
