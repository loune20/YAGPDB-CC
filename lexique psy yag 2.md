{{$args := parseArgs 0 "Syntaxe : <mot à définiri>"
    (carg "string" "Mot à définir")}}



{{$col1 := "RSD \nSED \nSPD \nTAG \nTCA \nTCC \nTCO/TOP \nTDAH \nTSPT"}}

{{$col2 := "Trigger \nTrigger Warning (TW) \nContent Warning (CW) \nFlash Warning (FW)"}}

{{$expli_lex := "Les mots définis dans le lexique sont les suivants : \n*(Utilisation : `-dico2 mot-a-definir`)*"}}

{{$embed_lexique := cembed

	"title" "Lexique collaboratif et  évolutif de la neuroatypie" 

	"description" "Lexique non-exhaustif et évolutif visant à éclaircir les termes et abréviations courantes liées au TDAH, à la neurodiversité, ainsi qu’aux sujets fréquemment abordés au sein de l’association en général. Ce lexique est collaboratif, car les membres du Discord peuvent y participer en suggérant des termes, ajoutant des définitions ou des précisions." 

	"color" 4645612

	"fields" (cslice
		(sdict "name" "Lexique" "value" $expli_lex "inline" false)
		(sdict "name" "AD ─ Psychologue" "value" "faire `-dico`" "inline" true)
		(sdict "name" "RSD ─ TSPT" "value" $col1 "inline" true)
		(sdict "name" "Trigger ─ TW, CW & FW" "value" $col2 "inline" true)
	)
}}


{{$def := ""}}
{{$envoi_def := true}}


