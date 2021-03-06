if [ -f $ZDOTDIR/.zshrc ]; then source $ZDOTDIR/.zshrc; fi
export PS1="[{{.Owner}}/{{.Name}}] $PS1"

precmd() { eval "$PROMPT_COMMAND" }

{{- range $K, $V := .Env}}
{{- if eq $K "PATH"}}
export {{$K}}="{{$V}}:$PATH"
{{- else}}
export {{$K}}="{{$V}}"
{{- end}}
{{- end}}

{{range $K, $CMD := .Scripts}}
alias {{$K}}='state run {{$CMD}}'
{{end}}

cd "{{.WD}}"

{{.UserScripts}}
