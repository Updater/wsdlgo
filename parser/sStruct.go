package parser

type sStruct struct {
	Name                 string
	Fields               mapofFields
	NillableRequiredType bool

	// Custom marshaler related fields.
	EncoderFields mapofFields
	Marshaler     bool

	w *wsdl
}

type mapofStructs map[string]sStruct

func (m mapofStructs) add(i string, s sStruct) string {
	if i == "" {
		return i
	}

	// TODO handle interfaces and slice of bytes
	if n := toGoType(makeUnexported(removeNS(s.Name))); n == "interface{}" || n == "[]byte" {
		return ""
	}

	s.resolveFields()
	s.resolveName()

	if _, ok := m[s.Name]; !ok {
		m[s.Name] = s
	}

	return s.Name
}

func (s *sStruct) resolveName() {
	if s.NillableRequiredType {
		s.Name = makeUnexported(lintName(normalize(removeNS(s.Name + "ReqNil"))))
		return
	}

	s.Name = toGoType(replaceReservedWords(makeUnexported(lintName(removeNS(s.Name)))))
}

func (s *sStruct) resolveFields() {
	if s.NillableRequiredType {
		s.Fields = mapofFields{s.Name: sField{
			Type: toGoPointerType(makeUnexported(lintName(removeNS(s.Name))))},
		}
		return
	}

	s.Fields = mapofFields{}
}
