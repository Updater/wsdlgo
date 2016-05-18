// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type seqInner struct {
	Str3 *string `xml:"str3"`
}

type typeOut struct {
	Str1 *string `xml:"str1"`
	Str2 *string `xml:"str2"`
	*seqInner
}
