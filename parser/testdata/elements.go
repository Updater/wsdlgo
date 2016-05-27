// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"time"
)

type echoStruct struct {
	MString   *string    `xml:"m_string"`
	MDatetime *time.Time `xml:"m_datetime"`

	MAttr *string `xml:"m_attr,attr"`
}

type echoele struct {
	Struct1 *echoStruct `xml:"struct1"`
	Struct2 *echoStruct `xml:"struct-2"`

	AttrString *string `xml:"attr_string,attr"`
	Attrint    *int32  `xml:"attr-int,attr"`
}

type echoResponse struct {
	Struct1 []echoStruct `xml:"struct1"`
	Struct2 *echoStruct  `xml:"struct-2"`

	AttrString *string `xml:"attr_string,attr"`
	Attrint    *int32  `xml:"attr-int,attr"`
}
