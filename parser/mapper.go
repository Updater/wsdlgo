package parser

// Mapper represents mapping related behaviors.
type mapper interface {
	doMap(interface{}) bool
}

// DoMap executes logic for mapping child elements of provided Mapper type.
func doMap(m []mapper, p interface{}) {
	for _, v := range m {
		v.doMap(p)
	}
}
