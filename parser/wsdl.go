package parser

import (
	"encoding/xml"
	"strings"
)

// resolveXMLTag is a helper function to solve the proper XML Name tag based for element.
func resolveXMLTag(s string, x string) (string, bool) {
	r := strings.Split(s, ":")
	if len(r) != 2 {
		return r[len(r)-1], false
	}

	rs := r[1]
	if x == "" {
		return rs, false
	}

	switch r[0] {
	case "partns":
		rs = x + " " + rs
	}

	return rs, true
}

func addPackageImport(s string, m map[string]bool) {
	if _, ok := m[s]; !ok {
		m[s] = true
	}
}

// wsdl represents the global structure of a wsdl file.
type wsdl struct {
	Name            string         `xml:"name,attr"`
	TNS             string         `xml:"tns,attr"`
	TargetNamespace string         `xml:"targetNamespace,attr"`
	Imports         []wsdlImport   `xml:"import"`
	Doc             string         `xml:"documentation"`
	Types           wsdlType       `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Messages        []wsdlMessage  `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
	PortTypes       []wsdlPortType `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Binding         []wsdlBinding  `xml:"http://schemas.xmlsoap.org/wsdl/ binding"`
	Service         []wsdlService  `xml:"http://schemas.xmlsoap.org/wsdl/ service"`

	PackageImport map[string]bool
}

// UnmarshalXML satisfies the XML Unmarshaler interface.
// Populates wsdl based on xml data.
func (w *wsdl) UnmarshalXML(d *xml.Decoder, s xml.StartElement) error {
	// wsdlAlias is used to disconnect struct methods and prevent potential loop.
	type wsdlAlias wsdl
	v := wsdlAlias(*w)

	if err := d.DecodeElement(&v, &s); err != nil {
		return err
	}

	v.PackageImport = make(map[string]bool)

	var wp []*wsdlPart
	for _, m := range v.Messages {
		for _, p := range m.Parts {
			wp = append(wp, p)
		}
	}

	var xe []*xsdElement
	var se []*xsdElement
	var ct []*xsdElement

	for _, s := range v.Types.Schemas {
		se = append(se, s.Elements...)
		xe = append(xe, se...)

		for _, e := range s.Elements {
			if e.ComplexType != nil {
				ct = append(ct, e.ComplexType.Sequence...)
				ct = append(ct, e.ComplexType.Choice...)
				ct = append(ct, e.ComplexType.SequenceChoice...)
				ct = append(ct, e.ComplexType.All...)
			}
		}
		xe = append(xe, ct...)

		for _, c := range s.ComplexTypes {
			xe = append(xe, c.Sequence...)
			xe = append(xe, c.Choice...)
			xe = append(xe, c.SequenceChoice...)
			xe = append(xe, c.All...)
		}
	}

	// Set proper XML Tag for field.
	func(wp []*wsdlPart, x []*xsdElement) {
		for _, p := range wp {
			for _, e := range x {
				if strings.HasSuffix(p.Element, e.Name) {
					if s, ok := resolveXMLTag(p.Element, p.Partns); ok {
						e.XMLTag = s
					}
				}
			}
		}
	}(wp, se)

	// Populate fields which are used by templates for adding package imports.
	func(m map[string]bool, x []*xsdElement) {
		for _, e := range x {
			if removePackage(toGoType(removeNS(e.Type))) == "Time" {
				addPackageImport("time", m)
			}

			if e.XMLTag != "" || (e.NameReqNil != "" && !e.TypeReqNilExists) {
				addPackageImport("encoding/xml", m)
			}
		}
	}(v.PackageImport, xe)

	*w = wsdl(v)
	return nil
}

// wsdlImport is the struct used for deserializing wsdl imports.
type wsdlImport struct {
	Namespace string `xml:"namespace,attr"`
	Location  string `xml:"location,attr"`
}

// wsdlType represents the entry point for deserializing XSD schemas used by the wsdl file.
type wsdlType struct {
	Doc     string       `xml:"documentation"`
	Schemas []*xsdSchema `xml:"schema"`
}

// wsdlPart defines the struct for a function parameter within a wsdl.
type wsdlPart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
	Type    string `xml:"type,attr"`
	Partns  string `xml:"partns,attr"`
}

// wsdlMessage represents a function, which in turn has one or more parameters.
type wsdlMessage struct {
	Name  string      `xml:"name,attr"`
	Doc   string      `xml:"documentation"`
	Parts []*wsdlPart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

// wsdlFault represents a wsdl fault message.
type wsdlFault struct {
	Name      string        `xml:"name,attr"`
	Message   string        `xml:"message,attr"`
	Doc       string        `xml:"documentation"`
	SOAPFault wsdlSOAPFault `xml:"http://schemas.xmlsoap.org/wsdl/soap/ fault"`
}

// wsdlInput represents a wsdl input message.
type wsdlInput struct {
	Name       string           `xml:"name,attr"`
	Message    string           `xml:"message,attr"`
	Doc        string           `xml:"documentation"`
	SOAPBody   wsdlSOAPBody     `xml:"http://schemas.xmlsoap.org/wsdl/soap/ body"`
	SOAPHeader []wsdlSOAPHeader `xml:"http://schemas.xmlsoap.org/wsdl/soap/ header"`
}

// wsdlOutput represents a wsdl output message.
type wsdlOutput struct {
	Name       string           `xml:"name,attr"`
	Message    string           `xml:"message,attr"`
	Doc        string           `xml:"documentation"`
	SOAPBody   wsdlSOAPBody     `xml:"http://schemas.xmlsoap.org/wsdl/soap/ body"`
	SOAPHeader []wsdlSOAPHeader `xml:"http://schemas.xmlsoap.org/wsdl/soap/ header"`
}

// wsdlOperation represents the contract of an entire operation or function.
type wsdlOperation struct {
	Name          string            `xml:"name,attr"`
	Doc           string            `xml:"documentation"`
	Input         wsdlInput         `xml:"input"`
	Output        wsdlOutput        `xml:"output"`
	Faults        []wsdlFault       `xml:"fault"`
	SOAPOperation wsdlSOAPOperation `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
}