{{if $args.IsSet 0 }}
	{{$mot := $args.Get 0}}

	{{if eq $mot "RSD"}}
		{{$def = "Rejection Sensitive Dysphoria, dysphorie sensible au rejet. Trouble comorbide du TDAH entraînant des réactions émotionnelles très fortes et disproportionnées aux jugements négatifs (ou perçus comme tels). Affecte beaucoup l'estime de soi."}}

	{{else if eq $mot "SED"}}
		{{$def = "Syndrome d'Ehlers-Danlos, trouble des tissus conjonctifs. Manifestations très variées."}}

	{{else if eq $mot "SPD"}}
		{{$def = "Sensory Processing Disorder (anglais), signifiant trouble du traitement sensoriel. Défini par la difficulté du cerveau à traiter correctement les informations sensorielles quelles qu'elles soient, ce qui provoque des hypersensibilités, des hyposensibilités (réaction à des stimuli trop fort ou trop faible) et/ou des des dysfonctionnement de la perception sensorielle (incapacité à comprendre un son si le bruit de fond est trop fort, vision excessivement mauvaise dans la pénombre, etc.). \nC'est une caractéristique de l’autisme, mais qu’on peut aussi retrouver dans le TDAH."}}

	{{else if eq $mot "TAG"}}
		{{$def = "Trouble Anxieux Généralisé"}}

	{{else if eq $mot "TCA"}}
		{{$def = "Trouble du Comportement Alimentaire"}}

	{{else if eq $mot "TCC"}}
		{{$def = "Thérapie Cognitivo-Comportementale, ensemble de traitements des troubles psychiatriques basés sur une approche fondée sur la psychologie scientifique et ayant la particularité de traiter les difficultés immédiates du patient."}}

	{{else if eq $mot "TCO"}}
		{{$def = "Trouble de la Conduite et/ou de l'opposition, ou TOP (Trouble Oppositionnel avec Provocation). Le trouble de la conduite est marqué par un comportement dans lequel le droit des autres ou la norme sociale est bafouée. Le trouble de l'opposition est marqué par une désobéissance, une opposition généralisée face à l'autorité."}}

	{{else if eq $mot "TOC"}}
		{{$def = "Trouble Obsessionnel Compulsif"}}

	{{else if eq $mot "TDAH"}}
		{{$def = "Trouble du déficit de de l’attention avec ou sans hyperactivité. \nLe terme anglais (ADHD - Attention Deficit Hyperactivity Disorder) est aussi fréquemment utilisé. \nOn donne aussi parfois au D le sens de dysfonctionnement au lieu de déficitaire, étant donné que le TDAH n’est pas un problème de manque d’attention, mais de régulation de l’attention. Cette erreur étant l’origine de clichés *(«tu ne peux pas avoir un TDAH, tu peux rester concentrer sur telle activité pendant très longtemps» ─ alors que l’hyperfocus est une manifestation typique du TDAH)* faisant obstacle au diagnostic et à la prise en charge du TDAH, cette appellation est intéressante."}}

	{{else if eq $mot "TSA"}}
		{{$def = "Troubles du Spectre Autistique (synonyme d'autisme)"}}

	{{else if eq $mot "TSPT"}}
		{{$def = "Trouble de Stress Post-Traumatique ; un trouble lié aux traumatismes se déclenchant à la suite d'un événement traumatisant vécu directement ou indirectement. Il est caractérisé par des symptômes d'évitements liés au trauma, d'intrusion (souvenirs, pensées...) et d'hyperstimulation (hypervigilance, hyperréactivité). \nDit PTSD (Post-Traumatic Stress Disorder) en anglais."}}

	{{else if eq $mot "Trigger"}}
		{{$def = "terme anglais signifiant “déclencheur” utilisé à l’origine pour les personnes atteintes de TSPT. Désigne l’élément déclencheur de résurgence traumatiques, donnant lieux à des flashback, crises d’angoisses, … \nIl est aussi utilisé de façon générale pour parler des déclencheurs d’autres troubles (notamment TCA). \nAttention, il ne désigne pas une simple gêne mais bien une véritable détresse ; les triggers sont spécifiques et propres à chacun."}}

	{{else if eq $mot "Trigger Warning"}}
		{{$def = "(ou **TW**, avertissement de déclencheur) : avertissement placé au début d’un contenu (texte, vidéo, conversation orale…) visant à prévenir les personnes souffrant d’un TSPT ou d’autres troubles que le contenu qui va suivre pourrait les trigger (déclencher une résurgence traumatique), parce qu’il contient un sujet précis."}}

	{{else if eq $mot "Content Warning"}}
		{{$def = "(ou **CW**, avertissement de contenu) : équivalent de TW pour la population générale ; vise à avertir de la présence de sujets fréquemment jugés comme sensibles."}}

	{{else if eq $mot "Flash Warning"}}
		{{$def = "(ou **FW**, avertissement de lumières stroboscopiques/flash) : avertissement utilisé les personnes épileptiques qu’une vidéo pourrait provoquer une crise."}}

	{{else if eq $mot "Stim"}}
		{{$def = "auto-stimulation (sensorielle), mouvement ou bruit fait de manière plutôt répétitive qui aide à réguler des émotions et gérer les situations de surstimulation (trop de stimulations) ou sous stimulation (manque de stimulation) sensorielle. \nExemples de stim : flapping (mouvement répétitif avec les mains, comme des applaudissements, secouement frénétique des mains, souvent dans des moments d’euphorie), rongeage d’ongles (ou peau des doigts), craquage d’os, etc. \nStimmer n’est pas quelque chose qu’il faut réprimer ; au contraire, il s’agit d’une action bénéfique et souvent nécessaire."}}




	{{else}}
		{{$envoi_def = false}}
	{{end}}


	{{$embed_def := cembed
		"title" (joinStr "" "Lexique : " $mot)
		"description" (joinStr "" $mot " : " $def) 
		"color" 4645612
	}}


	{{if eq $envoi_def true}}
		{{sendMessage nil $embed_def}}
	{{else}}
		{{sendMessage nil "Je ne connais pas encore ce mot :( \nTape `-dico2` pour la liste des mots que je connais \n Tu peux aussi ping Louis (@loune20) pour une suggestion de mot à ajouter !"}}
	{{end}}


{{else}}
	{{sendMessage nil $embed_lexique}}
{{end}}
