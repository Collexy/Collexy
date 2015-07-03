package xml_models

import
(
	"encoding/xml"
)

type Media struct {
	XMLName     xml.Name               `xml:"media"`
	Name        string                 `xml:"name"`
	ParentId    string                 `xml:parentId`
	MediaType   string                 `xml:"mediaType"`
	MediaTypeId *int                   `xml:"mediaTypeId"`
	Meta        map[string]interface{} `xml:"meta"`
	Children    []*Media               `xml:"children"`
}