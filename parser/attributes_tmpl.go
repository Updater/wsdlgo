package parser

const attributesTmpl = `
{{define "Attributes"}}
	{{range .}}	{{replaceReservedWords .Name | makeExported}} {{toGoPointerType .Type}} ` + "`" + `xml:"{{.Name}},attr"` + "`" + `
	{{end}}
{{end}}`
