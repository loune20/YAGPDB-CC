{{/*THREAD*/}}
{{$thread := ID}}
{{$x := sendMessageRetID $thread "Rien à signaler, je désarchive juste !"}}
{{deleteMessage $thread $x 5}}


Empêchement de l'archivage des threads utiles : fait :white_check_mark:
