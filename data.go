/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

type Coder interface {
	Decode([]byte)
	Encode() []byte
}

func DataTransform() bool {
	return false
}
