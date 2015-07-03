package xml_models

import (
//"encoding/xml"
)

type TabProperty struct {
	//Id int `json:"id"`
	Name        string `xml:"name" json:"name,omitempty"`
	Order       int    `xml:"order" json:"order,omitempty"`
	DataType 	*DataType `json:"data_type,omitempty"`
	DataTypeId  int    `xml:"dataTypeId" json:"data_type_id,omitempty"`
	HelpText    string `xml:"helpText" json:"help_text,omitempty"`
	Description string `xml:"description" json:"description,omitempty"`
}
