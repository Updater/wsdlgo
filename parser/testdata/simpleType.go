// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

// Definition of types
type (
	contracttype string
	iD           string
	myversion    string
	uniqueID     string
)

// Constants associated with types defined above
const (
	contracttypeAll          contracttype = "All"
	contracttypeMonthToMonth contracttype = "MonthToMonth"
	contracttypeNoContract   contracttype = "NoContract"
	contracttypeOneYear      contracttype = "OneYear"
	contracttypeSixMonths    contracttype = "SixMonths"
	contracttypeThreeYears   contracttype = "ThreeYears"
	contracttypeTwoYears     contracttype = "TwoYears"
	contracttypeUnKnown      contracttype = "UnKnown"
	myversion16              myversion    = "1.6"
	myversion18              myversion    = "1.8"
	myversion19              myversion    = "1.9"
	myversionhtml            myversion    = "html"
	uniqueID1                uniqueID     = "1"
	uniqueIDHTML             uniqueID     = "Html"
)

type ruby struct {
	Date      *time.Time `xml:"date"`
	Myversion *myversion `xml:"myversion"`
}

type pingIDInMessage struct {
	XMLName xml.Name `xml:"http://www.test.com/test/ ping_id_in"`

	*iD
}
