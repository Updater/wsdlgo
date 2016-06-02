// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"time"
)

type echoResponse struct {
	AttrString *string      `xml:"attr_string,attr"`
	Attrint    *int32       `xml:"attr-int,attr"`
	Struct1    []echoStruct `xml:"struct1"`
	Struct2    *echoStruct  `xml:"struct-2"`
}

type echoStruct struct {
	MAttr     *string    `xml:"m_attr,attr"`
	MDatetime *time.Time `xml:"m_datetime"`
	MString   *string    `xml:"m_string"`
}

type echoele struct {
	AttrString *string     `xml:"attr_string,attr"`
	Attrint    *int32      `xml:"attr-int,attr"`
	Struct1    *echoStruct `xml:"struct1"`
	Struct2    *echoStruct `xml:"struct-2"`
}
