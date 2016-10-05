package cctools

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// this will include utilities to assist with working with .docx files

type Plane struct {
	Plane string `xml:"plane"`
}

func ParseXML(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	/*
		b, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	*/
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		ct := xml.CopyToken(t)
		fmt.Println(ct)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "plane" {
				var plane Plane
				decoder.DecodeElement(&plane, &se)
				fmt.Println(plane.Plane)
			}
		}
	}
}
