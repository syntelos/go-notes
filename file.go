/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"os"
)

type FileName string
type FileList []FileName

func ListTextFiles(path string) (fileList FileList) {
	var fo *os.File
	var er error

	fo, er = os.Open(path)
	if nil != er {
		return fileList
	} else {
		defer fo.Close()

		var fi os.FileInfo
		fi, er = fo.Stat()
		if nil != er {
			return fileList
		} else {
			if fi.IsDir() {
				var nm FileName
				var dl []os.DirEntry
				dl, er = os.ReadDir(path)

				for _, de := range dl {
					nm = MakeFileName(path,de)
					if nm.IsText() && IsTableName(nm.TableName()) {
						
						fileList = append(fileList,nm)
					}
				}
			} else {
				var nm FileName = FileName(path)
				if nm.IsText() && IsTableName(nm.TableName()) {

					fileList = append(fileList,nm)
				}
			}
			return fileList
		}
	}
}

func MakeFileName(p string, de os.DirEntry) (fn FileName) {

	return FileName(FileCat(p,de.Name()))
}

func FileCat(a, b string) string {
	var end int = len(a)
	if 0 < end {
		var last int = (end-1)

		if '/' == a[last] {

			return (a+b)
		} else {
			return (a+"/"+b)
		}
	} else {
		return b
	}
}

func (this FileName) IsText() bool {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] && ".txt" == this[first:end]{

		return true
	} else {
		return false
	}
}

func (this FileName) Base() (that FileName) {

	var x, z = 0, len(this)

	for x = (z-1); 0 <= x; x-- {

		if '/' == this[x] {

			that = FileName(this[x+1:z])

			return that
		}
	}
	return this
}

func (this FileName) TableName() TableName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		var head FileName = this[0:first]

		var x, z = 0, len(head)

		for x = (z-1); 0 <= x; x-- {

			if '-' == head[x] {

				head = head[0:x]

				break
			}
		}

		z = len(head)

		for x = (z-1); 0 <= x; x-- {

			if '/' == head[x] {

				return TableName(head[x+1:z])
			}
		}
		return TableName(head)

	} else {
		return ""
	}
}

func (this FileName) Source(fext string) FileName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		var found string = string(this[(first+1):end])

		if ".txt" == found {

			return this
		} else {

			return FileName(string(this[0:first])+"."+fext)
		}
	} else {
		return ""
	}
}

func (this FileName) Target(fext string) FileName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {
		/*
		 * The target reflects the source.
		 */
		var reflection FileName

		var found string = string(this[(first+1):end])

		if ".svg" == found {

			reflection = this
		} else {

			reflection = FileName(string(this[0:first])+"."+fext)
		}

		if HaveObjective(ObjectiveKeyTargetWeb) {
			/*
			 * The target is a projection from
			 * the source into "notes".
			 */
			var projection FileName = IndexFile(reflection).Target().Target()
			var filename FileName = reflection.Base()

			return FileName(FileCat(string(projection),string(filename)))

		} else {

			return reflection
		}
	} else {
		return ""
	}
}
