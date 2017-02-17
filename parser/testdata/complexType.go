// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT
package types

import (
	"encoding/xml"
	"time"
)

type arrayOfProducts struct {
	Product []string `xml:"Product"`
}

type arrayOfProductsprod struct {
	Product []string `xml:"Product"`
}

type arrayOfProductsprodReqNil struct {
	*arrayOfProductsprod
}

// MarshalXML satisfies the XML Marshaler interface for type arrayOfProductsprodReqNil.
func (t arrayOfProductsprodReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.arrayOfProductsprod == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t.arrayOfProductsprod, s)
}

type baseRequestMessage struct {
	TransactionID *int32 `xml:"TransactionID,attr"`
}

type dateTimeReqNil struct {
	*time.Time
}

// MarshalXML satisfies the XML Marshaler interface for type dateTimeReqNil.
func (t dateTimeReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.Time == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t.Time, s)
}

type getOfferAvailabilityRequest struct {
	*baseRequestMessage
	Source stringReqNil `xml:"Source"`
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

	return e.EncodeElement(t.int32, s)
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
	MON           *string      `xml:"MON"`
	Package       *packageType `xml:"Package"`
	SessionID     *string      `xml:"sessionId"`
	Status        *bool        `xml:"Status"`
	StatusCode    *string      `xml:"StatusCode"`
	StatusMessage []string     `xml:"StatusMessage"`
}

type orderStatusResponse struct {
	Orderid      *string   `xml:"Orderid"`
	Version      []version `xml:"Version"`
	VoiceOrderid *string   `xml:"VoiceOrderid"`
}

type packageType struct {
	Usoc *string `xml:"usoc,attr"`
}

type pingResponseType struct {
	TransactionID *string      `xml:"TransactionId"`
	Version       stringReqNil `xml:"Version"`
}

type result struct {
	*string
	Code *int32 `xml:"Code,attr"`
}

type serviceProductType struct {
	ServiceProducts  *arrayOfProducts          `xml:"ServiceProducts"`
	ServiceProducts1 arrayOfProductsprodReqNil `xml:"ServiceProducts.1"`
	Version          stringReqNil              `xml:"Version"`
}

type serviceProductTypeExt struct {
	*serviceProductType
	Nonboth *string `xml:"nonboth"`
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

	return e.EncodeElement(t.string, s)
}

type timeType struct {
	Hour   *byte `xml:"Hour"`
	Minute *byte `xml:"Minute"`
}

type version struct {
	OrderStatus   []orderStatus `xml:"OrderStatus"`
	VersionNumber *string       `xml:"VersionNumber"`
}
