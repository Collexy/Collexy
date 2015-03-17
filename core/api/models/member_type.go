package models

import (
    //"fmt"
  "encoding/json"
  corehelpers "collexy/core/helpers"
  coreglobals "collexy/core/globals"
  "time"
  //"net/http"
  //"html/template"
  "strconv"
  "database/sql"
  "log"
)

type MemberType struct {
  Id int `json:"id,omitempty"`
  Node_id int `json:"node_id,omitempty"`
  Alias string `json:"alias,omitempty"`
  Description string `json:"description,omitempty"`
  Icon string `json:"icon,omitempty"`
  ParentMemberTypeNodeId int `json:"parent_member_type_node_id,omitempty"`
  Tabs []Tab `json:"tabs,omitempty"`
  Meta map[string]interface{} `json:"meta,omitempty"`
  ParentMemberTypes []MemberType `json:"parent_content_types,omitempty"`
  Node *Node `json:"node,omitempty"`
}

func GetMemberTypeExtendedByNodeId(nodeId int) (memberType MemberType){

  querystr := `SELECT my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date,
    res.id as mt_id, res.parent_member_type_node_id as mt_parent_member_type_node_id, res.alias as mt_alias,
    res.description as mt_description, res.icon as mt_icon, res.meta::json as mt_meta, res.mt_tabs as mt_tabs, res.parent_member_types as mt_parent_member_types
    FROM member_type 
    JOIN node as my_node 
    ON my_node.id = member_type.node_id 
    JOIN
    LATERAL
    (
      SELECT my_member_type.*,ffgd.*,gf2.*
      FROM member_type as my_member_type, node as my_member_type_node,
      LATERAL 
      (
          SELECT array_to_json(array_agg(okidoki)) as parent_member_types
          FROM (
            SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.parent_member_type_node_id, c.meta, gf.* as tabs
            FROM member_type as c, node,
          LATERAL (
              select json_agg(row1) as tabs from((
              select y.name, ss.properties
              from json_to_recordset(
            (
                select * 
                from json_to_recordset(
              (
                  SELECT json_agg(ggg)
                  from(
                SELECT tabs
                FROM 
                (   
                    SELECT *
                    FROM member_type as ct
                    WHERE ct.id=c.id
                ) dsfds

                  )ggg
              )
                ) as x(tabs json)
            )
              ) as y(name text, properties json),
              LATERAL (
            select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id',row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id',row.data_type_node_id, 'alias', row.data_type_alias,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
            from(
                select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
                from json_to_recordset(properties) 
                as k(name text, "order" int, data_type_node_id int, help_text text, description text)
                JOIN data_type
                ON data_type.node_id = k.data_type_node_id
                )row
              ) ss
              ))row1
          ) gf
            where path @> subpath(my_member_type_node.path,0,nlevel(my_member_type_node.path)-1) and c.node_id = node.id
          )okidoki
      ) ffgd,
      --
      LATERAL 
      (
          SELECT okidoki.tabs as mt_tabs
          FROM (
            SELECT c.id as cid, gf.* as tabs
            FROM member_type as c, node,
          LATERAL (
              select json_agg(row1) as tabs from((
          select y.name, ss.properties
          from json_to_recordset(
          (
        select * 
        from json_to_recordset(
            (
          SELECT json_agg(ggg)
          from(
        SELECT tabs
        FROM 
        (   
            SELECT *
            FROM member_type as ct
            WHERE ct.id=c.id
        ) dsfds

          )ggg
            )
        ) as x(tabs json)
          )
          ) as y(name text, properties json),
          LATERAL (
        select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id', row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id', row.data_type_node_id, 'alias', row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
        from(
      select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
      from json_to_recordset(properties) 
      as k(name text, "order" int, data_type_node_id int, help_text text, description text)
      JOIN data_type
      ON data_type.node_id = k.data_type_node_id
      )row
          ) ss
              ))row1
          ) gf
          WHERE c.id = my_member_type.id
          )okidoki
          limit 1
      ) gf2
      --
      WHERE my_member_type_node.id = my_member_type.node_id
    ) res
    ON res.node_id = member_type.node_id
    WHERE member_type.node_id=$1`

    // node
    var node_id, node_created_by, node_type int
    var node_path, node_name string
    var node_created_date time.Time

    var mt_id int
    var mt_parent_member_type_node_id sql.NullString

    var mt_alias, mt_description, mt_icon string
    var mt_tabs, mt_meta []byte
    var mt_parent_member_types []byte

    db := coreglobals.Db

    row := db.QueryRow(querystr, nodeId)

    err:= row.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &mt_id, &mt_parent_member_type_node_id, &mt_alias, &mt_description, &mt_icon, &mt_meta, &mt_tabs, &mt_parent_member_types)

    var parent_member_type_node_id int
    if mt_parent_member_type_node_id.Valid {
    // use s.String
        id, _ := strconv.Atoi(mt_parent_member_type_node_id.String)
        parent_member_type_node_id = id
    } else {
     // NULL value
    }

    var parent_member_types []MemberType
    var tabs []Tab
    var mt_metaMap map[string]interface{}


    json.Unmarshal(mt_parent_member_types, &parent_member_types)
    json.Unmarshal(mt_tabs, &tabs)
    json.Unmarshal(mt_meta, &mt_metaMap)

    //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
    //lol, _ := json.Marshal(ctTabs)
    //fmt.Println(string(lol))

    //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
    
    //fmt.Println("ksjdflk sdfkj: " + node_name)


    //helpers.PanicIf(err)
    switch {
        case err == sql.ErrNoRows:
                log.Printf("No node with that ID.")
        case err != nil:
                log.Fatal(err)
        default:
                node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil, ""}

                memberType = MemberType{mt_id, node_id, mt_alias, mt_description, mt_icon, parent_member_type_node_id, tabs, mt_metaMap, parent_member_types, &node}
                //memberType = MemberType{mt_id, mt_node_id, mt_alias, mt_description, mt_icon, mt_thumbnail, parent_member_type_node_id, ctTabs, x, nil, &node}
    }

    return
}

