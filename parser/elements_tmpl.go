package parser

const elementsTmpl = `
{{define "Elements"}}
	{{range .}}{{replaceReservedWords .Name | makeExported}} {{if eq .MaxOccurs "unbounded"}}[]{{.Type | toGoType}}{{else}}{{.Type | toGoPointerType}}{{end}} ` + "`" + `xml:"{{.Name}}"` + "`" + ` 
	{{end}}
{{end}}`
