package parser_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Bridgevine/wsdlgo/parser"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		wf  string
		pf  string
		err error
	}{
		{
			wf: "testdata/simpleType.wsdl",
			pf: "testdata/simpleType.go",
		},
	}

	for index, tt := range tests {
		if tt.wf != "" && tt.pf != "" {
			f, err := os.Open(tt.wf)
			if err != nil {
				t.Fatalf("Test %d errored while opening xml file %s: %v", index, tt.wf, err)
			}
			defer f.Close()

			g, err := parser.NewGenerator(f, "types")
			if err != nil {
				t.Fatalf("Test %d errored while opening xml file %s: %v", index, tt.wf, err)
			}

			b := new(bytes.Buffer)
			g.Write(b)

			s, err := ioutil.ReadFile(tt.pf)
			if err != nil {
				t.Fatalf("Test %d errored while Go source code file %s: %v", index, tt.pf, err)
			}

			if !bytes.Equal(b.Bytes(), s) {
				t.Fatalf("Test %d errored while comparing wsdl and Go source code file [%s] and [%s]", index, tt.pf, tt.wf)
			}
		}
	}
}
