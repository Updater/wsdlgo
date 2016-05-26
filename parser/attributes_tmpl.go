package parser

const attributesTmpl = `
{{define "Attributes"}}
	{{range .}}	{{replaceReservedWords .Name | makeExported}} {{removeNS .Type | toGoPointerType}} ` + "`" + `xml:"{{.Name}},attr"` + "`" + `
	{{end}}
{{end}}`
