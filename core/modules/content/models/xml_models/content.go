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
	//coremodulesettingsmodelsxmlmodels "collexy/core/modules/settings/models/xml_models"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
)

type MapContainer struct {
	InnerXML []byte `xml:",innerxml"`
}

type Content struct {
	XMLName       xml.Name `xml:"contentItem"`
	Id            *int
	Path          string
	ParentId      *int          `xml:"parentId,omitempty"`
	Name          string        `xml:"name"`
	Parent        string        `xml:"parent,omitempty"`
	ContentType   string        `xml:"contentType,omitempty"`
	ContentTypeId *int          `xml:"contentTypeId,omitempty"`
	Template      string        `xml:"template,omitempty"`
	TemplateId    *int          `xml:"templateId,omitempty"`
	MetaByte      *MapContainer `xml:"meta,omitempty"`
	Meta          map[string]interface{}
	Children      []*Content `xml:"children>contentItem,omitempty"`
}

func (this *Content) DoCustomXMLParsing() (res map[string]interface{}) {
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
		//this.Meta = m
		this.Meta = m["meta"].(map[string]interface{})
		res = m["meta"].(map[string]interface{})
	}

	return
}

func GetContentById(id int) (content Content) {
	fmt.Printf("is is &v: ", id)
	db := coreglobals.Db
	querystr := `SELECT path, parent_id, name, content_type_id, template_id, meta   
	FROM content 
WHERE id=$1`

	// node
	// var created_by int
	var path, name string
	// var created_date *time.Time
	// var is_partial bool
	// var parent_templates []byte
	var parent_id, template_id sql.NullInt64
	var content_type_id int

	var meta []byte

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&path, &parent_id, &name, &content_type_id, &template_id, &meta)

	var pid int

	if parent_id.Valid {
		// use s.String
		pid = int(parent_id.Int64)
	} else {
		// NULL value
	}
	var tid int
	if template_id.Valid {
		// use s.String
		tid = int(template_id.Int64)
	} else {
		// NULL value
	}

	var content_metaMap map[string]interface{}

	json.Unmarshal(meta, &content_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No template with that ID.")
	case err != nil:
		log.Fatal(err)
		//panic(err)
	default:
		content = Content{xml.Name{}, &id, path, &pid, name, "", "", &content_type_id, "", &tid, nil, content_metaMap, nil}
	}

	return
}

func (this *Content) Post(parentContent *Content, contentTypes []*coremodulesettingsmodels.ContentType, templates []*coremodulesettingsmodels.Template) {
	// do DB POST
	db := coreglobals.Db

	
	if parentContent != nil {
		if *parentContent.Id == 0 {
			parentContent.Id = nil
		} else {
			this.ParentId = parentContent.Id
		}
	} else {
		parentContent = &Content{}
		if this.ParentId != nil {
			if *this.ParentId != 0 {
				parentContent.Id = this.ParentId
			}
			
		}
	}
	

	// if this.ParentId == nil || *this.ParentId == 0 {
	// 	if parentContent != nil {
	// 		if *parentContent.Id == 0 {
	// 			parentContent.Id = nil
	// 		} else {
	// 			this.ParentId = parentContent.Id
	// 		}
	// 	} else {
	// 		parentContent = &Content{}
	// 	}
	// } else {
	// 	if parentContent == nil {
	// 		parentContent = &Content{}
	// 	}
	// 	parentContent.Id = this.ParentId
	// }

	if this.TemplateId == nil || *this.TemplateId == 0 {
		for i, t := range templates {
			fmt.Printf("templates is: %s (i: %d)\n", t.Alias, i)
			if this.Template == t.Alias {
				this.TemplateId = &t.Id
				break
			}
		}
	}
	if this.ContentTypeId == nil || *this.ContentTypeId == 0 {
		for i, ct := range contentTypes {
			fmt.Printf("contenttypes is: %s (i: %d)\n", ct.Alias, i)
			if this.ContentType == ct.Alias {
				this.ContentTypeId = &ct.Id
				break
			}
		}
	}

	var meta interface{} = nil

	if this.MetaByte.InnerXML != nil {
		mmap := this.DoCustomXMLParsing()

		if mmap != nil {
			j, _ := json.Marshal(mmap)
			meta = j
		}
	}

	c1 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sqlStr := `INSERT INTO content (name, parent_id, content_type_id, template_id, meta, created_by) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		err1 := db.QueryRow(sqlStr, this.Name, parentContent.Id, this.ContentTypeId,
			this.TemplateId, meta, -1).Scan(&id)
		corehelpers.PanicIf(err1)
		c1 <- int(id)
	}()

	go func() {
		for i := range c1 {
			fmt.Println(i)
			this.Id = &i
		}
	}()

	wg1.Wait()

	if this.ParentId != nil && *this.ParentId != 0 {

		// Channel c, is for getting the parent template
		// We need to append the id of the newly created template to the path of the parent id to create the new path
		c2 := make(chan *Content)

		var wg2 sync.WaitGroup

		wg2.Add(1)

		go func() {
			defer wg2.Done()

			tpl := GetContentById(*this.ParentId)
			c2 <- &tpl
		}()

		go func() {
			for i := range c2 {
				fmt.Println(i)
				parentContent = i
			}
		}()

		wg2.Wait()
	}

	// fmt.Println(parentTemplate.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE content 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if parentContent.Id != nil && *parentContent.Id != 0 {
		path = parentContent.Path + "." + strconv.FormatInt(id, 10)
	}
	this.Path = path
	fmt.Println(path)
	fmt.Println(id)

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	log.Println("content created successfully")

	// return the id of the newly inserted template (now this.ParentId)
	if this.Children != nil {
		if len(this.Children) > 0 {
			for _, c := range this.Children {
				c.Post(this, contentTypes, templates)
			}
		} else {
			//
		}
	} else {
		//
	}
}
