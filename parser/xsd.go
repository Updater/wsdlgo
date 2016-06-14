package parser

import (
	"encoding/xml"
	"strings"
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

func (x xsdSchema) doMap(p interface{}) bool {
	switch u := p.(type) {
	case *element:
		var m []mapper
		for _, v := range x.Elements {
			m = append(m, v)
		}
		for _, v := range x.ComplexTypes {
			m = append(m, v)
		}
		for _, v := range x.SimpleType {
			m = append(m, v)
		}

		doMap(m, u.Types)
		doMap(m, u.Consts)
		doMap(m, u.Structs)
		doMap(m, u.Messages)
		return true
	}

	return false
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
// https://www.w3.org/TR/2012/REC-xmlschema11-1-20120405/#cElement_Declarations
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

func (x xsdElement) doMap(p interface{}) bool {
	switch u := p.(type) {
	case mapofStructs:
		if x.ComplexType != nil {
			if x.ComplexType.isEmpty() {
				break
			}

			doMap([]mapper{x.ComplexType}, u)

			if s := u.add(x.Name, sStruct{Name: x.Name}); s != "" {
				if x.ComplexType != nil {
					m := u[s]
					doMap([]mapper{x.ComplexType}, &m)
				}
				return true
			}
		}

		if x.Nillable && !(x.MinOccurs == "0") {
			u.add(x.Type, sStruct{
				Name:                 x.Type,
				NillableRequiredType: true,
			})
			return true
		}

	case mapofMessages:
		if x.ComplexType != nil || x.Name == "" || x.Type == "" {
			break
		}

		if b := u.add(x.Name, x.Type); b {
			return true
		}

	case *sStruct:
		if x.Type == "" && x.ComplexType != nil && !x.ComplexType.isEmpty() {
			x.Type = x.Name
		}

		s := sField{
			Name: x.Name,
			Type: x.Type,
		}

		if x.Nillable && !(x.MinOccurs == "0") {
			s.required = true
			s.nillable = true
		}

		if strings.ToLower(x.MaxOccurs) == "unbounded" {
			s.array = true
		}

		if x.Type == "" && x.ComplexType != nil && !x.ComplexType.isEmpty() {
			s.pointer = true
		}

		if b := u.Fields.add(x.Name, s); b {
			return true
		}
	}

	return false
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
	Attributes     []xsdAttribute    `xml:"attribute"`
}

func (x xsdComplexType) doMap(p interface{}) bool {
	var e []xsdElement
	e = append(e, x.Sequence...)
	e = append(e, x.Choice...)
	e = append(e, x.SequenceChoice...)
	e = append(e, x.All...)
	e = append(e, x.ComplexContent.Extension.Sequence...)

	var a []xsdAttribute
	a = append(a, x.Attributes...)

	switch u := p.(type) {
	case mapofStructs:
		var m []mapper
		for _, v := range e {
			m = append(m, v)
		}
		for _, v := range a {
			m = append(m, v)
		}
		m = append(m, x.ComplexContent.Extension)
		doMap(m, u)

		if s := u.add(x.Name, sStruct{Name: x.Name}); s != "" {
			ps := u[s]
			doMap(m, &ps)
		}

		return true

	case *sStruct:
		var m []mapper
		for _, v := range e {
			m = append(m, v)
		}
		for _, v := range a {
			m = append(m, v)
		}
		doMap(m, u)

		return true
	}

	return false
}

func (x xsdComplexType) hasElement() bool {
	return (len(x.Sequence) + len(x.Choice) + len(x.SequenceChoice) + len(x.All)) > 0
}

func (x xsdComplexType) hasAttribute() bool {
	return len(x.Attributes) > 0
}

func (x xsdComplexType) isEmpty() bool {
	return !(x.hasElement() || x.hasAttribute())
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
	XMLName    xml.Name       `xml:"extension"`
	Base       string         `xml:"base,attr"`
	Attributes []xsdAttribute `xml:"attribute"`
	Sequence   []xsdElement   `xml:"sequence>element"`
}

func (x xsdExtension) doMap(p interface{}) bool {
	switch u := p.(type) {
	case *sStruct:
		u.Fields.add(x.Base, sField{
			Type:    x.Base,
			pointer: true,
		})

		return true
	}
	return false
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

func (x xsdAttribute) doMap(p interface{}) bool {
	switch u := p.(type) {
	case *sStruct:
		u.Fields.add(x.Name, sField{
			Name: x.Name,
			Type: x.Type,
			attr: true,
		})

		return true
	}

	return false
}

// xsdSimpleType element defines a simple type and specifies the constraints
// and information about the values of attributes or text-only elements.
type xsdSimpleType struct {
	Name        string         `xml:"name,attr"`
	Restriction xsdRestriction `xml:"restriction"`
	List        xsdList        `xml:"list"`
}

func (x xsdSimpleType) doMap(p interface{}) bool {
	switch u := p.(type) {
	case mapofTypes:

		for _, st := range x.List.SimpleType {
			if x.Restriction.Base != "" {
				break
			}

			st.Name = x.Name
			doMap([]mapper{st}, u)
		}

		u.add(sType{
			Name:           x.Name,
			UnderlyingType: x.Restriction.Base,
		})

	case mapofConsts:
		var rv []xsdRestrictionValue
		rv = append(rv, x.Restriction.Enumeration...)
		for _, l := range x.List.SimpleType {
			rv = append(rv, l.Restriction.Enumeration...)
		}

		for _, e := range rv {
			u.add(sConst{
				Type:  x.Name,
				Value: e.Value,
			})
		}

		return true
	}
	return false
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

// xsdList defines a list type.
type xsdList struct {
	SimpleType []xsdSimpleType `xml:"simpleType"`
}
