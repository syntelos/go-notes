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
					nm = FileName(de.Name())
					if nm.IsText() && IsTableName(nm.Base()) {
						
						fileList = append(fileList,nm)
					}
				}
				return fileList

			} else {
				var nm FileName = FileName(path)
				if nm.IsText() && IsTableName(nm.Base()) {

					fileList = append(fileList,nm)
					return fileList
				} else {
					return fileList
				}
			}
		}
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

func (this FileName) Base() TableName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		var head FileName = this[0:first]

		var x, z = 0, len(head)

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

func (this FileName) Source() FileName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		if ".txt" == this[first:end] {

			return this
		} else {

			return this[0:first]+".txt"
		}
	} else {
		return ""
	}
}

func (this FileName) Target() FileName {

	var first, last, end int = 0, 0, len(this)
	{
		last = (end-1)
		first = (last-3)
	}

	if 0 < first && '.' ==  this[first] {

		if ".svg" == this[first:end] {

			return this
		} else {

			return this[0:first]+".svg"
		}
	} else {
		return ""
	}
}
