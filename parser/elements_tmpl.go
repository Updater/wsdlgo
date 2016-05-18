package parser

const elementsTmpl = `
{{define "Elements"}}
	{{range .}}{{if .Type}}{{replaceReservedWords .Name | makeExported}} {{if eq .MaxOccurs "unbounded"}}[]{{.Type | toGoType}}{{else}}{{.Type | toGoPointerType}}{{end}} ` + "`" + `xml:"{{.Name}}"` + "`" + ` 
	{{else}}{{if .Name}}{{$type := replaceReservedWords .Name | makeUnexported}} {{$type | toGoPointerType}}Inner {{end}}
	{{end}}{{end}}
{{end}}


{{define "NestedElements"}}
	{{range .}}
		{{if not .Type}}
			{{if .Name}}
				{{$typeName := replaceReservedWords .Name | makeUnexported}}
				type {{$typeName}}Inner struct {
					{{with .ComplexType}}
						{{template "Elements" .Sequence}}
						{{template "Elements" .Choice}}
						{{template "Elements" .SequenceChoice}}
						{{template "Elements" .All}}
						{{template "Attributes" .Attributes}}
					{{end}}
				}
			{{end}}
		{{end}}
	{{end}}
{{end}}`
