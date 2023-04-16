/* 
duniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package distCalcPrint

import (
	
	BA	"git.duniter.org/gerard94/wotwizard/duniterClient/basicPrint"
	G	"git.duniter.org/gerard94/util/graphQL"
	GS	"git.duniter.org/gerard94/wotwizard/duniterClient/gqlSender"
	J	"git.duniter.org/gerard94/util/json"
	M	"git.duniter.org/gerard94/util/misc"
	SM	"git.duniter.org/gerard94/util/strMapping"
	W	"git.duniter.org/gerard94/wotwizard/duniterClient/web"
		"bufio"
		"fmt"
		"net/http"
		"strconv"
		"strings"
		"html/template"

)

const (
	
	calculatorName = "23calculator"
	
	defaultAnswerNb = 10
	
	queryGroup = `
		query MemberMissingGroupComplete ($group: [String!]!) {
			filterGroup (group: $group, status_list: [MEMBER, MISSING]) {
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
		query MemberMissingGroupErr ($group: [String!]!) {
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
		query MemberGroupComplete ($group: [String!]!) {
			filterGroup (group: $group, status_list: MEMBER) {
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
		query DistCalcNow {
			now {
				number
				bct
			}
		}
	`
	
	queryCalcDist = `
		query CalculateDist ($uid: String, $certs: [String!]) {
			now {
				number
				bct
			}
			valD: calcDist(uid: $uid, certifiers: $certs) {
				value {
					ratio
				}
				ok: dist_ok
			}
		}
	`
	
	queryCalcQual = `
		query CalculateQual ($uid: String, $certs: [String!]) {
			now {
				number
				bct
			}
			val: calcQual(uid: $uid, certifiers: $certs) {
				ratio
			}
		}
	`
	
	queryCalcBestDist = `
		query CalculateDist ($group: [String!], $uid: String, $certs: [String!], $ansNb: Int) {
			now {
				number
				bct
			}
			valD: calcDist(uid: $uid, certifiers: $certs) {
				value {
					ratio
				}
				ok: dist_ok
			}
			list: bestDist(group: $group, uid:$uid, certifiers: $certs, answerNb: $ansNb) {
				id {
					uid
				}
				valD: dist {
					value {
						ratio
					}
				}
			}
		}
	`
	
	queryCalcBestQual = `
		query CalculateQual ($group: [String!], $uid: String, $certs: [String!], $ansNb: Int) {
			now {
				number
				bct
			}
			val: calcQual(uid: $uid, certifiers: $certs) {
				ratio
			}
			list: bestQual(group: $group, uid:$uid, certifiers: $certs, answerNb: $ansNb) {
				id {
					uid
				}
				val: qual {
					ratio
				}
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
					<label for="uid">{{.UidL}} </label>
					<br>
					<input type="text" name="uid" id="uid" value="{{.Uid}}"/>
				</p>
				<p>
					<label for="certs">{{.CertsL}} </label>
					<br>
					<textarea id="certs" name="certs" rows=5 cols=33>{{.Certs}}</textarea>
				</p>
				<p>
					<label for="ansNb">{{.AnsNbL}} </label>
					<br>
					<input type="text" name="ansNb" id="ansNb" value="{{.AnsNb}}"/>
				</p>
				<p>
					<input type="radio" id="dist" name="distOrQual" value="0"{{if eq .DistOrQual "0"}} checked{{end}}>
					<label for="dist">{{.DistL}}</label>
					<input type="radio" id="qual" name="distOrQual" value="1"{{if eq .DistOrQual "1"}} checked{{end}}>
					<label for="qual">{{.QualL}}</label>
				</p>
				<p>
					<input type="submit" value="{{.SubmitL}}">
				</p>
			</form>
			{{if .Errors}}
				<p>
					{{range .Errors}}
						{{if .}}
							{{.}}
							<br>
						{{end}}
					{{end}}
				</p>
			{{end}}
			{{if .Value}}
				<p>
					{{.Value}}
				</p>
			{{end}}
			{{ if .List}}
				<p>
					{{.ListHeads}}
				</p>
				<p>
					{{range .List}}
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
		Number int32
		Bct int64
	}
	
	IdT struct {
		Uid string
	}
	
	ValT struct {
		Ratio float64
	}
	
	ValDT struct {
		Value ValT
		Ok bool
	}
	
	IdVal struct {
		Id IdT
		Val *ValT
		ValD *ValDT
	}
	
	ListT []*IdVal 
	
	DataT struct {
		Now *NowT
		Val *ValT
		ValD *ValDT
		List ListT
	}
	
	PathT []interface{}
	
	ErrorT struct {
		Message string
		Location struct {
			Position,
			Line,
			Column int
		}
		Path PathT
	}
	
	ErrorsT []*ErrorT
	
	ResultT struct {
		Errors ErrorsT
		Data *DataT
	}
	
	// Outputs
	
	StringList []string
	
	Disp struct {
		Title,
		Now,
		GroupL,
		Group,
		NoGroupL,
		UidL,
		Uid,
		CertsL,
		Certs,
		AnsNbL,
		AnsNb,
		DistL,
		QualL,
		DistOrQual,
		SubmitL,
		Value,
		ListHeads string
		NoGroupC bool
		Errors,
		List StringList
	}

)

var (
	
	groupDoc = GS.ExtractDocument(queryGroup)
	nowDoc = GS.ExtractDocument(queryNow)
	distCDoc = GS.ExtractDocument(queryCalcDist)
	qualCDoc = GS.ExtractDocument(queryCalcQual)
	distCBDoc = GS.ExtractDocument(queryCalcBestDist)
	qualCBDoc = GS.ExtractDocument(queryCalcBestQual)

)

func printNow (now *NowT, lang *SM.Lang) string {
	return fmt.Sprint(lang.Map("#duniterClient:Block"), " ", now.Number, "\t", BA.Ts2s(now.Bct, lang))
} //printNow

func printErr (lsG, lsU, lsC *Lists, lang *SM.Lang) (err StringList) {
	
	prt := func (ls *Lists, mes, loc string, i *int, err StringList) {
		if ls == nil {
			return
		}
		loc = lang.Map("#duniterClient:" + loc)
		for _, grId := range ls.Others {
			err[*i] = lang.Map("#duniterClient:" + mes, grId.Id.Uid, loc, strconv.Itoa(grId.Index + 1))
			*i++
		}
		for _, grS := range ls.Unknown {
			err[*i] = lang.Map("#duniterClient:IsUnknownLoc", grS.Uid, loc, strconv.Itoa(grS.Index + 1))
			*i++
		}
		for _, grS := range ls.Duplicate {
			err[*i] = lang.Map("#duniterClient:DuplicateLoc", grS.Uid, loc, strconv.Itoa(grS.Index + 1))
			*i++
		}
	} //prt
	
	//printErr
	n := 0
	if lsG != nil {
		n += len(lsG.Others) + len(lsG.Unknown) + len(lsG.Duplicate)
	}
	if lsU != nil {
		n += len(lsU.Others) + len(lsU.Unknown) + len(lsU.Duplicate)
	}
	if lsC != nil {
		n += len(lsC.Others) + len(lsC.Unknown) + len(lsC.Duplicate)
	}
	
	err = make(StringList, n)
	i := 0
	prt(lsG, "IsntMemberNorMissingLoc", "group", &i, err)
	prt(lsU, "IsntMemberNorMissingLoc", "uid", &i, err)
	prt(lsC, "IsntMemberLoc", "certifiers", &i, err)
	return
} //printErr

func printVal (d *DataT, uid string, lang *SM.Lang) string {
	var (val, ok string; dq float64)
	if d.Val != nil {
		if uid == "" {
			val = "QualOfUnknownIs"
		} else {
			val = "QualOfIs"
		}
		dq = d.Val.Ratio
		ok = ""
	} else {
		M.Assert(d.ValD != nil, 100)
		if uid == "" {
			val = "DistOfUnknownIs"
		} else {
			val = "DistOfIs"
		}
		dq = d.ValD.Value.Ratio
		if d.ValD.Ok {
			ok = "OK"
		} else {
			ok = "KO"
		}
	}
	ss := fmt.Sprintf("%6.2f", dq)
	return lang.Map("#duniterClient:" + val, uid, ss, ok)
}

func print (group string, noGroup bool, lsG, lsU, lsC *Lists, uid, certs, ansNb, distOrQual string, res *ResultT, lang *SM.Lang) *Disp {
	t := lang.Map("#duniterClient:Calculator")
	gr := lang.Map("#duniterClient:Group")
	noGrL := lang.Map("#duniterClient:noGroupL")
	id := lang.Map("#duniterClient:Uid")
	c := lang.Map("#duniterClient:Certifiers")
	an := lang.Map("#duniterClient:AnswerNb")
	dl := lang.Map("#duniterClient:Distance")
	ql := lang.Map("#duniterClient:Quality")
	s := lang.Map("#duniterClient:OK")
	lh := lang.Map("#duniterClient:ValIdHeads")
	
	err := printErr(lsG, lsU, lsC, lang)
	
	var (now string; val = ""; sl StringList = nil)
	if res == nil {
		j := GS.Send(nil, nowDoc, "")
		res = new(ResultT)
		J.ApplyTo(j, res)
		now = printNow(res.Data.Now, lang)
	} else {
		d := res.Data
		now = printNow(d.Now, lang)
		val = printVal(d, uid, lang)
		
		l := d.List
		if l != nil {
			sl = make(StringList, len(l))
			for i, iv := range l {
				w := new(strings.Builder)
				if iv.Val != nil {
					fmt.Fprintf(w, "%6.2f%%", iv.Val.Ratio)
				} else {
					M.Assert(iv.ValD != nil, 101)
					fmt.Fprintf(w, "%6.2f%%", iv.ValD.Value.Ratio)
				}
				fmt.Fprint(w, BA.SpL, iv.Id.Uid, BA.SpS)
				sl[i] = w.String()
			}
		}
	}
	
	return &Disp{Title: t, Now: now, GroupL: gr, NoGroupL: noGrL, UidL: id, CertsL: c, AnsNbL: an, DistL: dl, QualL: ql, SubmitL: s, Group: group, NoGroupC: noGroup, Uid: uid, Certs: certs, AnsNb: ansNb, DistOrQual: distOrQual, Value: val, ListHeads: lh, Errors: err, List: sl}
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

func makeListFromText (text string) J.Json {
	mk := J.NewMaker()
	mk.StartObject()
	pushListFromText(mk, text)
	mk.BuildField("group")
	mk.BuildObject()
	return mk.GetJson()
} //makeListFromText

func stringList2Group (mk *J.Maker, g IdListT) {
	mk.StartArray()
	for _, grId := range g {
		mk.PushString(grId.Id.Uid)
	}
	mk.BuildArray()
}

func null2Group (mk *J.Maker) {
	mk.PushNull()
}

func end (name string, temp *template.Template, r *http.Request, w http.ResponseWriter, data W.SharedData) {
	M.Assert(name == calculatorName, 20)
	lang := data.Lang()
	
	if  r.Method == "GET" {
		group := data.Group()
		temp.ExecuteTemplate(w, name, print(group, false, nil, nil, nil, "", "", strconv.Itoa(defaultAnswerNb), "0", nil, lang))
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
		uid := r.PostFormValue("uid")
		certs := r.PostFormValue("certs")
		distOrQual := r.PostFormValue("distOrQual")
		
		mk := J.NewMaker()
		mk.StartObject()
		
		var (j J.Json; lsG = new(Lists); groupEmpty = false)
		if group != "" && !noGroup {
			jG := makeListFromText(group)
			j = GS.Send(jG, groupDoc, "MemberMissingGroupErr")
			gRes := new(GroupT)
			J.ApplyTo(j, gRes)
			lsG = gRes.Data.FilterGroup
			j = GS.Send(jG, groupDoc, "MemberGroup")
			gRes = new(GroupT)
			J.ApplyTo(j, gRes)
			groupEmpty = len(gRes.Data.FilterGroup.Selected) == 0
			stringList2Group(mk, gRes.Data.FilterGroup.Selected)
		} else {
			null2Group(mk)
		}
		mk.BuildField("group")
		
		jU := makeListFromText(uid)
		j = GS.Send(jU, groupDoc, "MemberMissingGroupComplete")
		gRes := new(GroupT)
		J.ApplyTo(j, gRes)
		lsU := gRes.Data.FilterGroup
		if len(lsU.Selected) > 0 {
			M.Assert(len(lsU.Selected) == 1, 100)
			mk.PushString(lsU.Selected[0].Id.Uid)
		} else {
			mk.PushNull()
		}
		mk.BuildField("uid")
		
		var lsC = new(Lists)
		if certs != "" {
			jC := makeListFromText(certs)
			j = GS.Send(jC, groupDoc, "MemberGroupComplete")
			gRes := new(GroupT)
			J.ApplyTo(j, gRes)
			lsC = gRes.Data.FilterGroup
			stringList2Group(mk, lsC.Selected)
		} else {
			null2Group(mk)
		}
		mk.BuildField("certs")
		
		ansNb := r.PostFormValue("ansNb")
		an, err := strconv.Atoi(ansNb)
		if err != nil {
			an = 0
		}
		mk.PushInteger(int64(an))
		mk.BuildField("ansNb")
		
		mk.BuildObject()
		j = mk.GetJson()
		
		var doc *G.Document
		if groupEmpty {
			if distOrQual == "0" {
				doc = distCDoc
			} else {
				M.Assert(distOrQual == "1", 101)
				doc = qualCDoc
			}
		} else {
			if distOrQual == "0" {
				doc = distCBDoc
			} else {
				M.Assert(distOrQual == "1", 101)
				doc = qualCBDoc
			}
		}
		j = GS.Send(j, doc, "")
		res := new(ResultT)
		J.ApplyTo(j, res)
		temp.ExecuteTemplate(w, name, print(group, noGroup, lsG, lsU, lsC, uid, certs, ansNb, distOrQual, res, lang))
	}
} //end

func init() {
	W.RegisterPackage(calculatorName, html, end, true)
} //init
