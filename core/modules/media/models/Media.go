package models

import (
	//"fmt"
	"encoding/json"
	//"collexy/globals"
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	//"fmt"
	"time"
	//"net/http"
	"database/sql"
	"log"
	"strconv"
	"strings"
	//"reflect"
	//"errors"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	coremoduleuser "collexy/core/modules/user/models"
	//"github.com/kennygrant/sanitize"
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

type Media struct {
	Id                       int                        `json:"id"`
	Path                     string                     `json:"path"`
	ParentId                 *int                       `json:"parent_id,omitempty"`
	Name                     string                     `json:"name"`
	CreatedBy                int                        `json:"created_by"`
	CreatedDate              *time.Time                 `json:"created_date"`
	MediaTypeId              int                        `json:"media_type_id"`
	Meta                     map[string]interface{}     `json:"meta,omitempty"`
	PublicAccessMembers      map[string]interface{}     `json:"public_access_members,omitempty"`
	PublicAccessMemberGroups map[string]interface{}     `json:"public_access_member_groups,omitempty"`
	UserPermissions          map[string]*PermissionTest `json:"user_permissions,omitempty"`
	UserGroupPermissions     map[string]*PermissionTest `json:"user_group_permissions,omitempty"`
	// UserPermissions      []PermissionsContainer `json:"user_permissions,omitempty"`
	// UserGroupPermissions []PermissionsContainer `json:"user_group_permissions,omitempty"`
	// Additional fields (not persisted in db)
	Url              string                              `json:"url,omitempty"`
	FilePath         string                              `json:"file_path,omitempty"`
	Domains          []string                            `json:"domains,omitempty"`
	ParentMediaItems []*Media                            `json:"parent_media_items,omitempty"`
	ChildMediaItems  []*Media                            `json:"child_media_items,omitempty"`
	MediaType        *coremodulesettingsmodels.MediaType `json:"media_type,omitempty"`
	// Show bool `json:"show,omitempty"`
	// OldName string `json:"old_name,omitempty"`
}

func GetMedia(queryStringParams url.Values, user *coremoduleuser.User) (mediaSlice []Media) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta,
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions,
media_type.id AS ct_id, media_type.path AS ct_path, media_type.parent_id AS ct_parent_id,
media_type.name AS ct_name, media_type.alias AS ct_alias, media_type.created_by AS ct_created_by,
media_type.created_date AS ct_created_date, media_type.description AS ct_description,
media_type.icon AS ct_icon, media_type.thumbnail AS ct_thumbnail, media_type.meta AS ct_meta,
media_type.tabs AS ct_tabs 
FROM media 
   JOIN media_type ON media.media_type_id = media_type.id`

	// if ?type-id=x&levels=x(,x..)
	// else if ?type-id=x
	// else if ?levels=x(,x..)
	if queryStringParams.Get("levels") != "" {
		sqlStr = sqlStr + ` WHERE media.path ~ '*.*{` + queryStringParams.Get("levels") + `}'`
	}

	// if((queryStringParams.Get("type-id")!="" || queryStringParams.Get("type-id")!="") && queryStringParams.Get("media-type")!=""){
	if queryStringParams.Get("media-type") != "" {
		sqlStr = sqlStr + ` and media.media_type_id=` + queryStringParams.Get("media-type")
	}

	rows, err := db.Query(sqlStr)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var media_type_icon_str, media_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
			&media_created_date, &media_media_type_id, &media_meta,
			&media_user_permissions, &media_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			media_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			media_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		var parent_media_id_pointer *int = nil
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
			parent_media_id_pointer = &cpid
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		// var user_perm, user_group_perm []PermissionsContainer // map[string]PermissionsContainer
		var user_perm, user_group_perm map[string]*PermissionTest

		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(media_user_permissions, &user_perm)
		json.Unmarshal(media_user_group_permissions, &user_group_perm)

		var media_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(media_public_access, &public_access)

		json.Unmarshal(media_meta, &media_metaMap)

		var tabs []coremodulesettingsmodels.Tab
		var ct_metaMap map[string]interface{}

		json.Unmarshal(ct_tabs, &tabs)
		json.Unmarshal(ct_meta, &ct_metaMap)

		var accessGranted bool = false
		var accessDenied bool = false

		// if(err1 != nil){
		//   log.Println("Unmarshal Error: " + err1.Error())
		//   user_perm = nil
		// }
		userIdStr := strconv.Itoa(user.Id)
		// if permissions are set on the node for a specific user
		if media_user_permissions != nil {
			if user_perm[userIdStr] != nil {
				for i := 0; i < len(user_perm[userIdStr].Permissions); i++ {
					if accessGranted {
						break
					}
					if user_perm[userIdStr].Permissions[i] == "node_browse" {
						//fmt.Println("woauw it worked!")
						accessGranted = true
						media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
							media_media_type_id, media_metaMap, nil, nil, user_perm, nil, "", "", nil, nil, nil, &media_type}
						mediaSlice = append(mediaSlice, media)
						break
					}
				}
				if !accessGranted {
					accessDenied = true
				}
			}
		}

		// // if permissions are set on the node for a specific user
		// if media_user_permissions != nil {
		// 	for i := 0; i < len(user_perm); i++ {
		// 		if accessGranted {
		// 			break
		// 		}
		// 		if user_perm[i].Id == user.Id {
		// 			if accessGranted {
		// 				break
		// 			}
		// 			for j := 0; j < len(user_perm[i].Permissions); j++ {
		// 				if accessGranted {
		// 					break
		// 				}
		// 				if user_perm[i].Permissions[j] == "node_browse" {
		// 					//fmt.Println("woauw it worked!")
		// 					accessGranted = true
		// 					media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 					// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
		// 					media := Media{media_id, media_path, cpid, media_name, media_alias, media_created_by, media_created_date,
		// 						media_media_type_id, media_metaMap, public_access, user_perm, nil, media_type_id, "", nil, nil, nil, nil, &media_type}
		// 					mediaSlice = append(mediaSlice, media)
		// 					break
		// 				}
		// 			}
		// 			if !accessGranted {
		// 				accessDenied = true
		// 			}
		// 		}
		// 	}
		// }

		if !accessGranted && !accessDenied {
			// if no specific user node access has been specified, check node access per user_group
			if media_user_group_permissions != nil {
				for i := 0; i < len(user.UserGroupIds); i++ {
					if accessGranted {
						break
					}
					// for j := 0; j < len(user_group_perm); j++ {
					// 	if accessGranted {
					// 		break
					// 	}
					userGroupIdStr := strconv.Itoa(user.UserGroupIds[i])
					if user_group_perm[userGroupIdStr] != nil {
						if accessGranted {
							break
						}
						for j := 0; j < len(user_group_perm[userGroupIdStr].Permissions); j++ {
							if accessGranted {
								break
							}
							if user_group_perm[userGroupIdStr].Permissions[j] == "node_browse" {
								//fmt.Println("woauw it worked!")
								accessGranted = true
								media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
								media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
									media_media_type_id, media_metaMap, nil, nil, nil, user_group_perm, "", "", nil, nil, nil, &media_type}
								mediaSlice = append(mediaSlice, media)
								break
							}
						}
						if !accessGranted {
							accessDenied = true
						}
					}
					// }
				}
			}

		}

		// if !accessGranted && !accessDenied {
		// 	// if no specific user node access has been specified, check node access per user_group
		// 	if media_user_group_permissions != nil {
		// 		for i := 0; i < len(user.UserGroupIds); i++ {
		// 			if accessGranted {
		// 				break
		// 			}
		// 			for j := 0; j < len(user_group_perm); j++ {
		// 				if accessGranted {
		// 					break
		// 				}
		// 				if user_group_perm[j].Id == user.UserGroupIds[i] {
		// 					if accessGranted {
		// 						break
		// 					}
		// 					for k := 0; k < len(user_group_perm[j].Permissions); k++ {
		// 						if accessGranted {
		// 							break
		// 						}
		// 						if user_group_perm[j].Permissions[k] == "node_browse" {
		// 							//fmt.Println("woauw it worked!")
		// 							accessGranted = true
		// 							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 							media := Media{media_id, media_path, cpid, media_name, media_alias, media_created_by, media_created_date,
		// 								media_media_type_id, media_metaMap, public_access, nil, user_group_perm, media_type_id, "", nil, nil, nil, nil, &media_type}
		// 							mediaSlice = append(mediaSlice, media)
		// 							break
		// 						}
		// 					}
		// 					if !accessGranted {
		// 						accessDenied = true
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		// if no specific access has been granted per user_group either, use user groups default permissions
		if !accessGranted && !accessDenied {
			if user.UserGroups != nil {
				for i := 0; i < len(user.UserGroups); i++ {
					if accessGranted {
						break
					}
					for j := 0; j < len(user.UserGroups[i].Permissions); j++ {
						if user.UserGroups[i].Permissions[j] == "node_browse" {
							accessGranted = true
							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
							media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
								media_media_type_id, media_metaMap, nil, nil, nil, nil, "", "", nil, nil, nil, &media_type}
							mediaSlice = append(mediaSlice, media)
							break
						}
					}

				}
			}

		}
	}
	return
}

func GetMediaById(id int) (media Media) {

	db := coreglobals.Db

	sqlStr := `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta, 
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions,
  modified_media_type.id AS ct_id, modified_media_type.path AS ct_path, modified_media_type.parent_id AS ct_parent_id, modified_media_type.name as ct_name, modified_media_type.alias AS ct_alias,
  modified_media_type.created_by as ct_created_by, modified_media_type.description AS ct_description, modified_media_type.icon AS ct_icon, modified_media_type.thumbnail AS ct_thumbnail, 
  modified_media_type.meta::json AS ct_meta, modified_media_type.tabs AS ct_tabs, modified_media_type.allowed_media_types AS ct_allowed_media_types
FROM media
JOIN
LATERAL
(
  SELECT ct.*,pct.*  
  FROM media_type AS ct,
  -- Parent media types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS allowed_media_types
    FROM 
    (
      SELECT c.id, c.path, c.parent_id, c.name, c.alias, c.created_by, c.description, c.icon, c.thumbnail, c.meta
      FROM media_type AS c
      --where path @> subpath(ct.path,0,nlevel(ct.path)-1)
      --WHERE ct.meta->'allowed_media_type_ids' @> ('' || c.id || '')::jsonb
      WHERE c.id = ANY(ct.allowed_media_type_ids::int[]) 
    )res1
  ) pct
  
) modified_media_type
ON modified_media_type.id = media.media_type_id
WHERE media.id=$1`

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64

	var ct_path, ct_name, ct_alias, ct_description, ct_icon, ct_thumbnail string
	var ct_tabs, ct_meta []byte
	var ct_allowed_media_types []byte

	row := db.QueryRow(sqlStr, id)

	err := row.Scan(
		&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
		&media_created_date, &media_media_type_id, &media_meta,
		&media_user_permissions, &media_user_group_permissions,
		&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by,
		&ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_allowed_media_types)

	corehelpers.PanicIf(err)

	var ctpid int
	if ct_parent_id.Valid {
		// use s.String
		ctpid = int(ct_parent_id.Int64)
	} else {
		// NULL value
	}

	var cpid int
	var parent_media_id_pointer *int = nil
	if media_parent_id.Valid {
		cpid = int(media_parent_id.Int64)
		parent_media_id_pointer = &cpid
	}

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(media_user_permissions, &user_perm)
	json.Unmarshal(media_user_group_permissions, &user_group_perm)

	var allowed_media_types []coremodulesettingsmodels.MediaType
	var tabs []coremodulesettingsmodels.Tab
	var ct_metaMap map[string]interface{}
	var media_metaMap map[string]interface{}

	// var public_access *PublicAccess

	// json.Unmarshal(media_public_access, &public_access)

	json.Unmarshal(ct_allowed_media_types, &allowed_media_types)
	json.Unmarshal(ct_tabs, &tabs)
	json.Unmarshal(ct_meta, &ct_metaMap)
	json.Unmarshal(media_meta, &media_metaMap)

	media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, &time.Time{}, ct_description, ct_icon, ct_thumbnail, ct_metaMap, nil, nil, allowed_media_types, false, false, false, nil, nil, nil}

	media = Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
		media_media_type_id, media_metaMap, nil, nil, user_perm, user_group_perm, "", "", nil, nil, nil, &media_type}

	return
}

func GetMediaByIdChildren(id int, user *coremoduleuser.User) (mediaSlice []Media) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta, 
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions,
media_type.id AS ct_id, media_type.path AS ct_path, media_type.parent_id AS ct_parent_id,
media_type.name AS ct_name, media_type.alias AS ct_alias, media_type.created_by AS ct_created_by,
media_type.created_date AS ct_created_date, media_type.description AS ct_description,
media_type.icon AS ct_icon, media_type.thumbnail AS ct_thumbnail, media_type.meta AS ct_meta,
media_type.tabs AS ct_tabs 
FROM media
JOIN media_type ON media.media_type_id = media_type.id
WHERE media.parent_id=$1`

	rows, err := db.Query(sqlStr, id)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var media_type_icon_str, media_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
			&media_created_date, &media_media_type_id, &media_meta,
			&media_user_permissions, &media_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			media_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			media_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		var parent_media_id_pointer *int = nil
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
			parent_media_id_pointer = &cpid
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(media_user_permissions, &user_perm)
		json.Unmarshal(media_user_group_permissions, &user_group_perm)

		var media_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(media_public_access, &public_access)

		json.Unmarshal(media_meta, &media_metaMap)

		var tabs []coremodulesettingsmodels.Tab
		var ct_metaMap map[string]interface{}

		json.Unmarshal(ct_tabs, &tabs)
		json.Unmarshal(ct_meta, &ct_metaMap)

		var accessGranted bool = false
		var accessDenied bool = false

		// if(err1 != nil){
		//   log.Println("Unmarshal Error: " + err1.Error())
		//   user_perm = nil
		// }

		userIdStr := strconv.Itoa(user.Id)

		// if permissions are set on the node for a specific user
		if media_user_permissions != nil {
			if user_perm[userIdStr] != nil {
				for i := 0; i < len(user_perm[userIdStr].Permissions); i++ {
					if accessGranted {
						break
					}
					if user_perm[userIdStr].Permissions[i] == "node_browse" {
						//fmt.Println("woauw it worked!")
						accessGranted = true
						media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
							media_media_type_id, media_metaMap, nil, nil, user_perm, nil, "", "", nil, nil, nil, &media_type}
						mediaSlice = append(mediaSlice, media)
						break
					}
				}
				if !accessGranted {
					accessDenied = true
				}
			}
		}

		// if permissions are set on the node for a specific user
		// if media_user_permissions != nil {
		// 	for i := 0; i < len(user_perm); i++ {
		// 		if accessGranted {
		// 			break
		// 		}
		// 		if user_perm[i].Id == user.Id {
		// 			if accessGranted {
		// 				break
		// 			}
		// 			for j := 0; j < len(user_perm[i].Permissions); j++ {
		// 				if accessGranted {
		// 					break
		// 				}
		// 				if user_perm[i].Permissions[j] == "node_browse" {
		// 					//fmt.Println("woauw it worked!")
		// 					accessGranted = true
		// 					media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 					// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
		// 					media := Media{media_id, media_path, cpid, media_name, media_alias, media_created_by, media_created_date,
		// 						media_media_type_id, media_metaMap, public_access, user_perm, nil, media_type_id, "", nil, nil, nil, nil, &media_type}
		// 					mediaSlice = append(mediaSlice, media)
		// 					break
		// 				}
		// 			}
		// 			if !accessGranted {
		// 				accessDenied = true
		// 			}
		// 		}
		// 	}
		// }
		if !accessGranted && !accessDenied {
			// if no specific user node access has been specified, check node access per user_group
			if media_user_group_permissions != nil {
				for i := 0; i < len(user.UserGroupIds); i++ {
					if accessGranted {
						break
					}
					// for j := 0; j < len(user_group_perm); j++ {
					// 	if accessGranted {
					// 		break
					// 	}
					userGroupIdStr := strconv.Itoa(user.UserGroupIds[i])
					if user_group_perm[userGroupIdStr] != nil {
						if accessGranted {
							break
						}
						for j := 0; j < len(user_group_perm[userGroupIdStr].Permissions); j++ {
							if accessGranted {
								break
							}
							if user_group_perm[userGroupIdStr].Permissions[j] == "node_browse" {
								//fmt.Println("woauw it worked!")
								accessGranted = true
								media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
								media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
									media_media_type_id, media_metaMap, nil, nil, nil, user_group_perm, "", "", nil, nil, nil, &media_type}
								mediaSlice = append(mediaSlice, media)
								break
							}
						}
						if !accessGranted {
							accessDenied = true
						}
					}
					// }
				}
			}

		}
		// if !accessGranted && !accessDenied {
		// 	// if no specific user node access has been specified, check node access per user_group
		// 	if media_user_group_permissions != nil {
		// 		for i := 0; i < len(user.UserGroupIds); i++ {
		// 			if accessGranted {
		// 				break
		// 			}
		// 			for j := 0; j < len(user_group_perm); j++ {
		// 				if accessGranted {
		// 					break
		// 				}
		// 				if user_group_perm[j].Id == user.UserGroupIds[i] {
		// 					if accessGranted {
		// 						break
		// 					}
		// 					for k := 0; k < len(user_group_perm[j].Permissions); k++ {
		// 						if accessGranted {
		// 							break
		// 						}
		// 						if user_group_perm[j].Permissions[k] == "node_browse" {
		// 							//fmt.Println("woauw it worked!")
		// 							accessGranted = true
		// 							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 							media := Media{media_id, media_path, cpid, media_name, media_alias, media_created_by, media_created_date,
		// 								media_media_type_id, media_metaMap, public_access, nil, user_group_perm, media_type_id, "", nil, nil, nil, nil, &media_type}
		// 							mediaSlice = append(mediaSlice, media)
		// 							break
		// 						}
		// 					}
		// 					if !accessGranted {
		// 						accessDenied = true
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		// if no specific access has been granted per user_group either, use user groups default permissions
		if !accessGranted && !accessDenied {
			if user.UserGroups != nil {
				for i := 0; i < len(user.UserGroups); i++ {
					if accessGranted {
						break
					}
					for j := 0; j < len(user.UserGroups[i].Permissions); j++ {
						if user.UserGroups[i].Permissions[j] == "node_browse" {
							accessGranted = true
							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
							media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
								media_media_type_id, media_metaMap, nil, nil, nil, nil, "", "", nil, nil, nil, &media_type}
							mediaSlice = append(mediaSlice, media)
							break
						}
					}

				}
			}

		}
	}
	return
}

