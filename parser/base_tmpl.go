package parser

const baseTmpl = `
// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package {{.Name}}

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

{{range .WSDL.Types.Schemas}}
	{{ $targetNamespace := .TargetNamespace }}

	{{with .SimpleType}}
		// Definition of simple types
		type (
			{{range .}}	{{template "SimpleType" .}}
		{{end}})

		// Constants associated with simple types defined above
		const (
			{{range .}}	{{template "Const" .}}
		{{end}})
	{{end}}

	{{range .ComplexTypes}}
		{{/* ComplexTypeGlobal */}}
		{{$name := replaceReservedWords .Name | makeUnexported}}
		type {{$name}} struct {
			{{if ne .ComplexContent.Extension.Base ""}}
				{{template "ComplexContent" .ComplexContent}}
			{{else if ne .SimpleContent.Extension.Base ""}}
				{{template "SimpleContent" .SimpleContent}}
			{{else}}
				{{template "Elements" .Sequence}}
				{{template "Elements" .Choice}}
				{{template "Elements" .SequenceChoice}}
				{{template "Elements" .All}}
			{{end}}
		}
	{{end}}

	{{range .Elements}}
		{{if not .Type}}
			{{/* ComplexTypeLocal */}}
			{{$name := .Name}}
			{{with .ComplexType}}
				type {{$name | replaceReservedWords}} struct {
					{{template "Elements" .Sequence}}
					{{template "Elements" .Choice}}
					{{template "Elements" .SequenceChoice}}
					{{template "Elements" .All}}
				}
			{{end}}
		{{end}}
	{{end}}

{{end}}`
