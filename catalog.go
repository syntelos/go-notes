/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	"fmt"
	"strings"
)

type Index []Catalog

type Catalog struct {
	id, icon, path, link, name, embed string
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
	that.path = location
	that.link = string(table.Path())
	that.name = this.BaseName()
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
    }`,this.id,this.icon,this.path,this.link,this.name,this.embed)
}

func (this Catalog) LineString() string {
	return fmt.Sprintf(`{ "id": "%s", "icon": "%s", "path": "%s", "link": "%s", "name": "%s", "embed": "%s" }`,this.id,this.icon,this.path,this.link,this.name,this.embed)
}

func (this Index) String() string {
	var str strings.Builder

	str.WriteString(`[
`)
	for _, catalog := range this {

		str.WriteString(catalog.String())
	}
	str.WriteString(`
]
`)
	return str.String()
}
