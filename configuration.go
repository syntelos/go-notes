/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import "io/fs"

type Class uint16

const (
	ClassNotes   Class  = 0b0100000000000000
	ClassRecent   Class = 0b0010000000000000

	ClassSource   Class = 0b0001000000000000
	ClassTarget   Class = 0b0000100000000000

	ClassEncode   Class = 0b0000010000000000
	ClassUpdate   Class = 0b0000001000000000
	ClassContents Class = 0b0000000100000000
	ClassTabulate Class = 0b0000000010000000
)

const Class_Context = (ClassRecent|ClassNotes)
const Class_Operation = (ClassSource|ClassTarget)
const Class_Transform = (ClassEncode|ClassUpdate|ClassContents|ClassTabulate)

var Configuration Class = 0
var Context string
var Operator string
var Operands []string = nil

func Configure(argv []string) bool {
	Configuration = 0
	var argc int = len(argv)
	if 3 <= argc {
		var argx int = 0
		/*
		 * Extensible contextual wwweb root target
		 * directory.
		 */
		var cx string = argv[argx]
		switch cx {
		case "not", "notes":
			Configuration |= ClassNotes
			Context = "notes"
		case "rec", "recent":
			Configuration |= ClassRecent
			Context = "recent"
		default:
			if fs.ValidPath(cx) {
				Configuration |= ClassNotes
				Context = cx
			} else {
				return false
			}
		}
		argx += 1
		/*
		 * Operational wwweb production argumentation.
		 *
		 * This algorithm assures that operands
		 * follow operators, but is ambivalent to
		 * the ordering of multiple operators
		 * (i.e. "source encode", "target update").
		 */
		for rix, arg := range argv[argx:] {
			switch arg {

			case "src", "source", "tgt", "target":
				if 0 != (Configuration & Class_Operation) {
					return false
				} else {
					switch arg {
					case "src", "source":
						Configuration |= ClassSource

					case "tgt", "target":
						Configuration |= ClassTarget
					}
				}
			case "enc", "encode", "upd", "update", "con", "contents", "tab", "tabulate":
				if 0 != (Configuration & Class_Transform) {
					return false
				} else {
					switch arg {
					case "enc", "encode":
						Configuration |= ClassEncode
						Operator = "encode"

					case "upd", "update":
						Configuration |= ClassUpdate
						Operator = "update"

					case "con", "contents":
						Configuration |= ClassContents
						Operator = "contents"

					case "tab", "tabulate":
						Configuration |= ClassTabulate
						Operator = "tabulate"
					}
				}
			default:
				if 0 != (Configuration & Class_Context) &&
					0 != (Configuration & Class_Transform) {

					Operands = argv[rix+argx:]

					return SourceDefine() && TargetDefine()
				}
			}
		}
	}
	return false
}

func ConfigurationContext() Class {
	return (Configuration & Class_Context)
}

func ConfigurationOperation() Class {
	return (Configuration & Class_Operation)
}

func ConfigurationTransform() Class {
	return (Configuration & Class_Transform)
}

func ConfigurationContextDirectory() string {
	return Context
}

func ConfigurationSource() FileTypeClass {

	switch ConfigurationContext() {

	case ClassNotes:

		switch ConfigurationTransform() {

		case ClassEncode:
			return FileClassTable|FileTypeTXT

		case ClassUpdate:
			return FileClassTable|FileTypeSVG

		default:
			return 0
		}

	default:
		return 0
	}
}

func ConfigurationTarget() FileTypeClass {

	switch ConfigurationContext() {

	case ClassNotes:

		switch ConfigurationTransform() {

		case ClassEncode:
			return FileClassTable|FileTypeSVG

		case ClassUpdate:
			return FileClassTable|FileTypeJSN

		default:
			return 0
		}


	case ClassRecent:

		switch ConfigurationTransform() {

		case ClassUpdate:
			return FileClassIndex|FileTypeJSN

		default:
			return 0
		}


	default:
		return 0
	}
}

func HaveContext() bool {
	if 0 != (Configuration & Class_Context) && 0 != len(Context) {

		return true
	} else {
		return false
	}
}

func HaveOperand(index int) bool {

	return index < len(Operands)
}

func Operand(index int) string {

	if index < len(Operands) {

		return Operands[index]
	} else {
		return ""
	}
}

func OperandTarget() string {

	if HaveOperand(0) {

		return Operand(0)
	} else {
		return Context
	}
}
