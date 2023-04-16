/* 
duniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package qualitiesPrint

// Print qualitiesw

import (
	
	BA	"git.duniter.org/gerard94/wotwizardclient/basicPrint"
	G	"git.duniter.org/gerard94/util/graphQL"
	GS	"git.duniter.org/gerard94/wotwizardclient/gqlSender"
	J	"git.duniter.org/gerard94/util/json"
	M	"git.duniter.org/gerard94/util/misc"
	S	"git.duniter.org/gerard94/util/sort"
	SM	"git.duniter.org/gerard94/util/strMapping"
	W	"git.duniter.org/gerard94/wotwizardclient/web"
		"bufio"
		"fmt"
		"net/http"
		"html/template"
		"strconv"
		"strings"

)

const (
	
	distancesName = "20distances"
	qualitiesName  = "21qualities"
	centralitiesName = "22centralities"
	
	queryGroup = `
		query MemberMissingGroup ($group: [String!]!) {
			filterGroup (group: $group, status_list: [MEMBER, MISSING]) {
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
		query MemberGroup ($group: [String!]!) {
			filterGroup (group: $group, status_list: MEMBER) {
				selected {
					id {
						uid
					}
				}
			}
		}
		fragment dispUid on GroupString {
			uid
			index
		}
	`
	
	queryNow = `
		query MembersNow {
			now {
				number
				bct
			}
		}
	`
	
	queryDistances = `
		query Distances ($group: [String!]) {
			now {
				number
				bct
			}
			identities(group: $group, status_list: MEMBER) {
				uid
				distance {
					value {
						num
						denom
						ratio
					}
				}
			}
		}
	`
	
	queryQualities = `
		query Qualities ($group: [String!]) {
			now {
				number
				bct
			}
			identities(group: $group, status_list: MEMBER) {
				uid
				quality {
					num
					denom
					ratio
				}
			}
		}
	`
	
	queryCentralities = `
		query Centralities ($group: [String!]) {
			now {
				number
				bct
			}
			identities(group: $group, status_list: MEMBER) {
				uid
				centrality
			}
		}
	`
	
	html = `
		{{define "head"}}<title>{{.Title}}</title>{{end}}
		{{define "body"}}
			<p>
				<a href = "/">{{Map "index"}}</a>
			</p>
			<h1>{{.Title}}</h1>
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
					<input type="submit" value="{{.SubmitL}}">
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
			{{if .Qs}}
				<p>
					{{range .Qs}}
						{{.}}
						<br>
					{{end}}
				</p>
				<p>
					{{range .QsId}}
						{{.}}
						<br>
					{{end}}
				</p>
			{{end}}
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
		Number int
		Bct int64
	}
	
	Fraction struct {
		Num,
		Denom int
		Ratio float64
	}
	
	IdentityT struct {
		Uid string
		Distance struct {
			Value *Fraction
		}
		Quality *Fraction
		Centrality float64
	}
	
	IdentitiesT []IdentityT
	
	DataT struct {
		Now *NowT
		Identities IdentitiesT
	}
	
	Qualities struct {
		Data *DataT
	}

	propT struct {
		id string
		f *Fraction
		prop float64
	}
	
	propsT []*propT
	
	propsSort struct { 
		t propsT
	}

	// Ouputs
	
	StringList = []string
	
	QualT struct {
		Title,
		Now,
		GroupL,
		Group,
		NoGroupL,
		SubmitL string
		NoGroupC bool
		Errors,
		Qs,
		QsId StringList
	}

)

var (
	
	groupDoc = GS.ExtractDocument(queryGroup)
	nowDoc = GS.ExtractDocument(queryNow)
	distancesDoc = GS.ExtractDocument(queryDistances)
	qualitiesDoc = GS.ExtractDocument(queryQualities)
	centralitiesDoc = GS.ExtractDocument(queryCentralities)

)

func (s *propsSort) Swap (i, j int) {
	s.t[i], s.t[j] = s.t[j], s.t[i]
} //Swap

func (s *propsSort) Less (p1, p2 int) bool {
	return s.t[p1].prop > s.t[p2].prop || s.t[p1].prop == s.t[p2].prop && BA.CompP(s.t[p1].id, s.t[p2].id) == BA.Lt
} //Less

func count (dist bool, ids IdentitiesT) (props, propsId propsT) {
	n := len(ids)
	props = make(propsT, n)
	for i, id := range ids {
		p := new(propT)
		p.id = id.Uid
		if dist {
			p.f = id.Distance.Value
		} else {
			p.f = id.Quality
		}
		if p.f == nil {
			p.prop = id.Centrality
		} else {
			p.prop = p.f.Ratio
		}
		props[i] = p
	}
	propsId = make(propsT, n)
	copy(propsId, props)
	s := &propsSort{t: props}
	var ts = S.TS{Sorter: s}
	ts.QuickSort(0, n - 1)
	return
} //count

func printNow (now *NowT, lang *SM.Lang) string {
	return fmt.Sprint(lang.Map("#duniterClient:Block"), " ", now.Number, "\t", BA.Ts2s(now.Bct, lang))
} //printNow

func printErr (ls *Lists, lang *SM.Lang) (err StringList) {
	if ls == nil {
		return nil
	}
	err = make(StringList, len(ls.Others) + len(ls.Unknown) + len(ls.Duplicate))
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

func print (qual string, group string, noGroup bool, ls *Lists, qualities *Qualities, lang *SM.Lang) *QualT {

/*
const (
	
	scatColor = Ports.red
	scatShape = HermesScat.point

)

var (
	
	i, m: int
	s: Views.Title
	axe: ARRAY 2 { rune
	t: TextModels.Model
	f: TextMappers.Formatter
	p: HermesViews.View
	ts: []HermesUtil.Dot
	tr: HermesViews.Trace
	d: Data
	props, propsId: Props
*/
	
	err := printErr(ls, lang)
	
	var t string
	switch qual {
	case distancesName:
		t = lang.Map("#duniterClient:distances")
		//axe := "d"
	case qualitiesName:
		t = lang.Map("#duniterClient:qualities")
		//axe := "q"
	case centralitiesName:
		t = lang.Map("#duniterClient:centralities")
		//axe := "c"
	}
	
	gr := lang.Map("#duniterClient:Group")
	noGrL := lang.Map("#duniterClient:noGroupL")
	s := lang.Map("#duniterClient:OK")
	
	var (qs, qsId StringList = nil, nil; now string)
	if qualities == nil {
		j := GS.Send(nil, nowDoc, "")
		qualities := new(Qualities)
		J.ApplyTo(j, qualities)
		now = printNow(qualities.Data.Now, lang)
	} else {
		d := qualities.Data
		now = printNow(d.Now, lang)
		props, propsId := count(qual == distancesName, d.Identities)
		m := len(props)
		qs = make(StringList, m)
		pp := -1.0
		ii := 0
		for i, p := range props {
			if p.prop != pp {
				pp = p.prop
				ii = i
			}
			if p.f == nil {
				qs[i] = fmt.Sprintf("%v%v%05.2f%v%v%v", ii + 1, BA.SpL, pp, BA.SpL, p.id, BA.SpS)
			} else {
				qs[i] = fmt.Sprintf("%v%v%05.2f (%v / %v)%v%v%v", ii + 1, BA.SpL, pp, p.f.Num, p.f.Denom, BA.SpL, p.id, BA.SpS)
			}
		}
		qsId = make(StringList, m)
		for i, p := range propsId {
			if p.f == nil {
				qsId[i] = fmt.Sprintf("%v%v%05.2f", p.id, BA.SpL, p.prop)
			} else {
				qsId[i] = fmt.Sprintf("%v%v%05.2f (%v / %v)", p.id, BA.SpL, p.f.Ratio, p.f.Num, p.f.Denom)
			}
		}
	}
	return &QualT{Title: t, Now: now, GroupL: gr, NoGroupL: noGrL, NoGroupC: noGroup, SubmitL: s, Group: group, Errors: err, Qs: qs, QsId: qsId}
	/*
	if props != nil {
		p := HermesViews.New()
		NEW(ts, m)
		for i := 0 TO m - 1 {
			ts[i].x := i
			ts[i].y := props[i].prop
		}
		tr := HermesScat.Insert(p, ts, scatShape, false, false, false)
		tr.ChangeColor(scatColor)
		tr := HermesAxes.Insert(p, false, true, false, false, true, true, false, false, 'n', '', 0, 3)
		tr := HermesAxes.Insert(p, true, true, false, false, true, true, false, false, axe, '%', 0, 3)
		p.FixY(0., 100.)
	}
	*/
} //print

