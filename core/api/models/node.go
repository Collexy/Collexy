package models

import (
  "fmt"
  "net/url"
  "encoding/json"
  "time"
  corehelpers "collexy/core/helpers"
  //"collexy/globals"
  coreglobals "collexy/core/globals"
  //"log"
  //"strconv"
  "database/sql"
)

type Node struct {
  Id int `json:"id,omitempty"`
  Path string `json:"path,omitempty"`
  CreatedBy int `json:"created_by,omitempty"`
  Name string `json:"name,omitempty"`
  NodeType int `json:"node_type,omitempty"`
  CreatedDate *time.Time `json:"created_date,omitempty"`
  ParentId int `json:"parent_id,omitempty"`
  ParentNodes []Node `json:"parent_nodes,omitempty"`
  ChildNodes []Node `json:"child_nodes,omitempty"`
  Show bool `json:"show,omitempty"`
  OldName string `json:"old_name,omitempty"`
  UserPermissions []PermissionsContainer `json:"user_permissions,omitempty"`
  UserGroupPermissions []PermissionsContainer `json:"user_group_permissions,omitempty"`
  Icon string `json:"icon,omitempty"`
  //UserGroupPermissions []map[string]interface{} `json:"user_group_permissions,omitempty"`
  // UserGroupPermissions []Permission `json:"user_group_permissions,omitempty"`
}

// func TempGetUserAllowedNodes(user *User) (nodes []Node){
//   db := coreglobals.Db
//   queryStr := `SELECT DISTINCT ON (my_node.id) my_node.*
// FROM user_group AS my_user_group,
// LATERAL
// (
//   SELECT node.*, elem->'permissions' AS user_group_node_permissions
//   FROM node
//   LEFT OUTER JOIN
//   jsonb_array_elements(my_user_group.node_permissions) elem
//   ON elem->>'id' = node.id::text
//   --ORDER BY node.id
// )my_node
// WHERE (my_user_group.id = ANY('{2,3}')) 
// AND (user_group_node_permissions @> '12' OR (user_group_node_permissions IS NULL AND 12 = ANY(my_user_group.default_node_permissions)));`

//   rows, err := db.Query(queryStr)
//   corehelpers.PanicIf(err)
//   defer rows.Close()

//   var id, created_by, node_type int
//   var path, name string
//   var created_date time.Time
//   var user_group_node_permissions []PermissionsContainer //globals.IntSlice // []map[string]interface{}

//   for rows.Next(){
//       err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date, &user_group_node_permissions)
//       corehelpers.PanicIf(err)
//       node := Node{id, path, created_by, name, node_type, &created_date, 0, nil,nil,false, "",user_group_node_permissions}
//       nodes = append(nodes,node)
//   }
//   return
// }

type PermissionsContainer struct {
  Id int `json:"id"`
  Permissions coreglobals.StringSlice `json:"permissions"` //map[string]struct{} `json:"permissions"` 
} 

