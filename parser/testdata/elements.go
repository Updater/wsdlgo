// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type echoStruct struct {
	MString   *string    `xml:"m_string"`
	MDatetime *time.Time `xml:"m_datetime"`
}

type echoele struct {
	Struct1 *echoStruct `xml:"struct1"`
	Struct2 *echoStruct `xml:"struct-2"`
}

type echoResponse struct {
	Struct1 []echoStruct `xml:"struct1"`
	Struct2 *echoStruct  `xml:"struct-2"`
}
