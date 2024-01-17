/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"bufio"
	json "github.com/syntelos/go-json"
	"fmt"
	"io/fs"
	"log"
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

var NotesTargetIndex map[IndexFile]IndexTarget = make(map[IndexFile]IndexTarget)

func indexListWalker(path string, d fs.DirEntry, er error) error {

	if ! d.IsDir() {
		var ixfil IndexFile = IndexFile(path)

		if IndexFileTypeSVG == ixfil.FileType() {

			var a IndexTarget = ixfil.Target()

			var b IndexTarget = NotesTargetIndex[a.yyyymm]

			if b.IsInvalid() || a.yyyymmdd > b.yyyymmdd {
				NotesTargetIndex[a.yyyymm] = a
			}
		}
	}
	return nil
}

func ListIndexFiles() (fileList IndexTargetList) {
	/*
	 * Collect index map
	 */
	if 0 == len(NotesTargetIndex) {

		var dir fs.FS = os.DirFS(".")

		fs.WalkDir(dir,string(ObjectiveDirectory(ObjectiveKeyTargetWeb)),indexListWalker)
	}
	/*
	 * Serialize index map
	 */
	{
		for _, v := range NotesTargetIndex {

			if v.IsValid() {

				fileList = append(fileList,v)
			}
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

func (this IndexFile) Target() (empty IndexTarget) {
	var that IndexTarget
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
				if -1 == prefix && -1 != infix {

					prefix = ofs
					that.dir = this[0:ofs]

					that.path = IndexFile(FileCat(string(that.dir),string(that.yyyymmdd))+"."+string(IndexFextJSN))
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
/*
 * Read objective target index.  N.B. The objective target
 * index may differ from the original index source list by
 * manual intervention.
 */
func (this IndexTarget) IndexRead() (list []IndexCatalog) {
	var er error
	var fo *os.File

	fo, er = os.Open(string(this.path))
	if nil == er {
		defer fo.Close()

		var reader json.Reader = json.ReadFile(fo)
		if reader.IsNotEmpty() {

			var array json.Reader = reader.HeadArray()
			if array.IsNotEmpty() {

				var object json.Reader = array.HeadObject()

				for object.IsNotEmpty() {

					var field_id json.Reader = object.HeadField()

					if field_id.IsNotEmpty() && object.Contains(field_id) {
						var field_ic json.Reader = field_id.TailField()
						if field_ic.IsNotEmpty() && object.Contains(field_ic) {
							var field_pa json.Reader = field_ic.TailField()
							if field_pa.IsNotEmpty() && object.Contains(field_pa) {
								var field_li json.Reader = field_pa.TailField()
								if field_li.IsNotEmpty() && object.Contains(field_li) {
									var field_na json.Reader = field_li.TailField()
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.TailField()
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										}
									}
								}
							}
						}
					}
					object = object.TailObject()
				}

			} else {
				log.Fatalf("Reading '%s': empty JSON Array",this.path)
			}
		} else {
			log.Fatalf("Reading '%s': %v",this.path,er)
		}
	} else {
		log.Fatalf("Reading '%s': %v",this.path,er)
	}
	return list
}

func (this IndexCatalog) String() string {
	return fmt.Sprintf(`    {
        "id": "%s",
        "icon": "%s",
        "path": "%s",
        "link": "%s",
        "name": "%s",
        "embed": "%s"
    }`,this.id,this.icon,this.path,this.link,this.name,this.embed)
}

func Trim(value string) (empty string) {
	var vz int = len(value)
	if 0 < vz {
		var first, last int = 0, (vz-1)

		if '"' == value[first] && '"' == value[last] {

			return value[first+1:last]
		}
	}
	return empty
}
