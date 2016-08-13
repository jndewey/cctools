package cctools

import (
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

// this will include utilities to assist with working with .docx files

type WordDocumentXML struct {
	Plane	string	`xml:"Plane"`
}

func ParseXML(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var doc WordDocumentXML
	doc.Plane = "default"
	xml.Unmarshal(b, &doc)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(doc.Plane)
}




