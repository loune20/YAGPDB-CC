{{$args := parseArgs 0 "Syntaxe : <mot à définiri>"
    (carg "string" "Mot à définir")}}



{{$col1 := "AD \nAlexithymie \nBPD \n Troubles dys \nDyscalculie \nDysgraphie \nDyslexie \nDysorthographie \nDysphasie \nDyspraxie \n Dysfonction exécutive"}}

{{$col2 := "Ergothérapie \nHyperfocus \nIntérêt spécifique \nMéthylphénidate \nNeuroatypie \nPsychiatre \nPsychologue"}}

{{$expli_lex := "Les mots définis dans le lexique sont les suivants : \n*(Utilisation : `-dico mot-a-definir`)*"}}

{{$embed_lexique := cembed

	"title" "Lexique collaboratif et  évolutif de la neuroatypie" 

	"description" "Lexique non-exhaustif et évolutif visant à éclaircir les termes et abréviations courantes liées au TDAH, à la neurodiversité, ainsi qu’aux sujets fréquemment abordés au sein de l’association en général. Ce lexique est collaboratif, car les membres du Discord peuvent y participer en suggérant des termes, ajoutant des définitions ou des précisions." 

	"color" 4645612

	"fields" (cslice
		(sdict "name" "Lexique" "value" $expli_lex "inline" false)
		(sdict "name" "AD ─ Dysfonction exécutive" "value" $col1 "inline" true)
		(sdict "name" "Ergothérapie ─ Psychologue" "value" $col2 "inline" true)
		(sdict "name" "RSD ─ Stim" "value" "faire `-dico2`" "inline" true)
	)
}}


{{$def := ""}}
{{$envoi_def := true}}