func GetNodes(queryStringParams url.Values, user *User) (nodes []Node){

  // test
  // defer corehelpers.Un(helpers.Trace("SOME_ARBITRARY_STRING_SO_YOU_CAN_KEEP_TRACK"))

  // test, _ := json.Marshal(user)
  // fmt.Println(string(test))

	db := coreglobals.Db
  sqlStr := ""
  if(queryStringParams.Get("node-type")=="1" || queryStringParams.Get("node-type")=="2"){
    sqlStr = `SELECT node.id, node.path, node.created_by, node.name, node.node_type, node.created_date, node.user_permissions, node.user_group_permissions, content_type.icon as content_type_icon FROM node
     JOIN content ON node.id = content.node_id
     JOIN content_type ON content.content_type_node_id = content_type.id`
  } else {
    sqlStr = `SELECT node.id, node.path, node.created_by, node.name, node.node_type, node.created_date, node.user_permissions, node.user_group_permissions FROM node`
  }
  // if ?node-type=x&levels=x(,x..)
  // else if ?node-type=x
  // else if ?levels=x(,x..)
  if(queryStringParams.Get("node-type") != "" && queryStringParams.Get("levels") != ""){
      sqlStr = sqlStr + ` WHERE node_type=` + queryStringParams.Get("node-type") + ` and node.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
  } else if(queryStringParams.Get("node-type") != "" && queryStringParams.Get("levels")==""){
      sqlStr = sqlStr + ` WHERE node_type=` + queryStringParams.Get("node-type")
  } else if(queryStringParams.Get("node-type") == "" && queryStringParams.Get("levels") != ""){
      sqlStr = sqlStr + ` WHERE node.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
  }

  if((queryStringParams.Get("node-type")=="1" || queryStringParams.Get("node-type")=="2") && queryStringParams.Get("content-type")!=""){
    sqlStr = sqlStr + ` and content_type_node_id=` + queryStringParams.Get("content-type")
  }
  
  rows, err := db.Query(sqlStr)
  corehelpers.PanicIf(err)
  defer rows.Close()

  var id, created_by, node_type int
  var path, name string
  var created_date time.Time
  var user_permissions, user_group_permissions []byte
  var user_perm, user_group_perm []PermissionsContainer // map[string]PermissionsContainer

  //var userId = strconv.Itoa(user.Id)
  var content_type_icon sql.NullString

  for rows.Next(){
      var content_type_icon_str string

      if(queryStringParams.Get("node-type")=="1" || queryStringParams.Get("node-type")=="2"){
        err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date, &user_permissions, &user_group_permissions, &content_type_icon)
        corehelpers.PanicIf(err)
        if(content_type_icon.Valid){
        content_type_icon_str = content_type_icon.String
        } else {
          // NULL
        }
      } else {
        err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date, &user_permissions, &user_group_permissions)
        corehelpers.PanicIf(err)
      }

      
      
      //fmt.Sprintf("Node id is: %v", id)
      user_perm = nil
      user_group_perm = nil
      json.Unmarshal(user_permissions, &user_perm)
      json.Unmarshal(user_group_permissions, &user_group_perm)

      var accessGranted bool = false
      var accessDenied bool = false

      // if(err1 != nil){
      //   log.Println("Unmarshal Error: " + err1.Error())
      //   user_perm = nil
      // }

      // if permissions are set on the node for a specific user
      if(user_permissions != nil){
        for i := 0; i < len(user_perm); i++ {
          if(accessGranted){
            break
          }
          if(user_perm[i].Id == user.Id){
            if(accessGranted){
              break
            }
            for j := 0; j < len(user_perm[i].Permissions); j++ {
              if(accessGranted){
                break
              }
              if(user_perm[i].Permissions[j] == "node_browse"){
                //fmt.Println("woauw it worked!")
                accessGranted = true
                node := Node{id, path, created_by, name, node_type, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
                nodes = append(nodes,node)
                break
              }
            }
            if(!accessGranted){
              accessDenied = true;
            }
          }
        } 
      } 
      if(!accessGranted && !accessDenied){
        // if no specific user node access has been specified, check node access per user_group
        if(user_group_permissions != nil){
          for i:= 0; i< len(user.UserGroupIds); i++ {
            if(accessGranted){
              break
            }
            for j := 0; j < len(user_group_perm); j++ {
              if(accessGranted){
                break
              }
              if(user_group_perm[j].Id == user.UserGroupIds[i]){
                if(accessGranted){
                  break
                }
                for k := 0; k < len(user_group_perm[j].Permissions); k++ {
                  if(accessGranted){
                    break
                  }
                  if(user_group_perm[j].Permissions[k] == "node_browse"){
                    //fmt.Println("woauw it worked!")
                    accessGranted = true
                    node := Node{id, path, created_by, name, node_type, &created_date, 0, nil,nil,false, "", nil, user_group_perm, content_type_icon_str}
                    nodes = append(nodes,node)
                    break
                  }
                }
                if(!accessGranted){
                  accessDenied = true;
                }
              }
            } 
          }
        }
      }

      // if no specific access has been granted per user_group either, use user groups default permissions
      if(!accessGranted && !accessDenied){
        if(user.UserGroups != nil){
          for i:= 0; i< len(user.UserGroups); i++ {
            if(accessGranted){
              break
            }
            for j:= 0; j< len(user.UserGroups[i].Permissions); j++ {
              if(user.UserGroups[i].Permissions[j] == "node_browse"){
                accessGranted = true
                node := Node{id, path, created_by, name, node_type, &created_date, 0, nil,nil,false, "", nil, nil, content_type_icon_str}
                nodes = append(nodes,node)
                break
              }
            }
            
          }
        }
        
      }
  }
  return
}

