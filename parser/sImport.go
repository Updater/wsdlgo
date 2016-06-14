package parser

type sImport string

type mapofImports map[string]sImport

func (m mapofImports) add(s string) bool {
	if s == "" {
		return false
	}

	if _, ok := m[s]; !ok {
		m[s] = sImport(s)
		return true
	}

	return false
}
