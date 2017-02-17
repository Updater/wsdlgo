// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT @ 2017-02-17 18:06:49.773106019 +0100 CET
package types

import (
	"encoding/xml"
)

type accessoryProductType struct {
	AccessoryProductCode *string `xml:"AccessoryProductCode"`
	Quantity             *int32  `xml:"Quantity"`
}

type additionalLine struct {
	LineNumber  *int32            `xml:"LineNumber"`
	LineOptions []lineOptionsType `xml:"LineOptions"`
	ProductCode *string           `xml:"ProductCode"`
}

type additionalLinesType struct {
	AdditionalLines []additionalLine `xml:"AdditionalLines"`
}

type appointment struct {
	AppointmentDate  *string           `xml:"AppointmentDate"`
	AppointmentTimes []appointmentTime `xml:"AppointmentTimes"`
}

type appointmentTime struct {
	AppointmentAfterTime  *string `xml:"AppointmentAfterTime"`
	AppointmentBeforeTime *string `xml:"AppointmentBeforeTime"`
}

type arrayOfFamiliesAndProducts struct {
	FamilyAndProducts []familyAndProducts `xml:"FamilyAndProducts"`
}

type arrayOfOrders struct {
	Order []order `xml:"Order"`
}

type arrayOfPotentialAppointments struct {
	Appointment []appointment `xml:"Appointment"`
}

type arrayOfProductsOrdered struct {
	ProductsOrdered []productsOrdered `xml:"ProductsOrdered"`
}

type arrayOfQwestOrders struct {
	QwestOrder []qwestOrder `xml:"QwestOrder"`
}

type billingData1411Type struct {
	AddressLine1  *string `xml:"AddressLine1"`
	AddressLine2  *string `xml:"AddressLine2"`
	AddressLine3  *string `xml:"AddressLine3"`
	BillFirstName *string `xml:"BillFirstName"`
	BillLastName  *string `xml:"BillLastName"`
	City          *string `xml:"City"`
	State         *string `xml:"State"`
	ZIP           *string `xml:"ZIP"`
}

type billingDataType struct {
	AddressLine1 *string `xml:"AddressLine1"`
	AddressLine2 *string `xml:"AddressLine2"`
	AddressLine3 *string `xml:"AddressLine3"`
	City         *string `xml:"City"`
	EmailAddress *string `xml:"EmailAddress"`
	NameLine1    *string `xml:"NameLine1"`
	NameLine2    *string `xml:"NameLine2"`
	State        *string `xml:"State"`
	ZIP          *string `xml:"ZIP"`
}

type comboAccountType struct {
	AccountID       *string `xml:"AccountId"`
	AccountProvider *string `xml:"AccountProvider"`
}

type contactInformation1411Type struct {
	ContactFirstName *string `xml:"ContactFirstName"`
	ContactLastName  *string `xml:"ContactLastName"`
	ContactNumber    *string `xml:"ContactNumber"`
	EmailAddress     *string `xml:"EmailAddress"`
}

type contactInformationType struct {
	ContactNumber *string `xml:"ContactNumber"`
	FirstName     *string `xml:"FirstName"`
	LastName      *string `xml:"LastName"`
}

type creditCardInfoType struct {
	CardExpirationDate   *string `xml:"CardExpirationDate"`
	CardHolderAddress1   *string `xml:"CardHolderAddress1"`
	CardHolderCity       *string `xml:"CardHolderCity"`
	CardHolderName       *string `xml:"CardHolderName"`
	CardHolderState      *string `xml:"CardHolderState"`
	CardHolderZip        *string `xml:"CardHolderZip"`
	CardNumber           *string `xml:"CardNumber"`
	CardType             *string `xml:"CardType"`
	CustomerRelationShip *string `xml:"CustomerRelationShip"`
}

type creditResponseType struct {
	AdvancePayment *float64 `xml:"AdvancePayment"`
	Deposit        *float64 `xml:"Deposit"`
	TotalPayment   *float64 `xml:"TotalPayment"`
}

type familyAndProducts struct {
	ArrayOfResponseProducts []responseProduct `xml:"ArrayOfResponseProducts"`
	ProductFamily           *string           `xml:"ProductFamily"`
}

type lineOptionsType struct {
	LineOptionName  *string `xml:"LineOptionName"`
	LineOptionValue *string `xml:"LineOptionValue"`
}

type order struct {
	OrderNumber *string `xml:"OrderNumber"`
	OrderType   *string `xml:"OrderType"`
}

type ping struct {
	PingRequest *pingRequestType `xml:"pingRequest"`
}

type pingRequestType struct {
	SalesCode *string `xml:"SalesCode"`
}

type pingResponse struct {
	PingResponse *pingResponseType `xml:"pingResponse"`
}

type pingResponseType struct {
	TransactionID *string      `xml:"TransactionId"`
	Version       stringReqNil `xml:"Version"`
}

type prepareOrder1411 struct {
	PrepareOrder1411Request *prepareOrder1411Request `xml:"prepareOrder1411Request"`
}

type prepareOrder1411Request struct {
	AccountVerification      *string                           `xml:"AccountVerification"`
	BillingInformation       *billingData1411Type              `xml:"BillingInformation"`
	ComboAccountInfo         *comboAccountType                 `xml:"ComboAccountInfo"`
	ContactInformation       *contactInformation1411Type       `xml:"ContactInformation"`
	CustomerIsMoving         *bool                             `xml:"CustomerIsMoving"`
	CustomerRequestedDueDate *string                           `xml:"CustomerRequestedDueDate"`
	PartnerOrderID           *string                           `xml:"PartnerOrderId"`
	PortTN                   *string                           `xml:"PortTN"`
	PortTNAddress            *string                           `xml:"PortTNAddress"`
	QualTransactionID        *string                           `xml:"QualTransactionId"`
	Remarks                  []remarks                         `xml:"Remarks"`
	RequestedAccessories     *requestedAccessoriesType         `xml:"RequestedAccessories"`
	RequestedProducts        *requestedProductsType            `xml:"RequestedProducts"`
	SalesCode                *string                           `xml:"SalesCode"`
	ServiceAddressStruct     *string                           `xml:"ServiceAddressStruct"`
	ShippingInformation      *shippingInformation1411Type      `xml:"ShippingInformation"`
	SupplementalCustomerInfo *supplementalCustomerInfo1411Type `xml:"SupplementalCustomerInfo"`
	WTN                      *string                           `xml:"WTN"`
}

