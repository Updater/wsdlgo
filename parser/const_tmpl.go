package parser

const constTmpl = `
{{define "Const"}}
	{{$type := replaceReservedWords .Name | makeUnexported}}
	{{with .Restriction}}
		{{range .Enumeration}}
			{{$type}}{{$value := normalize .Value}}{{$value | lint}} {{$type}} = "{{$value}}" {{end}}
	{{end}}
{{end}}
`
