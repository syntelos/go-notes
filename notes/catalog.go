/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

import (
	"bufio"
	"fmt"
	"io"
	json "github.com/syntelos/go-json"
	"log"
	"os"
)

type Catalog []Text

type Tabulation []Text

var catalog_head Catalog = Catalog{ Text{ 0x3C, 0x3F, 0x78, 0x6D, 0x6C, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6F, 0x6E, 0x3D, 0x22, 0x31, 0x2E, 0x30, 0x22, 0x20, 0x73, 0x74, 0x61, 0x6E, 0x64, 0x61, 0x6C, 0x6F, 0x6E, 0x65, 0x3D, 0x22, 0x6E, 0x6F, 0x22, 0x3F, 0x3E }, Text{ 0x3C, 0x3F, 0x78, 0x6D, 0x6C, 0x2D, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3D, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3A, 0x2F, 0x2F, 0x63, 0x64, 0x6E, 0x2E, 0x6A, 0x73, 0x64, 0x65, 0x6C, 0x69, 0x76, 0x72, 0x2E, 0x6E, 0x65, 0x74, 0x2F, 0x67, 0x68, 0x2F, 0x61, 0x61, 0x61, 0x61, 0x6B, 0x73, 0x68, 0x61, 0x74, 0x2F, 0x63, 0x6D, 0x2D, 0x77, 0x65, 0x62, 0x2D, 0x66, 0x6F, 0x6E, 0x74, 0x73, 0x40, 0x6C, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2F, 0x66, 0x6F, 0x6E, 0x74, 0x73, 0x2E, 0x63, 0x73, 0x73, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3D, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2F, 0x63, 0x73, 0x73, 0x22, 0x3F, 0x3E }, Text{ 0x3C, 0x21, 0x44, 0x4F, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x73, 0x76, 0x67, 0x20, 0x50, 0x55, 0x42, 0x4C, 0x49, 0x43, 0x20, 0x22, 0x2D, 0x2F, 0x2F, 0x57, 0x33, 0x43, 0x2F, 0x2F, 0x44, 0x54, 0x44, 0x20, 0x53, 0x56, 0x47, 0x20, 0x31, 0x2E, 0x31, 0x2F, 0x2F, 0x45, 0x4E, 0x22, 0x20, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3A, 0x2F, 0x2F, 0x77, 0x77, 0x77, 0x2E, 0x77, 0x33, 0x2E, 0x6F, 0x72, 0x67, 0x2F, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x2F, 0x53, 0x56, 0x47, 0x2F, 0x31, 0x2E, 0x31, 0x2F, 0x44, 0x54, 0x44, 0x2F, 0x73, 0x76, 0x67, 0x31, 0x31, 0x2E, 0x64, 0x74, 0x64, 0x22, 0x3E }, Text{ 0x3C, 0x73, 0x76, 0x67, 0x20, 0x78, 0x6D, 0x6C, 0x6E, 0x73, 0x3D, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3A, 0x2F, 0x2F, 0x77, 0x77, 0x77, 0x2E, 0x77, 0x33, 0x2E, 0x6F, 0x72, 0x67, 0x2F, 0x32, 0x30, 0x30, 0x30, 0x2F, 0x73, 0x76, 0x67, 0x22, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6F, 0x6E, 0x3D, 0x22, 0x31, 0x2E, 0x31, 0x22, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x76, 0x69, 0x65, 0x77, 0x42, 0x6F, 0x78, 0x3D, 0x22, 0x30, 0x20, 0x30, 0x20, 0x35, 0x30, 0x30, 0x20, 0x35, 0x30, 0x30, 0x22, 0x3E }, Text{ 0x20, 0x20, 0x3C, 0x64, 0x65, 0x66, 0x73, 0x3E }, Text{ 0x20, 0x20, 0x20, 0x20, 0x3C, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3D, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2F, 0x63, 0x73, 0x73, 0x22, 0x3E }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x7B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x66, 0x61, 0x6D, 0x69, 0x6C, 0x79, 0x3A, 0x20, 0x22, 0x43, 0x6F, 0x6D, 0x70, 0x75, 0x74, 0x65, 0x72, 0x20, 0x4D, 0x6F, 0x64, 0x65, 0x72, 0x6E, 0x20, 0x54, 0x79, 0x70, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x22, 0x2C, 0x20, 0x6D, 0x6F, 0x6E, 0x6F, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20, 0x31, 0x32, 0x3B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x69, 0x6C, 0x6C, 0x3A, 0x20, 0x62, 0x6C, 0x61, 0x63, 0x6B, 0x3B }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7D }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2E, 0x74, 0x69, 0x74, 0x6C, 0x65, 0x20, 0x7B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20, 0x32, 0x32, 0x3B }, Text{ 0x09, 0x20, 0x20, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3A, 0x20, 0x62, 0x6F, 0x6C, 0x64 }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7D }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2E, 0x6C, 0x69, 0x6E, 0x6B, 0x20, 0x7B }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x66, 0x69, 0x6C, 0x6C, 0x3A, 0x20, 0x23, 0x30, 0x30, 0x36, 0x36, 0x42, 0x33 }, Text{ 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7D }, Text{ 0x20, 0x20, 0x20, 0x20, 0x3C, 0x2F, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x3E }, Text{ 0x20, 0x20, 0x3C, 0x2F, 0x64, 0x65, 0x66, 0x73, 0x3E }, Text{ 0x20, 0x20, 0x3C, 0x72, 0x65, 0x63, 0x74, 0x20, 0x78, 0x3D, 0x22, 0x30, 0x22, 0x20, 0x79, 0x3D, 0x22, 0x30, 0x22, 0x20, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3D, 0x22, 0x35, 0x30, 0x30, 0x22, 0x20, 0x66, 0x69, 0x6C, 0x6C, 0x3D, 0x22, 0x23, 0x46, 0x46, 0x46, 0x22, 0x20, 0x73, 0x74, 0x72, 0x6F, 0x6B, 0x65, 0x3D, 0x22, 0x23, 0x30, 0x30, 0x30, 0x22, 0x20, 0x73, 0x74, 0x72, 0x6F, 0x6B, 0x65, 0x2D, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3D, 0x22, 0x33, 0x30, 0x22, 0x20, 0x2F, 0x3E }}

