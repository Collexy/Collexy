package xml_models

import (
	"encoding/xml"
	"fmt"
	//"encoding/json"
	//"github.com/clbanning/mxj"
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"database/sql"
	"log"
	"strconv"
	"sync"
	//coremodulesettingsmodels "collexy/core/modules/settings/models"
)

type Template struct {
	XMLName  xml.Name `xml:"template"`
	Id       *int     `xml:"id,omitempty"`
	Path     string   `xml:"path,omitempty"`
	ParentId *int     `xml:"parentId,omitempty"`
	Name     string   `xml:"name"`
	Alias    string   `xml:"alias"`
	// Parent string			`xml:"parent"`
	IsPartial bool        `xml:"isPartial"`
	Content   string      `xml:"content,omitempty"`
	Children  []*Template `xml:"children>template,omitempty"`
}

func Walk(templates []*Template, f func(*Template) bool) {
	for _, n := range templates {
		if f(n) {
			Walk(n.Children, f)
		}
	}
}

func GetTemplateById(id int) (template Template) {

	db := coreglobals.Db
	querystr := `SELECT path, parent_id, name, alias, is_partial  
	FROM template 
WHERE id=$1`

	// node
	// var created_by int
	var path, name, alias string
	// var created_date *time.Time
	// var is_partial bool
	// var parent_templates []byte
	var parent_id sql.NullInt64
	var is_partial bool

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&path, &parent_id, &name, &alias, &is_partial)

	var pid int

	if parent_id.Valid {
		// use s.String
		pid = int(parent_id.Int64)
	} else {
		// NULL value
	}

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No template with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		template = Template{xml.Name{}, &id, path, &pid, name, alias, is_partial, "", nil}
	}

	return
}

func (this *Template) Post(parentTemplate *Template) {
	// do DB POST
	db := coreglobals.Db

	if parentTemplate != nil {
		if *parentTemplate.Id == 0 {
			parentTemplate.Id = nil
		} else {
			this.ParentId = parentTemplate.Id
		}
	} else {
		parentTemplate = &Template{}
		if this.ParentId != nil {
			if *this.ParentId != 0 {
				parentTemplate.Id = this.ParentId
			}
			
		}
	}
	// if parentTemplate != nil {
	// 	if *parentTemplate.Id == 0 {
	// 		parentTemplate.Id = nil
	// 	} else {
	// 		this.ParentId = parentTemplate.Id
	// 	}
	// } else {
	// 	parentTemplate = &Template{}
	// }


	c1 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sqlStr := `INSERT INTO template (name, alias, parent_id, is_partial, created_by) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
		err1 := db.QueryRow(sqlStr, this.Name, this.Alias, parentTemplate.Id, this.IsPartial, -1).Scan(&id)
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

	if this.ParentId != nil {

		// Channel c, is for getting the parent template
		// We need to append the id of the newly created template to the path of the parent id to create the new path
		c2 := make(chan *Template)

		var wg2 sync.WaitGroup

		wg2.Add(1)

		go func() {
			defer wg2.Done()
			tpl := GetTemplateById(*this.ParentId)
			c2 <- &tpl
		}()

		go func() {
			for i := range c2 {
				fmt.Println(i)
				parentTemplate = i
			}
		}()

		wg2.Wait()
	}

	// fmt.Println(parentTemplate.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE template 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if parentTemplate.Id != nil && *parentTemplate.Id != 0 {
		path = parentTemplate.Path + "." + strconv.FormatInt(id, 10)
	}

	fmt.Println(path)
	fmt.Println(id)

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	log.Println("content created successfully")

	// return the id of the newly inserted template (now this.ParentId)
	if this.Children != nil {
		if len(this.Children) > 0 {
			for _, c := range this.Children {
				c.Post(this)
			}
		} else {
			//
		}
	} else {
		//
	}
}