type prepareOrder1411Response struct {
	AppointmentSchedulerUnavailable *bool                         `xml:"AppointmentSchedulerUnavailable"`
	AppointmentText                 *string                       `xml:"AppointmentText"`
	CreditResponse                  *creditResponseType           `xml:"CreditResponse"`
	PPOSessionID                    *string                       `xml:"PPOSessionId"`
	PotentialAppointments           *arrayOfPotentialAppointments `xml:"PotentialAppointments"`
	PrepareOrder1411Response        *prepareOrder1411Response     `xml:"prepareOrder1411Response"`
	TransactionID                   *string                       `xml:"TransactionId"`
}

type prepareOrderPort struct {
	PrepareOrderPortRequest *prepareOrderPortRequest `xml:"prepareOrderPortRequest"`
}

type prepareOrderPortRequest struct {
	AccountVerification       *string                        `xml:"AccountVerification"`
	BillingInformation        *billingDataType               `xml:"BillingInformation"`
	ContactInformation        *contactInformationType        `xml:"ContactInformation"`
	CustomerIsMoving          *bool                          `xml:"CustomerIsMoving"`
	CustomerRequestedDueDate  *string                        `xml:"CustomerRequestedDueDate"`
	MarketSegment             *string                        `xml:"MarketSegment"`
	PartnerOrderID            *string                        `xml:"PartnerOrderId"`
	PortTN                    *string                        `xml:"PortTN"`
	PortTNAddress             *string                        `xml:"PortTNAddress"`
	QualTransactionID         *string                        `xml:"QualTransactionId"`
	QwestAddressStruct        *string                        `xml:"QwestAddressStruct"`
	RequestedAccessories      *requestedAccessoriesType      `xml:"RequestedAccessories"`
	RequestedProducts         *requestedProductsType         `xml:"RequestedProducts"`
	SalesCode                 *string                        `xml:"SalesCode"`
	ShippingInformation       *shippingInformationType       `xml:"ShippingInformation"`
	StdIndustryClassification *stdIndustryClassificationType `xml:"StdIndustryClassification"`
	SupplementalCustomerInfo  *supplementalCustomerInfoType  `xml:"SupplementalCustomerInfo"`
	WTN                       *string                        `xml:"WTN"`
}

type prepareOrderPortResponse struct {
	PrepareOrderResponse *prepareOrderResponse `xml:"prepareOrderResponse"`
}

type prepareOrderResponse struct {
	AppointmentSchedulerUnavailable *bool                         `xml:"AppointmentSchedulerUnavailable"`
	AppointmentText                 *string                       `xml:"AppointmentText"`
	CreditResponse                  *creditResponseType           `xml:"CreditResponse"`
	PPOSessionID                    *string                       `xml:"PPOSessionId"`
	PotentialAppointments           *arrayOfPotentialAppointments `xml:"PotentialAppointments"`
	TransactionID                   *string                       `xml:"TransactionId"`
}

type previousAddressType struct {
	AddlRoutingInfo *string `xml:"AddlRoutingInfo"`
	City            *string `xml:"City"`
	RoutingInfo     *string `xml:"RoutingInfo"`
	State           *string `xml:"State"`
	StreetName      *string `xml:"StreetName"`
	StreetNumber    *string `xml:"StreetNumber"`
	Unit            *string `xml:"Unit"`
	Zip             *string `xml:"Zip"`
}

type previousAddressType1411 struct {
	City         *string `xml:"City"`
	State        *string `xml:"State"`
	StreetName   *string `xml:"StreetName"`
	StreetNumber *string `xml:"StreetNumber"`
	Unit         *string `xml:"Unit"`
	Zip          *string `xml:"Zip"`
}

type productsOrdered struct {
	FamiliesAndProducts *arrayOfFamiliesAndProducts `xml:"FamiliesAndProducts"`
	OrderNumber         *string                     `xml:"OrderNumber"`
	OrderType           *string                     `xml:"OrderType"`
}

type qwestOrder struct {
	OrderNumber *string `xml:"OrderNumber"`
	OrderType   *string `xml:"OrderType"`
}

type remarks struct {
	Remark     *string `xml:"Remark"`
	RemarkType *string `xml:"RemarkType"`
}

type requestedAccessoriesType struct {
	AccessoryProduct []accessoryProductType `xml:"AccessoryProduct"`
}

type requestedAppointment struct {
	AppointmentDate *string          `xml:"AppointmentDate"`
	AppointmentTime *appointmentTime `xml:"AppointmentTime"`
}

type requestedProductsType struct {
	Product []string `xml:"Product"`
}

type responseProduct struct {
	PartnerTrackingCode *string `xml:"PartnerTrackingCode"`
	ProductName         *string `xml:"ProductName"`
}

type sOHOBillingDataType struct {
	BillAddressLine1      *string `xml:"BillAddressLine1"`
	BillAddressLine2      *string `xml:"BillAddressLine2"`
	BillBusinessName      *string `xml:"BillBusinessName"`
	BillBusinessNameLine2 *string `xml:"BillBusinessNameLine2"`
	BillCity              *string `xml:"BillCity"`
	BillState             *string `xml:"BillState"`
	BillZIP               *string `xml:"BillZIP"`
}

type sOHOContactInformationType struct {
	AlternateContactCBR       *string `xml:"AlternateContactCBR"`
	AlternateContactFirstName *string `xml:"AlternateContactFirstName"`
	AlternateContactLastName  *string `xml:"AlternateContactLastName"`
	OrderPlacedByCBR          *string `xml:"OrderPlacedByCBR"`
	OrderPlacedByEmail        *string `xml:"OrderPlacedByEmail"`
	OrderPlacedByFirstName    *string `xml:"OrderPlacedByFirstName"`
	OrderPlacedByLastName     *string `xml:"OrderPlacedByLastName"`
	ResponsiblePartyCBR       *string `xml:"ResponsiblePartyCBR"`
	ResponsiblePartyFirstName *string `xml:"ResponsiblePartyFirstName"`
	ResponsiblePartyLastName  *string `xml:"ResponsiblePartyLastName"`
}