// func GetNodes(queryStringParams url.Values) (nodes []Node){

//   db := coreglobals.Db
//   sql := `SELECT id, path, created_by, name, node_type, created_date FROM node`

//   // if ?node-type=x&levels=x(,x..)
//   // else if ?node-type=x
//   // else if ?levels=x(,x..)
//   if(queryStringParams.Get("node-type") != "" && queryStringParams.Get("levels") != ""){
//       sql = sql + ` WHERE node_type=` + queryStringParams.Get("node-type") + ` and node.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
//   } else if(queryStringParams.Get("node-type") != "" && queryStringParams.Get("levels")==""){
//       sql = sql + ` WHERE node_type=` + queryStringParams.Get("node-type")
//   } else if(queryStringParams.Get("node-type") == "" && queryStringParams.Get("levels") != ""){
//       sql = sql + ` WHERE node.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
//   }
  
//   rows, err := db.Query(sql)
//   corehelpers.PanicIf(err)
//   defer rows.Close()

//   var id, created_by, node_type int
//   var path, name string
//   var created_date time.Time

//   for rows.Next(){
//       err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date)
//       corehelpers.PanicIf(err)
//       node := Node{id, path, created_by, name, node_type, &created_date, 0, nil,nil,false, "", nil, nil}
//       nodes = append(nodes,node)
//   }
//   return
// }


func GetNodeById(id int) (node Node){
	db := coreglobals.Db
  querystr := `SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`

  row := db.QueryRow(querystr,id)

  var created_by, node_type int
  var path, name string
  var created_date time.Time

  err := row.Scan(&id, &path, &created_by, &name, &node_type, &created_date)
  corehelpers.PanicIf(err)
  node = Node{id, path, created_by, name, node_type, &created_date, 0, nil, nil, false, "", nil, nil, ""}
    
    return
}

func GetNodeByIdChildren(id int) (nodes []Node){
  fmt.Println("GETNODEIDBYCHILDREN")
  db := coreglobals.Db

  //querystr := "SELECT id, path, created_by, name, node_type, created_date FROM node WHERE parent_id=$1"

  querystr := `SELECT node.id, node.path, node.created_by, node.name, node.node_type, node.created_date, content_type.icon
FROM node 
LEFT OUTER JOIN content
ON content.node_id = node.id
LEFT OUTER JOIN content_type
ON content.content_type_node_id = content_type.node_id
WHERE parent_id=$1`

  rows, err := db.Query(querystr, id)
  corehelpers.PanicIf(err)
  defer rows.Close()

  var created_by, node_type int
  var path, name string
  var created_date time.Time

  var content_type_icon sql.NullString

  for rows.Next(){
    var content_type_icon_str string

    

      err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date, &content_type_icon)
      corehelpers.PanicIf(err)

      if(content_type_icon.Valid){
        content_type_icon_str = content_type_icon.String
      } else {
        // NULL
      }

      node := Node{id,path,created_by, name, node_type, &created_date, 0, nil, nil, false, "", nil, nil, content_type_icon_str}

      nodes = append(nodes, node)

  }
  return
}