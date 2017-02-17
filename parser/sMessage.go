package parser

type sMessage struct {
	XMLField sField
	Type     string

	// Custom marshaler related fields.
	Struct     string
	StructType string
	XmlTag     string
	Fields     mapofFields
	NSFields   mapofFields
	Namespaces []string
	LocalName  string
	Marshaler  bool
}

type mapofMessages map[string]sMessage

func (m mapofMessages) add(i string, s string) bool {
	if i == "" {
		return false
	}

	// TODO handle interfaces and slice of bytes
	if n := toGoType(makeUnexported(removeNS(i))); n == "interface{}" || n == "[]byte" {
		return false
	}

	i = makeUnexported(lintName(removeNS(i + "Message")))

	if v, ok := m[i]; ok {
		m[i] = sMessage{
			XMLField: sField{
				Name: "XMLName",
				Type: "xml.Name",
				Tag:  v.XMLField.Tag,
			},
			Type: toGoPointerType(lintName(makeUnexported(removeNS(s)))),
		}
		return true
	}

	m[i] = sMessage{
		XMLField: sField{
			Tag: s,
		},
	}

	return true
}