type sOHOCustomerInfoType struct {
	BusinessType *string `xml:"BusinessType"`
	SSN          *string `xml:"SSN"`
	TaxID        *string `xml:"TaxId"`
}

type sOHOListingInformationType struct {
	OmitAddress              *bool   `xml:"OmitAddress"`
	SOHOListingCity          *string `xml:"SOHOListingCity"`
	SOHOListingName          *string `xml:"SOHOListingName"`
	SOHOListingState         *string `xml:"SOHOListingState"`
	SOHOListingStreetAddress *string `xml:"SOHOListingStreetAddress"`
}

type sOHOPrepareOrder1308Request struct {
	AccountVerification       *string                     `xml:"AccountVerification"`
	AdditionalLines           *additionalLinesType        `xml:"AdditionalLines"`
	AddressStruct             *string                     `xml:"AddressStruct"`
	CustomerIsMoving          *bool                       `xml:"CustomerIsMoving"`
	CustomerRequestedDueDate  *string                     `xml:"CustomerRequestedDueDate"`
	HasSecuritySystem         *bool                       `xml:"HasSecuritySystem"`
	IsKeySystem               *bool                       `xml:"IsKeySystem"`
	MarketSegment             *string                     `xml:"MarketSegment"`
	NumberOfComputers         *int32                      `xml:"NumberOfComputers"`
	PartnerOrderID            *string                     `xml:"PartnerOrderId"`
	PortTN                    *string                     `xml:"PortTN"`
	PortTNAddress             *string                     `xml:"PortTNAddress"`
	QualTransactionID         *string                     `xml:"QualTransactionId"`
	RequestedAccessories      *requestedAccessoriesType   `xml:"RequestedAccessories"`
	RequestedProducts         *requestedProductsType      `xml:"RequestedProducts"`
	SOHOBillingInformation    *sOHOBillingDataType        `xml:"SOHOBillingInformation"`
	SOHOContactInformation    *sOHOContactInformationType `xml:"SOHOContactInformation"`
	SOHOCustomerInfo          *sOHOCustomerInfoType       `xml:"SOHOCustomerInfo"`
	SOHOListingInformation    *sOHOListingInformationType `xml:"SOHOListingInformation"`
	SOHOShippingInformation   *shippingInformationType    `xml:"SOHOShippingInformation"`
	SalesCode                 *string                     `xml:"SalesCode"`
	WTN                       *string                     `xml:"WTN"`
	YellowPagesClassification *yPClassificationType       `xml:"YellowPagesClassification"`
}

type sOHOPrepareOrder1308Response struct {
	AppointmentSchedulerUnavailable *bool                         `xml:"AppointmentSchedulerUnavailable"`
	AppointmentText                 *string                       `xml:"AppointmentText"`
	CreditResponse                  *creditResponseType           `xml:"CreditResponse"`
	PPOSessionID                    *string                       `xml:"PPOSessionId"`
	PotentialAppointments           *arrayOfPotentialAppointments `xml:"PotentialAppointments"`
	TransactionID                   *string                       `xml:"TransactionId"`
}

type shippingInformation1411Type struct {
	ShippingAddress   *string `xml:"ShippingAddress"`
	ShippingAddress2  *string `xml:"ShippingAddress2"`
	ShippingCity      *string `xml:"ShippingCity"`
	ShippingFirstName *string `xml:"ShippingFirstName"`
	ShippingLastName  *string `xml:"ShippingLastName"`
	ShippingState     *string `xml:"ShippingState"`
	ShippingZipCode   *string `xml:"ShippingZipCode"`
}

type shippingInformationType struct {
	ShippingAddress  *string `xml:"ShippingAddress"`
	ShippingAddress2 *string `xml:"ShippingAddress2"`
	ShippingCity     *string `xml:"ShippingCity"`
	ShippingName     *string `xml:"ShippingName"`
	ShippingState    *string `xml:"ShippingState"`
	ShippingZipCode  *string `xml:"ShippingZipCode"`
}

type sohoPrepareOrder1308 struct {
	SohoPrepareOrder1308Request *sOHOPrepareOrder1308Request `xml:"sohoPrepareOrder1308Request"`
}

type sohoPrepareOrder1308Response struct {
	SohoPrepareOrder1308Response *sOHOPrepareOrder1308Response `xml:"sohoPrepareOrder1308Response"`
}

