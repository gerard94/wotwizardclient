/* 
duniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package infosCertifsPrint

// Print informations de certification

import (
	
	BA	"git.duniter.org/gerard94/wotwizardclient/basicPrint"
	B	"git.duniter.org/gerard94/wotwizardclient/blockchainPrint"
	GS	"git.duniter.org/gerard94/wotwizardclient/gqlSender"
	J	"git.duniter.org/gerard94/util/json"
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
	
	infosCertifName = "24infosCertif"
	
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
	
	queryInfosCertif = `
		query InfosCertif ($group: [String!]) {
			now {
				number
				bct
			}
			identities(group: $group, status_list: [MISSING, MEMBER]) {
				uid
				pubkey
				status
				sentry
				quality {
					ratio
				}
				limitDate
				minDate
				hash
				sent_certifications_number
			}
			maxCerts: parameter(name: sigStock) {
				value
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
			{{if .Ics}}
				<p>
					{{range .Ics}}
						{{.}}
						<br>
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
		Ratio float64
	}
	
	IdentityT struct {
		Uid string
		Pubkey B.Pubkey
		Status string
		Sentry bool
		Quality *Fraction
		LimitDate,
		MinDate int64
		Hash B.Hash
		Sent_certifications_number int
	}
	
	IdentitiesT []IdentityT
	
	DataT struct {
		Now *NowT
		Identities IdentitiesT
		MaxCerts struct {
			Value int
		}
	}
	
	InfosCertif struct {
		Data *DataT
	}

	// Ouputs
	
	StringList = []string
	
	InfCert struct {
		Title,
		Now,
		GroupL,
		Group,
		NoGroupL,
		SubmitL string
		NoGroupC bool
		Errors,
		Ics StringList
	}

)

var (
	
	groupDoc = GS.ExtractDocument(queryGroup)
	nowDoc = GS.ExtractDocument(queryNow)
	infosCertifDoc = GS.ExtractDocument(queryInfosCertif)

)

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

func print (group string, noGroup bool, ls *Lists, infos *InfosCertif, lang *SM.Lang) *InfCert {
	
	err := printErr(ls, lang)

	t := lang.Map("#duniterClient:infosCertif")
	
	gr := lang.Map("#duniterClient:Group")
	noGrL := lang.Map("#duniterClient:noGroupL")
	s := lang.Map("#duniterClient:OK")
	
	nicknameL := lang.Map("#duniterClient:Nickname")
	pubkeyL := lang.Map("#duniterClient:Pubkey")
	statusL := lang.Map("#duniterClient:Status")
	sentryL := lang.Map("#duniterClient:Sentry")
	qualL := lang.Map("#duniterClient:Quality")
	limitL := lang.Map("#duniterClient:AppMLimitDate")
	availL := lang.Map("#duniterClient:Availability")
	hashL := lang.Map("#duniterClient:Hash")
	quotaL := lang.Map("#duniterClient:remainingCerts")
	yesL := lang.Map("#duniterClient:yes")
	noL := lang.Map("#duniterClient:no")
	
	var (qs StringList = nil; now string)
	if infos == nil {
		j := GS.Send(nil, nowDoc, "")
		infos := new(InfosCertif)
		J.ApplyTo(j, infos)
		now = printNow(infos.Data.Now, lang)
	} else {
		d := infos.Data
		now = printNow(d.Now, lang)
		ids := d.Identities
		maxC := d.MaxCerts.Value
		m := len(ids)
		qs = make(StringList, m)
		for i, id := range ids {
			var sentry string
			if id.Sentry {
				sentry = yesL
			} else {
				sentry = noL
			}
			qs[i] = fmt.Sprintf("%v: %v,  %v: %v,  %v: %v,  %v: %v,  %v: %05.2f%%,  %v: %v,  %v: %v,  %v: %v,  %v: %v", nicknameL, id.Uid, pubkeyL, id.Pubkey, statusL, lang.Map("#duniterClient:" + id.Status), sentryL, sentry, qualL, id.Quality.Ratio, limitL, BA.Ts2s(id.LimitDate, lang), availL, BA.Ts2s(id.MinDate, lang), hashL, id.Hash, quotaL, maxC - id.Sent_certifications_number)
		}
	}
	return &InfCert{Title: t, Now: now, GroupL: gr, NoGroupL: noGrL, NoGroupC: noGroup, SubmitL: s, Group: group, Errors: err, Ics: qs}
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
		temp.ExecuteTemplate(w, name, print(group, false, nil, nil, lang))
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
				temp.ExecuteTemplate(w, name, print(group, noGroup, ls, nil, lang))
				return
			}
			j = stringList2Group(gRes.Data.FilterGroup.Selected)
		} else {
			j = null2Group()
		}
		j = GS.Send(j, infosCertifDoc, "")
		infos := new(InfosCertif)
		J.ApplyTo(j, infos)
		temp.ExecuteTemplate(w, name, print(group, noGroup, ls, infos, lang))
	}
} //end

func init() {
	W.RegisterPackage(infosCertifName, html, end, true)
} //init
