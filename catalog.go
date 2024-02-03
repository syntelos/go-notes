/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	json "github.com/syntelos/go-json"
	"strings"
)

type Index []Catalog

type Catalog struct {
	id, icon, path, link, name, embed string
}

func CatalogIndex(this []FileLocation) (that Index) {
	for _, file := range this {
		that = append(that, file.FileCatalog())
	}
	return that
}

/*
 * Structural analogue to File#TableAnchor: <ID> =
 * <YYYYMMDD_HHMMSS>.
 */
func (this FileLocation) FileCatalog() (that Catalog) {

	var location string = this.WebLocation()

	var table TableName = this.TableName()

	that.id = this.YYYYMMDD_HHMMSS()
	that.icon = "syntelos-catalog"
	that.path = string(table.Path())
	that.link = string(table.Link())
	that.name = this.basename
	that.embed = location

	return that
}

func (this Catalog) String() string {
	return fmt.Sprintf(`    {
        "id": "%s",
        "icon": "%s",
        "path": "%s",
        "link": "%s",
        "name": "%s",
        "embed": "%s"
    }`, this.id, this.icon, this.path, this.link, this.name, this.embed)
}

func (this Catalog) LineString() string {
	return fmt.Sprintf(`{ "id": "%s", "icon": "%s", "path": "%s", "link": "%s", "name": "%s", "embed": "%s" }`, this.id, this.icon, this.path, this.link, this.name, this.embed)
}

func (this Catalog) Encode() []byte {

	return []byte(this.String())
}

func (this Catalog) Decode(content []byte) {
	var rdr json.Reader = json.NewReader("", content)
	if rdr.IsNotEmpty() {

		var object json.Reader = rdr.HeadObject()
		for object.IsNotEmpty() {

			var field_id json.Reader = object.CondHeadField("id")
			if field_id.IsNotEmpty() && object.Contains(field_id) {

				this.id = field_id.HeadString().TailString().StringUnquote()

				var field_ic = field_id.CondTailField("icon")
				if field_ic.IsNotEmpty() && object.Contains(field_ic) {

					this.icon = field_id.HeadString().TailString().StringUnquote()

					var field_pa = field_ic.CondTailField("path")
					if field_pa.IsNotEmpty() && object.Contains(field_pa) {

						this.path = field_id.HeadString().TailString().StringUnquote()

						var field_li = field_pa.CondTailField("link")
						if field_li.IsNotEmpty() && object.Contains(field_li) {

							this.link = field_id.HeadString().TailString().StringUnquote()

							var field_na = field_li.CondTailField("name")
							if field_na.IsNotEmpty() && object.Contains(field_na) {

								this.name = field_id.HeadString().TailString().StringUnquote()

								var field_em = field_na.CondTailField("embed")
								if field_em.IsNotEmpty() && object.Contains(field_em) {

									this.embed = field_id.HeadString().TailString().StringUnquote()
								}
							}
						}
					}
				}
			}
		}
	}
}

func (this Index) String() string {
	var str strings.Builder

	str.WriteString(`[
`)
	for ix, catalog := range this {

		if 0 != ix {

			str.WriteString(",\n")
		}

		str.WriteString(catalog.String())
	}
	str.WriteString(`
]
`)
	return str.String()
}

func (this Index) Encode() []byte {

	return []byte(this.String())
}

func (this Index) Decode(content []byte) {
	var rdr json.Reader = json.NewReader("", content)
	if rdr.IsNotEmpty() {

		var array json.Reader = rdr.HeadArray()
		if array.IsNotEmpty() {

			var object json.Reader = rdr.HeadObject()
			for object.IsNotEmpty() {

				var catalog Catalog = Catalog{}

				catalog.Decode([]byte(object.String()))

				object = object.TailObject()
			}
		}
	}

}