type stdIndustryClassificationType struct {
	SICCode *string `xml:"SICCode"`
	SICDesc *string `xml:"SICDesc"`
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

type submitOrder1411 struct {
	SubmitOrder1411Request *submitOrder1411Request `xml:"submitOrder1411Request"`
}

type submitOrder1411Request struct {
	ContactCustomer      *bool                 `xml:"ContactCustomer"`
	CreditCardInfo       *creditCardInfoType   `xml:"CreditCardInfo"`
	PPOSessionID         *string               `xml:"PPOSessionId"`
	PrepareTransactionID *string               `xml:"PrepareTransactionId"`
	RequestedAppointment *requestedAppointment `xml:"RequestedAppointment"`
	SalesCode            *string               `xml:"SalesCode"`
}

type submitOrder1411Response struct {
	AccountNumber           *string                  `xml:"AccountNumber"`
	AppointmentTime         *appointmentTime         `xml:"AppointmentTime"`
	AuthorizationCode       *string                  `xml:"AuthorizationCode"`
	CustomerResponseMessage *string                  `xml:"CustomerResponseMessage"`
	DueDate                 *string                  `xml:"DueDate"`
	OrderList               *arrayOfOrders           `xml:"OrderList"`
	OrderOutcome            *string                  `xml:"OrderOutcome"`
	PartnerOrderID          *string                  `xml:"PartnerOrderId"`
	PartnerResponseMessage  *string                  `xml:"PartnerResponseMessage"`
	ProductsOrderedList     *arrayOfProductsOrdered  `xml:"ProductsOrderedList"`
	ResponseCode            *int32                   `xml:"ResponseCode"`
	SubmitOrder1411Response *submitOrder1411Response `xml:"submitOrder1411Response"`
	TelephoneNumber         *string                  `xml:"TelephoneNumber"`
	TransactionID           *string                  `xml:"TransactionId"`
}

type submitOrderRequest struct {
	ContactCustomer      *bool                 `xml:"ContactCustomer"`
	CreditCardInfo       *creditCardInfoType   `xml:"CreditCardInfo"`
	PPOSessionID         *string               `xml:"PPOSessionId"`
	RequestedAppointment *requestedAppointment `xml:"RequestedAppointment"`
	SalesCode            *string               `xml:"SalesCode"`
	SubmitOrderRequest   *submitOrderRequest   `xml:"submitOrderRequest"`
}

type submitOrderRequestResponse struct {
	SubmitOrderResponse *submitOrderResponse `xml:"submitOrderResponse"`
}

type submitOrderResponse struct {
	AppointmentTime     *appointmentTime        `xml:"AppointmentTime"`
	AuthorizationCode   *string                 `xml:"AuthorizationCode"`
	DueDate             *string                 `xml:"DueDate"`
	OrderOutcome        *string                 `xml:"OrderOutcome"`
	PartnerOrderID      *string                 `xml:"PartnerOrderId"`
	ProductsOrderedList *arrayOfProductsOrdered `xml:"ProductsOrderedList"`
	QwestOrderList      *arrayOfQwestOrders     `xml:"QwestOrderList"`
	ResponseCode        *int32                  `xml:"ResponseCode"`
	ResponseMessage     *string                 `xml:"ResponseMessage"`
	TN                  *string                 `xml:"TN"`
	TransactionID       *string                 `xml:"TransactionId"`
}

type supplementalCustomerInfo1411Type struct {
	DOB             *string                  `xml:"DOB"`
	PersonalID      *string                  `xml:"PersonalId"`
	PersonalIDState *string                  `xml:"PersonalIdState"`
	PreviousAddress *previousAddressType1411 `xml:"PreviousAddress"`
	SSN             *string                  `xml:"SSN"`
}

type supplementalCustomerInfoType struct {
	DOB             *string              `xml:"DOB"`
	PersonalID      *string              `xml:"PersonalId"`
	PersonalIDState *string              `xml:"PersonalIdState"`
	PreviousAddress *previousAddressType `xml:"PreviousAddress"`
	SSN             *string              `xml:"SSN"`
}

type yPClassificationType struct {
	YPCode        *string `xml:"YPCode"`
	YPDescription *string `xml:"YPDescription"`
}

type pingMessage struct {
	XMLName xml.Name `xml:" ping"`

	*ping
}

func (t pingMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.ping == nil {
		return nil
	}

	x := t.ping.PingRequest

	start.Name = xml.Name{Local: "con:ping"}

	return e.EncodeElement(struct {
		XMLcon    string  `xml:"xmlns:con,attr"`
		XMLppo    string  `xml:"xmlns:ppo,attr"`
		SalesCode *string `xml:"pingRequest>ppo:SalesCode"`
	}{
		XMLcon:    "http://localhost/ContentOrderService",
		XMLppo:    "http://www.qwest.com/ppo/",
		SalesCode: x.SalesCode,
	}, start)
}

type pingResponseMessage struct {
	XMLName xml.Name `xml:" pingResponse"`

	*pingResponse
}

func (t pingResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.pingResponse == nil {
		return nil
	}

	x := t.pingResponse.PingResponse

	start.Name = xml.Name{Local: "con:pingResponse"}

	return e.EncodeElement(struct {
		XMLcon        string       `xml:"xmlns:con,attr"`
		XMLppo        string       `xml:"xmlns:ppo,attr"`
		TransactionID *string      `xml:"pingResponse>ppo:TransactionID"`
		Version       stringReqNil `xml:"pingResponse>ppo:Version"`
	}{
		XMLcon:        "http://localhost/ContentOrderService",
		XMLppo:        "http://www.qwest.com/ppo/",
		TransactionID: x.TransactionID,
		Version:       x.Version,
	}, start)
}

type prepareOrder1411Message struct {
	XMLName xml.Name `xml:" prepareOrder1411"`

	*prepareOrder1411
}

func (t prepareOrder1411Message) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.prepareOrder1411 == nil {
		return nil
	}

	x := t.prepareOrder1411.PrepareOrder1411Request

	start.Name = xml.Name{Local: "con:prepareOrder1411"}

	return e.EncodeElement(struct {
		XMLcon                   string                            `xml:"xmlns:con,attr"`
		XMLppo                   string                            `xml:"xmlns:ppo,attr"`
		AccountVerification      *string                           `xml:"prepareOrder1411Request>ppo:AccountVerification"`
		BillingInformation       *billingData1411Type              `xml:"prepareOrder1411Request>ppo:BillingInformation"`
		ComboAccountInfo         *comboAccountType                 `xml:"prepareOrder1411Request>ppo:ComboAccountInfo"`
		ContactInformation       *contactInformation1411Type       `xml:"prepareOrder1411Request>ppo:ContactInformation"`
		CustomerIsMoving         *bool                             `xml:"prepareOrder1411Request>ppo:CustomerIsMoving"`
		CustomerRequestedDueDate *string                           `xml:"prepareOrder1411Request>ppo:CustomerRequestedDueDate"`
		PartnerOrderID           *string                           `xml:"prepareOrder1411Request>ppo:PartnerOrderID"`
		PortTN                   *string                           `xml:"prepareOrder1411Request>ppo:PortTN"`
		PortTNAddress            *string                           `xml:"prepareOrder1411Request>ppo:PortTNAddress"`
		QualTransactionID        *string                           `xml:"prepareOrder1411Request>ppo:QualTransactionID"`
		Remarks                  []remarks                         `xml:"prepareOrder1411Request>ppo:Remarks"`
		RequestedAccessories     *requestedAccessoriesType         `xml:"prepareOrder1411Request>ppo:RequestedAccessories"`
		RequestedProducts        *requestedProductsType            `xml:"prepareOrder1411Request>ppo:RequestedProducts"`
		SalesCode                *string                           `xml:"prepareOrder1411Request>ppo:SalesCode"`
		ServiceAddressStruct     *string                           `xml:"prepareOrder1411Request>ppo:ServiceAddressStruct"`
		ShippingInformation      *shippingInformation1411Type      `xml:"prepareOrder1411Request>ppo:ShippingInformation"`
		SupplementalCustomerInfo *supplementalCustomerInfo1411Type `xml:"prepareOrder1411Request>ppo:SupplementalCustomerInfo"`
		WTN                      *string                           `xml:"prepareOrder1411Request>ppo:WTN"`
	}{
		XMLcon:                   "http://localhost/ContentOrderService",
		XMLppo:                   "http://www.qwest.com/ppo/",
		AccountVerification:      x.AccountVerification,
		BillingInformation:       x.BillingInformation,
		ComboAccountInfo:         x.ComboAccountInfo,
		ContactInformation:       x.ContactInformation,
		CustomerIsMoving:         x.CustomerIsMoving,
		CustomerRequestedDueDate: x.CustomerRequestedDueDate,
		PartnerOrderID:           x.PartnerOrderID,
		PortTN:                   x.PortTN,
		PortTNAddress:            x.PortTNAddress,
		QualTransactionID:        x.QualTransactionID,
		Remarks:                  x.Remarks,
		RequestedAccessories:     x.RequestedAccessories,
		RequestedProducts:        x.RequestedProducts,
		SalesCode:                x.SalesCode,
		ServiceAddressStruct:     x.ServiceAddressStruct,
		ShippingInformation:      x.ShippingInformation,
		SupplementalCustomerInfo: x.SupplementalCustomerInfo,
		WTN: x.WTN,
	}, start)
}

