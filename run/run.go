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
	_	"git.duniter.org/gerard94/wotwizardclient/static"
	
	W	"git.duniter.org/gerard94/wotwizardclient/web"
	
	//_	"git.duniter.org/gerard94/wotwizardclient/certificationsPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/distCalcPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/eventsPrint"
	//_	"git.duniter.org/gerard94/wotwizardclient/identitiesPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/identitySearch"
	_	"git.duniter.org/gerard94/wotwizardclient/language"
	_	"git.duniter.org/gerard94/wotwizardclient/membersPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/parametersPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/qualitiesPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/infosCertifPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/pseudosPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/sentriesPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/wotWizardPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/wwViews"

	_	"git.duniter.org/gerard94/wotwizardclient/tellLimitsPrint"
	_	"git.duniter.org/gerard94/wotwizardclient/stats"
	
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