// func GetMemberTypes() (memberTypes []MemberType){
//   querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
//     mt.id as mt_id, mt.node_id as mt_node_id, mt.parent_member_type_node_id as mt_parent_member_type_node_id, mt.alias as mt_alias,
//     mt.description as mt_description, mt.icon as mt_icon, mt.meta::json as mt_meta, res.mt_tabs as mt_tabs
//     FROM node
//     JOIN member_type as mt
//     ON mt.node_id = node.id
//     JOIN
//     LATERAL
//     (
//       SELECT my_member_type.*,gf2.*
//       FROM member_type as my_member_type, node as my_member_type_node,
//       LATERAL 
//       (
//           SELECT okidoki.tabs as mt_tabs
//           FROM (
//             SELECT c.id as cid, gf.* as tabs
//             FROM member_type as c, node,
//           LATERAL (
//               select json_agg(row1) as tabs from((
//           select y.name, ss.properties
//           from json_to_recordset(
//           (
//         select * 
//         from json_to_recordset(
//             (
//           SELECT json_agg(ggg)
//           from(
//         SELECT tabs
//         FROM 
//         (   
//             SELECT *
//             FROM member_type as mt
//             WHERE mt.id=c.id
//         ) dsfds

//           )ggg
//             )
//         ) as x(tabs json)
//           )
//           ) as y(name text, properties json),
//           LATERAL (
//         select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'alias',row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//         from(
//       select name, "order", data_type, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
//       from json_to_recordset(properties) 
//       as k(name text, "order" int, data_type int, help_text text, description text)
//       JOIN data_type
//       ON data_type.id = k.data_type
//       )row
//           ) ss
//               ))row1
//           ) gf
//           WHERE c.id = my_member_type.id
//           )okidoki
//           limit 1
//       ) gf2
//       --
//       WHERE my_member_type_node.id = my_member_type.node_id
//     ) res
//     ON res.node_id = mt.node_id
//     WHERE node.node_type=$1`

//     // node
//     var node_id, node_created_by, node_type int
//     var node_path, node_name string
//     var node_created_date time.Time

//     var mt_id, mt_node_id int
//     var mt_parent_member_type_node_id sql.NullString
//     var mt_alias, mt_description, mt_icon string
//     var mt_tabs, mt_meta []byte

//     db := coreglobals.Db

//     rows, err := db.Query(querystr)
//     corehelpers.PanicIf(err)
//     defer rows.Close()

//     for rows.Next(){
//       err:= rows.Scan(
//         &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
//         &mt_id, &mt_node_id, &mt_parent_member_type_node_id, &mt_alias, &mt_description, &mt_icon, &mt_meta, &mt_tabs)

//       var parent_member_type_node_id int
//       if mt_parent_member_type_node_id.Valid {
//       // use s.String
//           id, _ := strconv.Atoi(mt_parent_member_type_node_id.String)
//           parent_member_type_node_id = id
//       } else {
//        // NULL value
//       }

//       mt_tabs_str := string(mt_tabs)
//       //fmt.Println(":::::::::::::::::::::::::::::::::::1 ")
//       //fmt.Println(mt_tabs_str)

//       //fmt.Println(mt_tabs_str + " dsfjldskfj skdf")
//       mt_meta_str := string(mt_meta)
//       var x map[string]interface{}
//       json.Unmarshal([]byte(string(mt_meta_str)), &x)
//       //fmt.Println(mt_meta_str + " dsfjldskfj skdf")

//       // Decode the json object

//       var ctTabs []Tab
//       //var tab Tab

//       errlol := json.Unmarshal([]byte(mt_tabs_str), &ctTabs)
//       corehelpers.PanicIf(errlol)

//       //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
//       //lol, _ := json.Marshal(ctTabs)
//       //fmt.Println(string(lol))

//       //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
      
//       //fmt.Println("ksjdflk sdfkj: " + node_name)