type prepareOrder1411ResponseMessage struct {
	XMLName xml.Name `xml:" prepareOrder1411Response"`

	*prepareOrder1411Response
}

func (t prepareOrder1411ResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.prepareOrder1411Response == nil {
		return nil
	}

	x := t.prepareOrder1411Response.PrepareOrder1411Response

	start.Name = xml.Name{Local: "con:prepareOrder1411Response"}

	return e.EncodeElement(struct {
		XMLcon                          string                        `xml:"xmlns:con,attr"`
		XMLppo                          string                        `xml:"xmlns:ppo,attr"`
		AppointmentSchedulerUnavailable *bool                         `xml:"prepareOrder1411Response>ppo:AppointmentSchedulerUnavailable"`
		AppointmentText                 *string                       `xml:"prepareOrder1411Response>ppo:AppointmentText"`
		CreditResponse                  *creditResponseType           `xml:"prepareOrder1411Response>ppo:CreditResponse"`
		PPOSessionID                    *string                       `xml:"prepareOrder1411Response>ppo:PPOSessionID"`
		PotentialAppointments           *arrayOfPotentialAppointments `xml:"prepareOrder1411Response>ppo:PotentialAppointments"`
		PrepareOrder1411Response        *prepareOrder1411Response     `xml:"prepareOrder1411Response>ppo:PrepareOrder1411Response"`
		TransactionID                   *string                       `xml:"prepareOrder1411Response>ppo:TransactionID"`
	}{
		XMLcon: "http://localhost/ContentOrderService",
		XMLppo: "http://www.qwest.com/ppo/",
		AppointmentSchedulerUnavailable: x.AppointmentSchedulerUnavailable,
		AppointmentText:                 x.AppointmentText,
		CreditResponse:                  x.CreditResponse,
		PPOSessionID:                    x.PPOSessionID,
		PotentialAppointments:           x.PotentialAppointments,
		PrepareOrder1411Response:        x.PrepareOrder1411Response,
		TransactionID:                   x.TransactionID,
	}, start)
}

type prepareOrderPortMessage struct {
	XMLName xml.Name `xml:" prepareOrderPort"`

	*prepareOrderPort
}

func (t prepareOrderPortMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.prepareOrderPort == nil {
		return nil
	}

	x := t.prepareOrderPort.PrepareOrderPortRequest

	start.Name = xml.Name{Local: "con:prepareOrderPort"}

	return e.EncodeElement(struct {
		XMLcon                    string                         `xml:"xmlns:con,attr"`
		XMLppo                    string                         `xml:"xmlns:ppo,attr"`
		AccountVerification       *string                        `xml:"prepareOrderPortRequest>ppo:AccountVerification"`
		BillingInformation        *billingDataType               `xml:"prepareOrderPortRequest>ppo:BillingInformation"`
		ContactInformation        *contactInformationType        `xml:"prepareOrderPortRequest>ppo:ContactInformation"`
		CustomerIsMoving          *bool                          `xml:"prepareOrderPortRequest>ppo:CustomerIsMoving"`
		CustomerRequestedDueDate  *string                        `xml:"prepareOrderPortRequest>ppo:CustomerRequestedDueDate"`
		MarketSegment             *string                        `xml:"prepareOrderPortRequest>ppo:MarketSegment"`
		PartnerOrderID            *string                        `xml:"prepareOrderPortRequest>ppo:PartnerOrderID"`
		PortTN                    *string                        `xml:"prepareOrderPortRequest>ppo:PortTN"`
		PortTNAddress             *string                        `xml:"prepareOrderPortRequest>ppo:PortTNAddress"`
		QualTransactionID         *string                        `xml:"prepareOrderPortRequest>ppo:QualTransactionID"`
		QwestAddressStruct        *string                        `xml:"prepareOrderPortRequest>ppo:QwestAddressStruct"`
		RequestedAccessories      *requestedAccessoriesType      `xml:"prepareOrderPortRequest>ppo:RequestedAccessories"`
		RequestedProducts         *requestedProductsType         `xml:"prepareOrderPortRequest>ppo:RequestedProducts"`
		SalesCode                 *string                        `xml:"prepareOrderPortRequest>ppo:SalesCode"`
		ShippingInformation       *shippingInformationType       `xml:"prepareOrderPortRequest>ppo:ShippingInformation"`
		StdIndustryClassification *stdIndustryClassificationType `xml:"prepareOrderPortRequest>ppo:StdIndustryClassification"`
		SupplementalCustomerInfo  *supplementalCustomerInfoType  `xml:"prepareOrderPortRequest>ppo:SupplementalCustomerInfo"`
		WTN                       *string                        `xml:"prepareOrderPortRequest>ppo:WTN"`
	}{
		XMLcon:                    "http://localhost/ContentOrderService",
		XMLppo:                    "http://www.qwest.com/ppo/",
		AccountVerification:       x.AccountVerification,
		BillingInformation:        x.BillingInformation,
		ContactInformation:        x.ContactInformation,
		CustomerIsMoving:          x.CustomerIsMoving,
		CustomerRequestedDueDate:  x.CustomerRequestedDueDate,
		MarketSegment:             x.MarketSegment,
		PartnerOrderID:            x.PartnerOrderID,
		PortTN:                    x.PortTN,
		PortTNAddress:             x.PortTNAddress,
		QualTransactionID:         x.QualTransactionID,
		QwestAddressStruct:        x.QwestAddressStruct,
		RequestedAccessories:      x.RequestedAccessories,
		RequestedProducts:         x.RequestedProducts,
		SalesCode:                 x.SalesCode,
		ShippingInformation:       x.ShippingInformation,
		StdIndustryClassification: x.StdIndustryClassification,
		SupplementalCustomerInfo:  x.SupplementalCustomerInfo,
		WTN: x.WTN,
	}, start)
}

