/* 
duniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package web

import (
	
	A	"git.duniter.org/gerard94/util/avl"
	BA	"git.duniter.org/gerard94/wotwizardclient/basicPrint"
	F	"path/filepath"
	GS	"git.duniter.org/gerard94/wotwizardclient/gqlSender"
	J	"git.duniter.org/gerard94/util/json"
	M	"git.duniter.org/gerard94/util/misc"
	R	"git.duniter.org/gerard94/util/resources"
	SM	"git.duniter.org/gerard94/util/strMapping"
		"errors"
		"fmt"
		"net/http"
		"os"
		"strings"
		"text/scanner"
		"html/template"

)

const (
	
	queryVersion = `
		query Version {
			version
		}
	`
	
	base = `
		<html>
			<head>{{template "head" .}}</head>
			<body>{{template "body" .}}</body>
		</html>
	`
	
	htmlIndex = `
		{{define "head"}}<title>{{Map "Index"}}</title>{{end}}
		{{define "body"}}
			<h1>{{Map "Index"}}</h1>
			<p>
				{{Map "Server"}}{{.VersionS}}, {{Map "Client"}}{{.VersionC}}, {{Map "Tools"}}{{.VersionU}}
			</p>
			{{range $name, $temp := .P}}
				<p>
					<a href="/{{$name}}">{{Map $name}}</a> 
				</p>
			{{end}}
		{{end}}
	`
	
	authorizationsName = "Authorizations.txt"
	
	language_cookie_name = "v59ipaE3MoQDDtpt9edL"
	group_cookie_name = "9bhJrcdmRB9iGs0JLl82p9"
	
	cookie_original_chars = "\r\n"
	cookie_replace_chars = "~"

)

type (
	
	SharedData interface {
		Group () string
		SetGroup (string)
		Lang () *SM.Lang
		SetLang (string)
	}
	
	sharedData struct {
		r *http.Request
		w http.ResponseWriter
	}
	
	executeFunc func (name string, temp *template.Template, r *http.Request, w http.ResponseWriter, data SharedData)
	
	pack struct {
		tmp string
		call executeFunc
	}
	
	stringComp string

)

var (
	
	packages = make(map[string] *pack)
	packagesD = make(map[string] *pack)
	
	Version string
	
	versionDoc = GS.ExtractDocument(queryVersion)

)

func (sd *sharedData) Lang () *SM.Lang {
	if c, err := sd.r.Cookie(language_cookie_name); err == nil {
		return SM.NewLanguage(c.Value)
	}
	return SM.NewStdLanguage()
}

func (sd *sharedData) SetLang (l string) {
	http.SetCookie(sd.w, &http.Cookie{Name: language_cookie_name, Value: l})
}

func (sd *sharedData) Group () string {
	if c, err := sd.r.Cookie(group_cookie_name); err == nil {
		return strings.ReplaceAll(c.Value, cookie_replace_chars, cookie_original_chars)
	}
	return ""
}

func (sd *sharedData) SetGroup (group string) {
	http.SetCookie(sd.w, &http.Cookie{Name: group_cookie_name, Value: strings.ReplaceAll(group, cookie_original_chars, cookie_replace_chars)})
}

func RegisterPackage (name, temp string, call executeFunc, displayed bool) {
	p := new(pack)
	p.tmp = temp
	p.call = call
	packages[name] = p
	if displayed {
		packagesD[name] = p
	}
}

func (s1 stringComp) Compare (s2 A.Comparer) A.Comp {
	ss2 := s2.(stringComp)
	switch {
	case s1 < ss2:
		return A.Lt
	case s1 > ss2:
		return A.Gt
	default:
		return A.Eq
	}
}

func initAuthorizations () {
	name := F.Join(R.FindDir("duniterClient"), authorizationsName)
	f, err := os.Open(name)
	if err == nil {
		defer f.Close()
		s := new(scanner.Scanner)
		s.Init(f)
		s.Error = func(s *scanner.Scanner, msg string) {M.Halt(errors.New("File " + name + " incorrect"), 100)}
		s.Mode = scanner.ScanStrings
		auth := make(map[string] *int)
		for s.Scan() != scanner.EOF {
			ss := s.TokenText()
			M.Assert(ss[0] == '"' && ss[len(ss) - 1] == '"', ss, 101)
			auth[ss[1:len(ss) - 1]] = nil
		}
		for view, _ := range packagesD {
			if _, ok := auth[view]; !ok {
				delete(packages, view)
				delete(packagesD, view)
			}
		}
	} else {
		f, err := os.Create(name)
		M.Assert(err == nil, err, 102)
		defer f.Close()
		t := A.New()
		for view, _ := range packagesD {
			t.SearchIns(stringComp(view))
		}
		e := t.Next(nil)
		for e != nil {
			fmt.Fprintln(f, "\"" + e.Val().(stringComp) + "\"")
			e = t.Next(e)
		}
	}
}

func getHandler (name string, p *pack) http.HandlerFunc {
	
	return func (w http.ResponseWriter, r *http.Request) {
		data := &sharedData{r: r, w: w}
		lang := data.Lang()
		funcMap := make(template.FuncMap)
		funcMap["Map"] = func (name string) string {return lang.Map("#duniterClient:" + name)}
		temp := template.New(name)
		temp = temp.Funcs(funcMap)
		temp = template.Must(temp.Parse(p.tmp))
		temp = template.Must(temp.Parse(base))
		p.call(name, temp, r, w, data)
	}
	
}

func Start () {
	initAuthorizations()
	r := http.NewServeMux()
	for name, p := range packages {
		if name == "index" {
			r.HandleFunc("/", getHandler(name, p))
		} else {
			r.HandleFunc("/" + name, getHandler(name, p))
		}
	}
	htmlAddress := BA.HtmlAddress()
	server := &http.Server{
		Addr: htmlAddress,
		Handler: r,
	}
	fmt.Println("Listening on", htmlAddress, "...")
	server.ListenAndServe()
}

func manageIndex (name string, temp *template.Template, _ *http.Request, w http.ResponseWriter, _ SharedData) {
	
	type
		output struct {
			VersionU,
			VersionS,
			VersionC string
			P map[string] *pack
		}
	
	j := GS.Send(nil, versionDoc, ""); M.Assert(j != nil, 100)
	err := temp.ExecuteTemplate(w, name, &output{VersionU: M.Version(),VersionS: j.(*J.Object).Fields[0].Value.(*J.JsonVal).Json.(*J.Object).Fields[0].Value.(*J.String).S, VersionC: Version, P: packagesD})
	M.Assert(err == nil, err, 101)
}

func init () {
	RegisterPackage("index", htmlIndex, manageIndex, false)
}
