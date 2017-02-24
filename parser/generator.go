package parser

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"text/template"
	"time"
	"unicode"
)

// Generator defines behavior for sending generated code to a writer interface.
type Generator interface {
	Write(io.Writer) error
}

// generator defines the struct for the WSDL code generator.
type generator struct {
	Name    string
	Element element

	sources   []string
	content   []byte
	cert      string
	certKey   string
	marshaler bool
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
		"removePackage":         removePackage,
		"convertPointerToValue": convertPointerToValue,
		"timestamp":             time.Now,
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

// loadCert is a helper function to load certificates from the specified file location.
func loadCert(certFile, keyFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		MinVersion:   tls.VersionTLS10,
		MaxVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cert},
	}, nil
}

func fetchXSDSchema(url, certFile, keyFile string) ([]byte, error) {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	var b []byte
	c, err := loadCert(certFile, keyFile)
	if err != nil {
		return b, err
	}
	if c != nil {
		transport.TLSClientConfig = c
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}

	res, err := client.Get(url)
	if err != nil {
		return b, err
	}

	return ioutil.ReadAll(res.Body)
}

func (g *generator) parse() error {
	g.Element.Imports = make(mapofImports)
	g.Element.Types = make(mapofTypes)
	g.Element.Consts = make(mapofConsts)
	g.Element.Structs = make(mapofStructs)
	g.Element.Messages = make(mapofMessages)
	for _, f := range g.sources {
		xb, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		gwsdl := new(wsdl)
		if err := xml.Unmarshal(xb, &gwsdl); err != nil {
			return err
		}

		// Fetch schemas from the remote location.
		var xsdSchemas []xsdSchema
		var localNS, fieldNS string
		for _, s := range gwsdl.Types.Schemas {
			for _, i := range s.Imports {
				b, err := fetchXSDSchema(i.SchemaLocation, g.cert, g.certKey)
				if strings.Contains(i.Namespace, "localhost") {
					localNS = i.Namespace
				}
				if strings.Contains(i.Namespace, "qwest") {
					fieldNS = i.Namespace
				}
				if err != nil {
					return err
				}
				var xs xsdSchema
				if err := xml.Unmarshal(b, &xs); err != nil {
					return err
				}
				xsdSchemas = append(xsdSchemas, xs)
			}
		}

		if len(xsdSchemas) > 0 {
			gwsdl.Types.Schemas = xsdSchemas
		}

		doMap([]mapper{gwsdl}, &g.Element)

		// Update namespaces.
		for m, sMsg := range g.Element.Messages {
			sMsg.Namespaces = []string{localNS, fieldNS}
			sMsg.Marshaler = g.marshaler
			g.Element.Messages[m] = sMsg
		}

		for s, sStr := range g.Element.Structs {
			if sStr.NillableRequiredType {
				continue
			}

			sStr.Marshaler = g.marshaler
			sStr.EncoderFields = sStr.Fields
			g.Element.Structs[s] = sStr
		}

		if !g.marshaler {
			continue
		}
		// Update a message a struct is related with,
		// so we can build custom XMLMarshaler implementation on it.
		if err := addMarshalerToMessages(&g.Element, localNS, fieldNS); err != nil {
			return err
		}

		if err := addMarshalerToXMLStructs(&g.Element, localNS, fieldNS); err != nil {
			return err
		}
	}

	return g.populateElement()
}

func addFields(allFields mapofFields, e *element, fields mapofFields) {
	for f, fv := range fields {
		if f != "" && fv.Type != "" {
			allFields[f] = fv
		} else {
			n := convertPointerToValue(fv.Type)
			str, ok := e.Structs[n]
			if !ok {
				continue
			}
			addFields(allFields, e, str.Fields)
		}
	}
}

func addMarshalerToXMLStructs(e *element, localNS, fieldNS string) error {
	// Expand embedded fields.
	for s, sStr := range e.Structs {
		allFields := make(mapofFields)

		addFields(allFields, e, sStr.Fields)

		sStr.EncoderFields = allFields
		e.Structs[s] = sStr

	}

	for s, sStr := range e.Structs {
		var fields mapofFields
		b, err := json.Marshal(sStr.EncoderFields)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &fields); err != nil {
			return err
		}

		// Update fields' XML tags.
		for k, fld := range fields {
			fld.Tag = fmt.Sprintf("`xml:\"%s:%s\"`", namespaceToStr(fieldNS), fld.InitialName)
			fields[k] = fld
		}
		sStr.EncoderFields = fields

		e.Structs[s] = sStr
	}

	return nil
}

