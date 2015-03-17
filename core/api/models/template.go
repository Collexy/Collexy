package models

import (
  //"fmt"
  "encoding/json"
  corehelpers "collexy/core/helpers"
  coreglobals "collexy/core/globals"
  "time"
  "strconv"
  "database/sql"
  "log"
  "io/ioutil"
  "path/filepath"
  "database/sql/driver"
  //"encoding/binary"
  // "reflect"
   "fmt"
  "strings"
  "os"
)

type Template struct {
  Id int `json:"id,omitempty"`
  NodeId int `json:"node_id,omitempty"`
  Alias string `json:"alias,omitempty"`
  ParentTemplateNodeId int `json:"parent_template_node_id,omitempty"`
  Html string `json:"html,omitempty"`
  PartialTemplates []Template `json:"partial_templates,omitempty"`
  PartialTemplateNodes []Node `json:"partial_template_nodes,omitempty"`
  PartialTemplateNodeIds IntArray `json:"partial_template_node_ids,omitempty"`
  IsPartial bool `json:"is_partial,omitempty"`
  Node *Node `json:"node,omitempty"`
}

func (t *Template) Post(){
  tm, err := json.Marshal(t)
  corehelpers.PanicIf(err)
  fmt.Println("tm:::: ")
  fmt.Println(string(tm))
  db := coreglobals.Db

  tx, err := db.Begin()
  corehelpers.PanicIf(err)
  //defer tx.Rollback()
  var parentNode Node
  var id, created_by, node_type int
  var path, name string
  var created_date *time.Time
  err = tx.QueryRow(`SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`, t.ParentTemplateNodeId).Scan(&id, &path, &created_by, &name, &node_type, &created_date)
  switch {
    case err == sql.ErrNoRows:
      log.Printf("No user with that ID.")
    case err != nil:
      log.Fatal(err)
    default:
      parentNode = Node{id, path, created_by, name, node_type, created_date, 0, nil,nil, false, "", nil, nil, ""}
      //fmt.Printf("Username is %s\n", username)
  }

  // http://godoc.org/github.com/lib/pq
  // pq does not support the LastInsertId() method of the Result type in database/sql. 
  // To return the identifier of an INSERT (or UPDATE or DELETE), 
  // use the Postgres RETURNING clause with a standard Query or QueryRow call:
  
  var node_id int64
  err = db.QueryRow(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`, t.Node.Name, 3, 1, t.ParentTemplateNodeId).Scan(&node_id)
  //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateNodeId)
  //helpers.PanicIf(err)
  //node_id, err := res.LastInsertId()
  fmt.Println(strconv.FormatInt(node_id, 10))
  if err != nil {
    //log.Println(string(res))
    log.Fatal(err.Error())
  } else {
    _, err = tx.Exec("UPDATE node SET path=$1 WHERE id=$2", parentNode.Path + "." + strconv.FormatInt(node_id, 10), node_id)
    corehelpers.PanicIf(err)
    //println("LastInsertId:", node_id)
  }
  //defer r1.Close()

  _, err = tx.Exec("INSERT INTO template (node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES ($1, $2, $3, $4, $5)", node_id, t.Alias, false, t.PartialTemplateNodeIds, t.ParentTemplateNodeId)
  corehelpers.PanicIf(err)
  //defer r2.Close()
  err1 := tx.Commit()
  corehelpers.PanicIf(err1)

  tplNodeName := t.Node.Name + ".tmpl"
  absPath, _ := filepath.Abs("/views/" + tplNodeName)

  // write whole the body - maybe use bufio/os/io packages for buffered read/write on big files
  err = ioutil.WriteFile(absPath, []byte(t.Html), 0644)
  if err != nil {
      panic(err)
  }
}


func (t *Template) Update(){
  db := coreglobals.Db

  tx, err := db.Begin()
  corehelpers.PanicIf(err)
  //defer tx.Rollback()

  _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", t.Node.Name, t.Node.Id)
  corehelpers.PanicIf(err)
  //defer r1.Close()

  fmt.Println("partial template node ids (array): ")
  fmt.Println(t.PartialTemplateNodeIds)

  fmt.Println("partial template node ids (postgres format): ")
  partial_template_node_ids_pgs_format, _ := t.PartialTemplateNodeIds.Value()
  fmt.Println(partial_template_node_ids_pgs_format)

  _, err = tx.Exec(`UPDATE template SET alias = $1, parent_template_node_id = $2, partial_template_node_ids = $3 WHERE node_id = $4`, t.Alias, t.ParentTemplateNodeId, partial_template_node_ids_pgs_format, t.Node.Id)
  corehelpers.PanicIf(err)
  //defer r2.Close()
  err1 := tx.Commit()
  corehelpers.PanicIf(err1)

  name := t.Node.Name + ".tmpl"
  absPath, _ := filepath.Abs("/views/" + name)

  // write whole the body - maybe use bufio/os/io packages for buffered read/write on big files
  err = ioutil.WriteFile(absPath, []byte(t.Html), 0644)
  if err != nil {
      panic(err)
  }
}

/*
TODO: Fetch node for each parent template - for use in aliasOrNode in template edit controller
*/
  type IntArray []int 

func (b *IntArray) Scan(src interface{}) error { 
        switch src := src.(type) { 
        case nil: 
                *b = nil 
                return nil 

        case []byte: 
                // TODO: parse src into *b
          var intArr []int
          intArrString := string(src)
          intArrString = strings.Replace(intArrString, "{", "", -1)
          intArrString = strings.Replace(intArrString, "}", "", -1)
          var lol []string
          lol = strings.Split(intArrString, ",")
          for i := 0; i < len(lol); i++ {
            someval, _ := strconv.Atoi(lol[i])
             intArr = append(intArr, someval)
          }
          *b = intArr

        default: 
                return fmt.Errorf(`unsupported driver -> Scan pair: %T -> *IntArray`, src) 
        }
        return nil
}

func (b IntArray) Value() (driver.Value, error) {
  var str string = "{"
  var myarr []int = b
  fmt.Println("driver.Value 1: ")
  fmt.Println(b)
  for i := 0; i < len(myarr); i++ {
    str = str + strconv.Itoa(myarr[i])
    if(i<len(myarr)-1){
      str = str+","
    }
  }
  str = str + "}"
  //fmt.Println("driver.Value 2: ")
  //fmt.Println(str)
  return str, nil
  //return "{23,24}", nil
  //return "20,21", nil
    // Format b in PostgreSQL's array input format {1,2,3} and return it as as string or []byte.
    // if(b == nil){
    //   return nil, nil
    // } else if(len(*b)>0){
    //   var str string = "{"
    //   for i := 0; i < len(*b); i++ {
    //     str = str + string(*b[i])
    //     if(i<len(b-1)){
    //       str = str+", "
    //     }
    //   }
    //   str = str+"}"
    //   return str
    //   } else {
    //         return fmt.Errorf(`unsupported driver -> Scan pair: %T -> *IntArray`, src) 
    //   }
    //   return nil
}

func GetTemplates() (templates []Template) {
  
  db := coreglobals.Db

    querystr := `SELECT 
    node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
  t.id as template_id, t.node_id as template_node_id, t.alias as template_alias, t.partial_template_node_ids as partial_template_node_ids, parent_template_node_id as parent_template_node_id, t.is_partial as template_is_partial 
  FROM node 
  JOIN template as t
  ON t.node_id = node.id
  WHERE node.node_type=3`

  // node
  var node_id, node_created_by, node_type int
  var node_path, node_name string
  var node_created_date time.Time

  // data type
  var template_id, template_node_id int
  var template_alias, parent_template_node_id sql.NullString
  var template_is_partial bool
  //var partial_template_node_ids sql.NullString
  var partial_template_node_ids IntArray



  rows, err := db.Query(querystr)
    corehelpers.PanicIf(err)
    defer rows.Close()

    for rows.Next(){
      err:= rows.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &template_id, &template_node_id, &template_alias, &partial_template_node_ids, &parent_template_node_id, &template_is_partial)

      var template_alias_string string

      if template_alias.Valid {
        // use s.String
        template_alias_string = template_alias.String
      } else {
         // NULL value
      }
      var parent_template_node_id_int int
      if parent_template_node_id.Valid {
        // use s.String
        parent_template_node_id_int, _ = strconv.Atoi(parent_template_node_id.String)
      } else {
         // NULL value
      }

      // var partial_templates_slice []Template
      // if partial_template_node_ids.Valid {
      //   // use s.String
      //   temp := []byte(partial_template_node_ids.String)
      //   err := json.Unmarshal(temp, &partial_templates_slice)
      //   corehelpers.PanicIf(err)
      // } else {
      //    // NULL value
      // }

      // b := make([]int, len(partial_template_node_ids))
      // for i, v := range partial_template_node_ids {
      //     b[i] = v.(int)
      // }

      switch {
          case err == sql.ErrNoRows:
              log.Printf("No node with that ID.")
          case err != nil:
              log.Fatal(err)
          default:
            node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil, ""}
            template := Template{template_id, node_id, template_alias_string, parent_template_node_id_int, "", nil, nil, partial_template_node_ids, template_is_partial, &node}
            templates = append(templates,template)
      }
    }

  return
}

func GetTemplateByNodeId(nodeId int) (template Template){

  db := coreglobals.Db
  querystr := `select my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date, 
    ffgd.parent_template_nodes,
template.id as template_id, template.node_id as template_node_id, template.alias as template_alias, template.parent_template_node_id as parent_template_node_id, template.partial_template_node_ids as template_partial_templates, template.is_partial as template_is_partial
from node as my_node,
LATERAL 
(
    SELECT array_to_json(array_agg(node)) as parent_template_nodes
    from node
    where path @> subpath(my_node.path,0,nlevel(my_node.path)-1) and node_type=3 
    order by my_node.path asc
) ffgd,
template
where my_node.id=$1 and template.node_id = my_node.id`

  // node
  var node_id, node_created_by, node_type int
  var node_path, node_name string
  var node_created_date time.Time
  var parent_template_nodes []byte

  // template
  var template_id, template_node_id int
  var parent_template_node_id sql.NullString
  var template_alias string
  var template_partial_templates IntArray
  var template_is_partial bool

  var template_parent_template_node_id int

  row := db.QueryRow(querystr, nodeId)

  err:= row.Scan(
      &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date, &parent_template_nodes,
      &template_id, &template_node_id, &template_alias, &parent_template_node_id, &template_partial_templates, &template_is_partial)

  if parent_template_node_id.Valid {
    // use s.String
    id, _ := strconv.Atoi(parent_template_node_id.String)
    template_parent_template_node_id = id
  } else {
     // NULL value
  }
  
  name := node_name + ".tmpl"
  //absPath, _ := filepath.Abs("/views/" + name)
  absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/views/" + name)
  //fmt.Println("FILEPATH:: " + absPath)
  
  bs, err7 := ioutil.ReadFile(absPath)
  corehelpers.PanicIf(err7)
  str := string(bs)

  //var tplSlice []Template
  var parentTemplateNodesSlice []Node

  //json.Unmarshal(template_partial_templates, &tplSlice)
  json.Unmarshal(parent_template_nodes, &parentTemplateNodesSlice)

  switch {
    case err == sql.ErrNoRows:
      log.Printf("No node with that ID.")
    case err != nil:
      log.Fatal(err)
    default:
      node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, parentTemplateNodesSlice, nil, false, "", nil, nil, ""}
      template = Template{template_id, template_node_id, template_alias, template_parent_template_node_id, str, nil, nil, template_partial_templates, template_is_partial, &node}
  }

  return
}