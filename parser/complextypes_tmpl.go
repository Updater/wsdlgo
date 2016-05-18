package parser

const complexTypesTmpl = `
{{define "ComplexContent"}}
	{{$baseType := toGoType .Extension.Base}}
	{{ if $baseType }}
		{{$baseType}}
	{{end}}
	{{template "Elements" .Extension.Sequence}}
	{{template "Attributes" .Extension.Attributes}}
{{end}}

{{define "ComplexTypeNested"}}
	{{if .Name}}
		{{with .ComplexType}}
			{{template "NestedElements" .Sequence}}
			{{template "NestedElements" .Choice}}
			{{template "NestedElements" .SequenceChoice}}
			{{template "NestedElements" .All}}
		{{end}}
	{{end}}
{{end}}

{{define "ComplexTypeElements"}}
	{{if .Name}}
		type {{replaceReservedWords .Name | makeUnexported}} struct {
			{{with .ComplexType}}
				{{template "Elements" .Sequence}}
				{{template "Elements" .Choice}}
				{{template "Elements" .SequenceChoice}}
				{{template "Elements" .All}}
				{{template "Attributes" .Attributes}}
			{{end}}
		}
	{{end}}
{{end}}`
