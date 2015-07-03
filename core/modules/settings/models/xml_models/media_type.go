package xml_models

import (
// 	coreglobals "collexy/core/globals"
// 	corehelpers "collexy/core/helpers"
// 	"database/sql"
// 	"encoding/json"
"encoding/xml"
// 	"fmt"
// 	"github.com/clbanning/mxj"
// 	"log"
// 	"strconv"
// 	"sync"
// 	//coremodulesettingsmodels "collexy/core/modules/settings/models"
)

type MediaType struct {
	XMLName xml.Name `xml:"mediaType"`
	Name    string   `xml:"name"`
	Alias   string   `xml:"alias"`
	// Parent string			`xml:"parent"`
	ParentId              *int                   `xml:"parentId"`
	Icon                  string                 `xml:"icon"`
	Thumbnail             string                 `xml:"thumbnail"`
	Meta                  map[string]interface{} `xml:"meta"`
	Tabs                  map[string]interface{} `xml:"tabs"`
	AllowAtRoot           bool                   `xml:"allowAtRoot"`
	IsContainer           bool                   `xml:"isContainer"`
	IsAbstract            bool                   `xml:"isAbstract"`
	AllowedMediaTypes     []string               `xml:"allowedMediaTypes"`
	AllowedMediaTypesIds  []int                  `xml:"allowedMediaTypeIds"`
	CompositeMediaTypes   []string               `xml:"compositeMediaTypes"`
	CompositeMediaTypeIds []int                  `xml:"compositeMediaTypeIds"`
	Children              []*MediaType           `xml:"children"`
}