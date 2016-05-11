package parser

const constTmpl = `
{{define "Const"}}
{{$type := replaceReservedWords .Name | makeUnexported}}
{{with .Restriction}}
	{{range .Enumeration}}
		{{$type}}{{$value := normalize .Value | lint}}{{$value}} {{$type}} = "{{$value}}" {{end}}
{{end}}
{{end}}
`