func GetMediaByIdParents(id int, user *coremoduleuser.User) (mediaSlice []Media) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta, 
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions,
media_type.id AS ct_id, media_type.path AS ct_path, media_type.parent_id AS ct_parent_id,
media_type.name AS ct_name, media_type.alias AS ct_alias, media_type.created_by AS ct_created_by,
media_type.created_date AS ct_created_date, media_type.description AS ct_description,
media_type.icon AS ct_icon, media_type.thumbnail AS ct_thumbnail, media_type.meta AS ct_meta,
media_type.tabs AS ct_tabs
FROM media 
JOIN media_type 
ON media.media_type_id = media_type.id
WHERE media.path @> 
(
	SELECT path
	FROM
	media
	WHERE
	id = $1
) ORDER BY media.path`

	rows, err := db.Query(sqlStr, id)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var media_type_icon_str, media_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
			&media_created_date, &media_media_type_id, &media_meta,
			&media_user_permissions, &media_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			media_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			media_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		var parent_media_id_pointer *int = nil
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
			parent_media_id_pointer = &cpid
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(media_user_permissions, &user_perm)
		json.Unmarshal(media_user_group_permissions, &user_group_perm)

		var media_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(media_public_access, &public_access)

		json.Unmarshal(media_meta, &media_metaMap)

		var tabs []coremodulesettingsmodels.Tab
		var ct_metaMap map[string]interface{}

		json.Unmarshal(ct_tabs, &tabs)
		json.Unmarshal(ct_meta, &ct_metaMap)

		var accessGranted bool = false
		var accessDenied bool = false

		// if(err1 != nil){
		//   log.Println("Unmarshal Error: " + err1.Error())
		//   user_perm = nil
		// }

		userIdStr := strconv.Itoa(user.Id)

		// if permissions are set on the node for a specific user
		if media_user_permissions != nil {
			if user_perm[userIdStr] != nil {
				for i := 0; i < len(user_perm[userIdStr].Permissions); i++ {
					if accessGranted {
						break
					}
					if user_perm[userIdStr].Permissions[i] == "node_browse" {
						//fmt.Println("woauw it worked!")
						accessGranted = true
						media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
							media_media_type_id, media_metaMap, nil, nil, user_perm, nil, "", "", nil, nil, nil, &media_type}
						mediaSlice = append(mediaSlice, media)
						break
					}
				}
				if !accessGranted {
					accessDenied = true
				}
			}
		}

		if !accessGranted && !accessDenied {
			// if no specific user node access has been specified, check node access per user_group
			if media_user_group_permissions != nil {
				for i := 0; i < len(user.UserGroupIds); i++ {
					if accessGranted {
						break
					}
					// for j := 0; j < len(user_group_perm); j++ {
					// 	if accessGranted {
					// 		break
					// 	}
					userGroupIdStr := strconv.Itoa(user.UserGroupIds[i])
					if user_group_perm[userGroupIdStr] != nil {
						if accessGranted {
							break
						}
						for j := 0; j < len(user_group_perm[userGroupIdStr].Permissions); j++ {
							if accessGranted {
								break
							}
							if user_group_perm[userGroupIdStr].Permissions[j] == "node_browse" {
								//fmt.Println("woauw it worked!")
								accessGranted = true
								media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
								media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
									media_media_type_id, media_metaMap, nil, nil, nil, user_group_perm, "", "", nil, nil, nil, &media_type}
								mediaSlice = append(mediaSlice, media)
								break
							}
						}
						if !accessGranted {
							accessDenied = true
						}
					}
					// }
				}
			}

		}

		// // if permissions are set on the node for a specific user
		// if media_user_permissions != nil {
		// 	for i := 0; i < len(user_perm); i++ {
		// 		if accessGranted {
		// 			break
		// 		}
		// 		if user_perm[i].Id == user.Id {
		// 			if accessGranted {
		// 				break
		// 			}
		// 			for j := 0; j < len(user_perm[i].Permissions); j++ {
		// 				if accessGranted {
		// 					break
		// 				}
		// 				if user_perm[i].Permissions[j] == "node_browse" {
		// 					//fmt.Println("woauw it worked!")
		// 					accessGranted = true
		// 					media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 					// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
		// 					media := Media{media_id, media_path, cpid, media_name, media_alias, media_created_by, media_created_date,
		// 						media_media_type_id, media_metaMap, public_access, user_perm, nil, media_type_id, "", nil, nil, nil, nil, &media_type}
		// 					mediaSlice = append(mediaSlice, media)
		// 					break
		// 				}
		// 			}
		// 			if !accessGranted {
		// 				accessDenied = true
		// 			}
		// 		}
		// 	}
		// }
		// if !accessGranted && !accessDenied {
		// 	// if no specific user node access has been specified, check node access per user_group
		// 	if media_user_group_permissions != nil {
		// 		for i := 0; i < len(user.UserGroupIds); i++ {
		// 			if accessGranted {
		// 				break
		// 			}
		// 			for j := 0; j < len(user_group_perm); j++ {
		// 				if accessGranted {
		// 					break
		// 				}
		// 				if user_group_perm[j].Id == user.UserGroupIds[i] {
		// 					if accessGranted {
		// 						break
		// 					}
		// 					for k := 0; k < len(user_group_perm[j].Permissions); k++ {
		// 						if accessGranted {
		// 							break
		// 						}
		// 						if user_group_perm[j].Permissions[k] == "node_browse" {
		// 							//fmt.Println("woauw it worked!")
		// 							accessGranted = true
		// 							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 							media := Media{media_id, media_path, cpid, media_name, media_alias, media_created_by, media_created_date,
		// 								media_media_type_id, media_metaMap, public_access, nil, user_group_perm, media_type_id, "", nil, nil, nil, nil, &media_type}
		// 							mediaSlice = append(mediaSlice, media)
		// 							break
		// 						}
		// 					}
		// 					if !accessGranted {
		// 						accessDenied = true
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		// if no specific access has been granted per user_group either, use user groups default permissions
		if !accessGranted && !accessDenied {
			if user.UserGroups != nil {
				for i := 0; i < len(user.UserGroups); i++ {
					if accessGranted {
						break
					}
					for j := 0; j < len(user.UserGroups[i].Permissions); j++ {
						if user.UserGroups[i].Permissions[j] == "node_browse" {
							accessGranted = true
							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
							media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
								media_media_type_id, media_metaMap, nil, nil, nil, nil, "", "", nil, nil, nil, &media_type}
							mediaSlice = append(mediaSlice, media)
							break
						}
					}

				}
			}

		}
	}
	return
}

// func DeleteMedia(id int){
//   db := coreglobals.Db

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)

//   _, err1 := tx.Exec("DELETE FROM media where node_id=$1", id)
//   corehelpers.PanicIf(err1)
//   _, err2 := tx.Exec("DELETE FROM node where id=$1", id)
//   corehelpers.PanicIf(err2)
//   //defer r2.Close()
//   err3 := tx.Commit()
//   corehelpers.PanicIf(err3)
// }

// func (t *Media) Post(){

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
//   err = tx.QueryRow(`SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`, t.Node.ParentId).Scan(&id, &path, &created_by, &name, &node_type, &created_date)
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
//   err = tx.QueryRow(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`, t.Node.Name, t.Node.NodeType, 1, t.Node.ParentId).Scan(&node_id)
//   //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateId)
//   //corehelpers.PanicIf(err)
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

