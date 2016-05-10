package parser

import (
	"bytes"
	"encoding/xml"
	"errors"
	"go/format"
	"io"
	"io/ioutil"
	"strings"
	"text/template"
	"unicode"
)

// Errors that can be returned.
var (
	ErrMissingWSDLFile = errors.New("WSDL file is required to generate Go proxy")
)

// Generator defines behavior for sending generated code to a writer interface.
type Generator interface {
	Write(io.Writer) error
}

// generator defines the struct for the WSDL code generator.
type generator struct {
	Name string
	WSDL *wsdl

	content []byte
	file    string
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

func (g *generator) populateContent() error {
	funcMap := template.FuncMap{
		"toGoType":             toGoType,
		"replaceReservedWords": replaceReservedWords,
		"makeUnexported":       makeUnexported,
	}

	tb, err := template.New("base.tmpl").Funcs(funcMap).Parse(baseTmpl)
	if err != nil {
		return err
	}

	tt, err := template.Must(tb.Clone()).Parse(simpleTypesTmpl)
	if err != nil {
		return err
	}

	d := new(bytes.Buffer)
	err = tt.Execute(d, g)
	if err != nil {
		return err
	}

	g.content = d.Bytes()
	return nil
}

func (g *generator) parse() error {
	data, err := ioutil.ReadFile(g.file)
	if err != nil {
		return err
	}

	g.WSDL = new(wsdl)
	err = xml.Unmarshal(data, g.WSDL)
	if err != nil {
		return err
	}

	return g.populateContent()
}

// NewGenerator initializes a Generator interface implemented by generator type.
func NewGenerator(in string, pkg string) (Generator, error) {
	in = strings.TrimSpace(in)
	if in == "" {
		return nil, ErrMissingWSDLFile
	}

	pkg = strings.TrimSpace(pkg)
	if pkg == "" {
		pkg = "types"
	}

	g := generator{
		file: in,
		Name: pkg,
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

// replaceReservedWords replaces Go reserved keywords to avoid compilation issues
func replaceReservedWords(s string) string {
	v, ok := reservedWords[s]
	if ok {
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
}

func toGoType(s string) string {
	// Handles name space, ie. xsd:string, xs:string
	r := strings.Split(s, ":")

	t := r[0]
	if len(r) == 2 {
		t = r[1]
	}

	v, ok := xsd2GoTypes[strings.ToLower(t)]
	if ok {
		return v
	}

	return "*" + replaceReservedWords(makeUnexported(t))
}
