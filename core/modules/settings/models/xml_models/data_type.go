package xml_models

import
(
	"encoding/xml"
)

type DataType struct {
	XMLName xml.Name `xml:"dataType"`
	Id int	 `xml"dataTypeId,omitempty`
	Name    string   `xml:"name"`
	
}