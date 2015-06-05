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
	// "strings"
	//"reflect"
	//"errors"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	coremoduleuser "collexy/core/modules/user/models"
	//"github.com/kennygrant/sanitize"
	"net/url"
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"bufio"
)

type Media struct {
	Id                       int                        `json:"id"`
	Path                     string                     `json:"path"`
	ParentId                 int                        `json:"parent_id,omitempty"`
	Name                     string                     `json:"name"`
	CreatedBy                int                        `json:"created_by"`
	CreatedDate              *time.Time                 `json:"created_date"`
	MediaTypeId              int                        `json:"media_type_id"`
	Meta                     map[string]interface{}     `json:"meta,omitempty"`
	PublicAccessMembers      map[string]interface{}    `json:"public_access_members,omitempty"`
	PublicAccessMemberGroups map[string]interface{}     `json:"public_access_member_groups,omitempty"`
	UserPermissions          map[string]*PermissionTest `json:"user_permissions,omitempty"`
	UserGroupPermissions     map[string]*PermissionTest `json:"user_group_permissions,omitempty"`
	// UserPermissions      []PermissionsContainer `json:"user_permissions,omitempty"`
	// UserGroupPermissions []PermissionsContainer `json:"user_group_permissions,omitempty"`
	// Additional fields (not persisted in db)
	Url              string                              `json:"url,omitempty"`
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
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
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
						media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
							media_media_type_id, media_metaMap, nil, nil, user_perm, nil, "", nil, nil, nil, &media_type}
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
								media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
								media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
									media_media_type_id, media_metaMap, nil, nil, nil, user_group_perm, "", nil, nil, nil, &media_type}
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
							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
							media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
								media_media_type_id, media_metaMap, nil, nil, nil, nil, "", nil, nil, nil, &media_type}
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

	var media_type_parent_id int
	if ct_parent_id.Valid {
		// use s.String
		media_type_parent_id = int(ct_parent_id.Int64)
	} else {
		// NULL value
	}

	var cpid int
	if media_parent_id.Valid {
		cpid = int(media_parent_id.Int64)
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

	media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, media_type_parent_id, ct_name, ct_alias, ct_created_by, &time.Time{}, ct_description, ct_icon, ct_thumbnail, ct_metaMap, nil, nil, allowed_media_types, false, false, false, nil, nil, nil}

	media = Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
		media_media_type_id, media_metaMap, nil, nil, user_perm, user_group_perm, "", nil, nil, nil, &media_type}

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
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
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
						media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
							media_media_type_id, media_metaMap, nil, nil, user_perm, nil, "", nil, nil, nil, &media_type}
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
								media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
								media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
									media_media_type_id, media_metaMap, nil, nil, nil, user_group_perm, "", nil, nil, nil, &media_type}
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
							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
							media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
								media_media_type_id, media_metaMap, nil, nil, nil, nil, "", nil, nil, nil, &media_type}
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
		if media_parent_id.Valid {
			cpid = int(media_parent_id.Int64)
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
						media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
							media_media_type_id, media_metaMap, nil, nil, user_perm, nil, "", nil, nil, nil, &media_type}
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
								media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
								media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
									media_media_type_id, media_metaMap, nil, nil, nil, user_group_perm, "", nil, nil, nil, &media_type}
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
							media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, media_type_icon_str, media_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil}
							media := Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
								media_media_type_id, media_metaMap, nil, nil, nil, nil, "", nil, nil, nil, &media_type}
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

func (c *Media) Update() {

	// db := coreglobals.Db

	// meta, _ := json.Marshal(c.Meta)

	// sqlStr := `UPDATE media 
	// SET name=$1, meta=$3 
 // 	WHERE id=$4;`

	// _, err := db.Exec(sqlStr, c.Name, meta, c.Id)

	// corehelpers.PanicIf(err)

	UpdatePublicAccessForMedia(c)

	log.Println("media updated successfully")
}

