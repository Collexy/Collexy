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

type Media struct {
	XMLName     xml.Name `xml:"mediaItem"`
	Id          *int
	Path        string
	ParentId    *int          `xml:"parentId,omitempty"`
	Name        string        `xml:"name"`
	Parent      string        `xml:"parent,omitempty"`
	MediaType   string        `xml:"mediaType,omitempty"`
	MediaTypeId *int          `xml:"mediaTypeId,omitempty"`
	MetaByte    *MapContainer `xml:"meta,omitempty"`
	Meta        map[string]interface{}
	Children    []*Media `xml:"children>mediaItem,omitempty"`
}

func (this *Media) DoCustomXMLParsing() (res map[string]interface{}) {
	prefix := "<meta>"
	postfix := "</meta>"
	if this.MetaByte != nil {
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
	}

	return
}

func GetMediaById(id int) (media Media) {

	db := coreglobals.Db
	querystr := `SELECT path, parent_id, name, media_type_id, meta   
	FROM media 
WHERE id=$1`

	// node
	// var created_by int
	var path, name string
	// var created_date *time.Time
	// var is_partial bool
	// var parent_templates []byte
	var parent_id sql.NullInt64
	var media_type_id int

	var meta []byte

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&path, &parent_id, &name, &media_type_id, &meta)

	var pid int

	if parent_id.Valid {
		// use s.String
		pid = int(parent_id.Int64)
	} else {
		// NULL value
	}

	var media_metaMap map[string]interface{}

	json.Unmarshal(meta, &media_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No template with that ID.")
	case err != nil:
		log.Fatal(err)
		//panic(err)
	default:
		media = Media{xml.Name{}, &id, path, &pid, name, "", "", &media_type_id, nil, media_metaMap, nil}
	}

	return
}

func (this *Media) Post(parentMedia *Media, mediaTypes []*coremodulesettingsmodels.MediaType) {
	// do DB POST
	db := coreglobals.Db
	if parentMedia != nil {
		if *parentMedia.Id == 0 {
			parentMedia.Id = nil
		} else {
			this.ParentId = parentMedia.Id
		}
	} else {
		parentMedia = &Media{}
	}

	fmt.Println(this.Name)

	fmt.Println(parentMedia.Id)
	fmt.Println(&this.ParentId)

	if this.MediaTypeId == nil || *this.MediaTypeId == 0 {
		for i, ct := range mediaTypes {
			fmt.Printf("mediatypes is: %s (i: %d)\n", ct.Alias, i)
			if this.MediaType == ct.Alias {
				this.MediaTypeId = &ct.Id
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
		sqlStr := `INSERT INTO media (name, parent_id, media_type_id, meta, created_by) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err1 := db.QueryRow(sqlStr, this.Name, parentMedia.Id, this.MediaTypeId,
			meta, -1).Scan(&id)
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
		c2 := make(chan *Media)

		var wg2 sync.WaitGroup

		wg2.Add(1)

		go func() {
			defer wg2.Done()

			tpl := GetMediaById(*this.ParentId)
			c2 <- &tpl
		}()

		go func() {
			for i := range c2 {
				fmt.Println(i)
				parentMedia = i
			}
		}()

		wg2.Wait()
	}

	// fmt.Println(parentTemplate.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE media 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if parentMedia.Id != nil && *parentMedia.Id != 0 {
		path = parentMedia.Path + "." + strconv.FormatInt(id, 10)
	}
	this.Path = path
	fmt.Println(path)
	fmt.Println(id)

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	log.Println("media created successfully")

	// return the id of the newly inserted template (now this.ParentId)
	if this.Children != nil {
		if len(this.Children) > 0 {
			for _, c := range this.Children {
				c.Post(this, mediaTypes)
			}
		} else {
			//
		}
	} else {
		//
	}
}
