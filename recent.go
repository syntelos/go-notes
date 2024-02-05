/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

import (
	goauth "github.com/syntelos/go-auth"
	json "github.com/syntelos/go-json"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
	"path"
)
/*
 * RFC 3986 HTTP/S String.
 */
type URL = string

const GdriveSourceTableName TableName = TableName("google_drive_file_list")
/*
 * drive:files#list
 */
const GdriveSource URL = "https://www.googleapis.com/drive/v3/files?corpora=allDrives&orderBy=recency&includeItemsFromAllDrives=true&supportsAllDrives=true&mimeType=application%2Fpdf"
/*
 * Write catalog TABLE|JSON targets from ABSTRACT|GDR fetch
 * source.
 */
func (this FileLocation) RecentEncode() { // [TODO]
}
/*
 * Write catalog INDEX|JSON target from catalog TABLE|JSON
 * sources.
 */
func (this FileLocation) RecentUpdate() {
	var tgt FileLocation = this
	/*
	 * Do not overwrite existing target
	 */
	if tgt.IsValid() && tgt.NotExists() {
		var membership FileIx = tgt.FileIndex()
		/*
		 * Include source list as ordered members of {FileIx} `dirname`
		 */
		var src []FileLocation
		{
			for _, rev := range SourceList(ConfigurationSource()) {
				var rel FileIx = rev.FileIndex()
				if rel == membership {

					src = append(src, rev)
				}
			}
			src = FileSort(src)
		}

		if 0 < len(src) {

			var cat Index = CatalogIndex(src)

			tgt.Write(cat.Encode())
		}
	}
}

func (this FileLocation) RecentContents() {
}

func (this FileLocation) RecentTabulate() {
}
/*
 * Write TABLE|JSON target from SRC (Google Drive PDF files
 * list).
 */
func (this FileLocation) RecentFetch() {
	if HaveOperand(1) {
		var src []byte
		var err error
		src, err = os.ReadFile(Operand(1))
		if nil == err {
			var doc json.Reader = json.NewReader(Operand(1), src)
			var obj json.Reader = doc.HeadObject()

			if obj.IsNotEmpty() && doc.Contains(obj) {

				this.Write(src)
			}
		}
	} else {
		var GdriveScopes []string = []string{"drive"}

		var token *oauth2.Token = goauth.Token(GdriveScopes)
		if nil != token {
			var b io.Reader
			var q *http.Request
			var p *http.Response
			var r error

			q, r = http.NewRequest("GET", GdriveSource, b)
			if nil == r {
				q.Close = true

				token.SetAuthHeader(q)

				p, r = http.DefaultClient.Do(q)
				if nil == r && 200 == p.StatusCode && 0 < p.ContentLength {

					var c []byte
					c, r = io.ReadAll(p.Body)

					if p.ContentLength == int64(len(c)) {
						var doc json.Reader = json.NewReader(GdriveSource, c)
						var obj json.Reader = doc.HeadObject()

						if obj.IsNotEmpty() && doc.Contains(obj) {

							this.Write(c)
						}
					}
				}
			}
		}
	}
}

func RecentFetchTarget() (invalid FileLocation) {
	var tgt string = Operand(0)
	if 0 != len(tgt) {
		var fix FileIndex = FileClassify(tgt)
		if fix.IsValid() {

			return fix.Condense()
		} else {
			return invalid
		}
	} else {
		var cla FileTypeClass = (FileClassTable|FileTypeJSN)

		var src string = GdriveSource
		var tab TableName = GdriveSourceTableName
		var dat DateTime = NewDateTime()
		var dir, bas, loc string
		{
			dir = path.Join(tgt,dat.YYYY(),dat.MM())
			bas = string(tab)+"-"+string(dat)
			loc = path.Join(dir,bas)+".json"
		}
		return FileLocation{cla,src,dir,bas,loc,tab,dat}
	}
}

func RecentFetchSource(tgt,src string) (invalid FileLocation) {
	if 0 != len(src) {
		var fix FileIndex = FileClassify(src)
		if fix.IsValid() {

			return fix.Condense()
		} else {
			return invalid
		}
	} else {
		var cla FileTypeClass = (FileClassAbstract|FileTypeJSN)

		var src string = GdriveSource
		var tab TableName = GdriveSourceTableName
		var dat DateTime = NewDateTime()
		var dir, bas, loc string
		{
			dir = path.Join(tgt,dat.YYYY(),dat.MM())
			bas = string(tab)+"-"+string(dat)
			loc = path.Join(dir,bas)+".json"
		}
		return FileLocation{cla,src,dir,bas,loc,tab,dat}
	}
}