func UpdatePublicAccessForMedia(m *Media){
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
		vv := v
		for _, val := range vv.Items{
			//fmt.Printf("v.Items val is: %#v\n", val)
			
			if(val.MediaId == m.Id){

				if(m.PublicAccessMemberGroups != nil){
					fmt.Printf("val.MediaId:%d\n", val.MediaId)
					fmt.Printf("m.Id:%d\n", m.Id)
					fmt.Println("m.PublicAccessMemberGroups != nil\n")
					var memberGroups []int
					for key, _ := range m.PublicAccessMemberGroups {
						intVal,_ := strconv.Atoi(key)
						memberGroups = append(memberGroups, intVal)
					}
					val.MemberGroups = memberGroups
					fmt.Printf("val.memberGroups is: %d\n", val.MemberGroups)
					fmt.Printf("memberGroups is: %d\n", memberGroups)
				}
				if(m.PublicAccessMembers != nil){
					var members []int
					for key, _ := range m.PublicAccessMembers {
						intVal,_ := strconv.Atoi(key)
						members = append(members, intVal)
					}
					val.Members = members
				}
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
  modified_media_type.created_by as ct_created_by, modified_media_type.description AS ct_description, modified_media_type.icon AS ct_icon, modified_media_type.thumbnail AS ct_thumbnail, modified_media_type.meta::json AS ct_meta, modified_media_type.ct_tabs AS ct_tabs, modified_media_type.parent_media_types AS ct_parent_media_types, modified_media_type.composite_media_types AS ct_composite_media_types 
FROM media
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
			'path',row.data_type_path, 
			'parent_id', row.data_type_parent_id,
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
		SELECT k.name, "order",data_type_id, data_type.path as data_type_path, data_type.parent_id as data_type_parent_id, data_type.name as data_type_name, 
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
			'path',row.data_type_path, 
			'parent_id', row.data_type_parent_id,
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
		SELECT k.name, "order",data_type_id, data_type.path as data_type_path, data_type.parent_id as data_type_parent_id, data_type.name as data_type_name, 
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
			'path',row.data_type_path, 
			'parent_id', row.data_type_parent_id,
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
		SELECT k.name, "order",data_type_id, data_type.path as data_type_path, data_type.parent_id as data_type_parent_id, data_type.name as data_type_name, 
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

	row := db.QueryRow(queryStr, id)

	err := row.Scan(
		&media_id, &media_path, &media_parent_id, &media_name, &media_created_by,
		&media_created_date, &media_media_type_id, &media_meta,
		&media_user_permissions, &media_user_group_permissions,
		&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by,
		&ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_parent_media_types, &ct_composite_media_types)

	corehelpers.PanicIf(err)

	var media_type_parent_id int
	if ct_parent_id.Valid {
		// use s.String
		media_type_parent_id = int(ct_parent_id.Int64)
	} else {
		// NULL value
	}

	var cpid int
	if media_parent_id.Valid {
		cpid = int(media_parent_id.Int64)
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

	if protectedMedia != nil{
		//fmt.Printf("v.Items is: %#v\n", protectedMedia)
		public_access_member_groups = make(map[string]interface{})
		for _, mgId := range protectedMedia.MemberGroups{
			fmt.Printf("mgIdlolol: %d", mgId)
			mgIdStr := strconv.Itoa(mgId)
			//log.Println("mgidstr: " + mgIdStr)
			
			public_access_member_groups[mgIdStr] = true;
		}
	}

	media_type := coremodulesettingsmodels.MediaType{ct_id, ct_path, media_type_parent_id, ct_name, ct_alias, ct_created_by, &time.Time{}, ct_description, ct_icon, ct_thumbnail, ct_metaMap, tabs, parent_media_types, nil, false, false, false, nil, nil, composite_media_types}

	media = Media{media_id, media_path, cpid, media_name, media_created_by, media_created_date,
		media_media_type_id, media_metaMap, public_access_members, public_access_member_groups, user_perm, user_group_perm, "", nil, nil, nil, &media_type}

	return
}
