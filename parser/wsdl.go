package parser

import "strings"

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
}

func (w *wsdl) doMap(p interface{}) bool {
	switch u := p.(type) {
	case *element:
		var m []mapper
		for _, v := range w.Messages {
			m = append(m, v)
		}
		for _, v := range w.Types.Schemas {
			m = append(m, v)
		}
		doMap(m, u)

		// Import handling logic
		for _, s := range u.Structs {
			if s.NillableRequiredType {
				u.Imports.add("encoding/xml")
			}

			for _, f := range s.Fields {
				if strings.HasPrefix(convertPointerToValue(f.Type), "time.") {
					u.Imports.add("time")
				}

				if f.Type == "xml.Name" {
					u.Imports.add("encoding/xml")
				}
			}
		}

		for _, s := range u.Messages {
			if s.Type != "" {
				u.Imports.add("encoding/xml")
				break
			}
		}

		return true
	}

	return false
}

// wsdlImport is the struct used for deserializing wsdl imports.
type wsdlImport struct {
	Namespace string `xml:"namespace,attr"`
	Location  string `xml:"location,attr"`
}

// wsdlType represents the entry point for deserializing XSD schemas used by the wsdl file.
type wsdlType struct {
	Doc     string      `xml:"documentation"`
	Schemas []xsdSchema `xml:"schema"`
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
	Name  string     `xml:"name,attr"`
	Doc   string     `xml:"documentation"`
	Parts []wsdlPart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

func (x wsdlMessage) doMap(p interface{}) bool {
	switch u := p.(type) {
	case *element:
		if len(x.Parts) == 0 {
			break
		}

		n := removeNS(x.Parts[0].Element)
		tg := "`" + `xml:"` + x.Parts[0].Partns + ` ` + n + `"` + "`"
		u.Messages.add(n, tg)
		return true
	}

	return false
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