func addMarshalerToMessages(e *element, localNS, fieldNS string) error {
	for msg, sMsg := range e.Messages {
		sStr, ok := e.Structs[convertPointerToValue(sMsg.Type)]
		if !ok {
			continue
		}

		hasOne := len(sStr.Fields) == 1
		for k, f := range sStr.Fields {
			if hasOne {
				sMsg.Struct = k
				sMsg.XmlTag = extractTagName(f.Tag)
				sMsg.StructType = makeUnexported(convertPointerToValue(f.Type))
				break
			}

			if strings.ToLower(k) != strings.ToLower(convertPointerToValue(sMsg.Type)) {
				continue
			}
			sMsg.Struct = k
			sMsg.XmlTag = extractTagName(f.Tag)
			sMsg.StructType = makeUnexported(convertPointerToValue(f.Type))
		}

		sMsg.LocalName = fmt.Sprintf("%s:%s", namespaceToStr(localNS), convertPointerToValue(sMsg.Type))

		e.Messages[msg] = sMsg
	}

	for msg, sMsg := range e.Messages {
		sStr, ok := e.Structs[sMsg.StructType]
		if !ok {
			continue
		}

		var fields mapofFields
		b, err := json.Marshal(sStr.Fields)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &fields); err != nil {
			return err
		}

		// Update fields' XML tags.
		for k, fld := range fields {
			fld.Tag = fmt.Sprintf("`xml:\"%s>%s:%s\"`", sMsg.XmlTag, namespaceToStr(fieldNS), fld.InitialName)
			fields[k] = fld
		}
		sMsg.Fields = fields

		// Expand fields with namespaces.
		sMsg.NSFields = make(mapofFields)
		for _, n := range []string{localNS, fieldNS} {
			nstr := namespaceToStr(n)
			sMsg.NSFields["XML"+nstr] = sField{Name: n, Tag: fmt.Sprintf("`xml:\"xmlns:%s,attr\"`", nstr), Type: "string"}
		}

		e.Messages[msg] = sMsg
	}

	return nil
}

func extractTagName(xmltag string) string {
	matches := regexp.MustCompile(`\"(\w+)\"`).FindStringSubmatch(xmltag)
	if len(matches) == 2 {
		return matches[1]
	}

	return ""
}

func namespaceToStr(ns string) string {
	matches := regexp.MustCompile(`(\w+)`).FindAllString(ns, -1)
	if len(matches) == 0 {
		return ""
	}

	match := matches[len(matches)-1]

	return strings.ToLower(match[0:3])
}

// NewGenerator initializes a Generator interface implemented by generator type.
func NewGenerator(files []string, pkg, cert, certKey string, marshaler bool) (Generator, error) {
	var err error
	g := generator{
		sources:   files,
		Name:      pkg,
		cert:      cert,
		certKey:   certKey,
		marshaler: marshaler,
	}

	if err := g.parse(); err != nil {
		return &g, err
	}

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
	"break":       "breakType",
	"default":     "defaultType",
	"func":        "funcType",
	"interface":   "interfaceType",
	"select":      "selectType",
	"case":        "caseType",
	"defer":       "deferType",
	"go":          "goType",
	"map":         "mapType",
	"struct":      "structType",
	"chan":        "chanType",
	"else":        "elseType",
	"goto":        "gotoType",
	"package":     "packageType",
	"switch":      "switchType",
	"const":       "constType",
	"fallthrough": "fallthroughType",
	"if":          "ifType",
	"range":       "rangeType",
	"type":        "typeType",
	"continue":    "continueType",
	"for":         "forType",
	"import":      "importType",
	"return":      "returnType",
	"var":         "varType",
	"time":        "timeType",
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
	"string":             "string",
	"token":              "string",
	"float":              "float32",
	"double":             "float64",
	"decimal":            "float64",
	"integer":            "int32",
	"int":                "int32",
	"short":              "int16",
	"byte":               "int8",
	"long":               "int64",
	"boolean":            "bool",
	"datetime":           "time.Time",
	"date":               "time.Time",
	"time":               "time.Time",
	"base64binary":       "[]byte",
	"hexbinary":          "[]byte",
	"unsignedint":        "uint32",
	"unsignedshort":      "uint16",
	"unsignedbyte":       "byte",
	"unsignedlong":       "uint64",
	"anytype":            "interface{}",
	"anyuri":             "interface{}",
	"duration":           "string",
	"qname":              "string",
	"nonnegativeinteger": "uint",
	"gyearmonth":         "string",
	"language":           "string",
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
