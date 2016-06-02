// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

// Definition of types
type (
	iD        string
	myversion string
)

// Constants associated with types defined above
const (
	myversion16   myversion = "1.6"
	myversion18   myversion = "1.8"
	myversion19   myversion = "1.9"
	myversionHTML myversion = "html"
)

type ruby struct {
	Date      *time.Time `xml:"date"`
	Myversion *myversion `xml:"myversion"`
}

type ping_id_in struct {
	XMLName xml.Name `xml:"http://www.test.com/test/ ping_id_in"`

	*iD
}
