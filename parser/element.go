package parser

type element struct {
	Imports    mapofImports
	Types      mapofTypes
	Consts     mapofConsts
	Structs    mapofStructs
	Messages   mapofMessages
}

const elementTmpl = `
// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT
{{if ne .Name "" -}}
	package {{.Name}}
{{end}}

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
				{{range $e.Fields}}{{$t := .Type | removePackage | convertPointerToValue -}}
					if t.{{$t}} == nil {
						return e.EncodeElement("", s)
					}

					{{if eq .Type "Time"}}
						tt := time.Time(*t.Time)
						if tt.IsZero() {
							return e.EncodeElement("", s)
						}
					{{end}}
					return e.EncodeElement(t.{{$t}}, s)
				{{- end}}
			}
		{{end}}

		{{if $e.Marshaler}}
			func (t {{$i}}) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
				return e.EncodeElement(struct {
					{{range $e.EncoderFields -}}
						{{.Name}} {{.Type}} {{.Tag}}
					{{end}}
				}{
					{{range $e.EncoderFields -}}
						{{.Name}} : t.{{.Name}},
					{{end}}
				}, start)
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
			{{if $e.Marshaler}}
				func (t {{$i}}) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
					if t.{{$e.Type | convertPointerToValue}} == nil {
						return nil
					}

					x := t.{{$e.Type | convertPointerToValue}}.{{$e.Struct}}

					start.Name = xml.Name{Local: "{{$e.LocalName}}"}

					return e.EncodeElement(struct {
						{{range $i, $e := $e.NSFields -}}
							{{$i}} {{$e.Type}} {{$e.Tag}}
						{{end -}}
						{{range $e.Fields -}}
							{{.Name}} {{.Type}} {{.Tag}}
						{{end}}
					}{
						{{range $i, $e := $e.NSFields -}}
							{{$i}} : "{{$e.Name}}",
						{{end -}}
						{{range $e.Fields -}}
							{{.Name}} : x.{{.Name}},
						{{end}}
					}, start)
				}
			{{end}}
		{{end}}
	{{end}}
{{end}}`