//   _, err = tx.Exec("INSERT INTO media (node_id, media_type_node_id, meta) VALUES ($1, $2, $3)", node_id, t.MediaTypeId, meta)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()

//   if(t.Node.NodeType == 2){
//     var fi FileInfo
//     var fin FileNode
//     if(t.MediaTypeId == 40){
//       fi = FileInfo{t.Node.Name, 0, 0777 , time.Now(), true}
//       fin = FileNode{t.Meta["path"].(string), "", &fi, nil, "", true, ""}
//       //fin.Post()
//     } else {
//       fi = FileInfo{t.Node.Name, 0, 0777 , time.Now(), false}
//       fin = FileNode{t.Meta["path"].(string), "", &fi, nil, "", true, ""}
//     }
//     filePostErr := fin.Post()
//     if(filePostErr == nil){
//       err1 := tx.Commit()
//       corehelpers.PanicIf(err1)
//     }
//     // else {
//     //   fi = FileInfo{t.Node.Name, 0, 0777 , time.Time.Now(), false}
//     //   fin = FileNode{t.Meta.Path, "", fi, nil, "", true, ""}
//     // }
//   } else {
//       err1 := tx.Commit()
//       corehelpers.PanicIf(err1)

//   }

//   // // res, _ := json.Marshal(c)
//   // // log.Println(string(res))

//   // db := coreglobals.Db

//   // meta, _ := json.Marshal(c.Meta)

//   // tx, err := db.Begin()
//   // corehelpers.PanicIf(err)
//   // //defer tx.Rollback()

//   // _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", c.Node.Name, c.Node.Id)
//   // corehelpers.PanicIf(err)
//   // //defer r1.Close()

//   // _, err = tx.Exec(`UPDATE media
//   //   SET meta = $1
//   //   WHERE node_id = $2`, meta, c.Node.Id)
//   // corehelpers.PanicIf(err)
//   // //defer r2.Close()

//   // tx.Commit()
// }

// type Lol struct {
//   Id int64
//   NewPath string
// }

func (m *Media) Post() {
	var meta interface{} = nil

	var userPermissions interface{} = nil
	var userGroupPermissions interface{} = nil

	if m.Meta != nil {
		j, _ := json.Marshal(m.Meta)
		meta = j
	}

	if m.UserPermissions != nil {
		j, _ := json.Marshal(m.UserPermissions)
		userPermissions = j
	}

	if m.UserGroupPermissions != nil {
		j, _ := json.Marshal(m.UserGroupPermissions)
		userGroupPermissions = j
	}

	// http://godoc.org/github.com/lib/pq
	// pq does not support the LastInsertId() method of the Result type in database/sql.
	// To return the identifier of an INSERT (or UPDATE or DELETE),
	// use the Postgres RETURNING clause with a standard Query or QueryRow call:

	db := coreglobals.Db

	var parentMedia Media

	if m.ParentId != nil {

		// Channel c, is for getting the parent template
		// We need to append the id of the newly created template to the path of the parent id to create the new path
		c1 := make(chan Media)

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			c1 <- GetMediaById(*m.ParentId)
		}()

		go func() {
			for i := range c1 {
				fmt.Println(i)
				parentMedia = i
			}
		}()

		wg.Wait()
	}

	// This channel and WaitGroup is just to make sure the insert query is completed before we continue
	c2 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sqlStr := `INSERT INTO media ( 
			parent_id, name, created_by, media_type_id, 
			meta, 
			user_permissions, user_group_permissions) 
			VALUES (
				$1,$2,$3,$4,$5,$6,$7
				) RETURNING id`
		err1 := db.QueryRow(sqlStr, m.ParentId, m.Name, m.CreatedBy, m.MediaTypeId,
			meta,
			userPermissions, userGroupPermissions).Scan(&id)
		corehelpers.PanicIf(err1)

		c2 <- int(id)
	}()

	go func() {
		for i := range c2 {
			fmt.Println(i)
		}
	}()

	wg1.Wait()

	m.Id = int(id)

	// fmt.Println(parentTemplate.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE media 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if m.ParentId != nil {
		path = parentMedia.Path + "." + strconv.FormatInt(id, 10)
	}

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	c3 := make(chan []*Media)

	var wg2 sync.WaitGroup

	wg2.Add(1)

	go func() {
		defer wg2.Done()
		parents := GetMediaByIdParentsInternalUseOnly(m.Id)

		c3 <- parents
	}()

	go func() {
		for i := range c3 {
			fmt.Println(i)
			m.ParentMediaItems = append(m.ParentMediaItems, i...)
		}
	}()

	wg2.Wait()

	// add entries to media-access.xml if public access is set
	if m.PublicAccessMembers != nil || m.PublicAccessMemberGroups != nil {

		urlStr := "media"
		if len(m.ParentMediaItems) > 0 {
			for _, p := range m.ParentMediaItems {
				urlStr = urlStr + "/" + p.Name
			}
			m.Url = urlStr
		}

		fmt.Printf("url is: %s", m.Url)

		UpdatePublicAccessForMedia(m)
	}

	// get

	abspath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	abspath = abspath + "\\media"
	fmt.Println(abspath)

	if len(m.ParentMediaItems) > 0 {
		for _, p := range m.ParentMediaItems {
			abspath = abspath + "\\" + p.Name
		}
	}

	if m.Meta["attached_file"] == nil {
		// create directory 0777 permission too liberal?
		fmt.Println("creating directory: " + m.Name + "with path: " + abspath)
		err := os.Mkdir(abspath, 0644)
		if err != nil {
			panic(err)
		}
	}

	// if queryStringParams.Get("path") != "" {
	// 	path = path + "\\" + strings.Replace(queryStringParams.Get("path"),"%5C", "\\", -1)
	// }
	// fmt.Println(path)
	// if path != "" {}

	log.Println("media created successfully")

}

// Todo: If media is protected it should also be deleted in the XML file
func DeleteMedia(id int, queryStringParams url.Values) {
	// instead of sending query params with path, we could make a GetMediaById request here
	absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path := absPath
	fmt.Println(path)

	if queryStringParams.Get("path") != "" {
		path = path + "\\" + strings.Replace(queryStringParams.Get("path"), "%5C", "\\", -1)
	}
	fmt.Println(path)
	if path != "" {
		err1 := os.RemoveAll(path)

		if err1 != nil {
			fmt.Println(err1)

		}
	}

	var childMediaItems []Media

	c1 := make(chan []Media)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		c1 <- GetMediaByIdChildrenInternalUseOnly(id)
	}()

	go func() {
		for i := range c1 {
			fmt.Println(i)
			for _, child := range i {
				childMediaItems = append(childMediaItems, child)
			}
		}
	}()

	wg.Wait()

	db := coreglobals.Db

	sqlStr := `DELETE FROM media 
    WHERE id=any($1)`

	var deleteMediaIdsSlice coreglobals.IntSlice

	deleteMediaIdsSlice = append(deleteMediaIdsSlice, id)

	for _, child := range childMediaItems {
		deleteMediaIdsSlice = append(deleteMediaIdsSlice, child.Id)
	}

	fmt.Println(deleteMediaIdsSlice)

	// j, _ := json.Marshal(deleteMediaIdsSlice)

	//fmt.Println(string(j))
	j, _ := deleteMediaIdsSlice.Value()
	fmt.Println(j)

	_, err := db.Exec(sqlStr, j)

	corehelpers.PanicIf(err)

	//DeletePublicAccessForMedia(id)

	var deleteMediaIds map[int]int

	deleteMediaIds = make(map[int]int)
	deleteMediaIds[id] = id

	for _, child := range childMediaItems {
		deleteMediaIds[child.Id] = child.Id
	}

	DeletePublicAccessForMedia(deleteMediaIds)

	log.Printf("media with id %d was successfully deleted", id)
}

