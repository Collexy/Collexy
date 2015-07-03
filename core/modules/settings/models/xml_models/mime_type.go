package xml_models

import
(
	"encoding/xml"
)

type MimeType struct {
	XMLName     xml.Name `xml:"mimeType"`
	Name        string   `xml:"name"`
	MediaType   string   `xml:"mediaType"`
	MediaTypeId *int     `xml:"mediaTypeId"`
}