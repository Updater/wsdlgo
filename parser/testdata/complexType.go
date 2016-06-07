// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

type arrayOfProducts struct {
	Product []string `xml:"Product"`
}

type dateTimeReqNil struct {
	*time.Time
}

// MarshalXML satisfies the XML Marshaler interface for type dateTimeReqNil.
func (t dateTimeReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.Time == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t, s)
}

type getOrderStatusResponse struct {
	GetOrderStatusResult *getOrderStatusResult `xml:"GetOrderStatusResult"`
}

type getOrderStatusResult struct {
	OrderStatusResponse *orderStatusResponse `xml:"OrderStatusResponse"`
	StatusInfo          *orderStatusInfo     `xml:"StatusInfo"`
}

type intReqNil struct {
	*int32
}

// MarshalXML satisfies the XML Marshaler interface for type intReqNil.
func (t intReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.int32 == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t, s)
}

type myelements struct {
	DateOfBirth1    dateTimeReqNil   `xml:"DateOfBirth1"`
	DateOfBirth2    dateTimeReqNil   `xml:"DateOfBirth2"`
	Minzero         *int32           `xml:"minzero"`
	Minzeronil      *string          `xml:"minzeronil"`
	Nilint          intReqNil        `xml:"nilint"`
	Nilstring       stringReqNil     `xml:"nilstring"`
	Nonboth         *string          `xml:"nonboth"`
	ServiceProducts *arrayOfProducts `xml:"ServiceProducts"`
}

type orderStatus struct {
	Actualtime     *string         `xml:"actualtime"`
	Estimatedtime  *string         `xml:"estimatedtime"`
	IsTechAssigned *string         `xml:"isTechAssigned"`
	ServiceStatus  []serviceStatus `xml:"ServiceStatus"`
	Source         *string         `xml:"source"`
	Stage          *string         `xml:"stage"`
}

type orderStatusInfo struct {
	MON           *string  `xml:"MON"`
	SessionId     *string  `xml:"sessionId"`
	Status        *bool    `xml:"Status"`
	StatusCode    *string  `xml:"StatusCode"`
	StatusMessage []string `xml:"StatusMessage"`
}

type orderStatusResponse struct {
	Orderid      *string   `xml:"Orderid"`
	Version      []version `xml:"Version"`
	VoiceOrderid *string   `xml:"VoiceOrderid"`
}

type pingResponseType struct {
	TransactionId *string      `xml:"TransactionId"`
	Version       stringReqNil `xml:"Version"`
}

type serviceProductType struct {
	ServiceProducts *arrayOfProducts `xml:"ServiceProducts"`
	Version         stringReqNil     `xml:"Version"`
}

type serviceProductTypeExt struct {
	Nonboth *string `xml:"nonboth"`
	*serviceProductType
}

type serviceStatus struct {
	OrderNumber       *string `xml:"OrderNumber"`
	Provisionstatus   *string `xml:"provisionstatus"`
	Status            *string `xml:"status"`
	Statusdescription *string `xml:"Statusdescription"`
}

type stringReqNil struct {
	*string
}

// MarshalXML satisfies the XML Marshaler interface for type stringReqNil.
func (t stringReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.string == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t, s)
}

type version struct {
	OrderStatus   []orderStatus `xml:"OrderStatus"`
	VersionNumber *string       `xml:"VersionNumber"`
}