func cleanText (text string) string {
	
	IsNotUidChar := func (r rune) bool {
		return !('a' <= r  && r <= 'z' || 'A' <= r && r <= 'Z' || '0' <= r && r <= '9' || r == '-' || r == '_' || r == 13 || r == 10)
	} //IsNotUidChar
	
	//cleanText
	return strings.TrimFunc(text, IsNotUidChar)
} //cleanText

func makeListFromText (text string) J.Json {
	mk := J.NewMaker()
	mk.StartObject()
	mk.StartArray()
	if text != "" {
		sc := bufio.NewScanner(strings.NewReader(text))
		for sc.Scan() {
			mk.PushString(cleanText(sc.Text()))
		}
	}
	mk.BuildArray()
	mk.BuildField("group")
	mk.BuildObject()
	return mk.GetJson()
} //makeListFromText

func stringList2Group (g IdListT) J.Json {
	mk := J.NewMaker()
	mk.StartObject()
	mk.StartArray()
	for _, grId := range g {
		mk.PushString(grId.Id.Uid)
	}
	mk.BuildArray()
	mk.BuildField("group")
	mk.BuildObject()
	return mk.GetJson()
}

func null2Group () J.Json {
	mk := J.NewMaker()
	mk.StartObject()
	mk.PushNull()
	mk.BuildField("group")
	mk.BuildObject()
	return mk.GetJson()
}

