package xml_models

import (
	"encoding/xml"
)

type Tab struct {
	//Id int `json:"id"`
	XMLName  xml.Name `xml:"tab" json:"-"`
	Name       string         `xml:"name" json:"name"`
	Properties []*TabProperty `xml:"properties>property,omitempty" json:"properties,omitempty"`
}
