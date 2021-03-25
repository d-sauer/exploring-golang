package main

import (
	"fmt"
	"os"

	"github.com/antchfx/xmlquery"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Can't read file")
		os.Exit(1)
	}

	doc, err := xmlquery.Parse(file)
	if err != nil {
		fmt.Println("Can't parse XML")
		os.Exit(1)
	}

	// parsed := xmlquery.Find(doc, "//*:Attribute[@Name='https://aws.amazon.com/SAML/Attributes/Role']")
	list := xmlquery.Find(doc, "//ns1:Attribute[@Name='https://aws.amazon.com/SAML/Attributes/Role']")

	for i, n := range list {
		fmt.Printf("#%d  %s\n", i, n.InnerText())
	}
}
