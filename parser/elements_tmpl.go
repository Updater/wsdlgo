package parser

const elementsTmpl = `
{{define "ElementsRegular"}}{{replaceReservedWords .Name | makeExported}} {{if eq .MaxOccurs "unbounded"}}[]{{.Type | toGoType}}{{else}}{{.Type | toGoPointerType}}{{end}} ` + "`" + `xml:"{{.Name}}"` + "`" + `{{end}}

{{define "ElementsReqNil"}}{{replaceReservedWords .Name | makeExported}} {{.NameReqNil | toGoType}} ` + "`" + `xml:"{{.Name}}"` + "`" + `{{end}}

{{define "ElementsInner"}}{{if .Name}}{{$type := replaceReservedWords .Name | makeUnexported}} {{$type | toGoPointerType}}Inner {{end}}{{end}}

{{define "Elements"}}
	{{range .}}
		{{if .NameReqNil}}{{template "ElementsReqNil" .}}{{else}}{{if .Type}}{{template "ElementsRegular" .}}{{else}}{{template "ElementsInner" .}}{{end}}{{end}}{{end}}{{end}}

{{define "NestedElements"}}
	{{range .}}
		{{if not .Type}}
			{{if .Name}}
				{{$typeName := replaceReservedWords .Name | makeUnexported}}
				type {{$typeName}}Inner struct {
					{{with .ComplexType}}
						{{template "Elements" .Sequence}}
						{{template "Elements" .Choice}}
						{{template "Elements" .SequenceChoice}}
						{{template "Elements" .All}}
						{{template "Attributes" .Attributes}}
					{{end}}
				}
			{{end}}
		{{end}}
	{{end}}
{{end}}

{{define "NillableRequiredTypes"}}
	{{range .}}
		{{if .NameReqNil}}
			type {{.NameReqNil}} struct {
				{{.Type | toGoPointerType}}
			}

			// MarshalXML satisfies the XML Marshaler interface for type {{.NameReqNil}}.
			func (t {{.NameReqNil}}) MarshalXML(e *xml.Encoder, s xml.StartElement) error { {{$type := .Type | removeNS | toGoType | removePackage}}
				if t.{{$type}} == nil {
					return e.EncodeElement("", s)
				}

				{{if eq $type "Time"}}
					tt := time.Time(*t.Time)
					if tt.IsZero() {
						return e.EncodeElement("", s)
					}
				{{end}}
				return e.EncodeElement(t, s)
			}
		{{end}}
	{{end}}
{{end}}`