func DeletePublicAccessForMedia(mediaIds map[int]int) {
	if _, err := os.Stat("./config/media-access.xml"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("media-access.xml config file does not exist")
		} else {
			// other error
		}
	} else {

		configFile, err1 := os.Open("./config/media-access.xml")
		defer configFile.Close()
		if err1 != nil {
			log.Println("Error opening media-access.xml config file")
			//printError("opening config file", err1.Error())
		}

		XMLdata, err2 := ioutil.ReadAll(configFile) // use bufio intead since the xml can scale big

		fmt.Println(string(XMLdata))

		if err2 != nil {
			log.Println("Error reading from media-access.xml config file")
			fmt.Printf("error: %v", err2)
		}

		var v coreglobals.MediaAccessItems
		err := xml.Unmarshal(XMLdata, &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		//fmt.Printf("%#v\n", v)

		//coreglobals.MediaAccessConf = buildMap(v.Items...)
		mediaExists := false
		vv := v

		// reason for using a downward loop:
		// http://stackoverflow.com/questions/29005825/how-to-remove-element-of-struct-array-in-loop-in-golang
		for i := len(vv.Items) - 1; i >= 0; i-- {
			//for i, val := range vv.Items {

			if vv.Items[i].MediaId == mediaIds[vv.Items[i].MediaId] {
				fmt.Printf("v.Items[i].MediaId is: %#v\n", vv.Items[i].MediaId)
				fmt.Printf("mediaIds is: %#v\n", mediaIds)
				mediaExists = true
				vv.Items = append(vv.Items[:i], vv.Items[i+1:]...)
			}
		}

		if mediaExists {
			b, _ := xml.MarshalIndent(&vv, "", "    ")
			fmt.Println(string(b))

			// open output file
			fo, err := os.Create("./config/media-access.xml")
			if err != nil {
				panic(err)
			}
			// close fo on exit and check for its returned error
			defer func() {
				if err := fo.Close(); err != nil {
					panic(err)
				}
			}()

			w := bufio.NewWriter(fo)
			n4, err4 := w.WriteString(string(b))
			if err4 != nil {
				panic(err4)
			}
			fmt.Printf("wrote %d bytes\n", n4)
			w.Flush()
		}

		//fmt.Printf("%#v\n", v)
	}
}

func (m *Media) Put(oldMedia Media) {

	var meta interface{} = nil

	var userPermissions interface{} = nil
	var userGroupPermissions interface{} = nil

	if m.Meta != nil {
		j, _ := json.Marshal(m.Meta)
		meta = j
	}

	if m.UserPermissions != nil {
		j, _ := json.Marshal(m.UserPermissions)
		userPermissions = j
	}

	if m.UserGroupPermissions != nil {
		j, _ := json.Marshal(m.UserGroupPermissions)
		userGroupPermissions = j
	}

	var parentMedia Media

	if m.ParentId != nil {

		c1 := make(chan Media)

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			c1 <- GetMediaById(*m.ParentId)
		}()

		go func() {
			for i := range c1 {
				fmt.Println(i)
				parentMedia = i
			}
		}()

		wg.Wait()
	}

	db := coreglobals.Db

	sqlStr := `UPDATE media 
	SET path=$1, parent_id=$2, name=$3, created_by=$4, media_type_id=$5, 
	meta=$6, user_permissions=$7, user_group_permissions=$8
	WHERE id=$9;`

	path := strconv.Itoa(m.Id)
	if m.ParentId != nil {
		path = parentMedia.Path + "." + strconv.Itoa(m.Id)
	}

	_, err := db.Exec(sqlStr, path, m.ParentId, m.Name, m.CreatedBy, m.MediaTypeId,
		meta, userPermissions, userGroupPermissions, m.Id)

	corehelpers.PanicIf(err)

	// if m.PublicAccessMembers != nil || m.PublicAccessMemberGroups != nil{
	c3 := make(chan []*Media)

	var wg2 sync.WaitGroup

	wg2.Add(1)

	go func() {
		defer wg2.Done()
		parents := GetMediaByIdParentsInternalUseOnly(m.Id)

		c3 <- parents
	}()

	go func() {
		for i := range c3 {
			fmt.Println(i)
			m.ParentMediaItems = append(m.ParentMediaItems, i...)
		}
	}()

	wg2.Wait()

	urlStr := "media"
	if m.ParentMediaItems != nil {
		if len(m.ParentMediaItems) > 0 {
			for _, p := range m.ParentMediaItems {
				urlStr = urlStr + "/" + p.Name
			}
			m.Url = urlStr
		}
	}

	UpdatePublicAccessForMedia(m)

	// here it would be awesome to also be able to send originalData, so we had a from > to directory name
	// this could generally save us from a lot of additional database queries elsewhere
	// and old name field in the Media struct would in this instance be enough though
	if m.Name != oldMedia.Name {

		oldAbsFilePath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "\\" + oldMedia.FilePath)
		newAbsFilePath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "\\" + m.FilePath)

		fmt.Println("renaming fs entity: " + oldAbsFilePath + " to: " + newAbsFilePath)

		err := os.Rename(oldAbsFilePath, newAbsFilePath)
		if err != nil {
			panic(err)
		}
	}

	// }

	log.Println("media updated successfully")
}

// https://github.com/golang/go/issues/3688
func UpdatePublicAccessForMedia(m *Media) {
	if _, err := os.Stat("./config/media-access.xml"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("media-access.xml config file does not exist")
		} else {
			// other error
		}
	} else {

		configFile, err1 := os.Open("./config/media-access.xml")
		defer configFile.Close()
		if err1 != nil {
			log.Println("Error opening media-access.xml config file")
			//printError("opening config file", err1.Error())
		}

		XMLdata, err2 := ioutil.ReadAll(configFile) // use bufio intead since the xml can scale big

		fmt.Println(string(XMLdata))

		if err2 != nil {
			log.Println("Error reading from media-access.xml config file")
			fmt.Printf("error: %v", err2)
		}

		var v coreglobals.MediaAccessItems
		err := xml.Unmarshal(XMLdata, &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		//fmt.Printf("%#v\n", v)

		//coreglobals.MediaAccessConf = buildMap(v.Items...)
		mediaExists := false
		vv := v

		// reason for using a downward loop:
		// http://stackoverflow.com/questions/29005825/how-to-remove-element-of-struct-array-in-loop-in-golang
		for i := len(vv.Items) - 1; i >= 0; i-- {
			//for i, val := range vv.Items {
			//fmt.Printf("v.Items val is: %#v\n", val)

			if vv.Items[i].MediaId == m.Id {
				mediaExists = true
				if (m.PublicAccessMembers == nil || len(m.PublicAccessMembers) == 0) && (m.PublicAccessMemberGroups == nil || len(m.PublicAccessMemberGroups) == 0) {
					vv.Items = append(vv.Items[:i], vv.Items[i+1:]...)
					//vv.Items = append(vv.Items[:i], vv.Items[i+1:]...)
				} else {
					if m.PublicAccessMemberGroups != nil {
						fmt.Printf("val.MediaId:%d\n", vv.Items[i].MediaId)
						fmt.Printf("m.Id:%d\n", m.Id)
						fmt.Println("m.PublicAccessMemberGroups != nil\n")
						var memberGroups []int
						for key, _ := range m.PublicAccessMemberGroups {
							intVal, _ := strconv.Atoi(key)
							memberGroups = append(memberGroups, intVal)
						}
						vv.Items[i].MemberGroups = memberGroups
						fmt.Printf("val.memberGroups is: %d\n", vv.Items[i].MemberGroups)
						fmt.Printf("memberGroups is: %d\n", memberGroups)
					}
					if m.PublicAccessMembers != nil {
						var members []int
						for key, _ := range m.PublicAccessMembers {
							intVal, _ := strconv.Atoi(key)
							members = append(members, intVal)
						}
						vv.Items[i].Members = members
					}
				}
				url := strings.Replace(m.FilePath, "\\", "/", -1)
				if vv.Items[i].Url != url {
					vv.Items[i].Url = url
				}

			}
		}
		// if vv.Items[i].Path (9.16.17) contains m.path (9.16)
		// currentPathLevels = 3
		// mLevels = 2
		// mIsOnLevel = 1
		// url := strings.Replace(m.FilePath, "\\", "/", -1)
		// if vv.Items[i].Url substring until mIsOnLevel (1st dash) != url (regex, split?)
		//replace it

		// for i, val := range vv.Items {
		// 	//fmt.Printf("v.Items val is: %#v\n", val)

		// 	if val.MediaId == m.Id {
		// 		mediaExists = true
		// 		if m.PublicAccessMembers == nil && m.PublicAccessMemberGroups == nil{
		// 			vv.Items = append(vv.Items[:i], vv.Items[i+1:]...)
		// 		} else {
		// 			if m.PublicAccessMemberGroups != nil {
		// 				fmt.Printf("val.MediaId:%d\n", val.MediaId)
		// 				fmt.Printf("m.Id:%d\n", m.Id)
		// 				fmt.Println("m.PublicAccessMemberGroups != nil\n")
		// 				var memberGroups []int
		// 				for key, _ := range m.PublicAccessMemberGroups {
		// 					intVal, _ := strconv.Atoi(key)
		// 					memberGroups = append(memberGroups, intVal)
		// 				}
		// 				val.MemberGroups = memberGroups
		// 				fmt.Printf("val.memberGroups is: %d\n", val.MemberGroups)
		// 				fmt.Printf("memberGroups is: %d\n", memberGroups)
		// 			}
		// 			if m.PublicAccessMembers != nil {
		// 				var members []int
		// 				for key, _ := range m.PublicAccessMembers {
		// 					intVal, _ := strconv.Atoi(key)
		// 					members = append(members, intVal)
		// 				}
		// 				val.Members = members
		// 			}
		// 		}

		// 	}
		// }

		if !mediaExists {
			if (m.PublicAccessMembers != nil || len(m.PublicAccessMembers) > 0) || (m.PublicAccessMemberGroups != nil || len(m.PublicAccessMemberGroups) > 0) {
				var members []int
				for key, _ := range m.PublicAccessMembers {
					intVal, _ := strconv.Atoi(key)
					members = append(members, intVal)
				}

				var memberGroups []int
				for key, _ := range m.PublicAccessMemberGroups {
					intVal, _ := strconv.Atoi(key)
					memberGroups = append(memberGroups, intVal)
				}

				mai := coreglobals.MediaAccessItem{
					xml.Name{Local: "item"},
					m.Id,
					m.Url,
					0,
					0,
					members,
					memberGroups,
				}
				vv.Items = append(vv.Items, &mai)
			}

		}

		b, _ := xml.MarshalIndent(&vv, "", "    ")
		fmt.Println(string(b))

		// open output file
		fo, err := os.Create("./config/media-access.xml")
		if err != nil {
			panic(err)
		}
		// close fo on exit and check for its returned error
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		w := bufio.NewWriter(fo)
		n4, err4 := w.WriteString(string(b))
		if err4 != nil {
			panic(err4)
		}
		fmt.Printf("wrote %d bytes\n", n4)
		w.Flush()
		//fmt.Printf("%#v\n", v)
	}
}