{{if $args.IsSet 0 }}
	{{$mot := $args.Get 0}}

	{{if eq $mot "AD"}}
		{{$def = "antidépresseur"}}

	{{else if eq $mot "Alexithymie"}}
		{{$def = "trouble consistant en une difficulté à identifier et exprimer ses propres émotions et/ou celles des autres. Ce n'est pas une impossibilité à les ressentir."}}

	{{else if eq $mot "BPD"}}
		{{$def = "Borderline Personality Disorder (trouble de la personnalité borderline ou limite). Trouble caractérisé par une impulsivité majeure et une instabilité marquée des émotions, des relations interpersonnelles et de l'image de soi."}}

	{{else if eq $mot "DYS"}}
		{{$def = "Regroupe les troubles de la dyslexie, dyspraxie, dysphasie, ainsi que certaines manifestations induites de ces troubles comme la dyscalculie, la dysgraphie ou la dysorthographie."}}

	{{else if eq $mot "Dysfonction exécutive"}}
		{{$def = "trouble/dysfonctionnement des fonctions exécutives. Les fonctions exécutives gèrent l’organisation et la planification des actions, la gestion des émotions, l’impulsivité, etc. La dysfonction exécutive rend plus difficile voire empêche de débuter une action, de la faire dans l’ordre et de la terminer. \nLa dysfonction exécutive, parfois abrégée en “dysfex”, est caractéristique chez les personnes ayant un TDAH."}}
	
	{{else if eq $mot "Dyscalculie"}}
		{{$def = "trouble spécifique et durable ralentissant ou empêchant les acquisitions numériques et/ou du calcul nécessaires aux mathématiques, que ce soit l’accès à la numération, l’apprentissage des opérations comme l’addition, soustraction, multiplication et division, la résolution de problèmes ou la géométrie."}}

	{{else if eq $mot "Dysgraphie"}}
		{{$def = "trouble spécifique d’apprentissage affectant le geste graphique et l’aspect de l’écriture manuscrite."}}

	{{else if eq $mot "Dyslexie"}}
		{{$def = "trouble se caractérisant par des difficultés pour lire de façon correcte et fluide, pour décoder un texte et pour orthographier."}}

	{{else if eq $mot "Dysorthographie"}}
		{{$def = "trouble spécifique de l’écriture touchant aussi bien la production d’écrit que l’acquisition du lexique orthographique, des règles grammaticales et de conjugaison. Comme la dyslexie, à laquelle elle est étroitement liée, elle apparaît dès l’apprentissage du langage écrit."}}

	{{else if eq $mot "Dysphasie"}}
		{{$def = "un trouble structurel, inné et durable de l'apprentissage et du développement du langage oral."}}

	{{else if eq $mot "Dyspraxie"}}
		{{$def = "trouble neurologique qui apparaît dès l'enfance. Il affecte la planification, l'organisation, l'exécution et l'automatisation des gestes et des mouvements. Aussi appelée trouble développemental de la coordination."}}

	{{else if eq $mot "Ergothérapie"}}
		{{$def = "discipline paramédicale visant au bien-être, à la santé et à l'autonomie des individus, notamment handicapé·es, via l'activité des personnes. Aide fournie via de la rééducation, des conseils d'aménagements, etc., en collaboration avec des professionnel·les de santé (kiné, psy, orthophoniste, etc.)"}}

	{{else if eq $mot "Hyperfocus"}}
		{{$def = "forme de concentration très intense sur un sujet en particulier. La personne ne contrôle pas ou peu quand déclencher cet état ou le sujet sur lequel il porte. Une personne en état d'hyperfocus va fréquemment oublier ses besoins corporels (soif, douleur, fatigue, etc.) pendant de longues heures, et le sujet de son hyperfocus peut être présent à l'esprit de façon obsédante pendant longtemps. Cet état peut avoir des durées variables, de quelques heures/jour à plusieurs semaines, avec des fluctuations d'intensité. \nManifestation particulièrement typique du TDAH (problème de la régulation de l'attention), mais que l'on retrouve aussi dans le TSA par exemple. \nOn parle souvent d’hyperfixation pour désigner le sujet de l’hyperfocus."}}

	{{else if eq $mot "Intérêt spécifique"}}
		{{$def = "intérêt très fort envers un sujet chez une personne autiste;  souvent défini comme restreint et répétitif. Ce n'est pas l'équivalent exact d'une passion. \nOn l’abrège souvent en “IS”. \n[Pour en savoir plus](https://www.youtube.com/watch?v=kbd0wvMFj5U) (vidéo youtube, sous-titrée, 15 minutes)"}}

	{{else if eq $mot "Méthylphénidate"}}
		{{$def = "principe actif le plus utilisé dans le traitement médicamenteux du TDAH. On le retrouve dans la Ritaline et le Concerta, pour citer les plus fréquents. \nOn l’abrège souvent en *mph*."}}

	{{else if eq $mot "Neuroatypie"}}
		{{$def = "aussi neurodivergence ; définit un fonctionnement cognitif qui diffère de la norme. On utilise le plus souvent ce terme pour parler de TSA, de TDAH et de troubles dys mais les définitions (qui varient) incluent aussi d’autres troubles (schizophrénie, bipolarité, TOC, syndrôme de Tourette, etc.)"}}

	{{else if eq $mot "Psychiatre"}}
		{{$def = "médecin spécialiste des troubles psychologiques et des maladies mentales."}}

	{{else if eq $mot "Psychologue"}}
		{{$def = "professionnel·e de psychologie (comportements humains). Ce n'est pas un médecin, et ne peut donc ni diagnostiquer officiellement, ni prescrire de médicaments, ni être remboursé·e par la sécurité sociale."}}


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
		{{sendMessage nil "Je ne connais pas encore ce mot :( \nTape `-dico` pour la liste des mots que je connais \n Tu peux aussi ping Louis (@loune20) pour une suggestion de mot à ajouter !"}}
	{{end}}


{{else}}
	{{sendMessage nil $embed_lexique}}
{{end}}