var catalog_tail Catalog = Catalog{ Text{ 0x3C, 0x2F, 0x73, 0x76, 0x67, 0x3E }}

func (this Catalog) Read(src *os.File) (that Catalog, er error) {
	var fi os.FileInfo
	fi, er = src.Stat()
	if nil != er {
		return nil, fmt.Errorf("Read error file not found: %w",er)
	} else {
		var reader *bufio.Reader = bufio.NewReader(src)
		var inl []byte
		var isp bool
		inl, isp, er = reader.ReadLine()
		for true {
			if nil != er {
				if io.EOF == er {
					return this, nil
				} else {
					return nil, fmt.Errorf("Read error from '%s': %w",fi.Name(),er)
				}
			} else if isp {
				return nil, fmt.Errorf("Read error from '%s'.",fi.Name())
			} else {
				var line Text = inl
				this = append(this,line)
				inl, isp, er = reader.ReadLine()
			}
		}
		return nil, nil
	}
}

func (this Catalog) Write(tgt *os.File) (error) {
	var writer *bufio.Writer = bufio.NewWriter(tgt)

	for _, line := range this {

		writer.Write(line)
		writer.WriteByte('\n')
	}

	writer.Flush()

	return nil
}

func (this IndexTarget) CatalogEncode() (that Catalog) {
	var list []IndexCatalog = this.CatalogRead()
	if 0 != len(list) {
		/*
		 * Head
		 */
		for _, head := range catalog_head {
			that = append(that,head)
		}
		var px, py int = 30, 50
		var bhi int = 18
		{
			var str string = fmt.Sprintf("  <text class=\"title\" x=\"%d\" y=\"%d\">Contents</text>",px,py)

			that = append(that,[]byte(str))

			py += (bhi<<1)
		}
		/*
		 * Body
		 */
		var icat IndexCatalog

		for _, icat = range list {
			/*
			 * Prepend '#' for HREF (target top).
			 */
			var anchor string = icat.Anchor()

			var str string = fmt.Sprintf("  <a href=\"#%s\" target=\"_top\"><text class=\"link\" x=\"%d\" y=\"%d\">%s /%s</text></a>",anchor,px,py,icat.path,icat.id)

			that = append(that,[]byte(str))

			py += bhi
		}
		/*
		 * Tail
		 */
		for _, tail := range catalog_tail {
			that = append(that,tail)
		}
	}
	return that
}

