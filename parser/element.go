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

type sType struct {
	Name           string
	UnderlyingType string
}

type sConst struct {
	Name  string
	Type  string
	Value string
}

type sField struct {
	Name string
	Type string
	Tag  string
	attr bool
}

type sStruct struct {
	Fields               map[string]*sField
	NillableRequiredType bool
}

type sMessage struct {
	XMLField sField
	Type     string
}

type element struct {
	Imports  map[string]string
	Types    map[string]*sType
	Consts   map[string]*sConst
	Structs  map[string]*sStruct
	Messages map[string]*sMessage
}
