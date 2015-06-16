package models

import (
	// "fmt"
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"encoding/json"
	"fmt"
	"time"
	// "net/http"
	// "html/template"
	"database/sql"
	"log"
	"net/url"
	"strconv"
	"sync"
)

type MediaType struct {
	Id                    int                    `json:"id"`
	Path                  string                 `json:"path"`
	ParentId              *int                    `json:"parent_id,omitempty"`
	Name                  string                 `json:"name"`
	Alias                 string                 `json:"alias"`
	CreatedBy             int                    `json:"created_by"`
	CreatedDate           *time.Time             `json:"created_date"`
	Description           string                 `json:"description,omitempty"`
	Icon                  string                 `json:"icon,omitempty"`
	Thumbnail             string                 `json:"thumbnail,omitempty"`
	Meta                  map[string]interface{} `json:"meta,omitempty"`
	Tabs                  []Tab                  `json:"tabs,omitempty"`
	ParentMediaTypes      []MediaType            `json:"parent_media_types,omitempty"`
	AllowedMediaTypes     []MediaType            `json:"allowed_media_types,omitempty"`
	AllowAtRoot           bool                   `json:"allow_at_root"`
	IsContainer           bool                   `json:"is_container"`
	IsAbstract            bool                   `json:"is_abstract"`
	AllowedMediaTypeIds   []int                  `json:"allowed_media_type_ids,omitempty"`
	CompositeMediaTypeIds []int                  `json:"composite_media_type_ids,omitempty"`
	CompositeMediaTypes   []MediaType            `json:"composite_media_types,omitempty"`
}

