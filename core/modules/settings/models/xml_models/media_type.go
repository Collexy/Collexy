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

// type MapContainer struct {
// 	InnerXML []byte `xml:",innerxml"`
// }

type MediaType struct {
	XMLName xml.Name `xml:"mediaType"`
	Id      *int
	Path    string `xml:path,omitempty`
	Name    string `xml:"name"`
	Alias   string `xml:"alias"`
	// Parent string			`xml:"parent"`
	ParentId              *int          `xml:"parentId"`
	Description           string        `xml:"description"`
	Icon                  string        `xml:"icon"`
	Thumbnail             string        `xml:"thumbnail"`
	MetaByte              *MapContainer `xml:"meta,omitempty"` //map[string]interface{} `xml:"meta,omitempty"`
	Meta                  map[string]interface{}
	Tabs                  []Tab    `xml:"tabs>tab,omitempty"`
	AllowAtRoot           bool     `xml:"allowAtRoot"`
	IsContainer           bool     `xml:"isContainer"`
	IsAbstract            bool     `xml:"isAbstract"`
	AllowedMediaTypes     []string `xml:"allowedMediaTypes>mediaType,omitempty"`
	AllowedMediaTypeIds   []int
	CompositeMediaTypes   []string `xml:"compositeMediaTypes>mediaType,omitempty"`
	CompositeMediaTypeIds []int
	Children              []*MediaType `xml:"children>mediaType,omitempty"`
}

func (this *MediaType) DoCustomXMLParsing() (res map[string]interface{}) {
	prefix := "<meta>"
	postfix := "</meta>"
	if this.MetaByte.InnerXML != nil && len(this.MetaByte.InnerXML) > 0 {
		str := prefix + string(this.MetaByte.InnerXML) + postfix
		m, err := mxj.NewMapXml([]byte(str), true)
		fmt.Println("this.MetaByte.InnerXML")
		fmt.Println(this.MetaByte.InnerXML)
		fmt.Println("string(this.MetaByte.InnerXML)")
		fmt.Println(string(this.MetaByte.InnerXML))
		if err != nil {
			log.Println("ERROR")
			log.Println(err.Error())
		}
		this.Meta = m["meta"].(map[string]interface{})
		res = m["meta"].(map[string]interface{})
	}

	return
}

