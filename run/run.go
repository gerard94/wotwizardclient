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
	
	M	"git.duniter.org/gerard94/util/misc"
	W	"git.duniter.org/gerard94/wotwizardclient/web"
	
	// 55certificationsFrom & 56certificationsTo
	//_	"git.duniter.org/gerard94/wotwizardclient/certificationsPrint"
	
	// 23calculator
	_	"git.duniter.org/gerard94/wotwizardclient/distCalcPrint"
	
	// 30limitsMember & 31limitsMissing & 32limitsCerts
	_	"git.duniter.org/gerard94/wotwizardclient/eventsPrint"
	
	// 51revokedIdentities & 52missingIdentities & 53memberIdentities & 54newcomerIdentities
	//_	"git.duniter.org/gerard94/wotwizardclient/identitiesPrint"
	
	// 1explorer
	_	"git.duniter.org/gerard94/wotwizardclient/identitySearch"
	
	// language
	_	"git.duniter.org/gerard94/wotwizardclient/language"
	
	// 400membersCountT & 401membersFirstEntryT & 402membersLossT & 403membersCountG &
	// 404membersCountFlux & 405membersCountFluxPM &
	// 406membersFirstEntryG & 407membersFEFlux & 408membersFEFluxPM &
	// 409membersLossG & 410membersLossFlux & 411membersLossFluxPM
	_	"git.duniter.org/gerard94/wotwizardclient/membersPrint"
	
	// 50allParameters
	_	"git.duniter.org/gerard94/wotwizardclient/parametersPrint"
	
	// 20distances & 21qualities & 22centralities
	_	"git.duniter.org/gerard94/wotwizardclient/qualitiesPrint"
	
	// 24infosCertif
	_	"git.duniter.org/gerard94/wotwizardclient/infosCertifPrint"
	
	// 51pseudos
	_	"git.duniter.org/gerard94/wotwizardclient/pseudosPrint"
	
	// 57sentries
	_	"git.duniter.org/gerard94/wotwizardclient/sentriesPrint"
	
	// 01wwFile & 02wwPerms
	_	"git.duniter.org/gerard94/wotwizardclient/wotWizardPrint"
	
	// 00wwView
	_	"git.duniter.org/gerard94/wotwizardclient/wwViews"
	
	
	// 33memLim & 34certLim
	_	"git.duniter.org/gerard94/wotwizardclient/tellLimitsPrint"
	
	// 60sentCertsHistory
	_	"git.duniter.org/gerard94/wotwizardclient/stats"
	
		"fmt"

)

const (
	
	version = "5.12.4"

)

func Start () {
	fmt.Println("WotWizard Client version", version, "Tools version", M.Version(), "\n")
	W.Version = version
	W.Start()
}
