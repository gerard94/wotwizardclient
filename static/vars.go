/* 
DuniterClient: WotWizard.

Copyright (C) 2017 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package static

// Package strMapping implements a mapping between keys and longer strings.

var (
	
	strEn = `STRINGS

AllCertified	All Non-Revoked Certified Identities:
AllCertifiedIO	History of all sent certifications (start ↑ and end ↓):
AllCertifiers	All Known Non-Revoked Certifiers:
AllCertifiersIO	History of all received certifications (start ↑ and end ↓):
50allParameters	Informations: Blockchain Parameters
AnswerNb	Greatest number of suggested certifiers:
AppMLimitDate	Limit Date of Membership Renewal
AppNLimitDate	Limit Date of Membership Application
AppRLimitDate	Date of Revocation
Availability	Availability of next sent certification
Bct	Blockchain Median Time
Block	Block
Brackets	Brackets
Calculator	Distance and Quality Calculator
23calculator	Properties: Distance and Quality Calculator
22centralities	Properties: Centralities
centralities	Centralities
Centrality	Centrality
certification	certification:
Certifications	Certifications
certifications	certifications:
55certificationsFrom	Informations: Certifications From Members
certificationsFrom	Certifications From Members
56certificationsTo	Informations: Certifications To Members
certificationsTo	Certifications To Members
Certifiers	New certifiers:
certifiers	new certifiers
34certLim	Limits: Certifications Limits Dates by Dates
certsNb	^0 certifications
Client	Client Version: 
Computation_duration	Computation Duration
Day	Day
day	day(s)
Delay	Delay (days)
Distance	Distance Rule
DistanceE	Expected Distance Rule
distanceRule	Expected Distance Rule: ^0% (^1 / ^2)
distanceRule0	Expected Distance Rule: ^0%
20distances	Properties: Distances (rule)
distances	Distances (rule)
DistOfIs	The expected distance of ^0 is ^1%: ^2
DistOfUnknownIs	The expected distance is ^1%: ^2
Distribution	Distribution
dossiers	dossiers
Duplicate	^0 appears twice (line ^1)
DuplicateLoc	^0 appears twice (in ^1, line ^2)
en	English
Entry	Entry of 
es	Español
Exit	Exit of 
Explorer	Explorer
1explorer	Web of Trust Explorer
FirstEntries	First Entries
FirstEntriesFlux	Flux of First Entries
FirstEntriesFluxPM	Flux of First Entries per member
Forward	Number of days until forecast date
fr	Français
Group	Enter every nickname of your group, line by line, or let the list empty if you have no group:
group	group
Hash	Hash
history	History
hour	hour(s)
In	In
Index	Index
index	index
24infosCertif	Properties: Informations for Certification
infosCertif	Informations for Certification
intCertifs	internal certifications
IsntMemberLoc	^0 is not a member (in ^1, line ^2)
IsntMemberNorMissing	^0 is not a member nor excluded (line ^1)
IsntMemberNorMissingLoc	^0 is not a member nor excluded (in ^1, line ^2)
IsUnknown	^0 is unknown (line ^1)
IsUnknownLoc	^0 is unknown (in ^1, line ^2)
KO	KO
language	Choose Language
LastApplication	Last membership application
LCCertifiers	She was certified by 
LCComplement	This identity will lack certifications soon at the following date:
LimitsC	Certifications Limits
LimitDate	Limit Date
32limitsCerts	Limits: Limit Dates of Certifications
limitsCerts	Limit Dates of Certifications
LimitsM	Limits of Memberships
30limitsMember	Limits: Limit Dates of Memberships
limitsMember	Limit Dates of Memberships
31limitsMissing	Limits: Limit Dates of not-renewed memberships
limitsMissing	Limit Dates of not-renewed memberships
limitsSubCerts	The displayed identity will lack certifications at this date
limitsSubMember	The displayed member (with her Distance Rule) will loose its adhesion at this date
limitsSubMissing	The displayed excluded identity will be revoked once and for all at this date
LMCertifiers	certified by
LMComplement	To stay member, this identity must renew her membership:
Losses	Losses
LossesFlux	Flux of Losses
LossesFluxPM	Flux of Losses per Member
Mean	Mean
Median	Median
MEMBER	MEMBER
53memberIdentities	Informations: Members List
Members	Members
404membersCountFlux	Evolution: Flux of Members (graphics)
405membersCountFluxPM	Evolution: Flux of Members per Member (graphics)
403membersCountG	Evolution: Number of members (graphics)
400membersCountT	Evolution:  Number of members (list)
407membersFEFlux	Evolution: Flux of First Entries (graphics)
408membersFEFluxPM	Evolution: Flux of First Entries per Member (graphics)
406membersFirstEntryG	Evolution: First Entries (graphics)
401membersFirstEntryT	Evolution: First Entries (list)
MembersFlux	Flux of Members
MembersFluxPM	Flux of Members per Member
410membersLossFlux	Evolution: Flux of Losses (graphics)
411membersLossFluxPM	Evolution: Flux of Losses per Member (graphics)
409membersLossG	Evolution: Losses (graphics)
402membersLossT	Evolution: Losses (list)
MembersNb	Number of members
33memLim	Limits: Memberships Limits Date by Date
minApplicationDate	Wait at least two months after the last application (^0)
minute	minute(s)
MISSING	EXCLUDED
Missing	Excluded Identities - Needing a new application
52missingIdentities	Informations: List of Excluded Identities - Needing a new application
MissingNb	Number of Excluded Identities
Missings	Excluded
month	month(s)
neededCerts	needed certifications:
Never	Never
NEWCOMER	NEWCOMER
54newcomerIdentities	Informations: List of Newcomers
Newcomers	Newcomers
newcomers	Pending Dossiers
NewcomersNb	Number of Newcomers
Nickname	Nickname
no	no
noGroupL	Ignore group
NotMembers	Not Members
OK	OK
Out	Out
Para_avgGenTime	The average time for writing 1 block (wished time)
Para_c	The relative growth of the UD every [dtReeval] period
Para_dt	Time period between two UD
Para_dtDiffEval	The number of blocks required to evaluate again PoWMin value
Para_dtReeval	Time period between two re-evaluation of the UD
Para_idtyWindow	Maximum delay an identity can wait before being expired for non-writing
Para_medianTimeBlocks	Number of blocks used for calculating median time
Para_msPeriod	Minimum delay between 2 memberships of a same issuer
Para_msValidity	Maximum age of an active membership
Para_msWindow	Maximum delay a membership can wait before being expired for non-writing
Para_percentRot	The proportion of calculating members not excluded from the proof of work
Para_sigPeriod	Minimum delay between two certifications of a same issuer
Para_sigQty	Minimum quantity of signatures to be part of the WoT
Para_sigStock	Maximum quantity of active certifications made by member
Para_sigValidity	Maximum age of an active certification
Para_sigReplay	Minimum delay before replaying a certification
Para_sigWindow	Maximum delay a certification can wait before being expired for non-writing
Para_stepMax	Maximum distance between a WOT member and [xpercent] of sentries
Para_txWindow	Maximum delay a transaction can wait before being expired for non-writing
Para_ud0	UD(0), i.e. initial Universal Dividend
Para_udReevalTime0	Time of first reevaluation of the UD
Para_udTime0	Time of first UD
Para_xpercent	Minimum percent of sentries to reach to match the distance rule
Parameters	Parameters
pending	pending new application
PendingApplication	Pending membership application
Permutations	Permutations
permutations	permutations
PermutationsNb	Number of Permutations: 
PresentCertified	Currently or Soon (°) Certified Identities:
PresentCertifiers	Current or Coming (°) Certifiers:
Proba	Probability
Pseudos	Pseudonyms
51pseudos	Informations: List of All Pseudonyms
PseudosNb	Number of Pseudonyms
Pubkey	Public Key
21qualities	Properties: Qualities
qualities	Qualities
Quality	Quality
QualityE	Expected Quality
QualOfIs	The expected quality of ^0 is ^1%
QualOfUnknownIs	The expected quality is ^1%
remainingCerts	Remaining Certifications
requiredCertsNb	^0 certifications, ^1 needed for distance rule
REVOKED	REVOKED
Revoked	Revoked
51revokedIdentities	Informations: List of Revoked Identities
RevokedM	Revoked Identities
RevokedNb	Number of Revoked Identities
Revokeds	Revoked
SDev	Standard Deviation
second	second(s)
Select	Select
SentCertsHistory	Sent Certifications History
60sentCertsHistory	Statistics: Sent Certifications History
Sentries	Sentries
57sentries	Informations: List of Sentries
SentriesNb	Number of Sentries
Sentry	Sentry
Server	Server Version: 
ShowFile	File
SortedByCExpDates	Sorted by Expiration Dates of Certifications
SortedByCExpDatesL	Sorted by Expiration Dates of Certifications (→: limit date)
Status	Status
Threshold	Threshold
TypeUidOrPubkey	Start of Nickname or Public Key
Uid	Nickname of the identity under consideration (keep empty if she is a newcomer):
uid	the identity under consideration
Utc	Blockchain Actual Time
ValIdHeads	Values obtained for some extra certifiers, in decreasing order:
WarnNum	Number of Warnings:
Written_block	First Membership
wwByDate	Sorted by Dates
wwByName	Sorted by Names
01wwFile	Forecast: Preparatory File
wwMeta	Metadata
02wwPerms	Forecast: Permutations
00wwView	ForeCast: WotWizard View
year	year(s)
yes	yes
`
	strFr = `STRINGS

AllCertified	Toutes les identités certifiées non-révoquées :
AllCertifiedIO	Débuts (↑) et fins (↓) de validité de toutes les certifications émises :
AllCertifiers	Tous les certificateurs connus et non-révoqués :
AllCertifiersIO	Débuts (↑) et fins (↓) de validité de toutes les certifications reçues :
50allParameters	Informations : Paramètres de la chaîne de blocs
AnswerNb	Nombre maximal de propositions de certificateurs :
AppMLimitDate	Date limite de réadhésion
AppNLimitDate	Date limite de la demande d'adhésion
AppRLimitDate	Date de révocation
Availability	Disponibilité du prochain envoi de certification
Bct	Temps médian de la chaîne de blocs
Block	Bloc
Brackets	Tranches
Calculator	Calculateur de distance et de qualité
23calculator	Propriétés : Calculateur de distance et de qualité
22centralities	Propriétés : Centralités
centralities	Centralités
Centrality	Centralité
certification	certification :
Certifications	Certifications
certifications	certifications :
55certificationsFrom	Informations : Certifications depuis les membres
certificationsFrom	Certifications depuis les membres
56certificationsTo	Informations : Certifications vers les membres
certificationsTo	Certifications vers les membres
Certifiers	Nouveaux certificateurs :
certifiers	les nouveaux certificateurs
34certLim	Limites : Limites des certifications par date
certsNb	^0 certifications
Client	Version client : 
Computation_duration	Durée du calcul
Day	Jour
day	jour(s)
Delay	Durée (jours)
Distance	Règle de distance
DistanceE	Règle de distance attendue
distanceRule	Règle de distance attendue : ^0% (^1 / ^2)
distanceRule0	Règle de distance attendue : ^0%
20distances	Propriétés : Distances (règle de)
distances	Distances (règle de)
DistOfIs	La distance attendue de ^0 est de ^1% : ^2
DistOfUnknownIs	La distance attendue est de ^1% : ^2
Distribution	Distribution
dossiers	dossiers
Duplicate	^0 apparaît plusieurs fois (ligne ^1)
DuplicateLoc	^0 apparaît plusieurs fois (dans ^1, ligne ^2)
en	English
Entry	Entrée de 
es	Español
Exit	Sortie de 
Explorer	Explorateur
1explorer	Explorateur de la toile de confiance
FirstEntries	Premières entrées
FirstEntriesFlux	Flux des premières entrées
FirstEntriesFluxPM	Flux des premières entrées par membre
Forward	Nombre de jours jusqu'à la date de prévision 
fr	Français
Group	Entrez les pseudos de votre groupe, ligne par ligne, ou laissez la liste vide si vous n'avez pas de groupe :
group	le groupe
Hash	Hash
history	Entrées / sorties de la toile de confiance
hour	heure(s)
In	Entrée
Index	Menu
index	menu
24infosCertif	Propriétés : Informations de certification
infosCertif	Informations de certification
intCertifs	certifications internes
IsntMemberLoc	^0 n'est pas membre (dans ^1, ligne ^2)
IsntMemberNorMissing	^0 n'est ni membre ni exclu (ligne ^1)
IsntMemberNorMissingLoc	^0 n'est ni membre ni exclu (dans ^1, ligne ^2)
IsUnknown	^0 est inconnu(e) (ligne ^1)
IsUnknownLoc	^0 est inconnu(e) (dans ^1, ligne ^2)
KO	KO
language	Choix de la langue
LastApplication	Dernière demande d'adhésion
LCCertifiers	Elle a été certifiée par 
LCComplement	Cette identité va bientôt manquer de certifications à la date suivante :
LimitsC	Limites des certifications
LimitDate	Date limite
32limitsCerts	Limites : Dates limites des certifications
limitsCerts	Dates limites des certifications
LimitsM	Limites des adhésions
30limitsMember	Limites : Dates limites des adhésions
limitsMember	Dates limites des adhésions
31limitsMissing	Limites : Dates limites des adhésions non-renouvelées
limitsMissing	Dates limites des adhésions non-renouvelées
limitsSubCerts	Date à laquelle l'identité affichée manquera de certifications
limitsSubMember	Date à laquelle l'adhésion du membre affiché (avec sa règle de distance) ne sera plus valide
limitsSubMissing	Date à laquelle l'identité exclue affichée sera définitivement révoquée
LMCertifiers	certifié(e) par
LMComplement	Pour rester membre, cette identité doit renouveler son adhésion :
Losses	Pertes
LossesFlux	Flux de pertes
LossesFluxPM	Flux de pertes par membre
Mean	Moyenne
Median	Médiane
MEMBER	MEMBRE
53memberIdentities	Informations : Liste des membres
Members	Membres
404membersCountFlux	Evolution : Flux de membres (graphique)
405membersCountFluxPM	Evolution : Flux de membres par membre (graphique)
403membersCountG	Evolution : Nombre de membres (graphique)
400membersCountT	Evolution :  Nombre de membres (liste)
407membersFEFlux	Evolution : Flux des premières entrées (graphique)
408membersFEFluxPM	Evolution : Flux des premières entrées par membre (graphique)
406membersFirstEntryG	Evolution : Premières entrées (graphique)
401membersFirstEntryT	Evolution : Premières entrées (liste)
MembersFlux	Flux de membres
MembersFluxPM	Flux de membres par membre
410membersLossFlux	Evolution : Flux de pertes (graphique)
411membersLossFluxPM	Evolution : Flux de pertes par membre (graphique)
409membersLossG	Evolution : Pertes (graphique)
402membersLossT	Evolution : Pertes (liste)
MembersNb	Nombre de membres
33memLim	Limites : Limites des adhésions par date
minApplicationDate	Au moins deux mois d'attente après la dernière adhésion (^0)
minute	minute(s)
MISSING	EXCLU(E)
Missing	Exclu(e)s en attente de réadhésion
52missingIdentities	Informations : Liste des identités exclues en attente de réadhésion
MissingNb	Nombre des exclu(e)s
Missings	Exclu(e)s
month	mois
neededCerts	certifications nécessaires :
Never	Jamais
NEWCOMER	ARRIVANT(E)
54newcomerIdentities	Informations : Liste des arrivant(e)s
Newcomers	Arrivant(e)s
newcomers	dossiers en attente
NewcomersNb	Nombre des arrivant(e)s
Nickname	Pseudo
no	non
noGroupL	Ignorer le groupe
NotMembers	Non membres
OK	OK
Out	Sortie
Para_avgGenTime	Temps moyen d'écriture d'un bloc (temps attendu)
Para_c	Croissance relative du DU pour chaque période [dtReeval]
Para_dt	Durée entre deux DU
Para_dtDiffEval	Nombre de blocs exigé avant de pouvoir réévaluer la difficulté minimale
Para_dtReeval	Durée entre deux réévaluations du DU
Para_idtyWindow	Durée maximale d'attente en piscine d'une identité avant expiration
Para_medianTimeBlocks	Nombre de blocs pour le calcul du temps moyen
Para_msPeriod	Durée minimale entre deux demandes d'adhésion d'un membre
Para_msValidity	Durée maximale d'une adhésion
Para_msWindow	Durée maximale d'attente en piscine d'une demande d'adhésion avant expiration
Para_percentRot	Proportion de membres calculants non-exclus de la preuve de travail
Para_sigPeriod	Durée minimale entre deux certifications d'un même membre
Para_sigQty	Nombre minimal de certifications obtenues nécessaire pour être membre
Para_sigStock	Nombre maximal de certifications valides émises par membre
Para_sigValidity	Durée maximale de validité d'une certification
Para_sigReplay	Durée minimale avant le renouvellement d'une certification
Para_sigWindow	Durée maximale d'attente en piscine d'une certification avant expiration
Para_stepMax	Distance maximale entre un membre et [xpercent] des membres référents
Para_txWindow	Durée maximale d'attente en piscine d'une transaction avant expiration
Para_ud0	DU(0) : Dividende Universel initial
Para_udReevalTime0	Date de la première réévaluation du DU
Para_udTime0	Date du premier DU
Para_xpercent	Proportion minimale de membres référents à atteindre pour respecter la règle de distance
Parameters	Paramètres
pending	En cours de réadhésion
PendingApplication	Demande d'adhésion en attente
Permutations	Permutations
permutations	permutations
PermutationsNb	Nombre de permutations : 
PresentCertified	Identités actuellement ou prochainement (°) certifiées :
PresentCertifiers	Actuels ou prochains (°) certificateurs :
Proba	Probabilité
Pseudos	Pseudonymes
51pseudos	Informations : Liste de tous les pseudonymes
PseudosNb	Nombre de pseudonymes
Pubkey	Clef publique
21qualities	Propriétés : Qualités
qualities	Qualités
Quality	Qualité
QualityE	Qualité attendue
QualOfIs	La qualité attendue de ^0 est de ^1%
QualOfUnknownIs	La qualité attendue est de ^1%
remainingCerts	Certifications restantes
requiredCertsNb	^0 certifications, ^1 nécessaires pour la règle de distance
REVOKED	RÉVOQUÉ(E)
Revoked	Révoqué(e)
51revokedIdentities	Informations : Liste des identités révoquées
RevokedM	Identités révoquées
RevokedNb	Nombre d'identités révoquées
Revokeds	Révoqué(e)s
SDev	Écart type
second	seconde(s)
Select	Sélectionner
SentCertsHistory	Évolution des certifications émises
60sentCertsHistory	Statistiques : Évolution des certifications émises
Sentries	Membres référents
57sentries	Informations : Liste des membres référents
SentriesNb	Nombre de membres référents
Sentry	Référent
Server	Version serveur : 
ShowFile	Fichier
SortedByCExpDates	Tri par dates d'expiration des certifications
SortedByCExpDatesL	Tri par dates d'expiration des certifications (→ : date limite)
Status	Statut
Threshold	Seuil
TypeUidOrPubkey	Début de pseudo ou de clef publique
Uid	Pseudo de l'identité considérée (laissez vide s'il s'agit d'un arrivant) :
uid	l'identité considérée
Utc	Temps réel de la chaîne de blocs
ValIdHeads	Valeurs obtenues, par ordre décroissant, pour différents certificateurs supplémentaires :
WarnNum	Nombre d'alertes :
Written_block	Première adhésion
wwByDate	Tri par dates
wwByName	Tri par noms
01wwFile	Prévisions : Fichier préparatoire
wwMeta	Métadonnées
02wwPerms	Prévisions : Permutations
00wwView	Prévisions : Fenêtre WotWizard
year	année(s)
yes	oui
`
	strEs = `STRINGS

AllCertified	Toutes les identités certifiées non-révoquées :
AllCertifiedIO	Débuts (↑) et fins (↓) de validité de toutes les certifications émises :
AllCertifiers	Tous les certificateurs connus et non-révoqués :
AllCertifiersIO	Débuts (↑) et fins (↓) de validité de toutes les certifications reçues :
50allParameters	Informations : Paramètres de la chaîne de blocs
AnswerNb	Nombre maximal de propositions de certificateurs :
AppMLimitDate	Date limite de réadhésion
AppNLimitDate	Date limite de la demande d'adhésion
AppRLimitDate	Date de révocation
Availability	Disponibilité du prochain envoi de certification
Bct	Temps médian de la chaîne de blocs
Block	Bloc
Brackets	Tranches
Calculator	Calculateur de distance et de qualité
23calculator	Propriétés : Calculateur de distance et de qualité
22centralities	Propriétés : Centralités
centralities	Centralités
Centrality	Centralité
certification	certification :
Certifications	Certifications
certifications	certifications :
55certificationsFrom	Informations : Certifications depuis les membres
certificationsFrom	Certifications depuis les membres
56certificationsTo	Informations : Certifications vers les membres
certificationsTo	Certifications vers les membres
Certifiers	Nouveaux certificateurs :
certifiers	les nouveaux certificateurs
34certLim	Limites : Limites des certifications par date
certsNb	^0 certifications
Client	Version client : 
Computation_duration	Durée du calcul
Day	Jour
day	jour(s)
Delay	Durée (jours)
Distance	Règle de distance
DistanceE	Règle de distance attendue
distanceRule	Règle de distance attendue : ^0% (^1 / ^2)
distanceRule0	Règle de distance attendue : ^0%
20distances	Propriétés : Distances (règle de)
distances	Distances (règle de)
DistOfIs	La distance attendue de ^0 est de ^1% : ^2
DistOfUnknownIs	La distance attendue est de ^1% : ^2
Distribution	Distribution
dossiers	dossiers
Duplicate	^0 apparaît plusieurs fois (ligne ^1)
DuplicateLoc	^0 apparaît plusieurs fois (dans ^1, ligne ^2)
en	English
Entry	Entrée de 
es	Español
Exit	Sortie de 
Explorer	Explorateur
1explorer	Explorateur de la toile de confiance
FirstEntries	Premières entrées
FirstEntriesFlux	Flux des premières entrées
FirstEntriesFluxPM	Flux des premières entrées par membre
Forward	Nombre de jours jusqu'à la date de prévision 
fr	Français
Group	Entrez les pseudos de votre groupe, ligne par ligne, ou laissez la liste vide si vous n'avez pas de groupe :
group	le groupe
Hash	Hash
history	Entrées / sorties de la toile de confiance
hour	heure(s)
In	Entrée
Index	Menu
index	menu
24infosCertif	Propriétés : Informations de certification
infosCertif	Informations de certification
intCertifs	certifications internes
IsntMemberLoc	^0 n'est pas membre (dans ^1, ligne ^2)
IsntMemberNorMissing	^0 n'est ni membre ni exclu (ligne ^1)
IsntMemberNorMissingLoc	^0 n'est ni membre ni exclu (dans ^1, ligne ^2)
IsUnknown	^0 est inconnu(e) (ligne ^1)
IsUnknownLoc	^0 est inconnu(e) (dans ^1, ligne ^2)
KO	KO
language	Choix de la langue
LastApplication	Dernière demande d'adhésion
LCCertifiers	Elle a été certifiée par 
LCComplement	Cette identité va bientôt manquer de certifications à la date suivante :
LimitsC	Limites des certifications
LimitDate	Date limite
32limitsCerts	Limites : Dates limites des certifications
limitsCerts	Dates limites des certifications
LimitsM	Limites des adhésions
30limitsMember	Limites : Dates limites des adhésions
limitsMember	Dates limites des adhésions
31limitsMissing	Limites : Dates limites des adhésions non-renouvelées
limitsMissing	Dates limites des adhésions non-renouvelées
limitsSubCerts	Date à laquelle l'identité affichée manquera de certifications
limitsSubMember	Date à laquelle l'adhésion du membre affiché (avec sa règle de distance) ne sera plus valide
limitsSubMissing	Date à laquelle l'identité exclue affichée sera définitivement révoquée
LMCertifiers	certifié(e) par
LMComplement	Pour rester membre, cette identité doit renouveler son adhésion :
Losses	Pertes
LossesFlux	Flux de pertes
LossesFluxPM	Flux de pertes par membre
Mean	Moyenne
Median	Médiane
MEMBER	MEMBRE
53memberIdentities	Informations : Liste des membres
Members	Membres
404membersCountFlux	Evolution : Flux de membres (graphique)
405membersCountFluxPM	Evolution : Flux de membres par membre (graphique)
403membersCountG	Evolution : Nombre de membres (graphique)
400membersCountT	Evolution :  Nombre de membres (liste)
407membersFEFlux	Evolution : Flux des premières entrées (graphique)
408membersFEFluxPM	Evolution : Flux des premières entrées par membre (graphique)
406membersFirstEntryG	Evolution : Premières entrées (graphique)
401membersFirstEntryT	Evolution : Premières entrées (liste)
MembersFlux	Flux de membres
MembersFluxPM	Flux de membres par membre
410membersLossFlux	Evolution : Flux de pertes (graphique)
411membersLossFluxPM	Evolution : Flux de pertes par membre (graphique)
409membersLossG	Evolution : Pertes (graphique)
402membersLossT	Evolution : Pertes (liste)
MembersNb	Nombre de membres
33memLim	Limites : Limites des adhésions par date
minApplicationDate	Au moins deux mois d'attente après la dernière adhésion (^0)
minute	minute(s)
MISSING	EXCLU(E)
Missing	Exclu(e)s en attente de réadhésion
52missingIdentities	Informations : Liste des identités exclues en attente de réadhésion
MissingNb	Nombre des exclu(e)s
Missings	Exclu(e)s
month	mois
neededCerts	certifications nécessaires :
Never	Jamais
NEWCOMER	ARRIVANT(E)
54newcomerIdentities	Informations : Liste des arrivant(e)s
Newcomers	Arrivant(e)s
newcomers	dossiers en attente
NewcomersNb	Nombre des arrivant(e)s
Nickname	Pseudo
no	non
noGroupL	Ignorer le groupe
NotMembers	Non membres
OK	OK
Out	Sortie
Para_avgGenTime	Temps moyen d'écriture d'un bloc (temps attendu)
Para_c	Croissance relative du DU pour chaque période [dtReeval]
Para_dt	Durée entre deux DU
Para_dtDiffEval	Nombre de blocs exigé avant de pouvoir réévaluer la difficulté minimale
Para_dtReeval	Durée entre deux réévaluations du DU
Para_idtyWindow	Durée maximale d'attente en piscine d'une identité avant expiration
Para_medianTimeBlocks	Nombre de blocs pour le calcul du temps moyen
Para_msPeriod	Durée minimale entre deux demandes d'adhésion d'un membre
Para_msValidity	Durée maximale d'une adhésion
Para_msWindow	Durée maximale d'attente en piscine d'une demande d'adhésion avant expiration
Para_percentRot	Proportion de membres calculants non-exclus de la preuve de travail
Para_sigPeriod	Durée minimale entre deux certifications d'un même membre
Para_sigQty	Nombre minimal de certifications obtenues nécessaire pour être membre
Para_sigStock	Nombre maximal de certifications valides émises par membre
Para_sigValidity	Durée maximale de validité d'une certification
Para_sigReplay	Durée minimale avant le renouvellement d'une certification
Para_sigWindow	Durée maximale d'attente en piscine d'une certification avant expiration
Para_stepMax	Distance maximale entre un membre et [xpercent] des membres référents
Para_txWindow	Durée maximale d'attente en piscine d'une transaction avant expiration
Para_ud0	DU(0) : Dividende Universel initial
Para_udReevalTime0	Date de la première réévaluation du DU
Para_udTime0	Date du premier DU
Para_xpercent	Proportion minimale de membres référents à atteindre pour respecter la règle de distance
Parameters	Paramètres
pending	En cours de réadhésion
PendingApplication	Demande d'adhésion en attente
Permutations	Permutations
permutations	permutations
PermutationsNb	Nombre de permutations : 
PresentCertified	Identités actuellement ou prochainement (°) certifiées :
PresentCertifiers	Actuels ou prochains (°) certificateurs :
Proba	Probabilité
Pseudos	Pseudonymes
51pseudos	Informations : Liste de tous les pseudonymes
PseudosNb	Nombre de pseudonymes
Pubkey	Clef publique
21qualities	Propriétés : Qualités
qualities	Qualités
Quality	Qualité
QualityE	Qualité attendue
QualOfIs	La qualité attendue de ^0 est de ^1%
QualOfUnknownIs	La qualité attendue est de ^1%
remainingCerts	Certifications restantes
requiredCertsNb	^0 certifications, ^1 nécessaires pour la règle de distance
REVOKED	RÉVOQUÉ(E)
Revoked	Révoqué(e)
51revokedIdentities	Informations : Liste des identités révoquées
RevokedM	Identités révoquées
RevokedNb	Nombre d'identités révoquées
Revokeds	Révoqué(e)s
SDev	Écart type
second	seconde(s)
Select	Sélectionner
SentCertsHistory	Évolution des certifications émises
60sentCertsHistory	Statistiques : Évolution des certifications émises
Sentries	Membres référents
57sentries	Informations : Liste des membres référents
SentriesNb	Nombre de membres référents
Sentry	Référent
Server	Version serveur : 
ShowFile	Fichier
SortedByCExpDates	Tri par dates d'expiration des certifications
SortedByCExpDatesL	Tri par dates d'expiration des certifications (→ : date limite)
Status	Statut
Threshold	Seuil
TypeUidOrPubkey	Début de pseudo ou de clef publique
Uid	Pseudo de l'identité considérée (laissez vide s'il s'agit d'un arrivant) :
uid	l'identité considérée
Utc	Temps réel de la chaîne de blocs
ValIdHeads	Valeurs obtenues, par ordre décroissant, pour différents certificateurs supplémentaires :
WarnNum	Nombre d'alertes :
Written_block	Première adhésion
wwByDate	Tri par dates
wwByName	Tri par noms
01wwFile	Previsiones : Fichero preparatorio
wwMeta	Métadonnées
02wwPerms	Prévisions : Permutations
00wwView	Prévisions : Fenêtre WotWizard
year	année(s)
yes	oui
`

)
