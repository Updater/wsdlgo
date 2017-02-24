package parser

type sField struct {
	Name        string
	Type        string
	Tag         string
	Struct      sStruct
	InitialName string

	attr     bool
	pointer  bool
	array    bool
	required bool
	nillable bool
}

type mapofFields map[string]sField

func (m mapofFields) add(i string, s sField) bool {
	if i == "" || s.Type == "" {
		return false
	}

	s.resolveType()
	s.resolveTag()
	s.resolveName(i)
	m[s.Name] = s

	return true
}

func (s *sField) resolveName(i string) {
	if s.Name == "" {
		return
	}
	s.InitialName = makeExported(normalize(removeNS(i)))
	s.Name = makeExported(lintName(normalize(removeNS(i))))
}

func (s *sField) resolveType() {
	if s.nillable && s.required {
		s.Type = toGoType(makeUnexported(lintName(removeNS(s.Type + "ReqNil"))))
		return
	}

	tn := makeUnexported(lintName(removeNS(s.Type)))
	if s.array {
		s.Type = "[]" + toGoType(tn)
		return
	}
	if s.pointer {
		s.Type = toGoPointerType(tn)
		return
	}
	s.Type = toGoPointerType(tn)
}

func (s *sField) resolveTag() {
	if s.Name == "" {
		return
	}

	if s.attr {
		s.Tag = "`" + `xml:"` + s.Name + `,attr"` + "`"
		return
	}

	s.Tag = "`" + `xml:"` + s.Name + `"` + "`"
}
