package kit

import "encoding/xml"

type School struct {
	XMLName  xml.Name  `xml:"school"`
	Name     string    `xml:"name"`
	Students []Student `xml:"students>student"`
}

type Student struct {
	XMLName xml.Name `xml:"student"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}