// func (c *Media) Update(){

//   // res, _ := json.Marshal(c)
//   // log.Println(string(res))

//   db := coreglobals.Db

//   meta, _ := json.Marshal(c.Meta)

//   tx, err := db.Begin()
//   corehelpers.PanicIf(err)
//   //defer tx.Rollback()

//   _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", c.Node.Name, c.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r1.Close()

//   _, err = tx.Exec(`UPDATE media
//     SET meta = $1
//     WHERE node_id = $2`, meta, c.Node.Id)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()
//   if(c.Node.NodeType == 2){
//     //originalPath := "media\\Another Image Folder"
//     //originalNodeName := "Another Image Folder"
//     originalNodeName := c.Node.OldName
//     fmt.Println("Original Node Name: " + originalNodeName);

//     // rename filesystem folder that has this original url (btw make a hidden input field holding the old url) with c.Node.Name
//     folderNode := GetFilesystemNodeById("media", originalNodeName)
//     folderNode.FullPath = c.Meta["path"].(string)
//     //folderNode.OldPath = originalPath
//     //folderNode.FullPath = "media\\Another Image Folder1"
//     folderNode.Update()
//     fmt.Println("TEST ::: TEST ::: ERR (node_id: ")
//     fmt.Println(c.Node.Id)

//     // if media is of media type: folder
//     if(c.MediaTypeId == 16){

//       // check if node has children (SQL SELECT QUERY USING LTREE PATH)
//       rows, err101 := tx.Query(`SELECT media.node_id as node_id, meta as media_meta
//         FROM media
//         JOIN node
//         ON node.id = media.node_id
//         WHERE node.path <@ '1.` + strconv.Itoa(c.Node.Id) + `' AND node.path != '1.` + strconv.Itoa(c.Node.Id) + `'`)
//       //, strconv.Itoa(c.Node.Id), strconv.Itoa(c.Node.Id)
//       // if has children, iterate them
//       if err101 != nil {
//         log.Fatal(err101)
//       }
//       defer rows.Close()
//       var res []Lol
//       // foreach child node
//       fmt.Println("TEST ::: TEST ::: ERR1")
//       for rows.Next() {
//         fmt.Println("TEST ::: TEST ::: ERR2")
//         var node_id int64
//         var media_meta_byte_arr []byte

//         if err := rows.Scan(&node_id, &media_meta_byte_arr); err != nil {
//           log.Fatal(err)
//         }

//         var media_meta map[string]interface{}
//         json.Unmarshal([]byte(string(media_meta_byte_arr)), &media_meta)

//         var path string = media_meta["path"].(string)
//         var newPath string = strings.Replace(path, folderNode.OldPath, folderNode.FullPath, -1)
//         // update node's media.meta.url part where substing equals oldurl - with c.Meta.url
//         fmt.Println("TEST ::: TEST ::: ERR3")

//         res = append(res,Lol{node_id, newPath})
//         // _, err102 := tx.Exec(`UPDATE media
//         //   SET meta = json_object_update_key(meta::json, 'url', '$1'::text)::jsonb
//         //   WHERE node_id=$2`, newUrl, node_id)
//         // corehelpers.PanicIf(err102)
//       }
//       if err101 := rows.Err(); err101 != nil {
//         log.Fatal(err101)
//       }
//       fmt.Println("TEST ::: TEST ::: ERR4")
//       for i := 0; i < len(res); i++ {
//         fmt.Println(fmt.Sprintf("newpath: %s, node id: %v", res[i].NewPath, res[i].Id))
//         _, err102 := tx.Exec(`UPDATE media
//           SET meta = json_object_update_key(meta::json, 'path', $1::text)::jsonb
//           WHERE node_id=$2`, string(res[i].NewPath), res[i].Id)
//         corehelpers.PanicIf(err102)
//       }

//     }
//   }

//   tx.Commit()
// }

