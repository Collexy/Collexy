package models

import (
  "encoding/json"
  corehelpers "collexy/core/helpers"
  coreglobals "collexy/core/globals"
  "time"
  "fmt"
  //"net/http"
  //"html/template"
  "strconv"
  "database/sql"
  "log"
)

type DataType struct {
  //Id int `json:"id"`
  Id int `json:"id,omitempty"`
  NodeId int `json:"node_id,omitempty"`
  Alias string `json:"alias,omitempty"`
  Html string `json:"html,omitempty"`
  Node *Node `json:"node,omitempty"`
}

func (t *DataType) Post(){
  tm, err := json.Marshal(t)
  corehelpers.PanicIf(err)
  fmt.Println("tm:::: ")
  fmt.Println(string(tm))

  db := coreglobals.Db

  tx, err := db.Begin()
  corehelpers.PanicIf(err)
  //defer tx.Rollback()
  

  // http://godoc.org/github.com/lib/pq
  // pq does not support the LastInsertId() method of the Result type in database/sql. 
  // To return the identifier of an INSERT (or UPDATE or DELETE), 
  // use the Postgres RETURNING clause with a standard Query or QueryRow call:
  
  var node_id int64
  err = db.QueryRow(`INSERT INTO node (name, node_type, created_by) VALUES ($1, $2, $3) RETURNING id`, t.Node.Name, 11, 1).Scan(&node_id)
  //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateNodeId)
  //helpers.PanicIf(err)
  //node_id, err := res.LastInsertId()
  fmt.Println(strconv.FormatInt(node_id, 10))
  if err != nil {
    //log.Println(string(res))
    log.Fatal(err.Error())
  } else {
    _, err = tx.Exec("UPDATE node SET path=$1 WHERE id=$2", "1." + strconv.FormatInt(node_id, 10), node_id)
    corehelpers.PanicIf(err)
    //println("LastInsertId:", node_id)
  }
  //defer r1.Close()

  _, err = tx.Exec("INSERT INTO data_type (node_id, alias, html) VALUES ($1, $2, $3)", node_id, t.Alias, t.Html)
  corehelpers.PanicIf(err)
  //defer r2.Close()
  err1 := tx.Commit()
  corehelpers.PanicIf(err1)
}

func GetDataTypes() (dataTypes []DataType) {
  
  db := coreglobals.Db

    querystr := `SELECT 
    node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
  dt.id as data_type_id, dt.alias as data_type_alias, dt.node_id as data_type_node_id, dt.html as data_type_html
  FROM node 
  JOIN data_type as dt
  ON dt.node_id = node.id
  WHERE node.node_type=11`

  // node
  var node_id, node_created_by, node_type int
  var node_path, node_name string
  var node_created_date time.Time

  // data type
  var data_type_id, data_type_node_id int
  var data_type_alias sql.NullString
  var data_type_html []byte

  rows, err := db.Query(querystr)
    corehelpers.PanicIf(err)
    defer rows.Close()

    for rows.Next(){
      err:= rows.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &data_type_id, &data_type_alias, &data_type_node_id, &data_type_html)

      var data_type_alias_string string
      if data_type_alias.Valid {
        // use s.String
        data_type_alias_string = data_type_alias.String
      } else {
         // NULL value
      }

      switch {
          case err == sql.ErrNoRows:
              log.Printf("No node with that ID.")
          case err != nil:
              log.Fatal(err)
          default:
            node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil}
            dataType := DataType{data_type_id, node_id, data_type_alias_string, string(data_type_html), &node}
            dataTypes = append(dataTypes,dataType)
      }
    }

  return
}

func GetDataTypeByNodeId(nodeId int) (dt DataType) {
	
	db := coreglobals.Db

    querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
	dt.id as data_type_id, dt.alias as data_type_alias, dt.node_id as data_type_node_id, dt.html as data_type_html
	FROM node 
	JOIN data_type as dt
	ON dt.node_id = node.id
	WHERE node.id=$1`

  // node
  var node_id, node_created_by, node_type int
  var node_path, node_name string
  var node_created_date time.Time

  // data type
  var data_type_id, data_type_node_id int
  var data_type_alias sql.NullString
  var data_type_html []byte

  row := db.QueryRow(querystr, nodeId)

  err:= row.Scan(
      &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
      &data_type_id, &data_type_alias, &data_type_node_id, &data_type_html)

  var data_type_alias_string string
  if data_type_alias.Valid {
    // use s.String
    data_type_alias_string = data_type_alias.String
  } else {
     // NULL value
  }

  switch {
      case err == sql.ErrNoRows:
          log.Printf("No node with that ID.")
      case err != nil:
          log.Fatal(err)
      default:
      	node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil}
  		  dt = DataType{data_type_id, node_id, data_type_alias_string, string(data_type_html), &node}
  }

  return


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
}

func (dt *DataType) Update(){
  db := coreglobals.Db

  tx, err := db.Begin()
  corehelpers.PanicIf(err)
  //defer tx.Rollback()

  _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", dt.Node.Name, dt.Node.Id)
  corehelpers.PanicIf(err)
  //defer r1.Close()

  _, err = tx.Exec("UPDATE data_type SET alias = $1, html = $2 WHERE node_id = $3", dt.Alias, dt.Html, dt.Node.Id)
  corehelpers.PanicIf(err)
  //defer r2.Close()
  tx.Commit()
}



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