package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"regexp"

	"strings"

	"github.com/Bridgevine/wsdlgo/parser"
)

var (
	pkg     = flag.String("p", "types", "Package under which code will be generated")
	outFile = flag.String("o", "types.go", "File where the generated code will be saved")
	cert    = flag.String("cer", "", "TLS certificate, optional")
	certKey = flag.String("ck", "", "TLS certificate key, optional")
	m       = flag.String("m", "n", "y/n to turn on/off the customer XML marshaler")
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.SetPrefix("🍀  ")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s wsdl_file [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

	wp := os.Args[len(os.Args)-1]
	if *outFile == wp {
		log.Fatalln("Output file cannot be the same WSDL file")
	}

	var files []string
	for _, f := range os.Args {
		if !regexp.MustCompile(`\w+.(xml|wsdl|xsd|asmx)`).MatchString(f) {
			continue
		}
		files = append(files, f)
	}

	// generate code
	g, err := parser.NewGenerator(files, *pkg, *cert, *certKey, strings.ToLower(*m) == "y")
	if err != nil {
		log.Fatalln(err)
	}

	// Write .go source code file
	w := os.Stdout
	if *outFile != "" {
		w, err = os.OpenFile(*outFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0744)
		if err != nil {
			log.Fatalln(err)
		}
	}

	g.Write(w)

	log.Println("Done 💩")
}
