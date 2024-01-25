/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

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
	for argx, arg := range argv {
		switch arg {
		case "no", "not", "notes", "re", "rec", "recent":
			if 0 != (Configuration & Class_Context) {
				return false
			} else {
				switch arg {
				case "no", "not", "notes":
					Configuration |= ClassNotes
					Context = "notes"
					return true
				case "re", "rec", "recent":
					Configuration |= ClassRecent
					Context = "recent"
					return true
				}

			}
		case "src", "source", "tgt", "target":
			if 0 != (Configuration & Class_Operation) {
				return false
			} else {
				switch arg {
				case "src", "source":
					Configuration |= ClassSource
					return true
				case "tgt", "target":
					Configuration |= ClassTarget
					return true
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
					return true
				case "upd", "update":
					Configuration |= ClassUpdate
					Operator = "update"
					return true
				case "con", "contents":
					Configuration |= ClassContents
					Operator = "contents"
					return true
				case "tab", "tabulate":
					Configuration |= ClassTabulate
					Operator = "tabulate"
					return true
				}
			}
		default:
			if 0 != (Configuration & Class_Context) &&
				0 != (Configuration & Class_Transform) {

				if nil == Operands {

					Operands = argv[argx:]

					return true
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

func ConfigurationSource() FileTypeClass { // [TODO] (review)

	switch ConfigurationContext() {

	case ClassNotes:

		return FileClassTable|FileTypeTXT

	case ClassRecent:

		return FileClassIndex|FileTypeJSN

	default:
		return 0
	}
}

func ConfigurationTarget() FileTypeClass { // [TODO] (review)

	switch ConfigurationContext() {

	case ClassNotes:

		return FileClassTable|FileTypeSVG

	case ClassRecent:

		return FileClassIndex|FileTypeJSN

	default:
		return 0
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
