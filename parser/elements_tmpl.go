package parser

const elementsTmpl = `
{{define "Elements"}}
	{{range .}}
		{{replaceReservedWords .Name | makeExported}} {{if eq .MaxOccurs "unbounded"}}[]{{end}}{{.Type | toGoPointerType}} ` + "`" + `xml:"{{.Name}}"` + "`" + ` {{end}}
{{end}}`
