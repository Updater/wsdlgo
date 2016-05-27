// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

type seqInner struct {
	Str3 *string `xml:"str3"`
}

type typeOut struct {
	Str1 *string `xml:"str1"`
	Str2 *string `xml:"str2"`
	*seqInner
}
