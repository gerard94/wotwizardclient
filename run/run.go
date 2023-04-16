/* 
duniterClient: WotWizard

Copyright (C) 2017-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License  for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package run

import (
	
	_	"git.duniter.org/gerard94/util/babel/static"
	_	"git.duniter.org/gerard94/util/json/static"
	_	"git.duniter.org/gerard94/util/graphQL/static"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/static"
	
	W	"git.duniter.org/gerard94/wotwizard/duniterClient/web"
	
	//_	"git.duniter.org/gerard94/wotwizard/duniterClient/certificationsPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/distCalcPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/eventsPrint"
	//_	"git.duniter.org/gerard94/wotwizard/duniterClient/identitiesPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/identitySearch"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/language"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/membersPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/parametersPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/qualitiesPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/infosCertifPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/pseudosPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/sentriesPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/wotWizardPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/wwViews"

	_	"git.duniter.org/gerard94/wotwizard/duniterClient/tellLimitsPrint"
	_	"git.duniter.org/gerard94/wotwizard/duniterClient/stats"
	
		"fmt"

)

const (
	
	version = "5.11.0"

)

func Start () {
	fmt.Println("WotWizard version", version, "\n")
	W.Version = version
	W.Start()
}
