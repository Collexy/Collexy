package xml_models

import (
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/clbanning/mxj"
	"log"
	"strconv"
	"sync"
	//coremodulesettingsmodels "collexy/core/modules/settings/models"
)

type MapContainer struct {
	InnerXML []byte `xml:",innerxml"`
}

type ContentType struct {
	XMLName xml.Name `xml:"contentType"`
	Id      *int
	Path    string `xml:path,omitempty`
	Name    string `xml:"name"`
	Alias   string `xml:"alias"`
	// Parent string			`xml:"parent"`
	ParentId                *int          `xml:"parentId"`
	Description             string        `xml:"description"`
	Icon                    string        `xml:"icon"`
	Thumbnail               string        `xml:"thumbnail"`
	MetaByte                *MapContainer `xml:"meta,omitempty"` //map[string]interface{} `xml:"meta,omitempty"`
	Meta                    map[string]interface{}
	Tabs                    []Tab    `xml:"tabs"`
	AllowAtRoot             bool     `xml:"allowAtRoot"`
	IsContainer             bool     `xml:"isContainer"`
	IsAbstract              bool     `xml:"isAbstract"`
	AllowedContentTypes     []string `xml:"allowedContentTypes>contentType,omitempty"`
	AllowedContentTypeIds   []int
	CompositeContentTypes   []string `xml:"compositeContentTypes>contentType,omitempty"`
	CompositeContentTypeIds []int
	Template                string `xml:"template"`
	TemplateId              *int
	AllowedTemplates        []string `xml:"allowedTemplates>template,omitempty"`
	AllowedTemplateIds      []int
	Children                []*ContentType `xml:"children>contentType,omitempty"`
}

func (this *ContentType) DoCustomXMLParsing() (res map[string]interface{}) {
	prefix := "<meta>"
	postfix := "</meta>"
	if this.MetaByte.InnerXML != nil && len(this.MetaByte.InnerXML) > 0 {
		str := prefix + string(this.MetaByte.InnerXML) + postfix
		m, err := mxj.NewMapXml([]byte(str))
		fmt.Println("this.MetaByte.InnerXML")
		fmt.Println(this.MetaByte.InnerXML)
		fmt.Println("string(this.MetaByte.InnerXML)")
		fmt.Println(string(this.MetaByte.InnerXML))
		if err != nil {
			log.Println("ERROR")
			log.Println(err.Error())
		}
		this.Meta = m
		res = m
	}

	return
}

