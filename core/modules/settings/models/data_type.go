package models

import (
	//"encoding/json"
	//corehelpers "collexy/core/helpers"
	coreglobals "collexy/core/globals"
	"time"
	//"fmt"
	//"net/http"
	//"html/template"
	//"strconv"
	"database/sql"
	"log"
)

type DataType struct {
	Id          int        `json:"id"`
	Path        string     `json:"path"`
	ParentId    int        `json:"parent_id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Alias       string     `json:"alias,omitempty"`
	CreatedBy   int        `json:"created_by,omitempty"`
	CreatedDate *time.Time `json:"created_date,omitempty"`
	Html        string     `json:"html,omitempty"`
}

func GetDataTypes() (dataTypes []*DataType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, path, parent_id, name, alias, created_by, created_date, html 
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
		var html sql.NullString

		if err := rows.Scan(&id, &path, &parent_id, &name, &alias, &created_by, &created_date, &html); err != nil {
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

		dataType := &DataType{id, path, pid, name, alias, created_by, created_date, html_str}
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
	var html sql.NullString

	err := db.QueryRow(`SELECT id, path, parent_id, name, alias, created_by, created_date, html 
        FROM data_type WHERE id=$1`, id).Scan(&id, &path, &parent_id, &name, &alias, &created_by, &created_date, &html)
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

		dataType = &DataType{id, path, pid, name, alias, created_by, created_date, html_str}
	}
	return
}

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