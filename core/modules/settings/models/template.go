package models

import (
  //"fmt"
  //"encoding/json"
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
  Id int `json:"id"`
  Path string `json:"path"`
  ParentId int `json:"parent_id,omitempty"`
  Name string `json:"name"`
  Alias string `json:"alias"`
  CreatedBy int `json:"created_by"`
  CreatedDate *time.Time `json:"created_date"`
  IsPartial bool `json:"is_partial,omitempty"`
  Html string `json:"html,omitempty"`
  ParentTemplates []*Template `json:"parent_templates,omitempty"`
}

func GetTemplates() (templates []*Template){
    db := coreglobals.Db

    rows, err := db.Query(`SELECT id, path, parent_id, name, alias, created_by, created_date, is_partial 
        FROM template`)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id, created_by int
        var path, name, alias string
        var created_date time.Time
        var parent_id sql.NullInt64
        var is_partial bool

        if err := rows.Scan(&id,&path,&parent_id,&name,&alias,&created_by,&created_date,&is_partial); err != nil {
            log.Fatal(err)
        }

        var pid int

        if parent_id.Valid {
            pid = int(parent_id.Int64)
        }

        template := &Template{id,path,pid,name,alias,created_by,&created_date,is_partial,"",nil}
        templates = append(templates, template)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    return
}


func GetTemplateById(id int) (template Template){

  db := coreglobals.Db
  querystr := `SELECT path, parent_id, name, alias, created_by, created_date, is_partial, 
    ffgd.parent_templates
FROM template as my_template,
LATERAL 
(
    SELECT array_to_json(array_agg(template)) as parent_templates
    from template
    where path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    order by my_template.path asc
) ffgd
where my_template.id=$1`

  // node
  var created_by int
  var path, name, alias string
  var created_date *time.Time
  var is_partial bool
  var parent_templates []byte
  var parent_id sql.NullInt64


  row := db.QueryRow(querystr, id)

  err:= row.Scan(
      &path,&parent_id,&name,&alias,&created_by,&created_date,&is_partial,&parent_templates)

  var pid int

  if parent_id.Valid {
    // use s.String
    pid = int(parent_id.Int64)
  } else {
     // NULL value
  }
  
  tplName := name + ".tmpl"
  //absPath, _ := filepath.Abs("/views/" + name)
  absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/views/" + tplName)
  //fmt.Println("FILEPATH:: " + absPath)
  
  bs, err7 := ioutil.ReadFile(absPath)
  corehelpers.PanicIf(err7)
  str := string(bs)

  //var tplSlice []Template
  //var parentTemplatesSlice []Node

  //json.Unmarshal(template_partial_templates, &tplSlice)
  //json.Unmarshal(parent_templates, &parentTemplatesSlice)

  switch {
    case err == sql.ErrNoRows:
      log.Printf("No template with that ID.")
    case err != nil:
      log.Fatal(err)
    default:
      template = Template{id,path,pid,name,alias,created_by,created_date,is_partial,str,nil}
  }

  return
}


// func (t *Template) Post(){
//   tm, err := json.Marshal(t)
//   corehelpers.PanicIf(err)
//   fmt.Println("tm:::: ")
//   fmt.Println(string(tm))
//   db := coreglobals.Db

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)
//   //defer tx.Rollback()
//   var parentNode Node
//   var id, created_by, node_type int
//   var path, name string
//   var created_date *time.Time
//   err = tx.QueryRow(`SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`, t.ParentTemplateId).Scan(&id, &path, &created_by, &name, &node_type, &created_date)
//   switch {
//     case err == sql.ErrNoRows:
//       log.Printf("No user with that ID.")
//     case err != nil:
//       log.Fatal(err)
//     default:
//       parentNode = Node{id, path, created_by, name, node_type, created_date, 0, nil,nil, false, "", nil, nil, ""}
//       //fmt.Printf("Username is %s\n", username)
//   }

//   // http://godoc.org/github.com/lib/pq
//   // pq does not support the LastInsertId() method of the Result type in database/sql. 
//   // To return the identifier of an INSERT (or UPDATE or DELETE), 
//   // use the Postgres RETURNING clause with a standard Query or QueryRow call:
  
//   var node_id int64
//   err = db.QueryRow(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`, t.Node.Name, 3, 1, t.ParentTemplateId).Scan(&node_id)
//   //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateId)
//   //helpers.PanicIf(err)
//   //node_id, err := res.LastInsertId()
//   fmt.Println(strconv.FormatInt(node_id, 10))
//   if err != nil {
//     //log.Println(string(res))
//     log.Fatal(err.Error())
//   } else {
//     _, err = tx.Exec("UPDATE node SET path=$1 WHERE id=$2", parentNode.Path + "." + strconv.FormatInt(node_id, 10), node_id)
//     corehelpers.PanicIf(err)
//     //println("LastInsertId:", node_id)
//   }
//   //defer r1.Close()

//   _, err = tx.Exec("INSERT INTO template (node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES ($1, $2, $3, $4, $5)", node_id, t.Alias, false, t.PartialTemplateIds, t.ParentTemplateId)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()
//   err1 := tx.Commit()
//   corehelpers.PanicIf(err1)

//   tplNodeName := t.Node.Name + ".tmpl"
//   //absPath, _ := filepath.Abs("/views/" + tplNodeName)
//   absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/views/" + tplNodeName)
  
//   // write whole the body - maybe use bufio/os/io packages for buffered read/write on big files
//   createTemplateErr := ioutil.WriteFile(absPath, []byte(t.Html), 0644)
//   if createTemplateErr != nil {
//       panic(createTemplateErr)
//   }
// }


// func (t *Template) Update(){
//   db := coreglobals.Db

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)
//   //defer tx.Rollback()

//   _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", t.Node.Name, t.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r1.Close()

//   fmt.Println("partial template node ids (array): ")
//   fmt.Println(t.PartialTemplateIds)

//   fmt.Println("partial template node ids (postgres format): ")
//   partial_template_node_ids_pgs_format, _ := t.PartialTemplateIds.Value()
//   fmt.Println(partial_template_node_ids_pgs_format)

//   _, err = tx.Exec(`UPDATE template SET alias = $1, parent_template_node_id = $2, partial_template_node_ids = $3 WHERE node_id = $4`, t.Alias, t.ParentTemplateId, partial_template_node_ids_pgs_format, t.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()
//   err1 := tx.Commit()
//   corehelpers.PanicIf(err1)

//   name := t.Node.Name + ".tmpl"
//   absPath, _ := filepath.Abs("/views/" + name)

//   // write whole the body - maybe use bufio/os/io packages for buffered read/write on big files
//   err = ioutil.WriteFile(absPath, []byte(t.Html), 0644)
//   if err != nil {
//       panic(err)
//   }
// }

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

