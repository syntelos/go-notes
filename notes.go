/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	sort "github.com/syntelos/go-sort"
)

const NotesTarget FileName = FileName("notes")

var   notesTarget FileName

func IsNotes() bool {

	return (NotesTarget == notesTarget)
}

func Init() bool {

	var fo *os.File
	var er error

	fo, er = os.Open(string(NotesTarget))
	if nil == er {
		defer fo.Close()

		var fi os.FileInfo
		fi, er = fo.Stat()
		if nil == er {

			if fi.IsDir() {

				notesTarget = NotesTarget
				return true
			}
		}
	}
	return false
}

type IndexFile FileName

const (
	IndexFextTXT IndexFile = "txt"
	IndexFextSVG IndexFile = "svg"
	IndexFextJSN IndexFile = "json"
)

type IndexList []IndexFile

var notesTargetIndex map[IndexFile]IndexTarget = make(map[IndexFile]IndexTarget)

func indexListWalker(path string, d fs.DirEntry, er error) error {

	if ! d.IsDir() {
		var ixfil IndexFile = IndexFile(path)

		if IndexFileTypeSVG == ixfil.FileType() {

			var a IndexTarget = ixfil.Target()

			var b IndexTarget = notesTargetIndex[a.yyyymm]

			if b.IsInvalid() || a.yyyymmdd > b.yyyymmdd {
				notesTargetIndex[a.yyyymm] = a
			}
		}
	}
	return nil
}

func ListIndexFiles() (fileList IndexTargetList) {
	/*
	 * Collect index map
	 */
	if 0 == len(notesTargetIndex) {

		var dir fs.FS = os.DirFS(".")

		fs.WalkDir(dir,"notes",indexListWalker)
	}
	/*
	 * Serialize index map
	 */
	{
		for _, v := range notesTargetIndex {

			fileList = append(fileList,v)
		}
	}
	return fileList
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

func (this IndexFile) Target() (that IndexTarget) {
	that.dir = ""
	that.yyyymmdd = ""
	that.yyyymm = ""
	that.path = ""
	that.name = ""
	/*
	 * Parse filepath into directory, filename,
	 * elements, and filename extention.
	 */
	switch this.FileType() {
	case IndexFileTypeTXT, IndexFileTypeSVG:
		var prefix, infix, postfix, ppostfix int = -1, -1, -1, -1

		var ofs, len int = 0, len(this)

		for ofs = (len-2); 0 < ofs; ofs-- {

			switch this[ofs] {
			case '/':
				if -1 == prefix {
					prefix = ofs
					that.dir = this[0:ofs]

					that.path = that.dir+"/"+that.yyyymmdd+"."+IndexFextJSN
					that.name = TableName(this[prefix+1:infix])

					return that
				}
			case '-':
				if -1 == infix && 0 < postfix {
					infix = ofs
					that.yyyymmdd_hhmmss = this[infix+1:ppostfix]
					that.yyyymmdd = this[infix+1:postfix]
					that.yyyymm = that.yyyymmdd[0:6]
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
		return that

	default:
		return that
	}
}

type IndexTarget struct {
	dir IndexFile
	yyyymmdd_hhmmss IndexFile
	yyyymmdd IndexFile
	yyyymm IndexFile
	path IndexFile
	name TableName
}

type IndexTargetList []IndexTarget

func (this IndexTarget) IsInvalid() bool {

	return 0 == len(this.dir) || 0 == len(this.yyyymmdd_hhmmss) || 0 == len(this.yyyymmdd) || 0 == len(this.yyyymm) || 0 == len(this.path)
}

func (this IndexTarget) IsValid() bool {

	return 0 != len(this.dir) && 0 != len(this.yyyymmdd_hhmmss) && 0 != len(this.yyyymmdd) && 0 != len(this.yyyymm) && 0 != len(this.path)
}

func (this IndexTarget) IndexWrite() {

	var dl []os.DirEntry
	var er error

	dl, er = os.ReadDir(string(this.dir))

	if nil == er {
		var ordering IndexList
		var directory map[IndexFile]IndexFile = make(map[IndexFile]IndexFile)

		for _, de := range dl {

			var notes_svg IndexFile = this.dir+IndexFile("/")+IndexFile(de.Name())

			if notes_svg.IsSVG(){

				var key IndexFile = notes_svg.LongKey()

				directory[key] = notes_svg

				ordering = append(ordering,key)
			}
		}

		sort.Sort(ordering)
		{
			var tgt *os.File
			tgt, er = os.Create(string(this.path))
			if nil == er {
				var w *bufio.Writer = bufio.NewWriter(tgt)
				
				w.Write([]byte("[\n"))
				for x, key := range ordering {

					var notes_svg IndexFile = directory[key]
					/*
					 * N.B. "this name" may be unrelated
					 * to "notes name", so it is derived
					 * from the source.
					 */
					var notes_idx IndexTarget = notes_svg.Target()

					var name TableName = notes_idx.name
					var path TablePath = name.Path()
					var link TableLink = name.Link()

					if 0 != x {
						w.Write([]byte(",\n"))
					}
					var record string = fmt.Sprintf(`    {
        "id": "%s",
        "icon": "syntelos-catalog",
        "path": "%s",
        "link": "%s",
        "name": "%s",
        "embed": "/%s"
    }`,key,path,link,name,notes_svg)
					w.Write([]byte(record))
				}
				w.Write([]byte("\n]\n"))

				w.Flush()
				tgt.Close()
			}
		}
	}
}
