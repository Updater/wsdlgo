package parser

const baseTmpl = `
// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package {{.Name}}

{{range .WSDL.Types.Schemas}}
	{{ $targetNamespace := .TargetNamespace }}

// Definition of simple types
type (
	{{range .SimpleType}}	{{template "SimpleType" .}}
{{end}})

// Constants associated with simple types defined above
const (
	{{range .SimpleType}}	{{template "Const" .}}
{{end}})

{{end}}`
