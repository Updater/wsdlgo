// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"time"
)

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