type prepareOrderPortResponseMessage struct {
	XMLName xml.Name `xml:" prepareOrderPortResponse"`

	*prepareOrderPortResponse
}

func (t prepareOrderPortResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.prepareOrderPortResponse == nil {
		return nil
	}

	x := t.prepareOrderPortResponse.PrepareOrderResponse

	start.Name = xml.Name{Local: "con:prepareOrderPortResponse"}

	return e.EncodeElement(struct {
		XMLcon                          string                        `xml:"xmlns:con,attr"`
		XMLppo                          string                        `xml:"xmlns:ppo,attr"`
		AppointmentSchedulerUnavailable *bool                         `xml:"prepareOrderResponse>ppo:AppointmentSchedulerUnavailable"`
		AppointmentText                 *string                       `xml:"prepareOrderResponse>ppo:AppointmentText"`
		CreditResponse                  *creditResponseType           `xml:"prepareOrderResponse>ppo:CreditResponse"`
		PPOSessionID                    *string                       `xml:"prepareOrderResponse>ppo:PPOSessionID"`
		PotentialAppointments           *arrayOfPotentialAppointments `xml:"prepareOrderResponse>ppo:PotentialAppointments"`
		TransactionID                   *string                       `xml:"prepareOrderResponse>ppo:TransactionID"`
	}{
		XMLcon: "http://localhost/ContentOrderService",
		XMLppo: "http://www.qwest.com/ppo/",
		AppointmentSchedulerUnavailable: x.AppointmentSchedulerUnavailable,
		AppointmentText:                 x.AppointmentText,
		CreditResponse:                  x.CreditResponse,
		PPOSessionID:                    x.PPOSessionID,
		PotentialAppointments:           x.PotentialAppointments,
		TransactionID:                   x.TransactionID,
	}, start)
}

type sohoPrepareOrder1308Message struct {
	XMLName xml.Name `xml:" sohoPrepareOrder1308"`

	*sohoPrepareOrder1308
}

func (t sohoPrepareOrder1308Message) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.sohoPrepareOrder1308 == nil {
		return nil
	}

	x := t.sohoPrepareOrder1308.SohoPrepareOrder1308Request

	start.Name = xml.Name{Local: "con:sohoPrepareOrder1308"}

	return e.EncodeElement(struct {
		XMLcon                    string                      `xml:"xmlns:con,attr"`
		XMLppo                    string                      `xml:"xmlns:ppo,attr"`
		AccountVerification       *string                     `xml:"sohoPrepareOrder1308Request>ppo:AccountVerification"`
		AdditionalLines           *additionalLinesType        `xml:"sohoPrepareOrder1308Request>ppo:AdditionalLines"`
		AddressStruct             *string                     `xml:"sohoPrepareOrder1308Request>ppo:AddressStruct"`
		CustomerIsMoving          *bool                       `xml:"sohoPrepareOrder1308Request>ppo:CustomerIsMoving"`
		CustomerRequestedDueDate  *string                     `xml:"sohoPrepareOrder1308Request>ppo:CustomerRequestedDueDate"`
		HasSecuritySystem         *bool                       `xml:"sohoPrepareOrder1308Request>ppo:HasSecuritySystem"`
		IsKeySystem               *bool                       `xml:"sohoPrepareOrder1308Request>ppo:IsKeySystem"`
		MarketSegment             *string                     `xml:"sohoPrepareOrder1308Request>ppo:MarketSegment"`
		NumberOfComputers         *int32                      `xml:"sohoPrepareOrder1308Request>ppo:NumberOfComputers"`
		PartnerOrderID            *string                     `xml:"sohoPrepareOrder1308Request>ppo:PartnerOrderID"`
		PortTN                    *string                     `xml:"sohoPrepareOrder1308Request>ppo:PortTN"`
		PortTNAddress             *string                     `xml:"sohoPrepareOrder1308Request>ppo:PortTNAddress"`
		QualTransactionID         *string                     `xml:"sohoPrepareOrder1308Request>ppo:QualTransactionID"`
		RequestedAccessories      *requestedAccessoriesType   `xml:"sohoPrepareOrder1308Request>ppo:RequestedAccessories"`
		RequestedProducts         *requestedProductsType      `xml:"sohoPrepareOrder1308Request>ppo:RequestedProducts"`
		SOHOBillingInformation    *sOHOBillingDataType        `xml:"sohoPrepareOrder1308Request>ppo:SOHOBillingInformation"`
		SOHOContactInformation    *sOHOContactInformationType `xml:"sohoPrepareOrder1308Request>ppo:SOHOContactInformation"`
		SOHOCustomerInfo          *sOHOCustomerInfoType       `xml:"sohoPrepareOrder1308Request>ppo:SOHOCustomerInfo"`
		SOHOListingInformation    *sOHOListingInformationType `xml:"sohoPrepareOrder1308Request>ppo:SOHOListingInformation"`
		SOHOShippingInformation   *shippingInformationType    `xml:"sohoPrepareOrder1308Request>ppo:SOHOShippingInformation"`
		SalesCode                 *string                     `xml:"sohoPrepareOrder1308Request>ppo:SalesCode"`
		WTN                       *string                     `xml:"sohoPrepareOrder1308Request>ppo:WTN"`
		YellowPagesClassification *yPClassificationType       `xml:"sohoPrepareOrder1308Request>ppo:YellowPagesClassification"`
	}{
		XMLcon:                   "http://localhost/ContentOrderService",
		XMLppo:                   "http://www.qwest.com/ppo/",
		AccountVerification:      x.AccountVerification,
		AdditionalLines:          x.AdditionalLines,
		AddressStruct:            x.AddressStruct,
		CustomerIsMoving:         x.CustomerIsMoving,
		CustomerRequestedDueDate: x.CustomerRequestedDueDate,
		HasSecuritySystem:        x.HasSecuritySystem,
		IsKeySystem:              x.IsKeySystem,
		MarketSegment:            x.MarketSegment,
		NumberOfComputers:        x.NumberOfComputers,
		PartnerOrderID:           x.PartnerOrderID,
		PortTN:                   x.PortTN,
		PortTNAddress:            x.PortTNAddress,
		QualTransactionID:        x.QualTransactionID,
		RequestedAccessories:     x.RequestedAccessories,
		RequestedProducts:        x.RequestedProducts,
		SOHOBillingInformation:   x.SOHOBillingInformation,
		SOHOContactInformation:   x.SOHOContactInformation,
		SOHOCustomerInfo:         x.SOHOCustomerInfo,
		SOHOListingInformation:   x.SOHOListingInformation,
		SOHOShippingInformation:  x.SOHOShippingInformation,
		SalesCode:                x.SalesCode,
		WTN:                      x.WTN,
		YellowPagesClassification: x.YellowPagesClassification,
	}, start)
}

