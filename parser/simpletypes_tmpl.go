package parser

const simpleTypesTmpl = `
{{define "SimpleType"}}	{{$type := replaceReservedWords .Name | makeUnexported}}{{$type}} {{toGoType .Restriction.Base}}{{end}}

{{define "SimpleContent"}}	Value {{toGoType .Extension.Base}}{{end}}`
