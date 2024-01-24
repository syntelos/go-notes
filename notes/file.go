/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"os"
)

func DefineTextFiles(list []string) {
	/*
	 * Unique membership
	 */
	var unique map[IndexFile]bool = make(map[IndexFile]bool)
	{
		var filelist IndexFileList = sourceObjectiveIndex[IndexFileTypeTXT]
		for _, file := range filelist {
			unique[file] = true
		}
	}
	/*
	 * Review membership
	 */
	for _, path := range list {
		var fo *os.File
		var er error

		fo, er = os.Open(path)
		if nil == er {

			var fi os.FileInfo
			fi, er = fo.Stat()

			if nil == er {
				
				if fi.IsDir() {
					var nm IndexFile
					var dl []os.DirEntry
					dl, er = os.ReadDir(path)

					for _, de := range dl {
						nm = MakeFile(path,de)

						if unique[nm] && nm.IsFext("txt") && IsTableName(nm.TableName()) {

							unique[nm] = true

							var fileList IndexFileList = sourceObjectiveIndex[IndexFileTypeTXT]

							fileList = append(fileList,IndexFile(nm))
							sourceObjectiveIndex[IndexFileTypeTXT] = fileList
						}
					}
				} else {
					var nm IndexFile = IndexFile(path)

					if unique[nm] && nm.IsFext("txt") && IsTableName(nm.TableName()) {

						unique[nm] = true

						var fileList IndexFileList = sourceObjectiveIndex[IndexFileTypeTXT]

						fileList = append(fileList,IndexFile(nm))
						sourceObjectiveIndex[IndexFileTypeTXT] = fileList
					}
				}
			}
			fo.Close()
		}
	}
}

func MakeFile(p string, de os.DirEntry) (fn IndexFile) {

	return FileCat(IndexFile(p),IndexFile(de.Name()))
}

func FileCat(a, b IndexFile) IndexFile {
	var end int = len(a)
	if 0 < end {
		var last int = (end-1)

		if '/' == a[last] {

			return IndexFile(a+b)
		} else {
			return IndexFile(a+"/"+b)
		}
	} else {
		return IndexFile(b)
	}
}

func (this IndexFile) IsFext(fext string) bool {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] && fext == string(this[(first+1):end]) {

		return true
	} else {
		return false
	}
}

func (this IndexFile) Base() (that IndexFile) {

	var x, z = 0, len(this)

	for x = (z-1); 0 <= x; x-- {

		if '/' == this[x] {

			that = IndexFile(this[x+1:z])

			return that
		}
	}
	return this
}

func (this IndexFile) TableName() TableName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		var head IndexFile = this[0:first]

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

func (this IndexFile) FileSource(fext string) IndexFile {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		var found string = string(this[(first+1):end])

		if fext == found {

			return this
		} else {

			return IndexFile(string(this[0:first])+"."+fext)
		}
	} else {
		return ""
	}
}

func (this IndexFile) FileTarget(fext string) IndexFile {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {
		/*
		 * The target reflects the source.
		 */
		var reflection IndexFile

		var found string = string(this[(first+1):end])

		if fext == found {

			reflection = this
		} else {

			reflection = IndexFile(string(this[0:first])+"."+fext)
		}

		if HaveObjective(ObjectiveKeyTargetWeb) {
			/*
			 * The target is a projection from
			 * the source into target.
			 */
			var projection IndexFile = IndexFile(reflection).IndexTarget().Target()
			var filename IndexFile = reflection.Base()

			return FileCat(projection,filename)

		} else {

			return reflection
		}
	} else {
		return ""
	}
}

func (this IndexFile) IsValid() bool {
	return (0 != len(this))
}

func (this IndexFile) IsNotValid() bool {
	return (0 == len(this))
}
