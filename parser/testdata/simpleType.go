// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// Definition of simple types
type (
	myversion string
	iD        string
)

// Constants associated with simple types defined above
const (
	myversion16   myversion = "16"
	myversion18   myversion = "18"
	myversion19   myversion = "19"
	myversionHTML myversion = "html"
)

type ruby struct {
	Myversion *myversion `xml:"myversion"`
	Date      *time.Time `xml:"date"`
}