// wsdlPortType defines the service, operations that can be performed and the messages involved.
// A port type can be compared to a function library, module or class.
type wsdlPortType struct {
	Name       string          `xml:"name,attr"`
	Doc        string          `xml:"documentation"`
	Operations []wsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

// wsdlSOAPBinding represents a SOAP binding to the web service.
type wsdlSOAPBinding struct {
	Style     string `xml:"style,attr"`
	Transport string `xml:"transport,attr"`
}

// wsdlSOAPOperation represents a service operation in SOAP terms.
type wsdlSOAPOperation struct {
	SOAPAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

// wsdlSOAPHeader defines the header for a SOAP service.
type wsdlSOAPHeader struct {
	Message       string                `xml:"message,attr"`
	Part          string                `xml:"part,attr"`
	Use           string                `xml:"use,attr"`
	EncodingStyle string                `xml:"encodingStyle,attr"`
	Namespace     string                `xml:"namespace,attr"`
	HeadersFault  []wsdlSOAPHeaderFault `xml:"headerfault"`
}

// wsdlSOAPHeaderFault defines a SOAP fault header.
type wsdlSOAPHeaderFault struct {
	Message       string `xml:"message,attr"`
	Part          string `xml:"part,attr"`
	Use           string `xml:"use,attr"`
	EncodingStyle string `xml:"encodingStyle,attr"`
	Namespace     string `xml:"namespace,attr"`
}

// wsdlSOAPBody defines SOAP body characteristics.
type wsdlSOAPBody struct {
	Parts         string `xml:"parts,attr"`
	Use           string `xml:"use,attr"`
	EncodingStyle string `xml:"encodingStyle,attr"`
	Namespace     string `xml:"namespace,attr"`
}

// wsdlSOAPFault defines a SOAP fault message characteristics.
type wsdlSOAPFault struct {
	Parts         string `xml:"parts,attr"`
	Use           string `xml:"use,attr"`
	EncodingStyle string `xml:"encodingStyle,attr"`
	Namespace     string `xml:"namespace,attr"`
}

// wsdlSOAPAddress defines the location for the SOAP service.
type wsdlSOAPAddress struct {
	Location string `xml:"location,attr"`
}

// wsdlBinding defines only a SOAP binding and its operations
type wsdlBinding struct {
	Name        string          `xml:"name,attr"`
	Type        string          `xml:"type,attr"`
	Doc         string          `xml:"documentation"`
	SOAPBinding wsdlSOAPBinding `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
	Operations  []wsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

// wsdlPort defines the properties for a SOAP port only.
type wsdlPort struct {
	Name        string          `xml:"name,attr"`
	Binding     string          `xml:"binding,attr"`
	Doc         string          `xml:"documentation"`
	SOAPAddress wsdlSOAPAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

// wsdlService defines the list of SOAP services associated with the wsdl.
type wsdlService struct {
	Name  string     `xml:"name,attr"`
	Doc   string     `xml:"documentation"`
	Ports []wsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}
