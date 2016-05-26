package parser

const simpleTypesTmpl = `
{{define "SimpleTypeBase"}}	{{replaceReservedWords .Name | makeUnexported}} {{removeNS .Restriction.Base | toGoType}}{{end}}

{{define "SimpleType"}}
	{{if .Name}}
		type {{replaceReservedWords .Name | makeUnexported}} struct {
			{{if ne .XMLTag ""}}
				XMLName xml.Name ` + "`" + `xml:"{{.XMLTag}}"` + "`" + `
			{{end}}
			{{removeNS .Type | makeUnexported | toGoPointerType}}
		}
	{{end}}
{{end}}

{{define "SimpleContent"}}	Value {{removeNS .Extension.Base | toGoType}}{{end}}`
