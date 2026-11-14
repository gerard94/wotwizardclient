/* 
duniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package pseudosPrint

import (
	
	BA	"github.com/gerard94/wotwizardclient/basicPrint"
	G	"github.com/gerard94/util/graphQL"
	GS	"github.com/gerard94/wotwizardclient/gqlSender"
	J	"github.com/gerard94/util/json"
	M	"github.com/gerard94/util/misc"
	SM	"github.com/gerard94/util/strMapping"
	W	"github.com/gerard94/wotwizardclient/web"
		"fmt"
		"net/http"
		"html/template"

)

const (
	
	pseudosName = "51pseudos"
	
	queryPseudos = `
		query Pseudos {
			now {
				number
				bct
			}
			pseudos:identities(status_list:[REVOKED, MISSING, MEMBER]) {
				uid
			}
		}
	`
	
	htmlPseudos = `
		{{define "head"}}<title>{{.Title}}</title>{{end}}
		{{define "body"}}
			<p>
				<a href = "/">{{Map "index"}}</a>
			</p>
			<h1>{{.Title}}</h1>
			<h3>
				{{.Block}}
			</h3>
			<p>
				{{.Number}}
			</p>
			<p>
				{{range .Ids}}
					{{.}}
					<br>
				{{end}}
			</p>
			<p>
				<a href = "/">{{Map "index"}}</a>
			</p>
		{{end}}
	`

)

type (
	
	NowT struct {
		Number int
		Bct int64
	}
	
	Identity struct {
		Uid string
	}
	
	Identities []Identity
	
	DataT struct {
		Now *NowT
		Pseudos Identities
	}

	PseudosT struct {
		Data *DataT
	}
	
	//Outputs
	
	Disp0 []string
	
	Disp struct {
		Title,
		Block,
		Number string
		Ids Disp0
	}

)

var (
	
	pseudosDoc *G.Document

)

func printNow (now *NowT, lang *SM.Lang) string {
	return fmt.Sprint(lang.Map("#duniterClient:Block"), " ", now.Number, " ", BA.Ts2s(now.Bct, lang))
} //printNow

func print (pseudos *PseudosT, lang *SM.Lang) *Disp {
	d := pseudos.Data
	ids := d.Pseudos
	n := fmt.Sprint(lang.Map("#duniterClient:PseudosNb"), " = ", len(ids))
	dd := make(Disp0, len(ids))
	for i, id := range(ids) {
		dd[i] = id.Uid
	}
	return &Disp{Title: lang.Map("#duniterClient:Pseudos"), Block: printNow(d.Now, lang), Number: n, Ids: dd}
} //print

func end (name string, temp *template.Template, _ *http.Request, w http.ResponseWriter, data W.SharedData) {
	lang := data.Lang()
	M.Assert(name == pseudosName, name, 100)
	j := GS.Send(nil, pseudosDoc, "")
	pseudos := new(PseudosT)
	J.ApplyTo(j, pseudos)
	temp.ExecuteTemplate(w, name, print(pseudos, lang))
} //end

func init() {
	pseudosDoc = GS.ExtractDocument(queryPseudos)
	W.RegisterPackage(pseudosName, htmlPseudos, end, true)
} //init
