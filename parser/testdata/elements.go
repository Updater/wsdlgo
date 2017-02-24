// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT
package types

import (
	"encoding/xml"
	"time"
)

type decimalReqNil struct {
	*float64
}

// MarshalXML satisfies the XML Marshaler interface for type decimalReqNil.
func (t decimalReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.float64 == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t.float64, s)
}

type doubleReqNil struct {
	*float64
}

// MarshalXML satisfies the XML Marshaler interface for type doubleReqNil.
func (t doubleReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.float64 == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t.float64, s)
}

type echoResponse struct {
	Attrint    *int32       `xml:"attr-int,attr"`
	Attrstring *string      `xml:"attr_string,attr"`
	Struct1    []echoStruct `xml:"struct1"`
	Struct2    *echoStruct  `xml:"struct-2"`
}

type echoStruct struct {
	Mattr     *string    `xml:"m_attr,attr"`
	Mdatetime *time.Time `xml:"m_datetime"`
	Mstring   *string    `xml:"m_string"`
}

type echoele struct {
	Attrint    *int32      `xml:"attr-int,attr"`
	Attrstring *string     `xml:"attr_string,attr"`
	Struct1    *echoStruct `xml:"struct1"`
	Struct2    *echoStruct `xml:"struct-2"`
}
