package parser

type sConst struct {
	Name  string
	Type  string
	Value string
}

type mapofConsts map[string]sConst

func (m mapofConsts) add(s sConst) bool {
	s.resolveName()
	s.resolveType()

	if _, ok := m[s.Name]; !ok {
		m[s.Name] = s
		return true
	}

	return false
}

func (s *sConst) resolveName() {
	s.Name = makeUnexported(lintName(replaceReservedWords(s.Type + s.Value)))
}

func (s *sConst) resolveType() {
	s.Type = makeUnexported(lintName(s.Type))
}