func GetMediaTypeById(id int) (mediaType MediaType) {
	querystr := `SELECT media_type.id as media_type_id, media_type.path as media_type_path, 
    media_type.parent_id as media_type_parent_id, media_type.name as media_type_name, 
    media_type.alias as member_alias, media_type.description as media_type_description, 
    media_type.icon as media_type_icon, media_type.thumbnail as media_type_thumbnail, media_type.meta as media_type_meta, 
    media_type.tabs as media_type_tabs 
        FROM media_type
        WHERE media_type.id=$1`

	var media_type_id int
	var media_type_path, media_type_name, media_type_alias string
	var media_type_description, media_type_icon, media_type_thumbnail string

	var media_type_parent_id sql.NullInt64

	var media_type_tabs, media_type_meta []byte

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&media_type_id, &media_type_path, &media_type_parent_id, &media_type_name, &media_type_alias,
		&media_type_description, &media_type_icon, &media_type_thumbnail, &media_type_meta, &media_type_tabs)

	var parent_media_type_id int
	if media_type_parent_id.Valid {
		parent_media_type_id = int(media_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var tabs []Tab
	var media_type_metaMap map[string]interface{}

	json.Unmarshal(media_type_tabs, &tabs)
	json.Unmarshal(media_type_meta, &media_type_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No node with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		mediaType = MediaType{xml.Name{}, &media_type_id, media_type_path, media_type_name, media_type_alias, &parent_media_type_id, media_type_description, media_type_icon, media_type_thumbnail, nil, media_type_metaMap, tabs, false, false, false, nil, nil, nil, nil, nil}
	}

	return
}

var existingMediaTypes []MediaType

func (this *MediaType) Post(parentMediaType *MediaType, parentMediaTypes []MediaType, dataTypes []DataType) {

	db := coreglobals.Db
	if parentMediaType != nil {
		if *parentMediaType.Id == 0 {
			parentMediaType.Id = nil
		} else {
			this.ParentId = parentMediaType.Id
		}
	} else {
		parentMediaType = &MediaType{}
	}

	fmt.Printf("this.Name: %s\n", this.Name)
	fmt.Printf("this.Alias: %s\n", this.Alias)
	fmt.Printf("parentMediaType.Id: %v\n", parentMediaType.Id)
	fmt.Printf("this.ParentId: %d\n", &this.ParentId)

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
		// 			if prop.DataType != nil{
		// 				if dt.Alias == prop.DataType {
		// 					prop.DataTypeId = &dt.Id
		// 				}
		// 			}

		// 		}
		// 	}
		// }
		j, _ := json.Marshal(this.Tabs)
		tabs = j
	}

	c1 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sqlStr := `INSERT INTO media_type (parent_id, name, alias, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, 
            is_abstract, created_by) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
		err1 := db.QueryRow(sqlStr, parentMediaType.Id, this.Name, this.Alias,
			this.Description, this.Icon, this.Thumbnail, meta, tabs, this.AllowAtRoot,
			this.IsContainer, this.IsAbstract, -1).Scan(&id)
		corehelpers.PanicIf(err1)
		c1 <- int(id)
	}()

	go func() {
		for i := range c1 {
			fmt.Println("lolcat")
			fmt.Println(i)
			this.Id = &i
		}
	}()

	wg1.Wait()

	if this.ParentId != nil {

		// Channel c, is for getting the parent MediaType
		// We need to append the id of the newly created MediaType to the path of the parent id to create the new path
		c2 := make(chan *MediaType)

		var wg2 sync.WaitGroup

		wg2.Add(1)

		go func() {
			defer wg2.Done()
			tpl := GetMediaTypeById(*this.ParentId)
			c2 <- &tpl
		}()

		go func() {
			for i := range c2 {
				fmt.Println(i)
				parentMediaType = i
			}
		}()

		wg2.Wait()
	}

	// myCopy := *this

	parentMediaTypes = append(parentMediaTypes, *this)
	if parentMediaTypes != nil && len(parentMediaTypes) > 0 {
		for _, lol := range parentMediaTypes {
			var alreadyExists bool = false
			if existingMediaTypes != nil && len(existingMediaTypes) > 0 {
				for _, hehe := range existingMediaTypes {
					if lol.Alias == hehe.Alias {
						alreadyExists = true
						break
					}
				}
			}
			if !alreadyExists {
				existingMediaTypes = append(existingMediaTypes, lol)
			}
		}
	}

	// THIS SHOULD BE DEFERRED UNTIL THE END, perhaps in a separate update func,
	// since all content types need to be created before we can check for composites
	// OR maybe composites should just be defined first?
	// then what about 1 to 1 relationships or n to n?

	if existingMediaTypes != nil && len(existingMediaTypes) > 0 {
		for _, c := range existingMediaTypes {
			if this.CompositeMediaTypes != nil && len(this.CompositeMediaTypes) > 0 {
				for _, cc := range this.CompositeMediaTypes {
					if c.Alias == cc {
						if c.Id != nil && *c.Id != 0 {
							var alreadyExists bool = false
							for _, cctid := range this.CompositeMediaTypeIds {
								if cctid == *c.Id {
									alreadyExists = true
									break
								}
							}
							if !alreadyExists {
								id := c.Id
								this.CompositeMediaTypeIds = append(this.CompositeMediaTypeIds, *id)
							}

						}

					}
				}
			}
			if this.AllowedMediaTypes != nil && len(this.AllowedMediaTypes) > 0 {
				for _, cc := range this.AllowedMediaTypes {
					if c.Alias == cc {
						if c.Id != nil && *c.Id != 0 {
							var alreadyExists bool = false
							for _, actid := range this.AllowedMediaTypeIds {
								if actid == *c.Id {
									alreadyExists = true
									break
								}
							}
							if !alreadyExists {
								id := c.Id
								this.AllowedMediaTypeIds = append(this.AllowedMediaTypeIds, *id)
							}

						}

					}
				}
			}
		}
	} else {
		existingMediaTypes = append(existingMediaTypes, *this)
	}

	existingMediaTypes = append(existingMediaTypes, *this)

	allowedMediaTypeIds, err3 := coreglobals.IntSlice(this.AllowedMediaTypeIds).Value()
	corehelpers.PanicIf(err3)
	compositeMediaTypeIds, err4 := coreglobals.IntSlice(this.CompositeMediaTypeIds).Value()
	corehelpers.PanicIf(err4)

	fmt.Printf("parentMediaTypes: %v\n", parentMediaTypes)
	fmt.Printf("existingMediaTypes: %v\n", existingMediaTypes)
	fmt.Printf("this.CompositeMediaTypes: %v\n", this.CompositeMediaTypes)
	fmt.Printf("this.CompositeMediaTypeIds: %v\n", this.CompositeMediaTypeIds)

	// fmt.Println(parentMediaType.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE media_type 
    SET path=$1, allowed_media_type_ids=$2, composite_media_type_ids=$3 
    WHERE id=$4`

	path := strconv.FormatInt(id, 10)
	if parentMediaType.Id != nil && *parentMediaType.Id != 0 {
		path = parentMediaType.Path + "." + strconv.FormatInt(id, 10)
	}

	fmt.Println(path)
	fmt.Println(id)

	_, err6 := db.Exec(sqlStr, path, allowedMediaTypeIds, compositeMediaTypeIds, id)
	corehelpers.PanicIf(err6)

	log.Println("content created successfully")

	//myCopy := *this

	// return the id of the newly inserted MediaType (now this.ParentId)
	if this.Children != nil {
		if len(this.Children) > 0 {
			for _, c := range this.Children {
				c.Post(this, parentMediaTypes, dataTypes)
			}
		} else {
			//
		}
	} else {
		//
	}

}
