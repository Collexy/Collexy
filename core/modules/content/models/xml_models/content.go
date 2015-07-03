package xml_models

import
(
	"encoding/xml"
)

type Content struct {
	XMLName       xml.Name               `xml:"content"`
	Name          string                 `xml:"name"`
	ParentId      string                 `xml:parentId`
	ContentType   string                 `xml:"contentType"`
	ContentTypeId *int                   `xml:"contentTypeId"`
	Template      string                 `xml:"template"`
	TemplateId    *int                   `xml:"templateId`
	Meta          map[string]interface{} `xml:"meta"`
	Children      []*Content             `xml:"children"`
}