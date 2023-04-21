{{$args := parseArgs 1 "ID du message à save"

    (carg "int" "ID")}}

{{$id := $args.Get 0}}

{{$mess := getMessage (.Channel.ID) $id}}

 

{{$mess_total := joinStr "" "Message envoyé par : " $mess.Author.Username "\n À : " $mess.Timestamp.Parse.String "Dans :  \n" .Channel.Mention "\n> " $mess.Content}}

 

{{sendMessage 907795776804126770 $mess_total}}

{{deleteTrigger 0}} 
