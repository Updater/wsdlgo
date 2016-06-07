package parser

import (
	"bytes"
	"encoding/xml"
	"go/format"
	"io"
	"strings"
	"text/template"
	"unicode"
)

// Generator defines behavior for sending generated code to a writer interface.
type Generator interface {
	Write(io.Writer) error
}

// generator defines the struct for the WSDL code generator.
type generator struct {
	Name    string
	WSDL    *wsdl
	Element element

	content []byte
	reader  io.Reader
}

// Write implements the write behavior for Generator interface.
// Resulting in formatted code sent to a provided writer.
func (g *generator) Write(w io.Writer) error {
	// go fmt the generated code
	s, err := format.Source(g.content)
	if err != nil {
		_, err = w.Write(g.content)
		if err != nil {
			return err
		}
	}

	_, err = w.Write(s)
	return err
}

func (g *generator) populateElement() error {
	f := template.FuncMap{
		"toGoType":              toGoType,
		"toGoPointerType":       toGoPointerType,
		"replaceReservedWords":  replaceReservedWords,
		"makeUnexported":        makeUnexported,
		"makeExported":          makeExported,
		"normalize":             normalize,
		"lint":                  lint,
		"removeNS":              removeNS,
		"removePackage":         removePackage,
		"convertPointerToValue": convertPointerToValue,
	}

	t := template.New("types")
	for _, v := range []string{elementTmpl} {
		var err error
		t, err = template.Must(t.Clone()).Funcs(f).Parse(v)
		if err != nil {
			return err
		}
	}

	d := new(bytes.Buffer)
	err := t.Execute(d, g)
	if err != nil {
		return err
	}

	g.content = d.Bytes()
	return nil
}

func (g *generator) parse() error {
	b := new(bytes.Buffer)
	if _, err := b.ReadFrom(g.reader); err != nil {
		return err
	}

	g.WSDL = new(wsdl)
	if err := xml.Unmarshal(b.Bytes(), g.WSDL); err != nil {
		return err
	}

	g.Element.Imports = make(map[string]string)
	g.Element.Types = make(map[string]*sType)
	g.Element.Consts = make(map[string]*sConst)
	g.Element.Structs = make(map[string]*sStruct)
	g.Element.Messages = make(map[string]*sMessage)
	doMap([]mapper{g.WSDL}, &g.Element)

	return g.populateElement()
}

// NewGenerator initializes a Generator interface implemented by generator type.
func NewGenerator(r io.Reader, pkg string) (Generator, error) {
	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		pkg = "types"
	}

	g := generator{
		reader: r,
		Name:   pkg,
	}

	err := g.parse()
	return &g, err
}

func makeUnexported(s string) string {
	f := []rune(s)
	if len(f) == 0 {
		return s
	}

	f[0] = unicode.ToLower(f[0])
	return string(f)
}

func makeExported(s string) string {
	f := []rune(s)
	if len(f) == 0 {
		return s
	}

	f[0] = unicode.ToUpper(f[0])
	return string(f)
}

var reservedWords = map[string]string{
	"break":       "break_",
	"default":     "default_",
	"func":        "func_",
	"interface":   "interface_",
	"select":      "select_",
	"case":        "case_",
	"defer":       "defer_",
	"go":          "go_",
	"map":         "map_",
	"struct":      "struct_",
	"chan":        "chan_",
	"else":        "else_",
	"goto":        "goto_",
	"package":     "package_",
	"switch":      "switch_",
	"const":       "const_",
	"fallthrough": "fallthrough_",
	"if":          "if_",
	"range":       "range_",
	"type":        "type_",
	"continue":    "continue_",
	"for":         "for_",
	"import":      "import_",
	"return":      "return_",
	"var":         "var_",
	"time":        "time_",
}

// normalize, normalizes value to be used as a valid Go identifier, avoiding compilation issues
func normalize(s string) string {
	m := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return -1
	}

	return strings.Map(m, s)
}

func capitalizeMultipleWord(s string) string {
	var b bytes.Buffer
	ss := strings.Split(s, "_")

	if len(ss) < 2 {
		return s
	}

	b.WriteString(ss[0])
	for i := 1; i < len(ss); i++ {
		b.WriteString(makeExported(ss[i]))
	}

	return b.String()
}

// replaceReservedWords replaces Go reserved keywords to avoid compilation issues
func replaceReservedWords(s string) string {
	s = capitalizeMultipleWord(s)
	if v, ok := reservedWords[s]; ok {
		return v
	}

	return normalize(s)
}

var xsd2GoTypes = map[string]string{
	"string":        "string",
	"token":         "string",
	"float":         "float32",
	"double":        "float64",
	"decimal":       "float64",
	"integer":       "int32",
	"int":           "int32",
	"short":         "int16",
	"byte":          "int8",
	"long":          "int64",
	"boolean":       "bool",
	"datetime":      "time.Time",
	"date":          "time.Time",
	"time":          "time.Time",
	"base64binary":  "[]byte",
	"hexbinary":     "[]byte",
	"unsignedint":   "uint32",
	"unsignedshort": "uint16",
	"unsignedbyte":  "byte",
	"unsignedlong":  "uint64",
	"anytype":       "interface{}",
	"anyuri":        "interface{}",
	"duration":      "string",
	"qname":         "string",
}

func toGoType(s string) string {
	v, ok := xsd2GoTypes[strings.ToLower(s)]
	if ok {
		return v
	}

	return replaceReservedWords(s)
}

func toGoPointerType(s string) string {
	v := toGoType(s)
	if v == "interface{}" || strings.HasPrefix(v, "[]") {
		return v
	}

	return "*" + v
}

// lint returns a different name if it should be different.
func lint(s string) string {
	u := strings.ToUpper(s)
	_, ok := commonInitialisms[u]
	if ok {
		return u
	}

	return s
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
// Borrowwed from github.com/golang/lint
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}

func removeNS(s string) string {
	// Handles name space, ie. xsd:string, xs:string
	r := strings.Split(s, ":")

	if len(r) == 2 {
		return r[1]
	}

	return r[0]
}

func removePackage(s string) string {
	r := strings.Split(s, ".")

	if len(r) == 2 {
		return r[1]
	}

	return r[0]
}

func convertPointerToValue(s string) string {
	return strings.Replace(s, "*", "", -1)
}
