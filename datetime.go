/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"time"
)

func YYYY() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%04d",t.Year())
}

func MM() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%02d",t.Month())
}

func DD() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%02d",t.Day())
}

func YYYYMM() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%04d%02d",t.Year(),t.Month())
}

func YYYYMMDD() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%04d%02d%02d",t.Year(),t.Month(),t.Day())
}

func YYYYMMDD_HHMMSS() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%04d%02d%02d_%02d%02d%02d",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
}

func HHMMSS() string {
	var t time.Time = time.Now().Local()
	return fmt.Sprintf("%02d%02d%02d",t.Hour(),t.Minute(),t.Second())
}
/*
 * Date time "YYYYMMDD_HHMMSS" (TABLE) or "YYYYMMDD" (INDEX)
 * string.
 */
type DateTime string

func NewDateTime() DateTime {

	return DateTime(YYYYMMDD_HHMMSS())
}

func (this FileLocation) FileDateTime() DateTime {

	return DateTime(this.YYYYMMDD_HHMMSS())
}

func (this DateTime) IsValid() bool {

	return this.IsLong() || this.IsShort()
}

func (this DateTime) IsLong() bool {

	return (15 == len(this) && '_' == this[8])
}

func (this DateTime) IsShort() bool {

	return (8 == len(this))
}

func (this DateTime) YYYY() string {

	if 15 <= len(this) && '_' == this[8] {
		return string(this[0:4])
	} else {
		return ""
	}
}

func (this DateTime) MM() string {

	if 15 <= len(this) && '_' == this[8] {
		return string(this[4:6])
	} else {
		return ""
	}
}

func (this DateTime) YYYYMM() string {

	if 15 <= len(this) && '_' == this[8] {
		return string(this[0:6])
	} else {
		return ""
	}
}

func (this DateTime) YYYYMMDD() string {

	if 15 <= len(this) && '_' == this[8] {
		return string(this[0:8])
	} else {
		return ""
	}
}

func (this DateTime) YYYYMMDD_HHMMSS() string {

	if 15 <= len(this) && '_' == this[8] {
		return string(this[0:15])
	} else {
		return ""
	}
}

func (this DateTime) HHMMSS() string {

	if 15 <= len(this) && '_' == this[8] {
		return string(this[10:15])
	} else {
		return ""
	}
}
