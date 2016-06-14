package parser

type sType struct {
	Name           string
	UnderlyingType string
}

type mapofTypes map[string]sType

func (m mapofTypes) add(s sType) bool {
	if s.UnderlyingType == "" {
		return false
	}

	s.resolveName()
	s.resolveUnderlyingType()

	if _, ok := m[s.Name]; !ok {
		m[s.Name] = s
		return true
	}

	return false
}

func (s *sType) resolveName() {
	s.Name = makeUnexported(lintName(replaceReservedWords(s.Name)))
}

func (s *sType) resolveUnderlyingType() {
	s.UnderlyingType = makeUnexported(toGoType(removeNS(s.UnderlyingType)))
}