func (this IndexTarget) TabulateEncode() (that Tabulation) {
	var list []IndexCatalog = this.CatalogRead()
	if 0 != len(list) {

		var icat IndexCatalog

		for _, icat = range list {

			var anchor string = icat.Anchor()

			var str string = fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",icat.target.path,anchor,icat.id,icat.icon,icat.path,icat.link,icat.name,icat.embed)

			that = append(that,[]byte(str))
		}
	}
	return that
}
/*
 * Read objective target index.  N.B. The objective target
 * index may differ from the original index source list by
 * manual intervention.
 */
func (this IndexTarget) CatalogRead() (list []IndexCatalog) {
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

					var field_id json.Reader = object.CondHeadField("id")

					if field_id.IsNotEmpty() && object.Contains(field_id) {
						var field_ic json.Reader = field_id.CondTailField("icon")
						if field_ic.IsNotEmpty() && object.Contains(field_ic) {
							var field_pa json.Reader = field_ic.CondTailField("path")
							if field_pa.IsNotEmpty() && object.Contains(field_pa) {
								var field_li json.Reader = field_pa.CondTailField("link")
								if field_li.IsNotEmpty() && object.Contains(field_li) {
									var field_na json.Reader = field_li.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								} else {
									var field_na json.Reader = field_pa.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								}
							} else {
								var field_li json.Reader = field_ic.CondTailField("link")
								if field_li.IsNotEmpty() && object.Contains(field_li) {
									var field_na json.Reader = field_li.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								} else {
									var field_na json.Reader = field_pa.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								}
							}
						} else {
							var field_pa json.Reader = field_id.CondTailField("path")
							if field_pa.IsNotEmpty() && object.Contains(field_pa) {
								var field_li json.Reader = field_pa.CondTailField("link")
								if field_li.IsNotEmpty() && object.Contains(field_li) {
									var field_na json.Reader = field_li.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								} else {
									var field_na json.Reader = field_pa.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								}
							} else {
								var field_li json.Reader = field_ic.CondTailField("link")
								if field_li.IsNotEmpty() && object.Contains(field_li) {
									var field_na json.Reader = field_li.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								} else {
									var field_na json.Reader = field_pa.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {
										var field_em json.Reader = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									} else {
										var field_em json.Reader = field_li.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {


											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())
											var value_em string = Trim(field_em.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,value_em}

											list = append(list,catalog)
										} else {
											var value_id string = Trim(field_id.HeadString().TailString().String())
											var value_ic string = Trim(field_ic.HeadString().TailString().String())
											var value_pa string = Trim(field_pa.HeadString().TailString().String())
											var value_li string = Trim(field_li.HeadString().TailString().String())
											var value_na string = Trim(field_na.HeadString().TailString().String())

											var catalog IndexCatalog = IndexCatalog{this,value_id,value_ic,value_pa,value_li,value_na,""}

											list = append(list,catalog)
										}
									}
								}
							}
						}//
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

func (this IndexTarget) CatalogTarget() IndexFile {
	if this.IsValid() {
		var path IndexFile = this.Target()
		var file IndexFile = this.yyyymmdd+".svg"

		return FileCat(path,file)
	} else {
		return ""
	}
}

func (this IndexTarget) CatalogWrite() {
	var path IndexFile = this.CatalogTarget()
	var file *os.File
	var er error

	file, er = os.Create(string(path))
	if nil == er {
		var w *bufio.Writer = bufio.NewWriter(file)

		for _, line := range this.CatalogEncode() {

			w.Write(line)
			w.WriteByte('\n')
		}

		w.Flush()
		file.Close()
	}
}

func (this IndexTarget) TabulateTarget() IndexFile {
	if this.IsValid() {
		var path IndexFile = this.Target()
		var file IndexFile = this.yyyymmdd+".txt"

		return FileCat(path,file)
	} else {
		return ""
	}
}

func (this IndexTarget) TabulateWrite() {
	var path IndexFile = this.TabulateTarget()
	var file *os.File
	var er error

	file, er = os.Create(string(path))
	if nil == er {
		var w *bufio.Writer = bufio.NewWriter(file)

		for _, line := range this.TabulateEncode() {

			w.Write(line)
			w.WriteByte('\n')
		}

		w.Flush()
		file.Close()
	}
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
/*
 * User prepends '#' for use in web media HREF [HTML, XLINK, SVG11].
 */
func (this IndexCatalog) Anchor() string {
	return string(this.target.name)+"/"+string(this.target.yyyymmdd)+"/"+string(this.id)
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