func end (name string, temp *template.Template, r *http.Request, w http.ResponseWriter, data W.SharedData) {
	lang := data.Lang()
	if  r.Method == "GET" {
		group := data.Group()
		temp.ExecuteTemplate(w, name, print(name, group, false, nil, nil, lang))
	} else {
		r.ParseForm()
		oldgroup := r.PostFormValue("oldgroup")
		group := cleanText(r.PostFormValue("group"))
		if  group != oldgroup {
			data.SetGroup(group)
		} else {
			group = data.Group()
		}
		noGroup := r.PostFormValue("noGroup") == "on"
		var (j J.Json; ls = new(Lists))
		if group != "" && !noGroup {
			jG := makeListFromText(group)
			j = GS.Send(jG, groupDoc, "MemberMissingGroup")
			gRes := new(GroupT)
			J.ApplyTo(j, gRes)
			ls = gRes.Data.FilterGroup
			j = GS.Send(jG, groupDoc, "MemberGroup")
			gRes = new(GroupT)
			J.ApplyTo(j, gRes)
			if len(gRes.Data.FilterGroup.Selected) == 0 {
				temp.ExecuteTemplate(w, name, print(name, group, noGroup, ls, nil, lang))
				return
			}
			j = stringList2Group(gRes.Data.FilterGroup.Selected)
		} else {
			j = null2Group()
		}
		var doc *G.Document
		switch name {
		case distancesName:
			doc = distancesDoc
		case qualitiesName:
			doc = qualitiesDoc
		case centralitiesName:
			doc = centralitiesDoc
		default:
			M.Halt(name, 100)
		}
		j = GS.Send(j, doc, "")
		quals := new(Qualities)
		J.ApplyTo(j, quals)
		temp.ExecuteTemplate(w, name, print(name, group, noGroup, ls, quals, lang))
	}
} //end

func init() {
	W.RegisterPackage(distancesName, html, end, true)
	W.RegisterPackage(qualitiesName, html, end, true)
	W.RegisterPackage(centralitiesName, html, end, true)
} //init