type sohoPrepareOrder1308ResponseMessage struct {
	XMLName xml.Name `xml:" sohoPrepareOrder1308Response"`

	*sohoPrepareOrder1308Response
}

func (t sohoPrepareOrder1308ResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.sohoPrepareOrder1308Response == nil {
		return nil
	}

	x := t.sohoPrepareOrder1308Response.SohoPrepareOrder1308Response

	start.Name = xml.Name{Local: "con:sohoPrepareOrder1308Response"}

	return e.EncodeElement(struct {
		XMLcon                          string                        `xml:"xmlns:con,attr"`
		XMLppo                          string                        `xml:"xmlns:ppo,attr"`
		AppointmentSchedulerUnavailable *bool                         `xml:"sohoPrepareOrder1308Response>ppo:AppointmentSchedulerUnavailable"`
		AppointmentText                 *string                       `xml:"sohoPrepareOrder1308Response>ppo:AppointmentText"`
		CreditResponse                  *creditResponseType           `xml:"sohoPrepareOrder1308Response>ppo:CreditResponse"`
		PPOSessionID                    *string                       `xml:"sohoPrepareOrder1308Response>ppo:PPOSessionID"`
		PotentialAppointments           *arrayOfPotentialAppointments `xml:"sohoPrepareOrder1308Response>ppo:PotentialAppointments"`
		TransactionID                   *string                       `xml:"sohoPrepareOrder1308Response>ppo:TransactionID"`
	}{
		XMLcon: "http://localhost/ContentOrderService",
		XMLppo: "http://www.qwest.com/ppo/",
		AppointmentSchedulerUnavailable: x.AppointmentSchedulerUnavailable,
		AppointmentText:                 x.AppointmentText,
		CreditResponse:                  x.CreditResponse,
		PPOSessionID:                    x.PPOSessionID,
		PotentialAppointments:           x.PotentialAppointments,
		TransactionID:                   x.TransactionID,
	}, start)
}

type submitOrder1411Message struct {
	XMLName xml.Name `xml:" submitOrder1411"`

	*submitOrder1411
}

func (t submitOrder1411Message) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.submitOrder1411 == nil {
		return nil
	}

	x := t.submitOrder1411.SubmitOrder1411Request

	start.Name = xml.Name{Local: "con:submitOrder1411"}

	return e.EncodeElement(struct {
		XMLcon               string                `xml:"xmlns:con,attr"`
		XMLppo               string                `xml:"xmlns:ppo,attr"`
		ContactCustomer      *bool                 `xml:"submitOrder1411Request>ppo:ContactCustomer"`
		CreditCardInfo       *creditCardInfoType   `xml:"submitOrder1411Request>ppo:CreditCardInfo"`
		PPOSessionID         *string               `xml:"submitOrder1411Request>ppo:PPOSessionID"`
		PrepareTransactionID *string               `xml:"submitOrder1411Request>ppo:PrepareTransactionID"`
		RequestedAppointment *requestedAppointment `xml:"submitOrder1411Request>ppo:RequestedAppointment"`
		SalesCode            *string               `xml:"submitOrder1411Request>ppo:SalesCode"`
	}{
		XMLcon:               "http://localhost/ContentOrderService",
		XMLppo:               "http://www.qwest.com/ppo/",
		ContactCustomer:      x.ContactCustomer,
		CreditCardInfo:       x.CreditCardInfo,
		PPOSessionID:         x.PPOSessionID,
		PrepareTransactionID: x.PrepareTransactionID,
		RequestedAppointment: x.RequestedAppointment,
		SalesCode:            x.SalesCode,
	}, start)
}

type submitOrder1411ResponseMessage struct {
	XMLName xml.Name `xml:" submitOrder1411Response"`

	*submitOrder1411Response
}

