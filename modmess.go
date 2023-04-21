{{$args := parseArgs 2 "Mentionner un salon où envoyer le message et écrire le message"

	(carg "channel" "le chan où envoyer le message")

	(carg "string" "message à envoyer")}}

{{$chan := $args.Get 0}}

{{$message := $args.Get 1}}

{{$sender := (toString .Message.Author.Username)}} 

{{sendMessage $chan.ID $message}}

{{sendMessage 904709855716196422 (joinStr "" $sender " a envoyé dans " $chan.Mention " un message de modération lisant : \n> " $message)}}

{{deleteTrigger 0}}
