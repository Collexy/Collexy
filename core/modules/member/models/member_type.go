package models

import (
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type MemberType struct {
	Id                     int                            `json:"id"`
	Path                   string                         `json:"path"`
	ParentId               *int                           `json:"parent_id,omitempty"`
	Name                   string                         `json:"name"`
	Alias                  string                         `json:"alias"`
	CreatedBy              int                            `json:"created_by"`
	CreatedDate            *time.Time                     `json:"created_date"`
	Description            string                         `json:"description,omitempty"`
	Icon                   string                         `json:"icon,omitempty"`
	Thumbnail              string                         `json:"thumbnail,omitempty"`
	Meta                   map[string]interface{}         `json:"meta,omitempty"`
	Tabs                   []coremodulesettingsmodels.Tab `json:"tabs,omitempty"`
	IsAbstract             bool                           `json:"is_abstract"`
	CompositeMemberTypeIds []int                          `json:"composite_member_type_ids,omitempty"`
	CompositeMemberTypes   []MemberType                   `json:"composite_member_types,omitempty"`
	ParentMemberTypes      []MemberType                   `json:"parent_member_types,omitempty"`
}

func GetMemberTypes(queryStringParams url.Values) (memberTypes []*MemberType) {
	db := coreglobals.Db

	sqlStr := `SELECT member_type.id as member_type_id, member_type.path as member_type_path, 
        member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, 
        member_type.alias as member_alias, member_type.created_by as member_type_created_by, 
        member_type.created_date as member_type_created_date, member_type.description as member_type_description, 
        member_type.icon as member_type_icon, member_type.thumbnail as member_type_thumbnail, member_type.meta as member_type_meta, 
        member_type.tabs as member_type_tabs, member_type.is_abstract as member_type_is_abstract,
        member_type.composite_member_type_ids AS member_type_composite_member_type_ids
        FROM member_type`

	if queryStringParams.Get("levels") != "" {
		sqlStr = sqlStr + ` WHERE member_type.path ~ '*.*{` + queryStringParams.Get("levels") + `}'`
	}

	rows, err := db.Query(sqlStr)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var member_type_id, member_type_created_by int
		var member_type_path, member_type_name, member_type_alias string
		var member_type_description, member_type_icon, member_type_thumbnail string
		var member_type_created_date *time.Time

		var member_type_parent_id sql.NullInt64

		var member_type_tabs, member_type_meta []byte

		var member_type_is_abstract bool
		var member_type_composite_member_type_ids coreglobals.IntSlice

		if err := rows.Scan(&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name,
			&member_type_alias, &member_type_created_by, &member_type_created_date, &member_type_description,
			&member_type_icon, &member_type_thumbnail, &member_type_meta, &member_type_tabs,
			&member_type_is_abstract, &member_type_composite_member_type_ids); err != nil {
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

		memberType := &MemberType{member_type_id, member_type_path, &parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_thumbnail, member_type_metaMap, tabs, member_type_is_abstract, member_type_composite_member_type_ids, nil, nil}
		memberTypes = append(memberTypes, memberType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMemberTypesByIdChildren(id int) (memberTypes []*MemberType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT member_type.id as member_type_id, member_type.path as member_type_path, 
        member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, 
        member_type.alias as member_alias, member_type.created_by as member_type_created_by, 
        member_type.created_date as member_type_created_date, member_type.description as member_type_description, 
        member_type.icon as member_type_icon, member_type.thumbnail as member_type_thumbnail, member_type.meta as member_type_meta, 
        member_type.tabs as member_type_tabs, member_type.is_abstract as member_type_is_abstract,
        member_type.composite_member_type_ids AS member_type_composite_member_type_ids
        FROM member_type WHERE parent_id=$1`, id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var member_type_id, member_type_created_by int
		var member_type_path, member_type_name, member_type_alias string
		var member_type_description, member_type_icon, member_type_thumbnail string
		var member_type_created_date *time.Time

		var member_type_parent_id sql.NullInt64

		var member_type_tabs, member_type_meta []byte

		var member_type_is_abstract bool
		var member_type_composite_member_type_ids coreglobals.IntSlice

		if err := rows.Scan(&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name,
			&member_type_alias, &member_type_created_by, &member_type_created_date, &member_type_description,
			&member_type_icon, &member_type_thumbnail, &member_type_meta, &member_type_tabs,
			&member_type_is_abstract, &member_type_composite_member_type_ids); err != nil {
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

		memberType := &MemberType{member_type_id, member_type_path, &parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_thumbnail, member_type_metaMap, tabs, member_type_is_abstract, member_type_composite_member_type_ids, nil, nil}
		memberTypes = append(memberTypes, memberType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMemberTypeExtendedById(id int) (memberType MemberType) {

	querystr := `SELECT member_type.id as member_type_id, member_type.path as member_type_path, member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, member_type.alias as member_alias, member_type.created_by as member_type_created_by,  member_type.created_date as member_type_created_date, member_type.description as member_type_description, member_type.icon as member_type_icon, 
    member_type.thumbnail as member_type_thumbnail, member_type.meta as member_type_meta,
res.mt_tabs as member_type_tabs, res.parent_member_types as member_type_parent_member_types, member_type.is_abstract as member_type_is_abstract, 
member_type.composite_member_type_ids AS member_type_composite_member_type_ids, res.composite_member_types as member_type_composite_member_types
FROM member_type  
JOIN
LATERAL
(
    SELECT my_member_type.*,ffgd.*,cmt.*,gf2.*
    FROM member_type as my_member_type,
    -- parent member types
    LATERAL 
    (
        SELECT array_to_json(array_agg(okidoki)) AS parent_member_types
        FROM (
            SELECT mt.id, mt.path, mt.parent_id, mt.name, mt.alias, mt.created_by, mt.created_date, mt.description, mt.icon, mt.thumbnail, mt.meta, gf.* AS tabs, mt.is_abstract 
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
                        SELECT json_agg(json_build_object('name',row.name,'order',row."order",'data_type_id',row.data_type_id,'data_type', json_build_object('id',row.data_type_id, 'name',row.data_type_name, 'alias',row.data_type_alias, 'created_by',row.data_type_created_by,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) AS properties
                        FROM(
                            SELECT k.name, "order",data_type_id, data_type.name as data_type_name, data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.created_date as data_type_created_date, data_type.html AS data_type_html, help_text, description
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
    -- composite member types
    LATERAL 
    (
        SELECT array_to_json(array_agg(okidoki)) AS composite_member_types
        FROM (
            SELECT mt.id, mt.path, mt.parent_id, mt.name, mt.alias, mt.created_by, mt.created_date, mt.description, mt.icon, mt.thumbnail, mt.meta, gf.* AS tabs, mt.is_abstract 
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
                        SELECT json_agg(json_build_object('name',row.name,'order',row."order",'data_type_id',row.data_type_id,'data_type', json_build_object('id',row.data_type_id,'name',row.data_type_name, 'alias',row.data_type_alias, 'created_by',row.data_type_created_by,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) AS properties
                        FROM(
                            SELECT k.name, "order",data_type_id, data_type.name as data_type_name, data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.created_date as data_type_created_date, data_type.html AS data_type_html, help_text, description
                            FROM json_to_recordset(properties) 
                            AS k(name text, "order" int, data_type_id int, help_text text, description text)
                            JOIN data_type
                            ON data_type.id = k.data_type_id
                        )row
                    ) ss
                )
                )row1
            ) gf
            WHERE id = ANY(my_member_type.composite_member_type_ids)
        )okidoki
    ) cmt,
    -- tabs
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
                        SELECT json_agg(json_build_object('name',row.name,'order',row."order",'data_type_id',row.data_type_id,'data_type', json_build_object('id',row.data_type_id, 'name',row.data_type_name, 'alias',row.data_type_alias, 'created_by',row.data_type_created_by,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) AS properties
                        FROM(
                            SELECT k.name, "order", data_type_id, data_type.name as data_type_name, data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.created_date as data_type_created_date, data_type.html AS data_type_html, help_text, description
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
	var member_type_description, member_type_icon, member_type_thumbnail string
	var member_type_created_date *time.Time

	var member_type_parent_id sql.NullInt64

	var member_type_tabs, member_type_meta []byte
	var member_type_parent_member_types, member_type_composite_member_types []byte

	var member_type_is_abstract bool
	var member_type_composite_member_type_ids coreglobals.IntSlice

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name, &member_type_alias,
		&member_type_created_by, &member_type_created_date, &member_type_description, &member_type_icon, &member_type_thumbnail, &member_type_meta, &member_type_tabs,
		&member_type_parent_member_types, &member_type_is_abstract, &member_type_composite_member_type_ids, &member_type_composite_member_types)

	var parent_member_type_id int
	if member_type_parent_id.Valid {
		parent_member_type_id = int(member_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var parent_member_types, composite_member_types []MemberType
	var tabs []coremodulesettingsmodels.Tab
	var member_type_metaMap map[string]interface{}

	json.Unmarshal(member_type_parent_member_types, &parent_member_types)
	json.Unmarshal(member_type_composite_member_types, &composite_member_types)
	json.Unmarshal(member_type_tabs, &tabs)
	json.Unmarshal(member_type_meta, &member_type_metaMap)

	switch {
    	case err == sql.ErrNoRows:
    		log.Printf("No node with that ID.")
    	case err != nil:
    		log.Fatal(err)
            //panic(err)
    	default:
    		memberType = MemberType{member_type_id, member_type_path, &parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_thumbnail, member_type_metaMap, tabs, member_type_is_abstract, member_type_composite_member_type_ids, composite_member_types, parent_member_types}
    	}

	return
}

func GetMemberTypeById(id int) (memberType MemberType) {
	querystr := `SELECT member_type.id as member_type_id, member_type.path as member_type_path, 
    member_type.parent_id as member_type_parent_id, member_type.name as member_type_name, 
    member_type.alias as member_alias, member_type.created_by as member_type_created_by, 
    member_type.created_date as member_type_created_date, member_type.description as member_type_description, 
    member_type.icon as member_type_icon, member_type.thumbnail as member_type_thumbnail, member_type.meta as member_type_meta, 
    member_type.tabs as member_type_tabs, member_type.is_abstract as member_type_is_abstract
        FROM member_type
        WHERE member_type.id=$1`

	var member_type_id, member_type_created_by int
	var member_type_path, member_type_name, member_type_alias string
	var member_type_description, member_type_icon, member_type_thumbnail string
	var member_type_created_date *time.Time

	var member_type_parent_id sql.NullInt64

	var member_type_tabs, member_type_meta []byte

	var member_type_is_abstract bool

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&member_type_id, &member_type_path, &member_type_parent_id, &member_type_name, &member_type_alias,
		&member_type_created_by, &member_type_created_date, &member_type_description, &member_type_icon, &member_type_thumbnail, &member_type_meta, &member_type_tabs,
		&member_type_is_abstract)

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
    		memberType = MemberType{member_type_id, member_type_path, &parent_member_type_id, member_type_name, member_type_alias, member_type_created_by, member_type_created_date, member_type_description, member_type_icon, member_type_thumbnail, member_type_metaMap, tabs, member_type_is_abstract, nil, nil, nil}
	}

	return
}

func (ct *MemberType) Post() {
	var meta interface{} = nil
	var tabs interface{} = nil

	if ct.Meta != nil {
		j, _ := json.Marshal(ct.Meta)
		meta = j
	}

	if ct.Tabs != nil {
		j, _ := json.Marshal(ct.Tabs)
		tabs = j
	}

	// var parentId interface{} = nil

	// if ct.ParentId != nil && ct.ParentId != 0{
	// 	parentId = ct.ParentId
	// }

	// see template commented out post function and below
	// _pgs_format, _ := t.PartialTemplateIds.Value()
	// allowedMemberTypeIds, err3 := IntArray(ct.AllowedMemberTypeIds).Value()
	// corehelpers.PanicIf(err3)
	// compositeMemberTypeIds, err4 := IntArray(ct.CompositeMemberTypeIds).Value()
	// corehelpers.PanicIf(err4)
	// allowedTemplateIds, err5 := IntArray(ct.AllowedTemplateIds).Value()
	// corehelpers.PanicIf(err5)

	// http://godoc.org/github.com/lib/pq
	// pq does not support the LastInsertId() method of the Result type in database/sql.
	// To return the identifier of an INSERT (or UPDATE or DELETE),
	// use the Postgres RETURNING clause with a standard Query or QueryRow call:

	db := coreglobals.Db

	// Channel c, is for getting the parent template
	// We need to append the id of the newly created template to the path of the parent id to create the new path

	var parentMemberType MemberType

	if ct.ParentId != nil {
		c := make(chan MemberType)

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			c <- GetMemberTypeById(*ct.ParentId)
		}()

		go func() {
			for i := range c {
				fmt.Println(i)
				parentMemberType = i
			}
		}()

		wg.Wait()
	}

	// This channel and WaitGroup is just to make sure the insert query is completed before we continue
	c1 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		// sqlStr := `INSERT INTO member_type (parent_id, name, alias, created_by, description, icon, thumbnail, meta, tabs, allow_at_root, is_container,
		//           is_abstract, allowed_member_type_ids,composite_member_type_ids, template_id, allowed_template_ids)
		//           VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`
		// err1 := db.QueryRow(sqlStr, ct.ParentId, ct.Name, ct.Alias, ct.CreatedBy, ct.Description, ct.Icon, ct.Thumbnail, meta, tabs, ct.AllowAtRoot, ct.IsContainer,
		sqlStr := `INSERT INTO member_type (parent_id, name, alias, created_by, description, icon, thumbnail, meta, tabs) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
		err1 := db.QueryRow(sqlStr, ct.ParentId, ct.Name, ct.Alias, ct.CreatedBy, ct.Description, ct.Icon, ct.Thumbnail, meta, tabs).Scan(&id)
		corehelpers.PanicIf(err1)
		c1 <- int(id)
	}()

	go func() {
		for i := range c1 {
			fmt.Println(i)
		}
	}()

	wg1.Wait()

	// fmt.Println(parentTemplate.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE member_type 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if ct.ParentId != nil {
		path = parentMemberType.Path + "." + strconv.FormatInt(id, 10)
	}

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	log.Println("member_type created successfully")

}

func (ct *MemberType) Put() {
	var meta interface{} = nil
	var tabs interface{} = nil

	if ct.Meta != nil {
		j, _ := json.Marshal(ct.Meta)
		meta = j
	}

	if ct.Tabs != nil {
		j, _ := json.Marshal(ct.Tabs)
		tabs = j
	}

    compositeMemberTypeIds, err3 := coreglobals.IntSlice(ct.CompositeMemberTypeIds).Value()
    if err3 != nil {
        panic(err3)
    }
	// var parentId interface{} = nil

	// if ct.ParentId != nil && ct.ParentId != 0{
	// 	parentId = ct.ParentId
	// }

	// see template commented out post function and below
	// _pgs_format, _ := t.PartialTemplateIds.Value()
	// allowedMemberTypeIds, err3 := IntArray(ct.AllowedMemberTypeIds).Value()
	// corehelpers.PanicIf(err3)
	// compositeMemberTypeIds, err4 := IntArray(ct.CompositeMemberTypeIds).Value()
	// corehelpers.PanicIf(err4)
	// allowedTemplateIds, err5 := IntArray(ct.AllowedTemplateIds).Value()
	// corehelpers.PanicIf(err5)

	db := coreglobals.Db

	var parentMemberType MemberType

	if ct.ParentId != nil && *ct.ParentId != 0{
		// Channel c, is for getting the parent template
		// We need to append the id of the newly created template to the path of the parent id to create the new path
		c := make(chan MemberType)

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			c <- GetMemberTypeById(*ct.ParentId)
		}()

		go func() {
			for i := range c {
				fmt.Println(i)
				parentMemberType = i
			}
		}()

		wg.Wait()
	}

    if *ct.ParentId == 0 {
        ct.ParentId = nil
    }

	path := strconv.Itoa(ct.Id)
	if ct.ParentId != nil && *ct.ParentId != 0{
		path = parentMemberType.Path + "." + strconv.Itoa(ct.Id)
	}

	// sqlStr := `UPDATE member_type SET path=$1, parent_id=$2, name=$3, alias=$4, created_by=$5, description=$6, icon=$7, thumbnail=$8, meta=$9, tabs=$10, allow_at_root=$11, is_container=$12,
	//        is_abstract=$13, allowed_member_type_ids=$14,composite_member_type_ids=$15, template_id=$16, allowed_template_ids=$17
	//        WHERE id=$18`

	// _, err6 := db.Exec(sqlStr, path, ct.ParentId, ct.Name, ct.Alias, ct.CreatedBy, ct.Description, ct.Icon, ct.Thumbnail, meta, tabs, ct.AllowAtRoot, ct.IsContainer,
	// 	ct.IsAbstract, allowedMemberTypeIds, compositeMemberTypeIds, ct.TemplateId, allowedTemplateIds, ct.Id)

	sqlStr := `UPDATE member_type SET path=$1, parent_id=$2, name=$3, alias=$4, created_by=$5, description=$6, icon=$7, thumbnail=$8, meta=$9, tabs=$10, is_abstract=$11, composite_member_type_ids=$12 
        WHERE id=$13`

	_, err6 := db.Exec(sqlStr, path, ct.ParentId, ct.Name, ct.Alias, ct.CreatedBy, ct.Description, ct.Icon, ct.Thumbnail, meta, tabs, ct.IsAbstract, compositeMemberTypeIds, ct.Id)

	corehelpers.PanicIf(err6)

	log.Println("member_type updated successfully")

}

func DeleteMemberType(id int) {

	db := coreglobals.Db

	sqlStr := `DELETE FROM member_type 
    WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	corehelpers.PanicIf(err)

	log.Printf("member_type with id %d was successfully deleted", id)
}
