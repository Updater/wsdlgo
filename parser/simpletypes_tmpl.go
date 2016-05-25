package parser

const simpleTypesTmpl = `
{{define "SimpleType"}}	{{$type := replaceReservedWords .Name | makeUnexported}}{{$type}} {{removeNS .Restriction.Base | toGoType}}{{end}}

{{define "SimpleContent"}}	Value {{removeNS .Extension.Base | toGoType}}{{end}}`
