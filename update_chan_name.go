{{$args := parseArgs 1 "done, todo, yay"

    (carg "string" "done, todo, yay")}}

{{$status := $args.Get 0}}

 

{{if eq $status "done"}}

	{{$new_name := joinStr "" "✅" (slice .Channel.Name 1)}}

	{{editChannelName nil $new_name}}

        https://tenor.com/view/jim-carrey-yes-sir-you-got-it-you-rock-yes-boss-gif-15459239

 

{{else if eq $status "todo"}}

	{{$new_name := joinStr "" "❌" (slice .Channel.Name 1)}}

	{{editChannelName nil $new_name}}

        https://tenor.com/view/taking-notes-writing-down-jeremiah-buie-gif-16911555

 

{{else if eq $status "yay"}}

	{{$new_name := joinStr "" "✨" (slice .Channel.Name 1)}}

	{{editChannelName nil $new_name}}

	Ayé :relieved:

 

{{else}}

	{{sendMessage nil "lol non"}}

{{end}}
