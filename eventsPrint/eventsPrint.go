/* 
duniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package eventsPrint

import (
	
	BA	"github.com/gerard94/wotwizardclient/basicPrint"
	G	"github.com/gerard94/util/graphQL"
	GS	"github.com/gerard94/wotwizardclient/gqlSender"
	J	"github.com/gerard94/util/json"
	M	"github.com/gerard94/util/misc"
	SM	"github.com/gerard94/util/strMapping"
	W	"github.com/gerard94/wotwizardclient/web"
		"bufio"
		"fmt"
		"net/http"
		"strings"
		"strconv"
		"html/template"

)

const (
	
	limitsMemberName = "30limitsMember"
	limitsMissingName = "31limitsMissing"
	limitsCertsName = "32limitsCerts"
	
	queryGroup = `
		query Group ($group: [String!]!, $statuses: [Identity_Status!]!) {
			filterGroup (group: $group, status_list: $statuses) {
				selected {
					id {
						uid
					}
				}
				others {
					id {
						uid
					}
					index
				}
				unknown {
					... dispUid
				}
				duplicate {
					... dispUid
				}
			}
		}
		fragment dispUid on GroupString {
			uid
			index
		}
	`
	
	queryNow = `
		query EventsNow {
			now {
				number
				bct
			}
		}
	`
	
	queryLimitsMem = `
		query LimitsMem ($group: [String!], $includeDist: Boolean!) {
			now {
				number
				bct
			}
			identities: memEnds(group: $group) {
				uid
				limitDate
				distance @include(if: $includeDist) {
					value {
						ratio
					}
					dist_ok
				}
			}
		}
	`
	
	queryLimitsMissing = `
		query LimitsMissing ($group: [String!]) {
			now {
				number
				bct
			}
			identities: missEnds(group: $group) {
				uid
				limitDate
			}
		}
	`
	
	queryLimitsCerts = `
		query LimitsCerts ($group: [String!]) {
			now {
				number
				bct
			}
			identities: certEnds(group: $group) {
				uid
				status
				limitDate: certsLimit
			}
		}
	`
	
	htmlLimits = `
		{{define "head"}}<title>{{.Title}}</title>{{end}}
		{{define "body"}}
			<p>
				<a href = "/">{{Map "index"}}</a>
			</p>
			<h1>{{.Title}}</h1>
			<h2>{{.Subtitle}}</h2>
			<h3>
				{{.Now}}
			</h3>
			<form action="" method="post">
				<p>
					<label for="group">{{.GroupL}} </label>
					<br>
					<textarea id="group" name="group" rows=5 cols=33>{{.Group}}</textarea>
					<input type="checkbox" id="noGroup" name="noGroup"{{if .NoGroupC}} checked{{end}}>
						<label for="noGroup">{{.NoGroupL}}</label>
				</p>
				<textarea id="oldgroup" name="oldgroup" hidden>{{.Group}}</textarea>
				<p>
					<input type="submit" value="{{.Submit}}">
				</p>
			</form>
			{{if .Errors}}
				<p>
					{{range .Errors}}
						{{.}}
						<br>
					{{end}}
				</p>
			{{end}}
			<p>
				{{if .Ps}}
					{{range .Ps}}
						{{.}}
						<br>
					{{end}}
				{{end}}
			</p>
			<p>
				<a href = "/">{{Map "index"}}</a>
			</p>
		{{end}}
	`

)

type (
	
	Identity struct {
		Uid string
	}
	
	GroupId struct {
		Id *Identity
		Index int
	}
	
	GroupString struct {
		Uid string
		Index int
	}
	
	IdListT []*GroupId
	
	StringListT []*GroupString
	
	Lists struct {
		Selected,
		Others IdListT
		Unknown,
		Duplicate StringListT
	}
	
	GroupT struct {
		Data struct {
			FilterGroup *Lists
		}
	}
	
	NowT struct {
		Number int32
		Bct int64
	}
	
	IdentityT struct {
		Uid,
		Status string
		LimitDate int64
		Distance struct {
			Value struct {
				Ratio float64
			}
			Dist_ok bool
		}
	}
	
	IdentitiesT []*IdentityT
	
	DataT struct {
		Now *NowT
		Identities IdentitiesT
	}
	
	Limits struct {
		Data *DataT
	}

	propT struct {
		id string
		aux,
		aux2 bool
		prop int64
		prop2 float64
		prop3 bool
	}
	
	propsT []*propT
	
	propsSort struct {
		t propsT
	}
	
	// Outputs
	
	PropList []string
	
	Disp struct {
		Now,
		Title,
		Subtitle,
		GroupL,
		Group,
		NoGroupL,
		Submit string
		NoGroupC bool
		Errors,
		Ps PropList
	}

)

var (
	
	groupDoc = GS.ExtractDocument(queryGroup)
	nowDoc = GS.ExtractDocument(queryNow)
	limitsMemDoc = GS.ExtractDocument(queryLimitsMem)
	limitsMissDoc = GS.ExtractDocument(queryLimitsMissing)
	limitsCertsDoc = GS.ExtractDocument(queryLimitsCerts)

)

func count (d *DataT, what string) (props propsT) {
	props = make(propsT, len(d.Identities))
	for i, id := range d.Identities {
		p := new(propT)
		p.id = id.Uid
		p.aux = what != limitsCertsName || id.Status == "MEMBER"
		p.aux2 = what == limitsMemberName && id.Distance.Value.Ratio > 0
		p.prop = id.LimitDate
		p.prop2 = id.Distance.Value.Ratio
		p.prop3 = id.Distance.Dist_ok
		props[i] = p
	}
	return
} //count

func printNow (now *NowT, lang *SM.Lang) string {
	return fmt.Sprint(lang.Map("#duniterClient:Block"), " ", now.Number, "\t", BA.Ts2s(now.Bct, lang))
} //printNow

func printErr (ls *Lists, lang *SM.Lang) (err PropList) {
	if ls == nil {
		return nil
	}
	err = make(PropList, len(ls.Others) + len(ls.Unknown) + len(ls.Duplicate))
	i := 0
	for _, grId := range ls.Others {
		err[i] = lang.Map("#duniterClient:IsntMemberNorMissing", grId.Id.Uid, strconv.Itoa(grId.Index + 1))
		i++
	}
	for _, grS := range ls.Unknown {
		err[i] = lang.Map("#duniterClient:IsUnknown", grS.Uid, strconv.Itoa(grS.Index + 1))
		i++
	}
	for _, grS := range ls.Duplicate {
		err[i] = lang.Map("#duniterClient:Duplicate", grS.Uid, strconv.Itoa(grS.Index + 1))
		i++
	}
	return
} //printErr

func print (group string, noGroup bool, ls *Lists, limits *Limits, what, title, subtitle string, lang *SM.Lang) *Disp {
	
	err := printErr(ls, lang)
	
	var (ps PropList = nil; now string)
	t := lang.Map(title)
	st := lang.Map(subtitle)
	gr := lang.Map("#duniterClient:Group")
	noGrL := lang.Map("#duniterClient:noGroupL")
	s := lang.Map("#duniterClient:OK")
	if limits == nil {
		j := GS.Send(nil, nowDoc, "")
		limits := new(Limits)
		J.ApplyTo(j, limits)
		now = printNow(limits.Data.Now, lang)
		if what != "" {
			t += " (0)"
		}
	} else {
		d := limits.Data
		now = printNow(d.Now, lang)
		props := count(d, what)
		t = fmt.Sprint(t, " (", len(props), ")")
		ps = make(PropList, len(props))
		for i, p := range props {
			w := new(strings.Builder)
			fmt.Fprint(w, BA.Ts2s(p.prop, lang), BA.SpL)
			if p.aux2 {
				var ok string
				if p.prop3 {
					ok = "OK"
				} else {
					ok = "KO"
				}
				fmt.Fprintf(w, "(%v% 5.2f%%)%v", ok, p.prop2, BA.SpL)
			}
			if p.aux {
				fmt.Fprint(w, BA.SpS, BA.SpS)
			} else {
				fmt.Fprint(w, BA.OldIcon)
			}
			fmt.Fprint(w, p.id)
			fmt.Fprint(w, BA.SpS)
			ps[i] = w.String()
		}
	}
	return &Disp{Now: now, Title: t, Subtitle: st, GroupL: gr, NoGroupL: noGrL, NoGroupC: noGroup, Group: group, Submit: s, Errors: err, Ps: ps}
} //print

func cleanText (text string) string {
	
	IsNotUidChar := func (r rune) bool {
		return !('a' <= r  && r <= 'z' || 'A' <= r && r <= 'Z' || '0' <= r && r <= '9' || r == '-' || r == '_' || r == 13 || r == 10)
	} //IsNotUidChar
	
	//cleanText
	return strings.TrimFunc(text, IsNotUidChar)
} //cleanText

func pushListFromText (mk *J.Maker, text string) {
	mk.StartArray()
	if text != "" {
		sc := bufio.NewScanner(strings.NewReader(text))
		for sc.Scan() {
			mk.PushString(cleanText(sc.Text()))
		}
	}
	mk.BuildArray()
} //pushListFromText

func stringList2Group (g IdListT) J.Json {
	mk := J.NewMaker()
	mk.StartObject()
	mk.StartArray()
	for _, grId := range g {
		mk.PushString(grId.Id.Uid)
	}
	mk.BuildArray()
	mk.BuildField("group")
	mk.PushBoolean(true)
	mk.BuildField("includeDist")
	mk.BuildObject()
	return mk.GetJson()
}

func null2Group () J.Json {
	mk := J.NewMaker()
	mk.StartObject()
	mk.PushNull()
	mk.BuildField("group")
	mk.PushBoolean(false)
	mk.BuildField("includeDist")
	mk.BuildObject()
	return mk.GetJson()
}

func end (name string, temp *template.Template, r *http.Request, w http.ResponseWriter, data W.SharedData) {
	lang := data.Lang()
	var (doc *G.Document; title, subtitle string; j J.Json)
	switch name {
	case limitsMemberName:
		title = "#duniterClient:limitsMember"
		subtitle = "#duniterClient:limitsSubMember"
		doc = limitsMemDoc
	case limitsMissingName:
		title = "#duniterClient:limitsMissing"
		subtitle = "#duniterClient:limitsSubMissing"
		doc = limitsMissDoc
	case limitsCertsName:
		title = "#duniterClient:limitsCerts"
		subtitle = "#duniterClient:limitsSubCerts"
		doc = limitsCertsDoc
	default:
		M.Halt(name, 100)
	}
	if  r.Method == "GET" {
		group := data.Group()
		temp.ExecuteTemplate(w, name, print(group, false, nil, nil, "", title, subtitle, lang))
	} else {
		r.ParseForm()
		oldgroup := r.PostFormValue("oldgroup")
		group := cleanText(r.PostFormValue("group"))
		if group != oldgroup {
			data.SetGroup(group)
		} else {
			group = data.Group()
		}
		noGroup := r.PostFormValue("noGroup") == "on"
		var ls = new(Lists)
		if group != "" && !noGroup {
			mk := J.NewMaker()
			mk.StartObject()
			pushListFromText(mk, group)
			mk.BuildField("group")
			mk.StartArray()
			mk.PushString("MEMBER")
			mk.PushString("MISSING")
			mk.BuildArray()
			mk.BuildField("statuses")
			mk.BuildObject()
			j = mk.GetJson()
			j = GS.Send(j, groupDoc, "")
			gRes := new(GroupT)
			J.ApplyTo(j, gRes)
			ls = gRes.Data.FilterGroup
			mk = J.NewMaker()
			mk.StartObject()
			pushListFromText(mk, group)
			mk.BuildField("group")
			mk.StartArray()
			switch name {
			case limitsMemberName:
				mk.PushString("MEMBER")
			case limitsMissingName:
				mk.PushString("MISSING")
			case limitsCertsName:
				mk.PushString("MEMBER")
				mk.PushString("MISSING")
			default:
				M.Halt(name, 100)
			}
			mk.BuildArray()
			mk.BuildField("statuses")
			mk.BuildObject()
			j = mk.GetJson()
			j = GS.Send(j, groupDoc, "")
			gRes = new(GroupT)
			J.ApplyTo(j, gRes)
			if len(gRes.Data.FilterGroup.Selected) == 0 {
				temp.ExecuteTemplate(w, name, print(group, noGroup, ls, nil, name, title, subtitle, lang))
				return
			}
			j = stringList2Group(gRes.Data.FilterGroup.Selected)
		} else {
			j = null2Group()
		}
		j = GS.Send(j, doc, "")
		limits := new(Limits)
		J.ApplyTo(j, limits)
		temp.ExecuteTemplate(w, name, print(group, noGroup, ls, limits, name, title, subtitle, lang))
	}
} //end

func init() {
	W.RegisterPackage(limitsMemberName, htmlLimits, end, true)
	W.RegisterPackage(limitsMissingName, htmlLimits, end, true)
	W.RegisterPackage(limitsCertsName, htmlLimits, end, true)
} //init