func GetContentTypeById(id int) (contentType ContentType) {
	querystr := `SELECT content_type.id as content_type_id, content_type.path as content_type_path, 
    content_type.parent_id as content_type_parent_id, content_type.name as content_type_name, 
    content_type.alias as member_alias, content_type.description as content_type_description, 
    content_type.icon as content_type_icon, content_type.thumbnail as content_type_thumbnail, content_type.meta as content_type_meta, 
    content_type.tabs as content_type_tabs 
        FROM content_type
        WHERE content_type.id=$1`

	var content_type_id int
	var content_type_path, content_type_name, content_type_alias string
	var content_type_description, content_type_icon, content_type_thumbnail string

	var content_type_parent_id sql.NullInt64

	var content_type_tabs, content_type_meta []byte

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&content_type_id, &content_type_path, &content_type_parent_id, &content_type_name, &content_type_alias,
		&content_type_description, &content_type_icon, &content_type_thumbnail, &content_type_meta, &content_type_tabs)

	var parent_content_type_id int
	if content_type_parent_id.Valid {
		parent_content_type_id = int(content_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var tabs []Tab
	var content_type_metaMap map[string]interface{}

	json.Unmarshal(content_type_tabs, &tabs)
	json.Unmarshal(content_type_meta, &content_type_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No node with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		contentType = ContentType{xml.Name{}, &content_type_id, content_type_path, content_type_name, content_type_alias, &parent_content_type_id, content_type_description, content_type_icon, content_type_thumbnail, nil, content_type_metaMap, tabs, false, false, false, nil, nil, nil, nil, "", nil, nil, nil, nil}
	}

	return
}

var existingContentTypes []ContentType

func (this *ContentType) Post(parentContentType *ContentType, parentContentTypes []ContentType, templates []*Template) {
	fmt.Printf("len(templates) is: %d\n", len(templates))
	for i, t := range templates {
		fmt.Printf("t.Alias is: %s (i: %d)\n", t.Alias, i)
		if this.Template == t.Alias {
			this.TemplateId = t.Id
		}
		for _, at := range this.AllowedTemplates {
			if t.Alias == at {
				this.AllowedTemplateIds = append(this.AllowedTemplateIds, *t.Id)
			}
		}
	}

	db := coreglobals.Db
	if parentContentType != nil {
		if *parentContentType.Id == 0 {
			parentContentType.Id = nil
		} else {
			this.ParentId = parentContentType.Id
		}
	} else {
		parentContentType = &ContentType{}
	}

	fmt.Printf("this.Name: %s\n", this.Name)
	fmt.Printf("this.Alias: %s\n", this.Alias)
	fmt.Printf("parentContentType.Id: %v\n", parentContentType.Id)
	fmt.Printf("this.ParentId: %d\n", &this.ParentId)
	fmt.Printf("this.Template: %s\n", this.Template)
	fmt.Printf("this.TemplateId: %v\n", this.TemplateId)

	fmt.Printf("this.AllowedTemplateIds: %v\n", this.AllowedTemplateIds)

	var meta interface{} = nil
	var tabs interface{} = nil

	if this.MetaByte.InnerXML != nil {
		mmap := this.DoCustomXMLParsing()

		if mmap != nil {
			j, _ := json.Marshal(mmap)
			meta = j
		}
	}

	if this.Tabs != nil {
		// for _, tab := range this.Tabs {
		// 	for _, prop := range tab.Properties{
		// 		for _, dt := range dataTypes {
		// 			if dt.Id == prop.DataTypeId {
		// 				prop.DataType = dt
		// 			}
		// 		}
		// 	}
			
		// }
		j, _ := json.Marshal(this.Tabs)
		tabs = j
	}

	allowedTemplateIds, err5 := coreglobals.IntSlice(this.AllowedTemplateIds).Value()
	corehelpers.PanicIf(err5)

	c2 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sqlStr := `INSERT INTO content_type (parent_id, name, alias, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, 
            is_abstract, template_id, allowed_template_ids) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
		err1 := db.QueryRow(sqlStr, parentContentType.Id, this.Name, this.Alias,
			this.Description, this.Icon, this.Thumbnail, meta, tabs, this.AllowAtRoot,
			this.IsContainer, this.IsAbstract,
			this.TemplateId, allowedTemplateIds).Scan(&id)
		corehelpers.PanicIf(err1)
		c2 <- int(id)
	}()

	go func() {
		for i := range c2 {
			fmt.Println("lolcat")
			fmt.Println(i)
			this.Id = &i
		}
	}()

	wg1.Wait()

	if this.ParentId != nil {

		// Channel c, is for getting the parent ContentType
		// We need to append the id of the newly created ContentType to the path of the parent id to create the new path
		c2 := make(chan *ContentType)

		var wg2 sync.WaitGroup

		wg2.Add(1)

		go func() {
			defer wg2.Done()
			tpl := GetContentTypeById(*this.ParentId)
			c2 <- &tpl
		}()

		go func() {
			for i := range c2 {
				fmt.Println(i)
				parentContentType = i
			}
		}()

		wg2.Wait()
	}

	// myCopy := *this

	parentContentTypes = append(parentContentTypes, *this)
	if parentContentTypes != nil && len(parentContentTypes) > 0 {
		for _, lol := range parentContentTypes {
			var alreadyExists bool = false
			if existingContentTypes != nil && len(existingContentTypes) > 0 {
				for _, hehe := range existingContentTypes {
					if lol.Alias == hehe.Alias {
						alreadyExists = true
						break
					}
				}
			}
			if !alreadyExists {
				existingContentTypes = append(existingContentTypes, lol)
			}
		}
	}

	// THIS SHOULD BE DEFERRED UNTIL THE END, perhaps in a separate update func,
	// since all content types need to be created before we can check for composites
	// OR maybe composites should just be defined first?
	// then what about 1 to 1 relationships or n to n?

	if existingContentTypes != nil && len(existingContentTypes) > 0 {
		for _, c := range existingContentTypes {
			if this.CompositeContentTypes != nil && len(this.CompositeContentTypes) > 0 {
				for _, cc := range this.CompositeContentTypes {
					if c.Alias == cc {
						if c.Id != nil && *c.Id != 0 {
							var alreadyExists bool = false
							for _, cctid := range this.CompositeContentTypeIds {
								if cctid == *c.Id {
									alreadyExists = true
									break
								}
							}
							if !alreadyExists {
								id := c.Id
								this.CompositeContentTypeIds = append(this.CompositeContentTypeIds, *id)
							}

						}

					}
				}
			}
			if this.AllowedContentTypes != nil && len(this.AllowedContentTypes) > 0 {
				for _, cc := range this.AllowedContentTypes {
					if c.Alias == cc {
						if c.Id != nil && *c.Id != 0 {
							var alreadyExists bool = false
							for _, actid := range this.AllowedContentTypeIds {
								if actid == *c.Id {
									alreadyExists = true
									break
								}
							}
							if !alreadyExists {
								id := c.Id
								this.AllowedContentTypeIds = append(this.AllowedContentTypeIds, *id)
							}

						}

					}
				}
			}
		}
	} else {
		existingContentTypes = append(existingContentTypes, *this)
	}

	existingContentTypes = append(existingContentTypes, *this)

	allowedContentTypeIds, err3 := coreglobals.IntSlice(this.AllowedContentTypeIds).Value()
	corehelpers.PanicIf(err3)
	compositeContentTypeIds, err4 := coreglobals.IntSlice(this.CompositeContentTypeIds).Value()
	corehelpers.PanicIf(err4)

	fmt.Printf("parentContentTypes: %v\n", parentContentTypes)
	fmt.Printf("existingContentTypes: %v\n", existingContentTypes)
	fmt.Printf("this.CompositeContentTypes: %v\n", this.CompositeContentTypes)
	fmt.Printf("this.CompositeContentTypeIds: %v\n", this.CompositeContentTypeIds)

	// fmt.Println(parentContentType.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE content_type 
    SET path=$1, allowed_content_type_ids=$2, composite_content_type_ids=$3 
    WHERE id=$4`

	path := strconv.FormatInt(id, 10)
	if parentContentType.Id != nil && *parentContentType.Id != 0 {
		path = parentContentType.Path + "." + strconv.FormatInt(id, 10)
	}

	fmt.Println(path)
	fmt.Println(id)

	_, err6 := db.Exec(sqlStr, path, allowedContentTypeIds, compositeContentTypeIds, id)
	corehelpers.PanicIf(err6)

	log.Println("content created successfully")

	//myCopy := *this

	// return the id of the newly inserted ContentType (now this.ParentId)
	if this.Children != nil {
		if len(this.Children) > 0 {
			for _, c := range this.Children {
				c.Post(this, parentContentTypes, templates)
			}
		} else {
			//
		}
	} else {
		//
	}

}
