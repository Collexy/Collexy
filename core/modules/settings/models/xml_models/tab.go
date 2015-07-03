package xml_models

import (
//"encoding/xml"
)

type Tab struct {
	//Id int `json:"id"`
	Name       string         `xml:"name" json:"name"`
	Properties []*TabProperty `xml:"properties>property,omitempty" json:"properties,omitempty"`
}