func GetMediaTypes(queryStringParams url.Values) (mediaTypes []*MediaType) {
	db := coreglobals.Db
	sqlStr := `SELECT media_type.id as media_type_id, media_type.path as media_type_path, 
        media_type.parent_id as media_type_parent_id, media_type.name as media_type_name, 
        media_type.alias as member_alias, media_type.created_by as media_type_created_by, 
        media_type.created_date as media_type_created_date, media_type.description as media_type_description, 
        media_type.icon as media_type_icon, media_type.thumbnail as media_type_thumbnail,
        media_type.meta as media_type_meta, media_type.tabs as media_type_tabs, 
        media_type.allow_at_root AS media_type_allow_at_root,
        media_type.is_container AS media_type_is_container, media_type.is_abstract as media_type_is_abstract,
        media_type.allowed_media_type_ids AS media_type_allowed_media_type_ids 
        FROM media_type`

	if queryStringParams.Get("levels") != "" {
		sqlStr = sqlStr + ` WHERE media_type.path ~ '*.*{` + queryStringParams.Get("levels") + `}'`
	}

	// if queryStringParams.Get("type-id") != ""{
	//     sqlStr = sqlStr + ` WHERE media_type.type_id=` + queryStringParams.Get("type-id")
	// }

	rows, err := db.Query(sqlStr)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var media_type_id, media_type_created_by int
		var media_type_path, media_type_name, media_type_alias string
		var media_type_description, media_type_icon, media_type_thumbnail string
		var media_type_created_date *time.Time

		var media_type_allow_at_root, media_type_is_container, media_type_is_abstract bool
		var media_type_allowed_media_type_ids coreglobals.IntSlice

		var media_type_parent_id sql.NullInt64

		var media_type_tabs, media_type_meta []byte

		if err := rows.Scan(&media_type_id, &media_type_path, &media_type_parent_id, &media_type_name,
			&media_type_alias, &media_type_created_by, &media_type_created_date, &media_type_description,
			&media_type_icon, &media_type_thumbnail, &media_type_meta, &media_type_tabs,
			&media_type_allow_at_root, &media_type_is_container, &media_type_is_abstract, &media_type_allowed_media_type_ids); err != nil {
			log.Fatal(err)
		}

		var parent_media_type_id int
		if media_type_parent_id.Valid {
			parent_media_type_id = int(media_type_parent_id.Int64)
		} else {
			// NULL value
		}

		var tabs []Tab
		var media_type_metaMap map[string]interface{}

		json.Unmarshal(media_type_tabs, &tabs)
		json.Unmarshal(media_type_meta, &media_type_metaMap)

		mediaType := &MediaType{media_type_id, media_type_path, &parent_media_type_id, media_type_name, media_type_alias, media_type_created_by, media_type_created_date, media_type_description, media_type_icon, media_type_thumbnail, media_type_metaMap, tabs, nil, nil, media_type_allow_at_root, media_type_is_container, media_type_is_abstract, media_type_allowed_media_type_ids, nil, nil}
		mediaTypes = append(mediaTypes, mediaType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMediaTypesByIdChildren(id int) (mediaTypes []*MediaType) {
	db := coreglobals.Db
	sqlStr := `SELECT media_type.id as media_type_id, media_type.path as media_type_path, 
        media_type.parent_id as media_type_parent_id, media_type.name as media_type_name, 
        media_type.alias as member_alias, media_type.created_by as media_type_created_by, 
        media_type.created_date as media_type_created_date, media_type.description as media_type_description, 
        media_type.icon as media_type_icon, media_type.thumbnail as media_type_thumbnail,
        media_type.meta as media_type_meta, media_type.tabs as media_type_tabs 
        FROM media_type
        WHERE media_type.parent_id=$1`

	// if queryStringParams.Get("type-id") != ""{
	//     sqlStr = sqlStr + ` WHERE media_type.type_id=` + queryStringParams.Get("type-id")
	// }

	rows, err := db.Query(sqlStr, id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var media_type_id, media_type_created_by int
		var media_type_path, media_type_name, media_type_alias string
		var media_type_description, media_type_icon, media_type_thumbnail string
		var media_type_created_date *time.Time

		var media_type_parent_id sql.NullInt64

		var media_type_tabs, media_type_meta []byte

		if err := rows.Scan(&media_type_id, &media_type_path, &media_type_parent_id, &media_type_name,
			&media_type_alias, &media_type_created_by, &media_type_created_date, &media_type_description,
			&media_type_icon, &media_type_thumbnail, &media_type_meta, &media_type_tabs); err != nil {
			log.Fatal(err)
		}

		var parent_media_type_id int
		if media_type_parent_id.Valid {
			parent_media_type_id = int(media_type_parent_id.Int64)
		} else {
			// NULL value
		}

		var tabs []Tab
		var media_type_metaMap map[string]interface{}

		json.Unmarshal(media_type_tabs, &tabs)
		json.Unmarshal(media_type_meta, &media_type_metaMap)

		mediaType := &MediaType{media_type_id, media_type_path, &parent_media_type_id, media_type_name, media_type_alias, media_type_created_by, media_type_created_date, media_type_description, media_type_icon, media_type_thumbnail, media_type_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
		mediaTypes = append(mediaTypes, mediaType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMediaTypeExtendedById(id int) (mediaType MediaType) {

	querystr := `SELECT media_type.id as media_type_id, media_type.path as media_type_path, media_type.parent_id as media_type_parent_id, media_type.name as media_type_name, media_type.alias as member_alias, media_type.created_by as media_type_created_by,  media_type.created_date as media_type_created_date, media_type.description as media_type_description, media_type.icon as media_type_icon, media_type.thumbnail as media_type_thumbnail, media_type.meta as media_type_meta,
res.mt_tabs as media_type_tabs, res.parent_media_types as media_type_parent_media_types, res.composite_media_types as media_type_composite_media_types, 
media_type.allow_at_root AS media_type_allow_at_root, 
media_type.is_container AS media_type_is_container, media_type.is_abstract as media_type_is_abstract, 
media_type.allowed_media_type_ids AS media_type_allowed_media_type_ids, media_type.composite_media_type_ids AS media_type_composite_media_type_ids 
FROM media_type  
JOIN
LATERAL
(
    SELECT my_media_type.*,ffgd.*,cct.*, gf2.*
    FROM media_type as my_media_type,
    -- parent media types
    LATERAL 
    (
        SELECT array_to_json(array_agg(okidoki)) AS parent_media_types
        FROM (
            SELECT mt.id, mt.path, mt.parent_id, mt.name, mt.alias, mt.created_by, mt.description, mt.icon, mt.meta, gf.* AS tabs, mt.allow_at_root, mt.is_container, mt.is_abstract 
            FROM media_type AS mt,
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
                                    FROM media_type AS ct
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
            where path @> subpath(my_media_type.path,0,nlevel(my_media_type.path)-1)
        )okidoki
    ) ffgd,
    -- composite media types
    LATERAL 
    (
        SELECT array_to_json(array_agg(okidoki)) AS composite_media_types
        FROM (
            SELECT mt.id, mt.path, mt.parent_id, mt.name, mt.alias, mt.created_by, mt.description, mt.icon, mt.meta, gf.* AS tabs
            FROM media_type AS mt,
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
                                    FROM media_type AS ct
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
            --where path @> subpath(my_media_type.path,0,nlevel(my_media_type.path)-1)
            WHERE id = ANY(my_media_type.composite_media_type_ids)
        )okidoki
    ) cct,
    -- tabs
    LATERAL 
    (
        SELECT okidoki.tabs AS mt_tabs
        FROM (
            SELECT mt.id AS cid, gf.* AS tabs
            FROM media_type AS mt,
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
                                    FROM media_type AS ct
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
            WHERE mt.id = my_media_type.id
        )okidoki
        limit 1
    ) gf2
    --
) res
ON res.id = media_type.id
WHERE media_type.id=$1`

	// node
	var media_type_id, media_type_created_by int
	var media_type_path, media_type_name, media_type_alias string
	var media_type_description, media_type_icon, media_type_thumbnail string
	var media_type_created_date *time.Time

	var media_type_parent_id sql.NullInt64

	var media_type_allow_at_root, media_type_is_container, media_type_is_abstract bool
	var media_type_allowed_media_type_ids, media_type_composite_media_type_ids coreglobals.IntSlice

	var media_type_tabs, media_type_meta []byte
	var media_type_parent_media_types, media_type_composite_media_types []byte

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&media_type_id, &media_type_path, &media_type_parent_id, &media_type_name, &media_type_alias,
		&media_type_created_by, &media_type_created_date, &media_type_description, &media_type_icon, &media_type_thumbnail, &media_type_meta,
		&media_type_tabs, &media_type_parent_media_types, &media_type_composite_media_types, &media_type_allow_at_root, &media_type_is_container,
		&media_type_is_abstract, &media_type_allowed_media_type_ids, &media_type_composite_media_type_ids)

	var parent_media_type_id int
	if media_type_parent_id.Valid {
		parent_media_type_id = int(media_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var parent_media_types, composite_media_types []MediaType
	var tabs []Tab
	var media_type_metaMap map[string]interface{}

	json.Unmarshal(media_type_parent_media_types, &parent_media_types)
	json.Unmarshal(media_type_composite_media_types, &composite_media_types)
	json.Unmarshal(media_type_tabs, &tabs)
	json.Unmarshal(media_type_meta, &media_type_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No node with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		mediaType = MediaType{media_type_id, media_type_path, &parent_media_type_id, media_type_name, media_type_alias, media_type_created_by, media_type_created_date, media_type_description, media_type_icon, media_type_thumbnail, media_type_metaMap, tabs, parent_media_types, nil, media_type_allow_at_root, media_type_is_container, media_type_is_abstract, media_type_allowed_media_type_ids, media_type_composite_media_type_ids, composite_media_types}
	}

	return
}

func GetMediaTypeById(id int) (mediaType MediaType) {
	querystr := `SELECT media_type.id as media_type_id, media_type.path as media_type_path, 
    media_type.parent_id as media_type_parent_id, media_type.name as media_type_name, 
    media_type.alias as member_alias, media_type.created_by as media_type_created_by, 
    media_type.created_date as media_type_created_date, media_type.description as media_type_description, 
    media_type.icon as media_type_icon, media_type.thumbnail as media_type_thumbnail, media_type.meta as media_type_meta, 
    media_type.tabs as media_type_tabs  
        FROM media_type
        WHERE media_type.id=$1`

	var media_type_id, media_type_created_by int
	var media_type_path, media_type_name, media_type_alias string
	var media_type_description, media_type_icon, media_type_thumbnail string
	var media_type_created_date *time.Time

	var media_type_parent_id sql.NullInt64

	var media_type_tabs, media_type_meta []byte

	db := coreglobals.Db

	row := db.QueryRow(querystr, id)

	err := row.Scan(
		&media_type_id, &media_type_path, &media_type_parent_id, &media_type_name, &media_type_alias,
		&media_type_created_by, &media_type_created_date, &media_type_description, &media_type_icon, &media_type_thumbnail, &media_type_meta, &media_type_tabs)

	var parent_media_type_id int
	if media_type_parent_id.Valid {
		parent_media_type_id = int(media_type_parent_id.Int64)
	} else {
		// NULL value
	}

	var tabs []Tab
	var media_type_metaMap map[string]interface{}

	json.Unmarshal(media_type_tabs, &tabs)
	json.Unmarshal(media_type_meta, &media_type_metaMap)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No node with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		mediaType = MediaType{media_type_id, media_type_path, &parent_media_type_id, media_type_name, media_type_alias, media_type_created_by, media_type_created_date, media_type_description, media_type_icon, media_type_thumbnail, media_type_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
	}

	return
}

func (mt *MediaType) Post() {
	var meta interface{} = nil
    var tabs interface{} = nil

    if mt.Meta != nil {
        j, _ := json.Marshal(mt.Meta)
        meta = j
    }

    if mt.Tabs != nil {
        j, _ := json.Marshal(mt.Tabs)
        tabs = j
    }

	// see template commented out post function and below
	// _pgs_format, _ := t.PartialTemplateIds.Value()
	allowedMediaTypeIds, err3 := IntArray(mt.AllowedMediaTypeIds).Value()
	corehelpers.PanicIf(err3)
	compositeMediaTypeIds, err4 := IntArray(mt.CompositeMediaTypeIds).Value()
	corehelpers.PanicIf(err4)

	// http://godoc.org/github.com/lib/pq
	// pq does not support the LastInsertId() method of the Result type in database/sql.
	// To return the identifier of an INSERT (or UPDATE or DELETE),
	// use the Postgres RETURNING clause with a standard Query or QueryRow call:

	db := coreglobals.Db

    var parentMediaType MediaType

    if mt.ParentId != nil {
    	// Channel c, is for getting the parent template
    	// We need to append the id of the newly created template to the path of the parent id to create the new path
    	c := make(chan MediaType)
    	

    	var wg sync.WaitGroup

    	wg.Add(1)

    	go func() {
    		defer wg.Done()
    		c <- GetMediaTypeById(*mt.ParentId)
    	}()

    	go func() {
    		for i := range c {
    			fmt.Println(i)
    			parentMediaType = i
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
		sqlStr := `INSERT INTO media_type (parent_id, name, alias, created_by, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, 
            is_abstract, allowed_media_type_ids,composite_media_type_ids) 
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`
		err1 := db.QueryRow(sqlStr, mt.ParentId, mt.Name, mt.Alias, mt.CreatedBy, mt.Description, mt.Icon, mt.Thumbnail, meta, tabs, mt.AllowAtRoot, mt.IsContainer,
			mt.IsAbstract, allowedMediaTypeIds, compositeMediaTypeIds).Scan(&id)
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

	sqlStr := `UPDATE media_type 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if mt.ParentId != nil {
		path = parentMediaType.Path + "." + strconv.FormatInt(id, 10)
	}

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	log.Println("media type created successfully")

}

func (mt *MediaType) Put() {
	var meta interface{} = nil
    var tabs interface{} = nil

    if mt.Meta != nil {
        j, _ := json.Marshal(mt.Meta)
        meta = j
    }

    if mt.Tabs != nil {
        j, _ := json.Marshal(mt.Tabs)
        tabs = j
    }

	// see template commented out post function and below
	// _pgs_format, _ := t.PartialTemplateIds.Value()
	allowedMediaTypeIds, err3 := IntArray(mt.AllowedMediaTypeIds).Value()
	corehelpers.PanicIf(err3)
	compositeMediaTypeIds, err4 := IntArray(mt.CompositeMediaTypeIds).Value()
	corehelpers.PanicIf(err4)

	db := coreglobals.Db

    var parentMediaType MediaType

    if mt.ParentId != nil {

    	// Channel c, is for getting the parent template
    	// We need to append the id of the newly created template to the path of the parent id to create the new path
    	c := make(chan MediaType)

    	var wg sync.WaitGroup

    	wg.Add(1)

    	go func() {
    		defer wg.Done()
    		c <- GetMediaTypeById(*mt.ParentId)
    	}()

    	go func() {
    		for i := range c {
    			fmt.Println(i)
    			parentMediaType = i
    		}
    	}()

    	wg.Wait()
    }

	sqlStr := `UPDATE media_type SET path=$1, parent_id=$2, name=$3, alias=$4, created_by=$5, description=$6, icon=$7, thumbnail=$8, meta=$9, tabs=$10, allow_at_root=$11, is_container=$12, 
        is_abstract=$13, allowed_media_type_ids=$14,composite_media_type_ids=$15
        WHERE id=$16`

	path := strconv.Itoa(mt.Id)
	if mt.ParentId != nil {
		path = parentMediaType.Path + "." + strconv.Itoa(mt.Id)
	}
	_, err6 := db.Exec(sqlStr, path, mt.ParentId, mt.Name, mt.Alias, mt.CreatedBy, mt.Description, mt.Icon, mt.Thumbnail, meta, tabs, mt.AllowAtRoot, mt.IsContainer,
		mt.IsAbstract, allowedMediaTypeIds, compositeMediaTypeIds, mt.Id)

	corehelpers.PanicIf(err6)

	log.Println("media type updated successfully")

}

func DeleteMediaType(id int) {

	db := coreglobals.Db

	sqlStr := `DELETE FROM media_type 
    WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	corehelpers.PanicIf(err)

	log.Printf("media type with id %d was successfully deleted", id)
}

// func (t *MediaType) Post(){
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
//   err = tx.QueryRow(`SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`, t.ParentMediaTypeNodeId).Scan(&id, &path, &created_by, &name, &node_type, &created_date)
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
//   err = db.QueryRow(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`, t.Node.Name, t.Node.NodeType, 1, t.ParentMediaTypeNodeId).Scan(&node_id)
//   //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateNodeId)
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
//   meta, errMeta := json.Marshal(t.Meta)
//   corehelpers.PanicIf(errMeta)

//   tabs, errTabs := json.Marshal(t.Tabs)
//   corehelpers.PanicIf(errTabs)

//   _, err = tx.Exec("INSERT INTO media_type (node_id, alias, description, icon, thumbnail, parent_media_type_node_id, meta, tabs) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", node_id, t.Alias, t.Description, t.Icon, t.Thumbnail, t.ParentMediaTypeNodeId, meta, tabs)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()
//   err1 := tx.Commit()
//   corehelpers.PanicIf(err1)

// }

// func GetMediaTypes() (mediaTypes []MediaType){
//   querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
//     ct.id as ct_id, ct.node_id as ct_node_id, ct.parent_media_type_node_id as ct_parent_media_type_node_id, ct.alias as ct_alias,
//     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta, res.ct_tabs as ct_tabs
//     FROM node
//     JOIN media_type as ct
//     ON ct.node_id = node.id
//     JOIN
//     LATERAL
//     (
//       SELECT my_media_type.*,gf2.*
//       FROM media_type as my_media_type, node as my_media_type_node,
//       LATERAL
//       (
//           SELECT okidoki.tabs as ct_tabs
//           FROM (
//             SELECT c.id as cid, gf.* as tabs
//             FROM media_type as c, node,
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
//             FROM media_type as ct
//             WHERE ct.id=c.id
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
//           WHERE c.id = my_media_type.id
//           )okidoki
//           limit 1
//       ) gf2
//       --
//       WHERE my_media_type_node.id = my_media_type.node_id
//     ) res
//     ON res.node_id = ct.node_id
//     WHERE node.node_type=4`

//     // node
//     var node_id, node_created_by, node_type int
//     var node_path, node_name string
//     var node_created_date time.Time

//     var ct_id, ct_node_id int
//     var ct_parent_media_type_node_id sql.NullString
//     var ct_alias, ct_description, ct_icon, ct_thumbnail string
//     var ct_tabs, ct_meta []byte

//     db := coreglobals.Db

//     rows, err := db.Query(querystr)
//     corehelpers.PanicIf(err)
//     defer rows.Close()

//     for rows.Next(){
//       err:= rows.Scan(
//         &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
//         &ct_id, &ct_node_id, &ct_parent_media_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs)

//       var parent_media_type_node_id int
//       if ct_parent_media_type_node_id.Valid {
//       // use s.String
//           id, _ := strconv.Atoi(ct_parent_media_type_node_id.String)
//           parent_media_type_node_id = id
//       } else {
//        // NULL value
//       }

//       ct_tabs_str := string(ct_tabs)
//       //fmt.Println(":::::::::::::::::::::::::::::::::::1 ")
//       //fmt.Println(ct_tabs_str)

//       //fmt.Println(ct_tabs_str + " dsfjldskfj skdf")
//       ct_meta_str := string(ct_meta)
//       var x map[string]interface{}
//       json.Unmarshal([]byte(string(ct_meta_str)), &x)
//       //fmt.Println(ct_meta_str + " dsfjldskfj skdf")

//       // Decode the json object

//       var ctTabs []Tab
//       //var tab Tab

//       errlol := json.Unmarshal([]byte(ct_tabs_str), &ctTabs)
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
//                   node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil, ""}
//                   mediaType := MediaType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_media_type_node_id, ctTabs, x, nil, &node}
//                   mediaTypes = append(mediaTypes,mediaType)
//       }
//     }

//     return
// }

// func GetMediaTypeExtendedByNodeId(nodeId int) (mediaType MediaType){

//   querystr := `SELECT my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date,
//     res.id as ct_id, res.parent_media_type_node_id as ct_parent_media_type_node_id, res.alias as ct_alias,
//     res.description as ct_description, res.icon as ct_icon, res.thumbnail as ct_thumbnail, res.meta::json as ct_meta, res.ct_tabs as ct_tabs, res.parent_media_types as ct_parent_media_types
//     FROM media_type
//     JOIN node as my_node
//     ON my_node.id = media_type.node_id
//     JOIN
//     LATERAL
//     (
//       SELECT my_media_type.*,ffgd.*,gf2.*
//       FROM media_type as my_media_type, node as my_media_type_node,
//       LATERAL
//       (
//           SELECT array_to_json(array_agg(okidoki)) as parent_media_types
//           FROM (
//             SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.thumbnail, c.parent_media_type_node_id, c.meta, gf.* as tabs
//             FROM media_type as c, node,
//           LATERAL (
//               select json_agg(row1) as tabs from((
//               select y.name, ss.properties
//               from json_to_recordset(
//             (
//                 select *
//                 from json_to_recordset(
//               (
//                   SELECT json_agg(ggg)
//                   from(
//                 SELECT tabs
//                 FROM
//                 (
//                     SELECT *
//                     FROM media_type as ct
//                     WHERE ct.id=c.id
//                 ) dsfds

//                   )ggg
//               )
//                 ) as x(tabs json)
//             )
//               ) as y(name text, properties json),
//               LATERAL (
//             select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id',row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id',row.data_type_node_id, 'alias', row.data_type_alias,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//             from(
//                 select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
//                 from json_to_recordset(properties)
//                 as k(name text, "order" int, data_type_node_id int, help_text text, description text)
//                 JOIN data_type
//                 ON data_type.node_id = k.data_type_node_id
//                 )row
//               ) ss
//               ))row1
//           ) gf
//             where path @> subpath(my_media_type_node.path,0,nlevel(my_media_type_node.path)-1) and c.node_id = node.id
//           )okidoki
//       ) ffgd,
//       --
//       LATERAL
//       (
//           SELECT okidoki.tabs as ct_tabs
//           FROM (
//             SELECT c.id as cid, gf.* as tabs
//             FROM media_type as c, node,
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
//             FROM media_type as ct
//             WHERE ct.id=c.id
//         ) dsfds

//           )ggg
//             )
//         ) as x(tabs json)
//           )
//           ) as y(name text, properties json),
//           LATERAL (
//         select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id', row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id', row.data_type_node_id, 'alias', row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//         from(
//       select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
//       from json_to_recordset(properties)
//       as k(name text, "order" int, data_type_node_id int, help_text text, description text)
//       JOIN data_type
//       ON data_type.node_id = k.data_type_node_id
//       )row
//           ) ss
//               ))row1
//           ) gf
//           WHERE c.id = my_media_type.id
//           )okidoki
//           limit 1
//       ) gf2
//       --
//       WHERE my_media_type_node.id = my_media_type.node_id
//     ) res
//     ON res.node_id = media_type.node_id
//     WHERE media_type.node_id=$1`

//     // node
//     var node_id, node_created_by, node_type int
//     var node_path, node_name string
//     var node_created_date time.Time

//     var ct_id int
//     var ct_parent_media_type_node_id sql.NullString

//     var ct_alias, ct_description, ct_icon, ct_thumbnail string
//     var ct_tabs, ct_meta []byte
//     var ct_parent_media_types []byte

//     db := coreglobals.Db

//     row := db.QueryRow(querystr, nodeId)

//     err:= row.Scan(
//         &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
//         &ct_id, &ct_parent_media_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_parent_media_types)

//     var parent_media_type_node_id int
//     if ct_parent_media_type_node_id.Valid {
//     // use s.String
//         id, _ := strconv.Atoi(ct_parent_media_type_node_id.String)
//         parent_media_type_node_id = id
//     } else {
//      // NULL value
//     }

//     var parent_media_types []MediaType
//     var tabs []Tab
//     var ct_metaMap map[string]interface{}

//     json.Unmarshal(ct_parent_media_types, &parent_media_types)
//     json.Unmarshal(ct_tabs, &tabs)
//     json.Unmarshal(ct_meta, &ct_metaMap)

//     //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
//     //lol, _ := json.Marshal(ctTabs)
//     //fmt.Println(string(lol))

//     //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)

//     //fmt.Println("ksjdflk sdfkj: " + node_name)

//     //helpers.PanicIf(err)
//     switch {
//         case err == sql.ErrNoRows:
//                 log.Printf("No node with that ID.")
//         case err != nil:
//                 log.Fatal(err)
//         default:
//                 node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil, ""}

//                 mediaType = MediaType{ct_id, node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_media_type_node_id, tabs, ct_metaMap, parent_media_types, &node}
//                 //mediaType = MediaType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_media_type_node_id, ctTabs, x, nil, &node}
//     }

//     return
// }

// func GetMediaTypeByNodeId(nodeId int) (mediaType MediaType){
//   // querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
//   //   ct.id as ct_id, ct.node_id as ct_node_id, ct.parent_media_type_node_id as ct_parent_media_type_node_id, ct.alias as ct_alias,
//   //   ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta, res.ct_tabs as ct_tabs
//   //   FROM node
//   //   JOIN media_type as ct
//   //   ON ct.node_id = node.id
//   //   JOIN
//   //   LATERAL
//   //   (
//   //     SELECT my_media_type.*,gf2.*
//   //     FROM media_type as my_media_type, node as my_media_type_node,
//   //     LATERAL
//   //     (
//   //         SELECT okidoki.tabs as ct_tabs
//   //         FROM (
//   //           SELECT c.id as cid, gf.* as tabs
//   //           FROM media_type as c, node,
//   //         LATERAL (
//   //             select json_agg(row1) as tabs from((
//   //         select y.name, ss.properties
//   //         from json_to_recordset(
//   //         (
//   //       select *
//   //       from json_to_recordset(
//   //           (
//   //         SELECT json_agg(ggg)
//   //         from(
//   //       SELECT tabs
//   //       FROM
//   //       (
//   //           SELECT *
//   //           FROM media_type as ct
//   //           WHERE ct.id=c.id
//   //       ) dsfds

//   //         )ggg
//   //           )
//   //       ) as x(tabs json)
//   //         )
//   //         ) as y(name text, properties json),
//   //         LATERAL (
//   //       select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id', row.data_type_id,'node_id',row.data_type, 'alias',row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//   //       from(
//   //     select name, "order", data_type.id as data_type_id, data_type, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
//   //     from json_to_recordset(properties)
//   //     as k(name text, "order" int, data_type int, help_text text, description text)
//   //     JOIN data_type
//   //     ON data_type.node_id = k.data_type
//   //     )row
//   //         ) ss
//   //             ))row1
//   //         ) gf
//   //         WHERE c.id = my_media_type.id
//   //         )okidoki
//   //         limit 1
//   //     ) gf2
//   //     --
//   //     WHERE my_media_type_node.id = my_media_type.node_id
//   //   ) res
//   //   ON res.node_id = ct.node_id
//   //   WHERE node.id=$1`
//   querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
//     ct.id as ct_id, ct.node_id as ct_node_id, ct.parent_media_type_node_id as ct_parent_media_type_node_id, ct.alias as ct_alias,
//     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta, ct.tabs as ct_tabs
//     FROM node
//     JOIN media_type as ct
//     ON ct.node_id = node.id
//     WHERE node.id=$1`

//     // node
//     var node_id, node_created_by, node_type int
//     var node_path, node_name string
//     var node_created_date time.Time

//     var ct_id, ct_node_id int
//     var ct_parent_media_type_node_id sql.NullString
//     var ct_alias, ct_description, ct_icon, ct_thumbnail string
//     var ct_tabs, ct_meta []byte

//     db := coreglobals.Db

//     row := db.QueryRow(querystr, nodeId)

//     err:= row.Scan(
//         &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
//         &ct_id, &ct_node_id, &ct_parent_media_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs)

//     var parent_media_type_node_id int
//     if ct_parent_media_type_node_id.Valid {
//     // use s.String
//         id, _ := strconv.Atoi(ct_parent_media_type_node_id.String)
//         parent_media_type_node_id = id
//     } else {
//      // NULL value
//     }

//     ct_tabs_str := string(ct_tabs)
//     //fmt.Println(":::::::::::::::::::::::::::::::::::1 ")
//     //fmt.Println(ct_tabs_str)

//     //fmt.Println(ct_tabs_str + " dsfjldskfj skdf")
//     ct_meta_str := string(ct_meta)
//     var x map[string]interface{}
//     json.Unmarshal([]byte(string(ct_meta_str)), &x)
//     //fmt.Println(ct_meta_str + " dsfjldskfj skdf")

//     // Decode the json object

//     var ctTabs []Tab
//     //var tab Tab

//     errlol := json.Unmarshal([]byte(ct_tabs_str), &ctTabs)
//     corehelpers.PanicIf(errlol)

//     //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
//     //lol, _ := json.Marshal(ctTabs)
//     //fmt.Println(string(lol))

//     //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)

//     //fmt.Println("ksjdflk sdfkj: " + node_name)

//     //helpers.PanicIf(err)
//     switch {
//         case err == sql.ErrNoRows:
//                 log.Printf("No node with that ID.")
//         case err != nil:
//                 log.Fatal(err)
//         default:
//                 node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, "", nil, nil, ""}
//                 mediaType = MediaType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_media_type_node_id, ctTabs, x, nil, &node}
//     }

//     return
// }

// // STILL NEEDS SOME WORK REGARDING TABS vs THE DATA TYPE ID/WHOLE OBJECT PROBLEM

// func (ct *MediaType) Update(){
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

//   _, err = tx.Exec(`UPDATE media_type
//     SET alias = $1, description = $2, icon = $3, thumbnail = $4, meta = $5, tabs = $6
//     WHERE node_id = $7`, ct.Alias, ct.Description, ct.Icon, ct.Thumbnail, meta, tabs, ct.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()

//   tx.Commit()
// }
