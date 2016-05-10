package parser

const baseTmpl = `
package {{.Name}}

{{range .WSDL.Types.Schemas}}
	{{ $targetNamespace := .TargetNamespace }}

	type (
	{{range .SimpleType}}	{{template "SimpleType" .}}
	{{end}})
{{end}}`
