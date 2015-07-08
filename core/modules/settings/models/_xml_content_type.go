package models

import (
	"encoding/xml"
	//"fmt"
	"github.com/clbanning/mxj"
)

func (m *ContentType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName     xml.Name `xml:"contentType"`
		Path        string   `xml:path,omitempty`
		Name        string   `xml:"name"`
		Alias       string   `xml:"alias"`
		Description string   `xml:"description"`
		Icon        string   `xml:"icon"`
		Thumbnail   string   `xml:"thumbnail"`
		Meta        struct {
			Inner []byte `xml:",innerxml"`
		} `xml:"meta"`
		Tabs                  []Tab          `xml:"tabs"`
		AllowAtRoot           bool           `xml:"allowAtRoot"`
		IsContainer           bool           `xml:"isContainer"`
		IsAbstract            bool           `xml:"isAbstract"`
		AllowedContentTypes   []string       `xml:"allowedContentTypes>contentType,omitempty"`
		CompositeContentTypes []string       `xml:"compositeContentTypes>contentType,omitempty"`
		Template              string         `xml:"template"`
		AllowedTemplates      []string       `xml:"allowedTemplates>template,omitempty"`
		Children              []*ContentType `xml:"children>contentType,omitempty"`
	}

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	m.Id = nil
	m.Path = v.Path
	m.Name = v.Name
	m.Alias = v.Alias
	m.CreatedBy = -1
	m.CreatedDate = nil
	m.Description = v.Name
	m.Icon = v.Name
	m.Thumbnail = v.Name
	myMap := make(map[string]interface{})

	// ... do the mxj magic here ... -

	temp := v.Meta.Inner

	prefix := "<meta>"
	postfix := "</meta>"
	str := prefix + string(temp) + postfix
	//fmt.Println(str)
	myMxjMap, err := mxj.NewMapXml([]byte(str), true)
	myMap = myMxjMap

	// fill myMap
	m.Meta = myMap
	m.Tabs = v.Tabs

	m.AllowAtRoot = v.AllowAtRoot
	m.IsContainer = v.IsContainer
	m.IsAbstract = v.IsAbstract

	m.AllowedTemplateIds = nil
	m.CompositeContentTypeIds = nil
	m.CompositeContentTypes = nil
	m.TemplateId = nil
	m.AllowedTemplateIds = nil

	// allowedcontenttypes, compositecontenttypes, template, allowedtemplates, children

	return nil
}