func (t submitOrder1411ResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.submitOrder1411Response == nil {
		return nil
	}

	x := t.submitOrder1411Response.SubmitOrder1411Response

	start.Name = xml.Name{Local: "con:submitOrder1411Response"}

	return e.EncodeElement(struct {
		XMLcon                  string                   `xml:"xmlns:con,attr"`
		XMLppo                  string                   `xml:"xmlns:ppo,attr"`
		AccountNumber           *string                  `xml:"submitOrder1411Response>ppo:AccountNumber"`
		AppointmentTime         *appointmentTime         `xml:"submitOrder1411Response>ppo:AppointmentTime"`
		AuthorizationCode       *string                  `xml:"submitOrder1411Response>ppo:AuthorizationCode"`
		CustomerResponseMessage *string                  `xml:"submitOrder1411Response>ppo:CustomerResponseMessage"`
		DueDate                 *string                  `xml:"submitOrder1411Response>ppo:DueDate"`
		OrderList               *arrayOfOrders           `xml:"submitOrder1411Response>ppo:OrderList"`
		OrderOutcome            *string                  `xml:"submitOrder1411Response>ppo:OrderOutcome"`
		PartnerOrderID          *string                  `xml:"submitOrder1411Response>ppo:PartnerOrderID"`
		PartnerResponseMessage  *string                  `xml:"submitOrder1411Response>ppo:PartnerResponseMessage"`
		ProductsOrderedList     *arrayOfProductsOrdered  `xml:"submitOrder1411Response>ppo:ProductsOrderedList"`
		ResponseCode            *int32                   `xml:"submitOrder1411Response>ppo:ResponseCode"`
		SubmitOrder1411Response *submitOrder1411Response `xml:"submitOrder1411Response>ppo:SubmitOrder1411Response"`
		TelephoneNumber         *string                  `xml:"submitOrder1411Response>ppo:TelephoneNumber"`
		TransactionID           *string                  `xml:"submitOrder1411Response>ppo:TransactionID"`
	}{
		XMLcon:                  "http://localhost/ContentOrderService",
		XMLppo:                  "http://www.qwest.com/ppo/",
		AccountNumber:           x.AccountNumber,
		AppointmentTime:         x.AppointmentTime,
		AuthorizationCode:       x.AuthorizationCode,
		CustomerResponseMessage: x.CustomerResponseMessage,
		DueDate:                 x.DueDate,
		OrderList:               x.OrderList,
		OrderOutcome:            x.OrderOutcome,
		PartnerOrderID:          x.PartnerOrderID,
		PartnerResponseMessage:  x.PartnerResponseMessage,
		ProductsOrderedList:     x.ProductsOrderedList,
		ResponseCode:            x.ResponseCode,
		SubmitOrder1411Response: x.SubmitOrder1411Response,
		TelephoneNumber:         x.TelephoneNumber,
		TransactionID:           x.TransactionID,
	}, start)
}

type submitOrderRequestMessage struct {
	XMLName xml.Name `xml:" submitOrderRequest"`

	*submitOrderRequest
}

func (t submitOrderRequestMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.submitOrderRequest == nil {
		return nil
	}

	x := t.submitOrderRequest.SubmitOrderRequest

	start.Name = xml.Name{Local: "con:submitOrderRequest"}

	return e.EncodeElement(struct {
		XMLcon               string                `xml:"xmlns:con,attr"`
		XMLppo               string                `xml:"xmlns:ppo,attr"`
		ContactCustomer      *bool                 `xml:"submitOrderRequest>ppo:ContactCustomer"`
		CreditCardInfo       *creditCardInfoType   `xml:"submitOrderRequest>ppo:CreditCardInfo"`
		PPOSessionID         *string               `xml:"submitOrderRequest>ppo:PPOSessionID"`
		RequestedAppointment *requestedAppointment `xml:"submitOrderRequest>ppo:RequestedAppointment"`
		SalesCode            *string               `xml:"submitOrderRequest>ppo:SalesCode"`
		SubmitOrderRequest   *submitOrderRequest   `xml:"submitOrderRequest>ppo:SubmitOrderRequest"`
	}{
		XMLcon:               "http://localhost/ContentOrderService",
		XMLppo:               "http://www.qwest.com/ppo/",
		ContactCustomer:      x.ContactCustomer,
		CreditCardInfo:       x.CreditCardInfo,
		PPOSessionID:         x.PPOSessionID,
		RequestedAppointment: x.RequestedAppointment,
		SalesCode:            x.SalesCode,
		SubmitOrderRequest:   x.SubmitOrderRequest,
	}, start)
}

type submitOrderRequestResponseMessage struct {
	XMLName xml.Name `xml:" submitOrderRequestResponse"`

	*submitOrderRequestResponse
}

func (t submitOrderRequestResponseMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.submitOrderRequestResponse == nil {
		return nil
	}

	x := t.submitOrderRequestResponse.SubmitOrderResponse

	start.Name = xml.Name{Local: "con:submitOrderRequestResponse"}

	return e.EncodeElement(struct {
		XMLcon              string                  `xml:"xmlns:con,attr"`
		XMLppo              string                  `xml:"xmlns:ppo,attr"`
		AppointmentTime     *appointmentTime        `xml:"submitOrderResponse>ppo:AppointmentTime"`
		AuthorizationCode   *string                 `xml:"submitOrderResponse>ppo:AuthorizationCode"`
		DueDate             *string                 `xml:"submitOrderResponse>ppo:DueDate"`
		OrderOutcome        *string                 `xml:"submitOrderResponse>ppo:OrderOutcome"`
		PartnerOrderID      *string                 `xml:"submitOrderResponse>ppo:PartnerOrderID"`
		ProductsOrderedList *arrayOfProductsOrdered `xml:"submitOrderResponse>ppo:ProductsOrderedList"`
		QwestOrderList      *arrayOfQwestOrders     `xml:"submitOrderResponse>ppo:QwestOrderList"`
		ResponseCode        *int32                  `xml:"submitOrderResponse>ppo:ResponseCode"`
		ResponseMessage     *string                 `xml:"submitOrderResponse>ppo:ResponseMessage"`
		TN                  *string                 `xml:"submitOrderResponse>ppo:TN"`
		TransactionID       *string                 `xml:"submitOrderResponse>ppo:TransactionID"`
	}{
		XMLcon:              "http://localhost/ContentOrderService",
		XMLppo:              "http://www.qwest.com/ppo/",
		AppointmentTime:     x.AppointmentTime,
		AuthorizationCode:   x.AuthorizationCode,
		DueDate:             x.DueDate,
		OrderOutcome:        x.OrderOutcome,
		PartnerOrderID:      x.PartnerOrderID,
		ProductsOrderedList: x.ProductsOrderedList,
		QwestOrderList:      x.QwestOrderList,
		ResponseCode:        x.ResponseCode,
		ResponseMessage:     x.ResponseMessage,
		TN:                  x.TN,
		TransactionID:       x.TransactionID,
	}, start)
}
