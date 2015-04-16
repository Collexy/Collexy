package models

import (
	coreglobals "collexy/core/globals"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	"database/sql"
	"encoding/json"
	"log"
	"time"
)

type MemberType struct {
	Id                int                            `json:"id"`
	Path              string                         `json:"path"`
	ParentId          int                            `json:"parent_id,omitempty"`
	Name              string                         `json:"name"`
	Alias             string                         `json:"alias"`
	CreatedBy         int                            `json:"created_by"`
	CreatedDate       *time.Time                     `json:"created_date"`
	Description       string                         `json:"description,omitempty"`
	Icon              string                         `json:"icon,omitempty"`
	Meta              map[string]interface{}         `json:"meta,omitempty"`
	Tabs              []coremodulesettingsmodels.Tab `json:"tabs,omitempty"`
	ParentMemberTypes []MemberType                   `json:"parent_member_types,omitempty"`
}

func GetMemberTypes() (memberTypes []*MemberType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT member_type.id as member_type_id, member_type.path as member_type_path, 
        member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, 
        member_type.alias as member_alias, member_type.created_by as member_type_created_by, 
        member_type.created_date as member_type_created_date, member_type.description as member_type_description, 
        member_type.icon as member_type_icon, member_type.meta as member_type_meta, 
        member_type.tabs as member_type_tabs
        FROM member_type`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var member_type_id, member_type_created_by int
		var member_type_path, member_type_name, member_type_alias string
		var member_type_description, member_type_icon string
		var member_type_created_date *time.Time

		var member_type_parent_id sql.NullInt64

		var member_type_tabs, member_type_meta []byte

		if err := rows.Scan(&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name,
			&member_type_alias, &member_type_created_by, &member_type_created_date, &member_type_description,
			&member_type_icon, &member_type_meta, &member_type_tabs); err != nil {
			log.Fatal(err)
		}

		var parent_member_type_id int
		if member_type_parent_id.Valid {
			parent_member_type_id = int(member_type_parent_id.Int64)
		} else {
			// NULL value
		}

		var tabs []coremodulesettingsmodels.Tab
		var member_type_metaMap map[string]interface{}

		json.Unmarshal(member_type_tabs, &tabs)
		json.Unmarshal(member_type_meta, &member_type_metaMap)

		memberType := &MemberType{member_type_id, member_type_path, parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_metaMap, tabs, nil}
		memberTypes = append(memberTypes, memberType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMemberTypeExtendedById(id int) (memberType MemberType) {

	querystr := `SELECT member_type.id as member_type_id, member_type.path as member_type_path, member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, member_type.alias as member_alias, member_type.created_by as member_type_created_by,  member_type.created_date as member_type_created_date, member_type.description as member_type_description, member_type.icon as member_type_icon, member_type.meta as member_type_meta,
res.mt_tabs as member_type_tabs, res.parent_member_types as member_type_parent_member_types
FROM member_type  
JOIN
LATERAL
(
    SELECT my_member_type.*,ffgd.*,gf2.*
    FROM member_type as my_member_type,
    LATERAL 
    (
        SELECT array_to_json(array_agg(okidoki)) AS parent_member_types
        FROM (
            SELECT mt.id, mt.path, mt.parent_id, mt.name, mt.alias, mt.created_by, mt.created_date, mt.description, mt.icon, mt.meta, gf.* AS tabs
            FROM member_type AS mt,
            LATERAL 
            (
                SELECT json_agg(row1) AS tabs FROM(
                (
                    SELECT y.name, ss.properties
                    FROM json_to_recordset(
                    (
                        SELECT * 
                        FROM json_to_recordset(
                        (
                            SELECT json_agg(ggg)
                            FROM(
                                SELECT tabs
                                FROM 
                                (   
                                    SELECT *
                                    FROM member_type AS ct
                                    WHERE ct.id=mt.id
                                ) dsfds
                            )ggg
                        )
                        ) AS x(tabs json)
                    )
                    ) AS y(name text, properties json),
                    LATERAL (
                        SELECT json_agg(json_build_object('name',row.name,'order',row."order",'data_type_id',row.data_type_id,'data_type', json_build_object('id',row.data_type_id, 'path',row.data_type_path, 'parent_id', row.data_type_parent_id,'name',row.data_type_name, 'alias',row.data_type_alias, 'created_by',row.data_type_created_by,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) AS properties
                        FROM(
                            SELECT k.name, "order",data_type_id, data_type.path as data_type_path, data_type.parent_id as data_type_parent_id, data_type.name as data_type_name, data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.created_date as data_type_created_date, data_type.html AS data_type_html, help_text, description
                            FROM json_to_recordset(properties) 
                            AS k(name text, "order" int, data_type_id int, help_text text, description text)
                            JOIN data_type
                            ON data_type.id = k.data_type_id
                        )row
                    ) ss
                )
                )row1
            ) gf
            where path @> subpath(my_member_type.path,0,nlevel(my_member_type.path)-1)
        )okidoki
    ) ffgd,
    --
    LATERAL 
    (
        SELECT okidoki.tabs AS mt_tabs
        FROM (
            SELECT mt.id AS cid, gf.* AS tabs
            FROM member_type AS mt,
            LATERAL 
            (
                SELECT json_agg(row1) AS tabs FROM(
                (
                    SELECT y.name, ss.properties
                    FROM json_to_recordset(
                    (
                        SELECT * 
                        FROM json_to_recordset(
                        (
                            SELECT json_agg(ggg)
                            FROM(
                                SELECT tabs
                                FROM 
                                (   
                                    SELECT *
                                    FROM member_type AS ct
                                    WHERE ct.id=mt.id
                                ) dsfds
                            )ggg
                        )) AS x(tabs json)
                    )) AS y(name text, properties json),
                    LATERAL (
                        SELECT json_agg(json_build_object('name',row.name,'order',row."order",'data_type_id',row.data_type_id,'data_type', json_build_object('id',row.data_type_id, 'path',row.data_type_path, 'parent_id', row.data_type_parent_id,'name',row.data_type_name, 'alias',row.data_type_alias, 'created_by',row.data_type_created_by,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) AS properties
                        FROM(
                            SELECT k.name, "order", data_type_id, data_type.path as data_type_path, data_type.parent_id as data_type_parent_id, data_type.name as data_type_name, data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.created_date as data_type_created_date, data_type.html AS data_type_html, help_text, description
                            FROM json_to_recordset(properties) 
                            AS k(name text, "order" int, data_type_id int, help_text text, description text)
                            JOIN data_type
                            ON data_type.id = k.data_type_id
                        )row
                    ) ss
                ))row1
            ) gf
            WHERE mt.id = my_member_type.id
        )okidoki
        limit 1
    ) gf2
    --
) res
ON res.id = member_type.id
WHERE member_type.id=$1`

	// node
	var member_type_id, member_type_created_by int
	var member_type_path, member_type_name, member_type_alias string
	var member_type_description, member_type_icon string
	var member_type_created_date *time.Time

	var member_type_parent_id sql.NullInt64

	var member_type_tabs, member_type_meta []byte
	var member_type_parent_member_types []byte

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name, &member_type_alias,
		&member_type_created_by, &member_type_created_date, &member_type_description, &member_type_icon, &member_type_meta, &member_type_tabs, &member_type_parent_member_types)

	var parent_member_type_id int
	if member_type_parent_id.Valid {
		parent_member_type_id = int(member_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var parent_member_types []MemberType
	var tabs []coremodulesettingsmodels.Tab
	var member_type_metaMap map[string]interface{}

	json.Unmarshal(member_type_parent_member_types, &parent_member_types)
	json.Unmarshal(member_type_tabs, &tabs)
	json.Unmarshal(member_type_meta, &member_type_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No node with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		memberType = MemberType{member_type_id, member_type_path, parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_metaMap, tabs, parent_member_types}
	}

	return
}

func GetMemberTypeById(id int) (memberType MemberType) {
	querystr := `SELECT member_type.id as member_type_id, member_type.path as member_type_path, 
    member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, 
    member_type.alias as member_alias, member_type.created_by as member_type_created_by, 
    member_type.created_date as member_type_created_date, member_type.description as member_type_description, 
    member_type.icon as member_type_icon, member_type.meta as member_type_meta, 
    member_type.tabs as member_type_tabs
        FROM member_type
        WHERE member_type.id=$1`

	var member_type_id, member_type_created_by int
	var member_type_path, member_type_name, member_type_alias string
	var member_type_description, member_type_icon string
	var member_type_created_date *time.Time

	var member_type_parent_id sql.NullInt64

	var member_type_tabs, member_type_meta []byte

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name, &member_type_alias,
		&member_type_created_by, &member_type_created_date, &member_type_description, &member_type_icon, &member_type_meta, &member_type_tabs)

	var parent_member_type_id int
	if member_type_parent_id.Valid {
		parent_member_type_id = int(member_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var tabs []coremodulesettingsmodels.Tab
	var member_type_metaMap map[string]interface{}

	json.Unmarshal(member_type_tabs, &tabs)
	json.Unmarshal(member_type_meta, &member_type_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No node with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		memberType = MemberType{member_type_id, member_type_path, parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_metaMap, tabs, nil}
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

//       var ctTabs []coremodulesettingsmodels.Tab
//       //var tab coremodulesettingsmodels.Tab

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
