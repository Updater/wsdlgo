package parser

type sMessage struct {
	XMLField sField
	Type     string
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

	i = toGoType(makeUnexported(lintName(removeNS(i))))

	if v, ok := m[i]; ok {
		m[i] = sMessage{
			XMLField: sField{
				Name: "XMLName",
				Type: "xml.Name",
				Tag:  v.XMLField.Tag,
			},
			Type: toGoPointerType(makeUnexported(removeNS(s))),
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