//       //helpers.PanicIf(err)
//       switch {
//           case err == sql.ErrNoRows:
//                   log.Printf("No node with that ID.")
//           case err != nil:
//                   log.Fatal(err)
//           default:
//                   node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, ""}
//                   memberType := MemberType{mt_id, mt_node_id, mt_alias, mt_description, mt_icon, parent_member_type_node_id, ctTabs, x, nil, &node}
//                   memberTypes = append(memberTypes,memberType)
//       }
//     }

//     return
// }

func GetMemberTypeByNodeId(nodeId int) (memberType MemberType){
  
  querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
    mt.id as mt_id, mt.node_id as mt_node_id, mt.parent_member_type_node_id as mt_parent_member_type_node_id, mt.alias as mt_alias,
    mt.description as mt_description, mt.icon as mt_icon, mt.meta::json as mt_meta, mt.tabs as mt_tabs
    FROM node
    JOIN member_type as mt
    ON mt.node_id = node.id
    WHERE node.id=$1`

    // node
    var node_id, node_created_by, node_type int
    var node_path, node_name string
    var node_created_date time.Time

    var mt_id, mt_node_id int
    var mt_parent_member_type_node_id sql.NullString
    var mt_alias, mt_description, mt_icon string
    var mt_tabs, mt_meta []byte

    db := coreglobals.Db

    row := db.QueryRow(querystr, nodeId)

    err:= row.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &mt_id, &mt_node_id, &mt_parent_member_type_node_id, &mt_alias, &mt_description, &mt_icon, &mt_meta, &mt_tabs)

    var parent_member_type_node_id int
    if mt_parent_member_type_node_id.Valid {
    // use s.String
        id, _ := strconv.Atoi(mt_parent_member_type_node_id.String)
        parent_member_type_node_id = id
    } else {
     // NULL value
    }

    mt_tabs_str := string(mt_tabs)
    //fmt.Println(":::::::::::::::::::::::::::::::::::1 ")
    //fmt.Println(mt_tabs_str)

    //fmt.Println(mt_tabs_str + " dsfjldskfj skdf")
    mt_meta_str := string(mt_meta)
    var x map[string]interface{}
    json.Unmarshal([]byte(string(mt_meta_str)), &x)
    //fmt.Println(mt_meta_str + " dsfjldskfj skdf")

    // Decode the json object

    var ctTabs []Tab
    //var tab Tab

    errlol := json.Unmarshal([]byte(mt_tabs_str), &ctTabs)
    corehelpers.PanicIf(errlol)

    //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
    //lol, _ := json.Marshal(ctTabs)
    //fmt.Println(string(lol))

    //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
    
    //fmt.Println("ksjdflk sdfkj: " + node_name)


    //helpers.PanicIf(err)
    switch {
        case err == sql.ErrNoRows:
                log.Printf("No node with that ID.")
        case err != nil:
                log.Fatal(err)
        default:
                node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil, ""}
                memberType = MemberType{mt_id, mt_node_id, mt_alias, mt_description, mt_icon, parent_member_type_node_id, ctTabs, x, nil, &node}
    }

    return
}

// STILL NEEDS SOME WORK REGARDING TABS vs THE DATA TYPE ID/WHOLE OBJECT PROBLEM

// func (ct *MemberType) Update(){
//   db := coreglobals.Db

//   meta, _ := json.Marshal(ct.Meta)

//   tabs, _ := json.Marshal(ct.Tabs)

//   var testMapSlice []map[string]interface{}
//   err1 := json.Unmarshal(tabs,&testMapSlice)
//   corehelpers.PanicIf(err1)
  
//   // //tabs, _ := json.Marshal(ct.Tabs)
//   // for i := 0; i < len(testMapSlice); i++ {
//   //   var properties []interface{} = testMapSlice[i]["properties"].([]interface{})
//   //   for j := 0; j < len(properties); j++ {
//   //     //fmt.Println(properties[j])
//   //     var property map[string]interface{} = properties[j].(map[string]interface{})
//   //     //var dt interface{} = property.data_type
//   //     fmt.Println("property begin ---")
//   //     fmt.Println(property)
//   //     fmt.Println("property end ---\n")
//   //     var dt map[string]interface{} = property["data_type"].(map[string]interface{})
//   //     fmt.Println(dt)
//   //     //property["data_type"] = dt["id"]
//   //   }
    
//   // }

//   res, _ := json.Marshal(testMapSlice)
//   log.Println(string(res))

//   // //b, _ := json.Marshal(testMap)
//   // fmt.Println(testMapSlice)
//   // fmt.Println(testMapSlice[0]["name"])
//   // fmt.Println(testMapSlice[0]["properties"])
//   // //fmt.Println(testMapSlice[name])
//   // //fmt.Println(testMapSlice['name'])
//   // //fmt.Println(testMapSlice[[`name`])

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)
//   //defer tx.Rollback()

//   _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", ct.Node.Name, ct.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r1.Close()

//   _, err = tx.Exec(`UPDATE content_type 
//     SET alias = $1, description = $2, icon = $3, thumbnail = $4, meta = $5, tabs = $6 
//     WHERE node_id = $7`, ct.Alias, ct.Description, ct.Icon, meta, tabs, ct.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()

//   tx.Commit()
// }