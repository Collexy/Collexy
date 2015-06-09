package models

import (
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"encoding/json"
	"time"
	//"fmt"
	//"net/http"
	//"html/template"
	"database/sql"
	"log"
	"strconv"
)

type DataType struct {
	Id          int                    `json:"id"`
	Path        string                 `json:"path"`
	ParentId    int                    `json:"parent_id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Alias       string                 `json:"alias,omitempty"`
	CreatedBy   int                    `json:"created_by,omitempty"`
	CreatedDate *time.Time             `json:"created_date,omitempty"`
	Html        string                 `json:"html,omitempty"`
	EditorAlias string                 `json:"editor_alias,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

func GetDataTypes() (dataTypes []*DataType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta 
        FROM data_type`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, created_by int
		var path, name, alias string
		var created_date *time.Time
		var parent_id sql.NullInt64
		var html, editor_alias sql.NullString
		var meta []byte

		if err := rows.Scan(&id, &path, &parent_id, &name, &alias, &created_by, &created_date, &html, &editor_alias, &meta); err != nil {
			log.Fatal(err)
		}

		var pid int

		if parent_id.Valid {
			pid = int(parent_id.Int64)
		}

		var html_str string

		if html.Valid {
			html_str = html.String
		}

		var editor_alias_str string

		if editor_alias.Valid {
			editor_alias_str = editor_alias.String
		}

		var data_type_metaMap map[string]interface{}

		json.Unmarshal(meta, &data_type_metaMap)

		dataType := &DataType{id, path, pid, name, alias, created_by, created_date, html_str, editor_alias_str, data_type_metaMap}
		dataTypes = append(dataTypes, dataType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetDataTypeById(id int) (dataType *DataType) {
	db := coreglobals.Db

	var created_by int
	var path, name, alias string
	var created_date *time.Time
	var parent_id sql.NullInt64
	var html, editor_alias sql.NullString
	var meta []byte

	err := db.QueryRow(`SELECT id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta
        FROM data_type WHERE id=$1`, id).Scan(&id, &path, &parent_id, &name, &alias, &created_by, &created_date, &html, &editor_alias, &meta)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No data type with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		var pid int

		if parent_id.Valid {
			pid = int(parent_id.Int64)
		}

		var html_str string

		if html.Valid {
			html_str = html.String
		}

		var editor_alias_str string

		if editor_alias.Valid {
			editor_alias_str = editor_alias.String
		}

		var data_type_metaMap map[string]interface{}

		json.Unmarshal(meta, &data_type_metaMap)

		dataType = &DataType{id, path, pid, name, alias, created_by, created_date, html_str, editor_alias_str, data_type_metaMap}
	}
	return
}

func (d *DataType) Post() {

	//meta, err := json.Marshal(d.Meta)
	//corehelpers.PanicIf(err)

	db := coreglobals.Db

	var id int64

	// sqlStr := `INSERT INTO data_type (name, alias, created_by, html, editor_alias, meta)
	// VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// err1 := db.QueryRow(sqlStr, d.Name, d.Alias, d.CreatedBy, d.Html, d.EditorAlias, meta).Scan(&id)
	sqlStr := `INSERT INTO data_type (name, alias, html, editor_alias) 
	VALUES ($1, $2, $3, $4) RETURNING id`
	err1 := db.QueryRow(sqlStr, d.Name, d.Alias, d.Html, d.EditorAlias).Scan(&id)

	corehelpers.PanicIf(err1)

	sqlStr = `UPDATE data_type 
	SET path=$1  
	WHERE id=$2`

	_, err2 := db.Exec(sqlStr, strconv.FormatInt(id, 10), id)

	corehelpers.PanicIf(err2)

	log.Println("data type created successfully")
}

func (d *DataType) Update() {

	meta, err := json.Marshal(d.Meta)
	corehelpers.PanicIf(err)

	db := coreglobals.Db

	sqlStr := `UPDATE data_type 
	SET path=$1, parent_id=$2, name=$3, alias=$4, 
	created_by=$5, html=$6, editor_alias=$7, meta=$8 
	WHERE id=$9`

	_, err1 := db.Exec(sqlStr, d.Path, d.ParentId, d.Name, d.Alias, d.CreatedBy, d.Html, d.EditorAlias, meta, d.Id)

	corehelpers.PanicIf(err1)

	log.Println("data type updated successfully")
}

func DeleteDataType(id int) {

	db := coreglobals.Db

	sqlStr := `delete FROM data_type 
	WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	corehelpers.PanicIf(err)

	log.Printf("data type with id %d was successfully deleted", id)
}

// func (t *DataType) Post() {
// 	// tm, err := json.Marshal(t)
// 	// corehelpers.PanicIf(err)
// 	// fmt.Println("tm:::: ")
// 	// fmt.Println(string(tm))

// 	db := coreglobals.Db

// 	tx, err := db.Begin()
// 	corehelpers.PanicIf(err)
// 	//defer tx.Rollback()

// 	// http://godoc.org/github.com/lib/pq
// 	// pq does not support the LastInsertId() method of the Result type in database/sql.
// 	// To return the identifier of an INSERT (or UPDATE or DELETE),
// 	// use the Postgres RETURNING clause with a standard Query or QueryRow call:

// 	var id int64
// 	err = db.QueryRow(`INSERT INTO node (name, node_type, created_by) VALUES ($1, $2, $3) RETURNING id`, t.Node.Name, 11, 1).Scan(&id)
// 	//res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateNodeId)
// 	//helpers.PanicIf(err)
// 	//id, err := res.LastInsertId()
// 	fmt.Println(strconv.FormatInt(id, 10))
// 	if err != nil {
// 		//log.Println(string(res))
// 		log.Fatal(err.Error())
// 	} else {
// 		_, err = tx.Exec("UPDATE node SET path=$1 WHERE id=$2", "1."+strconv.FormatInt(id, 10), id)
// 		corehelpers.PanicIf(err)
// 		//println("LastInsertId:", id)
// 	}
// 	//defer r1.Close()

// 	_, err = tx.Exec("INSERT INTO data_type (id, alias, html) VALUES ($1, $2, $3)", id, t.Alias, t.Html)
// 	corehelpers.PanicIf(err)
// 	//defer r2.Close()
// 	err1 := tx.Commit()
// 	corehelpers.PanicIf(err1)
// }

// func (t *DataType) Post(){
//   tm, err := json.Marshal(t)
//   corehelpers.PanicIf(err)
//   fmt.Println("tm:::: ")
//   fmt.Println(string(tm))

//   db := coreglobals.Db

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)
//   //defer tx.Rollback()

//   // http://godoc.org/github.com/lib/pq
//   // pq does not support the LastInsertId() method of the Result type in database/sql.
//   // To return the identifier of an INSERT (or UPDATE or DELETE),
//   // use the Postgres RETURNING clause with a standard Query or QueryRow call:

//   var node_id int64
//   err = db.QueryRow(`INSERT INTO node (name, node_type, created_by) VALUES ($1, $2, $3) RETURNING id`, t.Node.Name, 11, 1).Scan(&node_id)
//   //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateNodeId)
//   //helpers.PanicIf(err)
//   //node_id, err := res.LastInsertId()
//   fmt.Println(strconv.FormatInt(node_id, 10))
//   if err != nil {
//     //log.Println(string(res))
//     log.Fatal(err.Error())
//   } else {
//     _, err = tx.Exec("UPDATE node SET path=$1 WHERE id=$2", "1." + strconv.FormatInt(node_id, 10), node_id)
//     corehelpers.PanicIf(err)
//     //println("LastInsertId:", node_id)
//   }
//   //defer r1.Close()

//   _, err = tx.Exec("INSERT INTO data_type (node_id, alias, html) VALUES ($1, $2, $3)", node_id, t.Alias, t.Html)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()
//   err1 := tx.Commit()
//   corehelpers.PanicIf(err1)
// }

// // FILES STUFF NOT USED RIGHT NOW
// files_array := []DataTypeTemp{}

// absPathDir, _ := filepath.Abs("/admin/public/views/settings/data-type/tmpl")
// files, _ := ioutil.ReadDir(absPathDir)
// for _, f := range files {
//         //fmt.Println(f.Name())

//     //var filesMap map[string]interface{}
//     path := absPathDir + "\\" + f.Name()
//     bs, _ := ioutil.ReadFile(path)
//     html := string(bs)
//     //json_obj := fmt.Sprintf("{\"name\": \"%s\", \"path\": \"%s\", \"html\": \"%s\"}", f.Name(),path, html)
//     //json.Unmarshal([]byte(string(json_obj)), &filesMap)
//     dtt := DataTypeTemp{f.Name(), path, html}
//     files_array = append(files_array, dtt)

// }
// files_array_str, _ := json.Marshal(files_array)

// name := "text-input.html"
// absPath, _ := filepath.Abs("/admin/public/views/settings/data-type/tmpl/text-input.html")

// bs, err7 := ioutil.ReadFile(absPath)
// corehelpers.PanicIf(err7)
// str := string(bs)
//}

// func (dt *DataType) Update(){
//   db := coreglobals.Db

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)
//   //defer tx.Rollback()

//   _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", dt.Node.Name, dt.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r1.Close()

//   _, err = tx.Exec("UPDATE data_type SET alias = $1, html = $2 WHERE node_id = $3", dt.Alias, dt.Html, dt.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()
//   tx.Commit()
// }

// func (dt *DataType) MarshalJSONUpdate() ([]byte, error) {
//   // if dt.Id {
//   //   return json.Marshal(map[string]interface{}{
//   //       "id": dt.Id,
//   //   })
//   // }
//   if t.Id != 0 {
//     return json.Marshal(t.Id)
//   }

//   return json.Marshal(nil)
// }
