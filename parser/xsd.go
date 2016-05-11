package parser

import (
	"encoding/xml"
)

// xsdSchema represents an entire Schema structure.
type xsdSchema struct {
	XMLName            xml.Name         `xml:"schema"`
	Tns                string           `xml:"xmlns tns,attr"`
	Xs                 string           `xml:"xmlns xs,attr"`
	Version            string           `xml:"version,attr"`
	TargetNamespace    string           `xml:"targetNamespace,attr"`
	ElementFormDefault string           `xml:"elementFormDefault,attr"`
	Includes           []xsdInclude     `xml:"include"`
	Imports            []xsdImport      `xml:"import"`
	Elements           []xsdElement     `xml:"element"`
	ComplexTypes       []xsdComplexType `xml:"complexType"` //global
	SimpleType         []xsdSimpleType  `xml:"simpleType"`
}

// xsdInclude represents schema includes.
type xsdInclude struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
}

// xsdImport represents xsd imports within the main schema.
type xsdImport struct {
	XMLName        xml.Name `xml:"import"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Namespace      string   `xml:"namespace,attr"`
}

// xsdElement represents a Schema element.
type xsdElement struct {
	XMLName     xml.Name        `xml:"element"`
	Name        string          `xml:"name,attr"`
	Doc         string          `xml:"annotation>documentation"`
	Nillable    bool            `xml:"nillable,attr"`
	Type        string          `xml:"type,attr"`
	Ref         string          `xml:"ref,attr"`
	MinOccurs   string          `xml:"minOccurs,attr"`
	MaxOccurs   string          `xml:"maxOccurs,attr"`
	ComplexType *xsdComplexType `xml:"complexType"` //local
	SimpleType  *xsdSimpleType  `xml:"simpleType"`
	Groups      []xsdGroup      `xml:"group"`
}

// xsdComplexType represents a Schema complex type.
type xsdComplexType struct {
	XMLName        xml.Name          `xml:"complexType"`
	Abstract       bool              `xml:"abstract,attr"`
	Name           string            `xml:"name,attr"`
	Mixed          bool              `xml:"mixed,attr"`
	Sequence       []xsdElement      `xml:"sequence>element"`
	Choice         []xsdElement      `xml:"choice>element"`
	SequenceChoice []xsdElement      `xml:"sequence>choice>element"`
	All            []xsdElement      `xml:"all>element"`
	ComplexContent xsdComplexContent `xml:"complexContent"`
	SimpleContent  xsdSimpleContent  `xml:"simpleContent"`
	Attributes     []*xsdAttribute   `xml:"attribute"`
}

// xsdGroup element is used to define a group of elements to be used in complex type definitions.
type xsdGroup struct {
	Name     string       `xml:"name,attr"`
	Ref      string       `xml:"ref,attr"`
	Sequence []xsdElement `xml:"sequence>element"`
	Choice   []xsdElement `xml:"choice>element"`
	All      []xsdElement `xml:"all>element"`
}

// xsdComplexContent element defines extensions or restrictions on a complex
// type that contains mixed content or elements only.
type xsdComplexContent struct {
	XMLName   xml.Name     `xml:"complexContent"`
	Extension xsdExtension `xml:"extension"`
}

// xsdSimpleContent element contains extensions or restrictions on a text-only
// complex type or on a simple type as content and contains no elements.
type xsdSimpleContent struct {
	XMLName   xml.Name     `xml:"simpleContent"`
	Extension xsdExtension `xml:"extension"`
}

// xsdExtension element extends an existing simpleType or complexType element.
type xsdExtension struct {
	XMLName    xml.Name        `xml:"extension"`
	Base       string          `xml:"base,attr"`
	Attributes []*xsdAttribute `xml:"attribute"`
	Sequence   []xsdElement    `xml:"sequence>element"`
}

// xsdAttribute represent an element attribute. Simple elements cannot have
// attributes. If an element has attributes, it is considered to be of a
// complex type. But the attribute itself is always declared as a simple type.
type xsdAttribute struct {
	Name       string         `xml:"name,attr"`
	Doc        string         `xml:"annotation>documentation"`
	Type       string         `xml:"type,attr"`
	SimpleType *xsdSimpleType `xml:"simpleType"`
}

// xsdSimpleType element defines a simple type and specifies the constraints
// and information about the values of attributes or text-only elements.
type xsdSimpleType struct {
	Name        string         `xml:"name,attr"`
	Restriction xsdRestriction `xml:"restriction"`
}

// xsdRestriction defines restrictions on a simpleType, simpleContent, or complexContent definition.
type xsdRestriction struct {
	Base         string                `xml:"base,attr"`
	Enumeration  []xsdRestrictionValue `xml:"enumeration"`
	Pattern      xsdRestrictionValue   `xml:"pattern"`
	MinInclusive xsdRestrictionValue   `xml:"minInclusive"`
	MaxInclusive xsdRestrictionValue   `xml:"maxInclusive"`
	WhiteSpace   xsdRestrictionValue   `xml:"whitespace"`
	Length       xsdRestrictionValue   `xml:"length"`
	MinLength    xsdRestrictionValue   `xml:"minLength"`
	MaxLength    xsdRestrictionValue   `xml:"maxLength"`
}

// xsdRestrictionValue represents a restriction value.
type xsdRestrictionValue struct {
	Doc   string `xml:"annotation>documentation"`
	Value string `xml:"value,attr"`
}