func GetBackendMediaById(id int, protectedMedia *coreglobals.MediaAccessItem) (media Media) {
	db := coreglobals.Db
	queryStr := `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta, 
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions, 
  modified_media_type.id AS ct_id, modified_media_type.path AS ct_path, modified_media_type.parent_id AS ct_parent_id, modified_media_type.name as ct_name, modified_media_type.alias AS ct_alias,
  modified_media_type.created_by as ct_created_by, modified_media_type.description AS ct_description, modified_media_type.icon AS ct_icon, modified_media_type.thumbnail AS ct_thumbnail, modified_media_type.meta::json AS ct_meta, modified_media_type.ct_tabs AS ct_tabs, modified_media_type.parent_media_types AS ct_parent_media_types, modified_media_type.composite_media_types AS ct_composite_media_types, 
  filepath.fpath AS media_file_path
FROM media
-- file path
JOIN LATERAL
  (
    select id, fpath from media AS m1,
    LATERAL
	  (
	    SELECT array_to_string(array_agg(m2.name),'\') AS fpath
	    FROM LATERAL ( 
		SELECT *
		FROM media as m3
		WHERE m3.path @> media.path
		ORDER BY m3.path ) m2 
	  ) AS lolcat
    WHERE m1.id = media.id
  ) filepath
ON filepath.id = media.id
JOIN
LATERAL
(
  SELECT ct.*,pct.*,cct.*,ct_tabs_with_dt.*
  FROM media_type AS ct,
  -- Parent media types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS parent_media_types
    FROM 
    (
      SELECT c.id, c.path, c.parent_id, c.name, c.alias, c.created_by, c.description, c.icon, c.thumbnail, c.meta, gf.* AS tabs 
      FROM media_type AS c,
      LATERAL 
      (
        SELECT json_agg(row1) AS tabs 
        FROM 
        (
          SELECT y.name, ss.properties
          FROM json_to_recordset (
            (
              SELECT * 
              FROM json_to_recordset(
                (
                  SELECT json_agg(ggg)
                  FROM 
                  (
                    SELECT ct.tabs
                    FROM media_type AS ct
                    WHERE ct.id=c.id
                  )ggg
                )
              ) AS x(tabs json)
            )
          ) AS y(name text, properties json),
          LATERAL 
          (
            SELECT json_agg(
		    json_build_object
		    (
		      'name',row.name,
		      'order',row."order",
		      'data_type_id',row.data_type_id,
		      'data_type', json_build_object
		      (
			'id',row.data_type_id, 
			'name',row.data_type_name, 
			'alias',row.data_type_alias, 
			'created_by',row.data_type_created_by, 
			--'created_date',row.data_type_created_date,
			'html', row.data_type_html,
			'editor_alias', row.data_type_editor_alias,
			'meta', row.data_type_meta
		      ), 
		      'help_text', row.help_text, 
		      'description', row.description
		    )
            ) AS properties
	    FROM(
		SELECT k.name, "order",data_type_id, data_type.name as data_type_name, 
		data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.html AS data_type_html, data_type.editor_alias as data_type_editor_alias,
		data_type.meta as data_type_meta, help_text, description
		FROM json_to_recordset(properties) 
		AS k(name text, "order" int, data_type_id int, help_text text, description text)
		JOIN data_type
		ON data_type.id = k.data_type_id
	    )row
          ) ss
        )row1
      ) gf
      where path @> subpath(ct.path,0,nlevel(ct.path)-1)
    )res1
  ) pct,
    -- Composite media types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS composite_media_types
    FROM 
    (
      SELECT c.id, c.path, c.parent_id, c.name, c.alias, c.created_by, c.description, c.icon, c.thumbnail, c.meta, gf.* AS tabs 
      FROM media_type AS c,
      LATERAL 
      (
        SELECT json_agg(row1) AS tabs 
        FROM 
        (
          SELECT y.name, ss.properties
          FROM json_to_recordset (
            (
              SELECT * 
              FROM json_to_recordset(
                (
                  SELECT json_agg(ggg)
                  FROM 
                  (
                    SELECT ct.tabs
                    FROM media_type AS ct
                    WHERE ct.id=c.id
                  )ggg
                )
              ) AS x(tabs json)
            )
          ) AS y(name text, properties json),
          LATERAL 
          (
            SELECT json_agg(
		    json_build_object
		    (
		      'name',row.name,
		      'order',row."order",
		      'data_type_id',row.data_type_id,
		      'data_type', json_build_object
		      (
			'id',row.data_type_id, 
			'name',row.data_type_name, 
			'alias',row.data_type_alias, 
			'created_by',row.data_type_created_by, 
			--'created_date',row.data_type_created_date,
			'html', row.data_type_html,
			'editor_alias', row.data_type_editor_alias,
			'meta', row.data_type_meta
		      ), 
		      'help_text', row.help_text, 
		      'description', row.description
		    )
            ) AS properties
	    FROM(
		SELECT k.name, "order",data_type_id, data_type.name as data_type_name, 
		data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.html AS data_type_html, data_type.editor_alias as data_type_editor_alias,
		data_type.meta as data_type_meta, help_text, description
		FROM json_to_recordset(properties) 
		AS k(name text, "order" int, data_type_id int, help_text text, description text)
		JOIN data_type
		ON data_type.id = k.data_type_id
	    )row
          ) ss
        )row1
      ) gf
      --where path @> subpath(ct.path,0,nlevel(ct.path)-1)
      WHERE id = ANY(ct.composite_media_type_ids)
    )res1
  ) cct,
  -- Tabs
  LATERAL 
  (
    SELECT res2.tabs AS ct_tabs
    FROM 
    (
      SELECT c.id AS cid, gf.* AS tabs
      FROM media_type AS c,
      LATERAL 
      (
        SELECT json_agg(row1) AS tabs 
        FROM
        (
          SELECT y.name, ss.properties
          FROM json_to_recordset
          (
            (
              SELECT * 
              FROM json_to_recordset(
                (
                  SELECT json_agg(ggg)
                  FROM
                  (
                    SELECT ct.tabs
                    FROM media_type AS ct
                    WHERE ct.id=c.id
                  )ggg
                )
              ) AS x(tabs json)
            )
          ) AS y(name text, properties json),
          LATERAL 
          (
            SELECT json_agg(
		    json_build_object
		    (
		      'name',row.name,
		      'order',row."order",
		      'data_type_id',row.data_type_id,
		      'data_type', json_build_object
		      (
			'id',row.data_type_id, 
			'name',row.data_type_name, 
			'alias',row.data_type_alias, 
			'created_by',row.data_type_created_by, 
			--'created_date',row.data_type_created_date,
			'html', row.data_type_html,
			'editor_alias', row.data_type_editor_alias,
			'meta', row.data_type_meta
		      ), 
		      'help_text', row.help_text, 
		      'description', row.description
		    )
            ) AS properties
	    FROM(
		SELECT k.name, "order",data_type_id, data_type.name as data_type_name, 
		data_type.alias AS data_type_alias, data_type.created_by as data_type_created_by, data_type.html AS data_type_html, data_type.editor_alias as data_type_editor_alias,
		data_type.meta as data_type_meta, help_text, description
		FROM json_to_recordset(properties) 
		AS k(name text, "order" int, data_type_id int, help_text text, description text)
		JOIN data_type
		ON data_type.id = k.data_type_id
	    )row
          ) ss
        )row1
      ) gf
      WHERE c.id = ct.id
    )res2
    limit 1
  ) ct_tabs_with_dt
  --
) modified_media_type
ON modified_media_type.id = media.media_type_id
WHERE media.id=$1`
	// queryStr :=
	// `SELECT my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date, my_node.parent_id as media_parent_id,
	//   media.id as media_id, media.node_id as media_node_id, media.media_type_node_id as media_media_type_node_id, media.meta as media_meta,
	//   res.id as ct_id, res.node_id as ct_node_id, res.parent_media_type_node_id as ct_parent_media_type_node_id, res.alias as ct_alias,
	//   res.description as ct_description, res.icon as ct_icon, res.thumbnail as ct_thumbnail, res.meta::json as ct_meta, res.ct_tabs as ct_tabs, res.parent_media_types as ct_parent_media_types
	//   FROM media
	//   JOIN node as my_node
	//   ON my_node.id = media.node_id
	//   JOIN
	//   LATERAL
	//   (
	//     SELECT my_media_type.*,ffgd.*,gf2.*
	//     FROM media_type as my_media_type, node as my_media_type_node,
	//     LATERAL
	//     (
	//         SELECT array_to_json(array_agg(okidoki)) as parent_media_types
	//         FROM (
	//           SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.thumbnail, c.parent_media_type_node_id, c.meta, gf.* as tabs
	//           FROM media_type as c, node,
	//         LATERAL (
	//             select json_agg(row1) as tabs from((
	//             select y.name, ss.properties
	//             from json_to_recordset(
	//           (
	//               select *
	//               from json_to_recordset(
	//             (
	//                 SELECT json_agg(ggg)
	//                 from(
	//               SELECT tabs
	//               FROM
	//               (
	//                   SELECT *
	//                   FROM media_type as ct
	//                   WHERE ct.id=c.id
	//               ) dsfds

	//                 )ggg
	//             )
	//               ) as x(tabs json)
	//           )
	//             ) as y(name text, properties json),
	//             LATERAL (
	//           select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id',row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id',row.data_type_node_id, 'alias', row.data_type_alias,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
	//           from(
	//               select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
	//               from json_to_recordset(properties)
	//               as k(name text, "order" int, data_type_node_id int, help_text text, description text)
	//               JOIN data_type
	//               ON data_type.node_id = k.data_type_node_id
	//               )row
	//             ) ss
	//             ))row1
	//         ) gf
	//           where path @> subpath(my_media_type_node.path,0,nlevel(my_media_type_node.path)-1) and c.node_id = node.id
	//         )okidoki
	//     ) ffgd,
	//     --
	//     LATERAL
	//     (
	//         SELECT okidoki.tabs as ct_tabs
	//         FROM (
	//           SELECT c.id as cid, gf.* as tabs
	//           FROM media_type as c, node,
	//         LATERAL (
	//             select json_agg(row1) as tabs from((
	//         select y.name, ss.properties
	//         from json_to_recordset(
	//         (
	//       select *
	//       from json_to_recordset(
	//           (
	//         SELECT json_agg(ggg)
	//         from(
	//       SELECT tabs
	//       FROM
	//       (
	//           SELECT *
	//           FROM media_type as ct
	//           WHERE ct.id=c.id
	//       ) dsfds

	//         )ggg
	//           )
	//       ) as x(tabs json)
	//         )
	//         ) as y(name text, properties json),
	//         LATERAL (
	//       select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id', row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id', row.data_type_node_id, 'alias', row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
	//       from(
	//     select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
	//     from json_to_recordset(properties)
	//     as k(name text, "order" int, data_type_node_id int, help_text text, description text)
	//     JOIN data_type
	//     ON data_type.node_id = k.data_type_node_id
	//     )row
	//         ) ss
	//             ))row1
	//         ) gf
	//         WHERE c.id = my_media_type.id
	//         )okidoki
	//         limit 1
	//     ) gf2
	//     --
	//     WHERE my_media_type_node.id = my_media_type.node_id
	//   ) res
	//   ON res.node_id = media.media_type_node_id
	//   WHERE my_node.id=$1`

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64

	var ct_path, ct_name, ct_alias, ct_description, ct_icon, ct_thumbnail string
	var ct_tabs, ct_meta []byte
	var ct_parent_media_types, ct_composite_media_types []byte

	var media_file_path sql.NullString

	row := db.QueryRow(queryStr, id)

	err := row.Scan(
		&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
		&media_created_date, &media_media_type_id, &media_meta,
		&media_user_permissions, &media_user_group_permissions,
		&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by,
		&ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_parent_media_types, &ct_composite_media_types,
		&media_file_path)

	corehelpers.PanicIf(err)

	var ctpid int
	if ct_parent_id.Valid {
		// use s.String
		ctpid = int(ct_parent_id.Int64)
	} else {
		// NULL value
	}

	var cpid int
	var parent_media_id_pointer *int = nil
	if media_parent_id.Valid {
		cpid = int(media_parent_id.Int64)
		parent_media_id_pointer = &cpid
	}

	var media_file_path_string string = "media"
	if media_file_path.Valid {
		media_file_path_string = media_file_path_string + "\\" + media_file_path.String
	}

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(media_user_permissions, &user_perm)
	json.Unmarshal(media_user_group_permissions, &user_group_perm)

	var parent_media_types, composite_media_types []coremodulesettingsmodels.MediaType
	var tabs []coremodulesettingsmodels.Tab
	var ct_metaMap map[string]interface{}
	var media_metaMap map[string]interface{}

	// var public_access *PublicAccess

	// json.Unmarshal(media_public_access, &public_access)

	json.Unmarshal(ct_parent_media_types, &parent_media_types)
	json.Unmarshal(ct_composite_media_types, &composite_media_types)
	json.Unmarshal(ct_tabs, &tabs)
	json.Unmarshal(ct_meta, &ct_metaMap)
	json.Unmarshal(media_meta, &media_metaMap)

	// if _, err := os.Stat("./config/media-access.json"); err != nil {
	// 	if os.IsNotExist(err) {
	// 		// file does not exist
	// 		log.Println("media-access.json config file does not exist")
	// 	} else {
	// 		// other error
	// 	}
	// } else {

	// 	configFile, err1 := os.Open("./config/media-access.json")
	// 	defer configFile.Close()
	// 	if err1 != nil {
	// 		log.Println("Error opening media-access.json config file")
	// 		//printError("opening config file", err1.Error())
	// 	}

	// 	jsonParser := json.NewDecoder(configFile)
	// 	if err1 = jsonParser.Decode(coreglobals.MediaAccessConf); err1 != nil {
	// 		log.Println("Error parsing media-access.json config file")
	// 		//printError("parsing config file", err1.Error())
	// 	}
	// 	fmt.Println(coreglobals.MediaAccessConf)
	// 	// log.Println(coreglobals.Maccess.Items[0].Domains[0])
	// 	// log.Println(coreglobals.Maccess.Items[0].Url)
	// 	// fmt.Println(coreglobals.Maccess.Items[0].MemberGroups)
	// }
	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	if protectedMedia != nil {
		//fmt.Printf("v.Items is: %#v\n", protectedMedia)
		public_access_member_groups = make(map[string]interface{})
		for _, mgId := range protectedMedia.MemberGroups {
			fmt.Printf("mgIdlolol: %d", mgId)
			mgIdStr := strconv.Itoa(mgId)
			//log.Println("mgidstr: " + mgIdStr)

			public_access_member_groups[mgIdStr] = true
		}

		public_access_members = make(map[string]interface{})
		for _, mgId := range protectedMedia.Members {
			fmt.Printf("mgIdlolol: %d", mgId)
			mgIdStr := strconv.Itoa(mgId)
			//log.Println("mgidstr: " + mgIdStr)

			public_access_members[mgIdStr] = true
		}
	}

	media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, &time.Time{}, ct_description, ct_icon, ct_thumbnail, ct_metaMap, tabs, parent_media_types, nil, false, false, false, nil, nil, composite_media_types}

	media = Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
		media_media_type_id, media_metaMap, public_access_members, public_access_member_groups, user_perm, user_group_perm, "", media_file_path_string, nil, nil, nil, &media_type}

	return
}

