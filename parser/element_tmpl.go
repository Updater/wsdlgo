package parser

const elementTmpl = `
// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package {{.Name}}

{{with .Element.Imports}}
	import (
		{{range . -}}
			"{{.}}"
		{{end}}
	)
{{end}}

{{with .Element.Types}}
	// Definition of types
	type (
		{{range . -}}
			{{.Name}} {{.UnderlyingType}}
		{{end}}
	)
{{end}}

{{with .Element.Consts}}
	// Constants associated with types defined above
	const  (
		{{range . -}}
			{{.Name}} {{.Type}} = "{{.Value}}"
		{{end}}
	)
{{end}}

{{with .Element.Structs}}
	{{range $i, $e := .}}
		type {{$i}} struct {
			{{range $e.Fields -}}	
				{{.Name}} {{.Type}} {{.Tag}}
			{{end}}
		}
		
		{{if $e.NillableRequiredType}}
			// MarshalXML satisfies the XML Marshaler interface for type {{$i}}.
			func (t {{$i}}) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
				{{range $e.Fields -}}	 
					if t.{{.Type | removePackage | convertPointerToValue}} == nil {
						return e.EncodeElement("", s)
					}
				
					{{if eq .Type "Time"}}
						tt := time.Time(*t.Time)
						if tt.IsZero() {
							return e.EncodeElement("", s)
						}
					{{end}}
					return e.EncodeElement(t, s)
				{{- end}}
			}
		{{end}}
	{{end}}
{{end}}

{{with .Element.Messages}}
	{{range $i, $e := .}}
		{{if $e.Type}}
			type {{$i}} struct {
				{{with $e.XMLField -}}
					{{.Name}} {{.Type}} {{.Tag}}
				{{end}}
				{{$e.Type}}
			}
		{{end}}
	{{end}}
{{end}}`