// INTERNAL USE ONLY FUNCTIONS (temporary, these should not be necessary  - just to get something working quickly)

func GetMediaByIdParentsInternalUseOnly(id int) (mediaSlice []*Media) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta, 
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions,
media_type.id AS ct_id, media_type.path AS ct_path, media_type.parent_id AS ct_parent_id,
media_type.name AS ct_name, media_type.alias AS ct_alias, media_type.created_by AS ct_created_by,
media_type.created_date AS ct_created_date, media_type.description AS ct_description,
media_type.icon AS ct_icon, media_type.thumbnail AS ct_thumbnail, media_type.meta AS ct_meta,
media_type.tabs AS ct_tabs
FROM media 
JOIN media_type 
ON media.media_type_id = media_type.id
WHERE media.path @> 
(
	SELECT path
	FROM
	media
	WHERE
	id = $1
) ORDER BY media.path`

	rows, err := db.Query(sqlStr, id)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var media_type_icon_str, media_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
			&media_created_date, &media_media_type_id, &media_meta,
			&media_user_permissions, &media_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			media_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			media_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		var parent_media_id_pointer *int = nil
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
			parent_media_id_pointer = &cpid
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(media_user_permissions, &user_perm)
		json.Unmarshal(media_user_group_permissions, &user_group_perm)

		var media_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(media_public_access, &public_access)

		json.Unmarshal(media_meta, &media_metaMap)

		var tabs []coremodulesettingsmodels.Tab
		var ct_metaMap map[string]interface{}

		json.Unmarshal(ct_tabs, &tabs)
		json.Unmarshal(ct_meta, &ct_metaMap)

		media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
		media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
			media_media_type_id, media_metaMap, nil, nil, nil, nil, "", "", nil, nil, nil, &media_type}
		mediaSlice = append(mediaSlice, &media)

	}
	return
}

func GetMediaByIdChildrenInternalUseOnly(id int) (mediaSlice []Media) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT media.id AS media_id, media.path AS media_path, media.parent_id AS media_parent_id,
media.name AS media_name, media.created_by AS media_created_by, 
media.created_date AS media_created_date, media.media_type_id AS media_media_type_id,
media.meta AS media_meta, 
media.user_permissions AS media_user_permissions, media.user_group_permissions AS media_user_group_permissions,
media_type.id AS ct_id, media_type.path AS ct_path, media_type.parent_id AS ct_parent_id,
media_type.name AS ct_name, media_type.alias AS ct_alias, media_type.created_by AS ct_created_by,
media_type.created_date AS ct_created_date, media_type.description AS ct_description,
media_type.icon AS ct_icon, media_type.thumbnail AS ct_thumbnail, media_type.meta AS ct_meta,
media_type.tabs AS ct_tabs 
FROM media
JOIN media_type ON media.media_type_id = media_type.id
WHERE media.parent_id=$1`

	rows, err := db.Query(sqlStr, id)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var media_id, media_created_by, media_media_type_id int
	var media_path, media_name string
	var media_parent_id sql.NullInt64
	var media_created_date *time.Time
	var media_meta, media_user_permissions, media_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var media_type_icon_str, media_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
			&media_created_date, &media_media_type_id, &media_meta,
			&media_user_permissions, &media_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			media_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			media_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		var parent_media_id_pointer *int = nil
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
			parent_media_id_pointer = &cpid
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(media_user_permissions, &user_perm)
		json.Unmarshal(media_user_group_permissions, &user_group_perm)

		var media_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(media_public_access, &public_access)

		json.Unmarshal(media_meta, &media_metaMap)

		var tabs []coremodulesettingsmodels.Tab
		var ct_metaMap map[string]interface{}

		json.Unmarshal(ct_tabs, &tabs)
		json.Unmarshal(ct_meta, &ct_metaMap)

		// if(err1 != nil){
		//   log.Println("Unmarshal Error: " + err1.Error())
		//   user_perm = nil
		// }

		media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
		media := Media{media_id, media_path, parent_media_id_pointer, media_name, media_created_by, media_created_date,
			media_media_type_id, media_metaMap, nil, nil, nil, nil, "", "", nil, nil, nil, &media_type}
		mediaSlice = append(mediaSlice, media)

	}
	return
}
