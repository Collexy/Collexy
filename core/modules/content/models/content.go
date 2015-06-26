package models

import (
	//"fmt"
	"encoding/json"
	//"collexy/globals"
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"fmt"
	"time"
	//"net/http"
	"database/sql"
	"html/template"
	"log"
	"strconv"
	// "strings"
	"reflect"
	//"errors"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	coremoduleuser "collexy/core/modules/user/models"
	"github.com/kennygrant/sanitize"
	"net/url"
	"sync"
)

type Content struct {
	Id            int                    `json:"id"`
	Path          string                 `json:"path"`
	ParentId      *int                   `json:"parent_id,omitempty"`
	Name          string                 `json:"name"`
	CreatedBy     int                    `json:"created_by"`
	CreatedDate   *time.Time             `json:"created_date"`
	ContentTypeId int                    `json:"content_type_id"`
	TemplateId    *int                   `json:"template_id,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
	// PublicAccess         *PublicAccess          `json:"public_access,omitempty"`
	PublicAccessMembers      map[string]interface{}     `json:"public_access_members,omitempty"`
	PublicAccessMemberGroups map[string]interface{}     `json:"public_access_member_groups,omitempty"`
	UserPermissions          map[string]*PermissionTest `json:"user_permissions,omitempty"`
	UserGroupPermissions     map[string]*PermissionTest `json:"user_group_permissions,omitempty"`
	// UserPermissions      []PermissionsContainer `json:"user_permissions,omitempty"`
	// UserGroupPermissions []PermissionsContainer `json:"user_group_permissions,omitempty"`
	// Additional fields (not persisted in db)
	Url                string                                `json:"url,omitempty"`
	Domains            []string                              `json:"domains,omitempty"`
	ParentContentItems []*Content                            `json:"parent_content_items,omitempty"`
	ChildContentItems  []*Content                            `json:"child_content_items,omitempty"`
	Template           *coremodulesettingsmodels.Template    `json:"template,omitempty"`
	ContentType        *coremodulesettingsmodels.ContentType `json:"content_type,omitempty"`
	// Show bool `json:"show,omitempty"`
	// OldName string `json:"old_name,omitempty"`
}

func GetContent(queryStringParams url.Values, user *coremoduleuser.User) (contentSlice []Content) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id, 
content.meta AS content_meta, 
content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
content_type.id AS ct_id, content_type.path AS ct_path, content_type.parent_id AS ct_parent_id,
content_type.name AS ct_name, content_type.alias AS ct_alias, content_type.created_by AS ct_created_by,
content_type.created_date AS ct_created_date, content_type.description AS ct_description,
content_type.icon AS ct_icon, content_type.thumbnail AS ct_thumbnail, content_type.meta AS ct_meta,
content_type.tabs AS ct_tabs
FROM content 
   JOIN content_type ON content.content_type_id = content_type.id`

	// if ?type-id=x&levels=x(,x..)
	// else if ?type-id=x
	// else if ?levels=x(,x..)
	if queryStringParams.Get("levels") != "" {
		sqlStr = sqlStr + ` WHERE content.path ~ '*.*{` + queryStringParams.Get("levels") + `}'`
	}

	// if((queryStringParams.Get("type-id")!="" || queryStringParams.Get("type-id")!="") && queryStringParams.Get("content-type")!=""){
	if queryStringParams.Get("content-type") != "" {
		sqlStr = sqlStr + ` and content.content_type_id=` + queryStringParams.Get("content-type")
	}

	rows, err := db.Query(sqlStr)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var content_type_icon_str, content_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			content_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			content_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var template_id int
		if content_template_id.Valid {
			template_id = int(content_template_id.Int64)
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		// var user_perm, user_group_perm []PermissionsContainer // map[string]PermissionsContainer
		var user_perm, user_group_perm map[string]*PermissionTest

		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		var content_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		json.Unmarshal(content_meta, &content_metaMap)

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
		if content_user_permissions != nil {
			if user_perm[userIdStr] != nil {
				for i := 0; i < len(user_perm[userIdStr].Permissions); i++ {
					if accessGranted {
						break
					}
					if user_perm[userIdStr].Permissions[i] == "node_browse" {
						//fmt.Println("woauw it worked!")
						accessGranted = true
						content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
							content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, user_perm, nil, "", nil, nil, nil, nil, &content_type}
						contentSlice = append(contentSlice, content)
						break
					}
				}
				if !accessGranted {
					accessDenied = true
				}
			}
		}

		// // if permissions are set on the node for a specific user
		// if content_user_permissions != nil {
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
		// 					content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 					// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
		// 					content := Content{content_id, content_path, cpid, content_name, content_created_by, content_created_date,
		// 						content_content_type_id, template_id, content_metaMap, public_access, user_perm, nil, content_type_id, "", nil, nil, nil, nil, &content_type}
		// 					contentSlice = append(contentSlice, content)
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
			if content_user_group_permissions != nil {
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
								content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
								content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
									content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, nil, user_group_perm, "", nil, nil, nil, nil, &content_type}
								contentSlice = append(contentSlice, content)
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
		// 	if content_user_group_permissions != nil {
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
		// 							content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 							content := Content{content_id, content_path, cpid, content_name, content_created_by, content_created_date,
		// 								content_content_type_id, template_id, content_metaMap, public_access, nil, user_group_perm, content_type_id, "", nil, nil, nil, nil, &content_type}
		// 							contentSlice = append(contentSlice, content)
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
							content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
							content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
								content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, nil, nil, "", nil, nil, nil, nil, &content_type}
							contentSlice = append(contentSlice, content)
							break
						}
					}

				}
			}

		}
	}
	return
}

func GetContentById(id int) (content Content) {

	db := coreglobals.Db

	sqlStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
  modified_content_type.id AS ct_id, modified_content_type.path AS ct_path, modified_content_type.parent_id AS ct_parent_id, modified_content_type.name as ct_name, modified_content_type.alias AS ct_alias,
  modified_content_type.created_by as ct_created_by, modified_content_type.description AS ct_description, modified_content_type.icon AS ct_icon, modified_content_type.thumbnail AS ct_thumbnail, 
  modified_content_type.meta::json AS ct_meta, modified_content_type.tabs AS ct_tabs, modified_content_type.allowed_content_types AS ct_allowed_content_types 
FROM content
JOIN
LATERAL
(
  SELECT ct.*,pct.*  
  FROM content_type AS ct,
  -- Parent content types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS allowed_content_types
    FROM 
    (
      SELECT c.id, c.path, c.parent_id, c.name, c.alias, c.created_by, c.description, c.icon, c.thumbnail, c.meta
      FROM content_type AS c
      --where path @> subpath(ct.path,0,nlevel(ct.path)-1)
      --WHERE ct.meta->'allowed_content_type_ids' @> ('' || c.id || '')::jsonb
      WHERE c.id = ANY(ct.allowed_content_type_ids::int[]) 
    )res1
  ) pct
  
) modified_content_type
ON modified_content_type.id = content.content_type_id
WHERE content.id=$1`

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64

	var ct_path, ct_name, ct_alias, ct_description, ct_icon, ct_thumbnail string
	var ct_tabs, ct_meta []byte
	var ct_allowed_content_types []byte

	row := db.QueryRow(sqlStr, id)

	err := row.Scan(
		&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
		&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
		&content_user_permissions, &content_user_group_permissions,
		&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by,
		&ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_allowed_content_types)

	corehelpers.PanicIf(err)

	var ctpid int
	if ct_parent_id.Valid {
		// use s.String
		ctpid = int(ct_parent_id.Int64)
	} else {
		// NULL value
	}

	var cpid int
	if content_parent_id.Valid {
		cpid = int(content_parent_id.Int64)
	}

	var template_id int
	if content_template_id.Valid {
		template_id = int(content_template_id.Int64)
	}

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(content_user_permissions, &user_perm)
	json.Unmarshal(content_user_group_permissions, &user_group_perm)

	var allowed_content_types []coremodulesettingsmodels.ContentType
	var tabs []coremodulesettingsmodels.Tab
	var ct_metaMap map[string]interface{}
	var content_metaMap map[string]interface{}

	// var public_access *PublicAccess

	// json.Unmarshal(content_public_access, &public_access)

	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	json.Unmarshal(content_public_access_members, &public_access_members)
	json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

	json.Unmarshal(ct_allowed_content_types, &allowed_content_types)
	json.Unmarshal(ct_tabs, &tabs)
	json.Unmarshal(ct_meta, &ct_metaMap)
	json.Unmarshal(content_meta, &content_metaMap)

	content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, &time.Time{}, ct_description, ct_icon, ct_thumbnail, ct_metaMap, nil, nil, allowed_content_types, false, false, false, nil, nil, nil, 0, nil}

	content = Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
		content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, user_perm, user_group_perm, "", nil, nil, nil, nil, &content_type}

	return
}

func GetContentByIdChildren(id int, user *coremoduleuser.User) (contentSlice []Content) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
content_type.id AS ct_id, content_type.path AS ct_path, content_type.parent_id AS ct_parent_id,
content_type.name AS ct_name, content_type.alias AS ct_alias, content_type.created_by AS ct_created_by,
content_type.created_date AS ct_created_date, content_type.description AS ct_description,
content_type.icon AS ct_icon, content_type.thumbnail AS ct_thumbnail, content_type.meta AS ct_meta,
content_type.tabs AS ct_tabs 
FROM content 
   JOIN content_type ON content.content_type_id = content_type.id
   WHERE content.parent_id=$1`

	rows, err := db.Query(sqlStr, id)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var content_type_icon_str, content_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			content_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			content_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var template_id int
		if content_template_id.Valid {
			template_id = int(content_template_id.Int64)
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		var content_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		json.Unmarshal(content_meta, &content_metaMap)

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
		if content_user_permissions != nil {
			if user_perm[userIdStr] != nil {
				for i := 0; i < len(user_perm[userIdStr].Permissions); i++ {
					if accessGranted {
						break
					}
					if user_perm[userIdStr].Permissions[i] == "node_browse" {
						//fmt.Println("woauw it worked!")
						accessGranted = true
						content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
							content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, user_perm, nil, "", nil, nil, nil, nil, &content_type}
						contentSlice = append(contentSlice, content)
						break
					}
				}
				if !accessGranted {
					accessDenied = true
				}
			}
		}

		// if permissions are set on the node for a specific user
		// if content_user_permissions != nil {
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
		// 					content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 					// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
		// 					content := Content{content_id, content_path, cpid, content_name, content_created_by, content_created_date,
		// 						content_content_type_id, template_id, content_metaMap, public_access, user_perm, nil, content_type_id, "", nil, nil, nil, nil, &content_type}
		// 					contentSlice = append(contentSlice, content)
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
			if content_user_group_permissions != nil {
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
								content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
								content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
									content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, nil, user_group_perm, "", nil, nil, nil, nil, &content_type}
								contentSlice = append(contentSlice, content)
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
		// 	if content_user_group_permissions != nil {
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
		// 							content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 							content := Content{content_id, content_path, cpid, content_name, content_created_by, content_created_date,
		// 								content_content_type_id, template_id, content_metaMap, public_access, nil, user_group_perm, content_type_id, "", nil, nil, nil, nil, &content_type}
		// 							contentSlice = append(contentSlice, content)
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
							content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
							content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
								content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, nil, nil, "", nil, nil, nil, nil, &content_type}
							contentSlice = append(contentSlice, content)
							break
						}
					}

				}
			}

		}
	}
	return
}

func GetContentByIdParents(id int, user *coremoduleuser.User) (contentSlice []Content) {

	db := coreglobals.Db
	sqlStr := ""
	// if(queryStringParams.Get("type-id") != nil){
	sqlStr = `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
content_type.id AS ct_id, content_type.path AS ct_path, content_type.parent_id AS ct_parent_id,
content_type.name AS ct_name, content_type.alias AS ct_alias, content_type.created_by AS ct_created_by,
content_type.created_date AS ct_created_date, content_type.description AS ct_description,
content_type.icon AS ct_icon, content_type.thumbnail AS ct_thumbnail, content_type.meta AS ct_meta,
content_type.tabs AS ct_tabs 
FROM content 
JOIN content_type 
ON content.content_type_id = content_type.id
WHERE content.path @> 
(
	SELECT path
	FROM
	content
	WHERE
	id = $1
) ORDER BY content.path`

	rows, err := db.Query(sqlStr, id)
	corehelpers.PanicIf(err)
	defer rows.Close()

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id sql.NullInt64
	var ct_created_date *time.Time
	var ct_path, ct_name, ct_alias, ct_description string
	var ct_tabs, ct_meta []byte
	var ct_icon, ct_thumbnail sql.NullString

	for rows.Next() {
		var content_type_icon_str, content_type_thumbnail_str string

		// if(queryStringParams.Get("type-id")!=nil){
		err := rows.Scan(&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by, &ct_created_date, &ct_description, &ct_icon,
			&ct_thumbnail, &ct_meta, &ct_tabs)

		corehelpers.PanicIf(err)

		if ct_icon.Valid {
			content_type_icon_str = ct_icon.String
		}

		if ct_thumbnail.Valid {
			content_type_thumbnail_str = ct_thumbnail.String
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var template_id int
		if content_template_id.Valid {
			template_id = int(content_template_id.Int64)
		}

		var ctpid int
		if ct_parent_id.Valid {
			ctpid = int(ct_parent_id.Int64)
		}

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		var content_metaMap map[string]interface{}

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		json.Unmarshal(content_meta, &content_metaMap)

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
		if content_user_permissions != nil {
			if user_perm[userIdStr] != nil {
				for i := 0; i < len(user_perm[userIdStr].Permissions); i++ {
					if accessGranted {
						break
					}
					if user_perm[userIdStr].Permissions[i] == "node_browse" {
						//fmt.Println("woauw it worked!")
						accessGranted = true
						content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
						// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
						content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
							content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, user_perm, nil, "", nil, nil, nil, nil, &content_type}
						contentSlice = append(contentSlice, content)
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
			if content_user_group_permissions != nil {
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
								content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
								content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
									content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, nil, user_group_perm, "", nil, nil, nil, nil, &content_type}
								contentSlice = append(contentSlice, content)
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
		// if content_user_permissions != nil {
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
		// 					content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 					// node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", user_perm, nil, ""}
		// 					content := Content{content_id, content_path, cpid, content_name, content_created_by, content_created_date,
		// 						content_content_type_id, template_id, content_metaMap, public_access, user_perm, nil, content_type_id, "", nil, nil, nil, nil, &content_type}
		// 					contentSlice = append(contentSlice, content)
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
		// 	if content_user_group_permissions != nil {
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
		// 							content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, ct_type_id, false, false, false, nil, nil, nil}
		// 							content := Content{content_id, content_path, cpid, content_name, content_created_by, content_created_date,
		// 								content_content_type_id, template_id, content_metaMap, public_access, nil, user_group_perm, content_type_id, "", nil, nil, nil, nil, &content_type}
		// 							contentSlice = append(contentSlice, content)
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
							content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, ct_created_date, ct_description, content_type_icon_str, content_type_thumbnail_str, ct_metaMap, tabs, nil, nil, false, false, false, nil, nil, nil, 0, nil}
							content := Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
								content_content_type_id, &template_id, content_metaMap, public_access_members, public_access_member_groups, nil, nil, "", nil, nil, nil, nil, &content_type}
							contentSlice = append(contentSlice, content)
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
//   sql := `SELECT id, path, created_by, name, type_id, created_date FROM node`

//   // if ?type-id=x&levels=x(,x..)
//   // else if ?type-id=x
//   // else if ?levels=x(,x..)
//   if(queryStringParams.Get("type-id") != "" && queryStringParams.Get("levels") != ""){
//       sql = sql + ` WHERE type_id=` + queryStringParams.Get("type-id") + ` and content.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
//   } else if(queryStringParams.Get("type-id") != "" && queryStringParams.Get("levels")==""){
//       sql = sql + ` WHERE type_id=` + queryStringParams.Get("type-id")
//   } else if(queryStringParams.Get("type-id") == "" && queryStringParams.Get("levels") != ""){
//       sql = sql + ` WHERE content.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
//   }

//   rows, err := db.Query(sql)
//   corehelpers.PanicIf(err)
//   defer rows.Close()

//   var id, created_by, type_id int
//   var path, name string
//   var created_date time.Time

//   for rows.Next(){
//       err := rows.Scan(&id, &path, &created_by, &name, &type_id, &created_date)
//       corehelpers.PanicIf(err)
//       node := Node{id, path, created_by, name, type_id, &created_date, 0, nil,nil,false, "", nil, nil}
//       nodes = append(nodes,node)
//   }
//   return
// }

// func GetNodeByIdChildren(id int) (nodes []Node){
//   fmt.Println("GETNODEIDBYCHILDREN")
//   db := coreglobals.Db

//   //querystr := "SELECT id, path, created_by, name, type_id, created_date FROM node WHERE parent_id=$1"

//   querystr := `SELECT node.id, node.path, node.created_by, node.name, node.type_id, node.created_date, content_type.icon
// FROM node
// LEFT OUTER JOIN content
// ON content.node_id = node.id
// LEFT OUTER JOIN content_type
// ON content.content_type_node_id = content_type.node_id
// WHERE parent_id=$1`

//   rows, err := db.Query(querystr, id)
//   corehelpers.PanicIf(err)
//   defer rows.Close()

//   var created_by, type_id int
//   var path, name string
//   var created_date time.Time

//   var content_type_icon sql.NullString

//   for rows.Next(){
//     var content_type_icon_str string

//       err := rows.Scan(&id, &path, &created_by, &name, &type_id, &created_date, &content_type_icon)
//       corehelpers.PanicIf(err)

//       if(content_type_icon.Valid){
//         content_type_icon_str = content_type_icon.String
//       } else {
//         // NULL
//       }

//       node := Node{id,path,created_by, name, type_id, &created_date, 0, nil, nil, false, "", nil, nil, content_type_icon_str}

//       nodes = append(nodes, node)

//   }
//   return
// }

//func (c *Content) TimeAgo(ti *time.Time) (t interface{}){
func (c *Content) TimeAgo() (t interface{}) {
	// See http://golang.org/pkg/time/#Parse
	//timeFormat := "2006-01-02 15:04 MST"

	var then time.Time = *c.CreatedDate
	//var then time.Time = *ti

	//fmt.Println(then.Format(time.RFC3339))
	// then, err := time.Parse(timeFormat, v)
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }

	duration := time.Since(then)
	if duration.Seconds() > 59 {
		fmt.Println("time >59 seconds")
		if duration.Minutes() > 59 {
			fmt.Println("time >59 minutes")
			if duration.Hours() > 72 {
				fmt.Println("time >72 hours")
				rofl := then.Format("Mon Jan _2, 2006")
				t = rofl
			} else {
				t = strconv.FormatFloat(duration.Hours(), 'f', 0, 64) + " hours ago"
			}
		} else {
			t = strconv.FormatFloat(duration.Minutes(), 'f', 0, 64) + " minutes ago"
		}
	} else {
		t = strconv.FormatFloat(duration.Seconds(), 'f', 0, 64) + " seconds ago"
	}

	return
}

func (c *Content) StripHtmlTags(str string) (strippedStr string) {
	strippedStr = sanitize.HTML(str)
	return
}

func (c *Content) GetSubstring(s string, start, offset int) (str string) {
	if offset < len(s) {
		str = s[start:offset]
	} else {
		str = s
	}
	return
}

func (c *Content) GetContentByDepth(start, offset, length int) (contentSlice []*Content) {
	db := coreglobals.Db

	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) > 1
  )oki1
) heh
ON heh.id = content.id
WHERE content.path ~ (ltree2text(subltree($1,$2,$3))||'.*{,'||$4::text||'}')::lquery`

	rows, err := db.Query(queryStr, c.Path, start, offset, length)
	corehelpers.PanicIf(err)
	defer rows.Close()

	//row := db.QueryRow(queryStr, paramId)
	for rows.Next() {
		var content_id, content_created_by, content_content_type_id int
		var content_path, content_name string
		var content_parent_id, content_template_id sql.NullInt64
		var content_created_date *time.Time
		var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
		var content_url sql.NullString
		var content_domains coreglobals.StringSlice

		var template_id, template_created_by int
		var template_path, template_name, template_alias string
		var template_parent_id sql.NullInt64
		var template_is_partial bool
		var template_parent_templates []byte

		rows.Scan(
			&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&content_url,
			&content_domains,
			&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
			&template_created_by, &template_is_partial, &template_parent_templates)

		/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
		//corehelpers.PanicIf(err)
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No content with that url.")
		case err != nil:
			log.Fatal(err)
		default:
			fmt.Printf("content domains is %v\n", content_domains)
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var contentTemplateId int
		if content_template_id.Valid {
			contentTemplateId = int(content_template_id.Int64)
		}

		var content_url_str string
		if content_url.Valid {
			content_url_str = content_url.String
		}

		var tpid int
		if template_parent_id.Valid {
			tpid = int(template_parent_id.Int64)
		}

		var parent_templates_final []*coremodulesettingsmodels.Template
		var meta map[string]interface{}

		json.Unmarshal(template_parent_templates, &parent_templates_final)
		json.Unmarshal(content_meta, &meta)

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

		content := &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
			content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
			content_url_str, content_domains, nil, nil, &template, nil}
		// content := &Content{content_id, content_node_id, content_content_type_node_id, meta, contentNode, coremodulesettingsmodels.ContentType{}, &template, nil, content_url_str, content_domains,nil}
		contentSlice = append(contentSlice, content)

	}
	return
}

func (c *Content) GetLinkedContent(metaKey string, metaValue int) (contentSlice []*Content) {
	metaValueStr := strconv.Itoa(metaValue)

	db := coreglobals.Db

	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) > 1
  )oki1
) heh
ON heh.id = content.id
WHERE content.meta->$1 @> $2;`

	rows, err := db.Query(queryStr, metaKey, metaValueStr)
	corehelpers.PanicIf(err)
	defer rows.Close()

	//row := db.QueryRow(queryStr, paramId)
	for rows.Next() {
		var content_id, content_created_by, content_content_type_id int
		var content_path, content_name string
		var content_parent_id, content_template_id sql.NullInt64
		var content_created_date *time.Time
		var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
		var content_url sql.NullString
		var content_domains coreglobals.StringSlice

		var template_id, template_created_by int
		var template_path, template_name, template_alias string
		var template_parent_id sql.NullInt64
		var template_is_partial bool
		var template_parent_templates []byte

		rows.Scan(
			&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&content_url,
			&content_domains,
			&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
			&template_created_by, &template_is_partial, &template_parent_templates)

		/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
		//corehelpers.PanicIf(err)
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No content with that url.")
		case err != nil:
			log.Fatal(err)
		default:
			fmt.Printf("content domains is %v\n", content_domains)
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var contentTemplateId int
		if content_template_id.Valid {
			contentTemplateId = int(content_template_id.Int64)
		}

		var content_url_str string
		if content_url.Valid {
			content_url_str = content_url.String
		}

		var tpid int
		if template_parent_id.Valid {
			tpid = int(template_parent_id.Int64)
		}

		var parent_templates_final []*coremodulesettingsmodels.Template
		var meta map[string]interface{}

		json.Unmarshal(template_parent_templates, &parent_templates_final)
		json.Unmarshal(content_meta, &meta)

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

		content := &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
			content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
			content_url_str, content_domains, nil, nil, &template, nil}
		// content := &Content{content_id, content_node_id, content_content_type_node_id, meta, contentNode, coremodulesettingsmodels.ContentType{}, &template, nil, content_url_str, content_domains,nil}
		contentSlice = append(contentSlice, content)

	}
	return
}

func (c *Content) HTML(str string) (html template.HTML) {
	html = template.HTML(fmt.Sprint(str))
	return
}

func (c *Content) GetByContentTypeId(contentTypeId int) (contentSlice []*Content) {
	db := coreglobals.Db

	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) > 1
  )oki1
) heh
ON heh.id = content.id
WHERE content.content_type_id = $1;`

	// master template
	//var master_template_name string

	rows, err := db.Query(queryStr, contentTypeId)
	corehelpers.PanicIf(err)
	defer rows.Close()

	//row := db.QueryRow(queryStr, paramId)
	for rows.Next() {
		var content_id, content_created_by, content_content_type_id int
		var content_path, content_name string
		var content_parent_id, content_template_id sql.NullInt64
		var content_created_date *time.Time
		var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
		var content_url sql.NullString
		var content_domains coreglobals.StringSlice

		var template_id, template_created_by int
		var template_path, template_name, template_alias string
		var template_parent_id sql.NullInt64
		var template_is_partial bool
		var template_parent_templates []byte

		rows.Scan(
			&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&content_url,
			&content_domains,
			&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
			&template_created_by, &template_is_partial, &template_parent_templates)

		/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
		//corehelpers.PanicIf(err)
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No content with that url.")
		case err != nil:
			log.Fatal(err)
		default:
			fmt.Printf("content domains is %v\n", content_domains)
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var contentTemplateId int
		if content_template_id.Valid {
			contentTemplateId = int(content_template_id.Int64)
		}

		var content_url_str string
		if content_url.Valid {
			content_url_str = content_url.String
		}

		var tpid int
		if template_parent_id.Valid {
			tpid = int(template_parent_id.Int64)
		}

		var parent_templates_final []*coremodulesettingsmodels.Template
		var meta map[string]interface{}

		json.Unmarshal(template_parent_templates, &parent_templates_final)
		json.Unmarshal(content_meta, &meta)

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

		content := &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
			content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
			content_url_str, content_domains, nil, nil, &template, nil}
		// content := &Content{content_id, content_node_id, content_content_type_node_id, meta, contentNode, coremodulesettingsmodels.ContentType{}, &template, nil, content_url_str, content_domains,nil}
		contentSlice = append(contentSlice, content)

	}
	return
}

func (c *Content) AppendSlice(orig []interface{}, elem interface{}) (slice []interface{}) {
	slice = append(orig, elem)
	return
}

func (c *Content) MkStruct() *Content {
	return &Content{}
}

func (c *Content) MkSlice(args ...interface{}) []interface{} {
	return args
}

// eq reports whether the first argument is equal to
// any of the remaining arguments.
func (c *Content) NotEq(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}
	x := args[0]
	switch x := x.(type) {
	case string, int, int64, byte, float32, float64:
		for _, y := range args[1:] {
			if x == y {
				return false
			}
		}
		return true
	}

	for _, y := range args[1:] {
		if reflect.DeepEqual(x, y) {
			return false
		}
	}
	return true
}

// eq reports whether the first argument is equal to
// any of the remaining arguments.
func (c *Content) Eq(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}
	x := args[0]
	switch x := x.(type) {
	case string, int, int64, byte, float32, float64:
		for _, y := range args[1:] {
			if x == y {
				return true
			}
		}
		return false
	}

	for _, y := range args[1:] {
		if reflect.DeepEqual(x, y) {
			return true
		}
	}
	return false
}

func (c *Content) GetHomeContentItem() (content *Content) {
	db := coreglobals.Db

	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) = 1
  )oki1
) heh
ON heh.id = content.id
--WHERE cn.path= subpath('1.42.46.47',0,nlevel(cn.path));
--WHERE cn.path <@ subltree($1,$2,$3)
  WHERE content.path ~ ltree2text(subltree($1,$2,$3))::lquery`

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
	var content_url sql.NullString
	var content_domains coreglobals.StringSlice

	var template_id, template_created_by int
	var template_path, template_name, template_alias string
	var template_parent_id sql.NullInt64
	var template_is_partial bool
	var template_parent_templates []byte

	err := db.QueryRow(queryStr, c.Path, 0, 1).Scan(
		&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
		&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
		&content_user_permissions, &content_user_group_permissions,
		&content_url,
		&content_domains,
		&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
		&template_created_by, &template_is_partial, &template_parent_templates)

	/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
	//corehelpers.PanicIf(err)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No content with that url.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("content domains is %v\n", content_domains)
	}

	var cpid int
	if content_parent_id.Valid {
		cpid = int(content_parent_id.Int64)
	}

	var contentTemplateId int
	if content_template_id.Valid {
		contentTemplateId = int(content_template_id.Int64)
	}

	var content_url_str string
	if content_url.Valid {
		content_url_str = content_url.String
	}

	var tpid int
	if template_parent_id.Valid {
		tpid = int(template_parent_id.Int64)
	}

	var parent_templates_final []*coremodulesettingsmodels.Template
	var meta map[string]interface{}

	json.Unmarshal(template_parent_templates, &parent_templates_final)
	json.Unmarshal(content_meta, &meta)

	// var public_access *PublicAccess

	// json.Unmarshal(content_public_access, &public_access)

	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	json.Unmarshal(content_public_access_members, &public_access_members)
	json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(content_user_permissions, &user_perm)
	json.Unmarshal(content_user_group_permissions, &user_group_perm)

	template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

	content = &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
		content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
		content_url_str, content_domains, nil, nil, &template, nil}

	return
}

func (c *Content) GetAncestors(offset, length int) (contentSlice []*Content) {
	db := coreglobals.Db

	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) > 1
  )oki1
) heh
ON heh.id = content.id
WHERE content.path <@ subltree($1,$2,$3);`

	// master template
	//var master_template_name string

	rows, err := db.Query(queryStr, c.Path, offset, length)
	corehelpers.PanicIf(err)
	defer rows.Close()

	//row := db.QueryRow(queryStr, paramId)
	for rows.Next() {
		var content_id, content_created_by, content_content_type_id int
		var content_path, content_name string
		var content_parent_id, content_template_id sql.NullInt64
		var content_created_date *time.Time
		var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
		var content_url sql.NullString
		var content_domains coreglobals.StringSlice

		var template_id, template_created_by int
		var template_path, template_name, template_alias string
		var template_parent_id sql.NullInt64
		var template_is_partial bool
		var template_parent_templates []byte

		rows.Scan(
			&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
			&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
			&content_user_permissions, &content_user_group_permissions,
			&content_url,
			&content_domains,
			&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
			&template_created_by, &template_is_partial, &template_parent_templates)

		/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
		//corehelpers.PanicIf(err)
		switch {
		case err == sql.ErrNoRows:
			log.Printf("No content with that url.")
		case err != nil:
			log.Fatal(err)
		default:
			fmt.Printf("content domains is %v\n", content_domains)
		}

		var cpid int
		if content_parent_id.Valid {
			cpid = int(content_parent_id.Int64)
		}

		var contentTemplateId int
		if content_template_id.Valid {
			contentTemplateId = int(content_template_id.Int64)
		}

		var content_url_str string
		if content_url.Valid {
			content_url_str = content_url.String
		}

		var tpid int
		if template_parent_id.Valid {
			tpid = int(template_parent_id.Int64)
		}

		var parent_templates_final []*coremodulesettingsmodels.Template
		var meta map[string]interface{}

		json.Unmarshal(template_parent_templates, &parent_templates_final)
		json.Unmarshal(content_meta, &meta)

		// var public_access *PublicAccess

		// json.Unmarshal(content_public_access, &public_access)

		var public_access_members map[string]interface{}
		var public_access_member_groups map[string]interface{}

		json.Unmarshal(content_public_access_members, &public_access_members)
		json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

		var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
		user_perm = nil
		user_group_perm = nil
		json.Unmarshal(content_user_permissions, &user_perm)
		json.Unmarshal(content_user_group_permissions, &user_group_perm)

		template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

		content := &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
			content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
			content_url_str, content_domains, nil, nil, &template, nil}
		// content := &Content{content_id, content_node_id, content_content_type_node_id, meta, contentNode, coremodulesettingsmodels.ContentType{}, &template, nil, content_url_str, content_domains,nil}
		contentSlice = append(contentSlice, content)

	}
	return
	//   db := coreglobals.Db

	//   queryStr := `SELECT cn.id AS node_id, cn.path AS node_path, cn.created_by AS node_created_by, cn.name AS node_name, cn.node_type AS node_type,
	//   cn.created_date AS node_created_date, cn.parent_id AS content_parent_id,
	//   content.id AS content_id, content.node_id AS content_node_id, content.content_type_node_id AS content_content_type_node_id, content.meta AS content_meta,
	//   okidoki.content_url as content_url,
	//   tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
	//   tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
	//   tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates,
	//   heh.domains
	// FROM content
	// JOIN node AS cn
	// ON content.node_id = cn.id
	// JOIN
	// (
	//   SELECT my_template.*, res1.*
	//   FROM template AS my_template,
	//   LATERAL
	//   (
	//     -- SELECT array_to_json(array_agg(node)) AS parent_template_nodes
	//     SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
	//     FROM template
	//     WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
	//     ORDER BY my_template.path ASC
	//   ) res1
	// ) AS tpl
	// --ON (content.meta->>'template_id')::int = tpl.id
	// ON content.template_id = tpl.id
	// JOIN
	// (
	//   SELECT *
	//   FROM node as mynode,
	//   LATERAL
	//   (
	//     SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
	//     FROM node
	// --    JOIN "domain"
	// --    ON "domain".node_id = node.id
	//     WHERE path @> mynode.path AND nlevel(path)>2
	//   ) ok
	// )okidoki
	// ON okidoki.id = cn.id
	// -- JOIN domain
	// -- ON ltree2text(subpath(cn.path,1,1)) = domain.node_id::text
	// JOIN
	// (
	//   SELECT mynode.*, oki1.*
	//   FROM node as mynode,
	//   LATERAL
	//   (
	//     SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
	//     FROM content, jsonb_array_elements_text(meta->'domains') elem
	//     WHERE ltree2text(subpath(mynode.path,1,1)) = content.node_id::text and nlevel(mynode.path) > 1
	//     -- SELECT array_agg(name)domains
	// --    FROM domain
	// --    WHERE ltree2text(subpath(mynode.path,1,1)) = domain.node_id::text and nlevel(mynode.path) > 1
	//   )oki1
	// ) heh
	// ON heh.id = cn.id
	// --WHERE cn.path= subpath('1.42.46.47',0,nlevel(cn.path));
	// WHERE cn.path <@ subltree($1,$2,$3)`

	//   rows, err := db.Query(queryStr, c.Node.Path, offset, length)
	//   corehelpers.PanicIf(err)
	//   defer rows.Close()

	//   //row := db.QueryRow(queryStr, paramId)
	//   for rows.Next(){

	//   // node
	//   var node_id, node_created_by, node_type int
	//   var node_path, node_name string
	//   var node_created_date time.Time
	//   var content_parent_id, content_template_id sql.NullString

	//   // content
	//   var content_id, content_node_id, content_content_type_node_id int
	//   var content_meta []byte

	//   // template node
	//   var template_id, template_created_by int
	//   var template_path, template_name, template_alias string
	//   var template_parent_id sql.NullInt64
	//   var template_is_partial bool
	//   var template_parent_templates []byte

	//   //
	//   var content_domains coreglobals.StringSlice
	//   var content_url sql.NullString

	//   // master template
	//   //var master_template_name string
	//     rows.Scan(
	//         &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date, &content_parent_id,
	//         &content_id, &content_node_id, &content_content_type_node_id, &content_meta, &content_url,
	//         &template_id,&template_path, &template_parent_id, &template_name, &template_alias, &template_created_by, &template_is_partial, &template_parent_templates,
	//         &content_domains)

	//     /* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
	//     //corehelpers.PanicIf(err)

	//     var content_url_str string
	//     if content_url.Valid {
	//       // use s.String
	//       content_url_str = content_url.String
	//     } else {
	//        // NULL value
	//     }

	//     var content_parent_node_id int
	//     if content_parent_id.Valid {
	//       // use s.String
	//       id, _ := strconv.Atoi(content_parent_id.String)
	//       content_parent_node_id = id
	//     } else {
	//        // NULL value
	//     }

	//     var tpid int
	//     if template_parent_id.Valid {
	//       tpid = int(template_parent_id.Int64)
	//     } else {
	//        // NULL value
	//     }

	//     var parent_templates_final []*coremodulesettingsmodels.Template
	//     var meta map[string]interface{}

	//     json.Unmarshal(template_parent_templates, &parent_templates_final)
	//     json.Unmarshal(content_meta, &meta)
	//     //json.Unmarshal(partial_template_nodes, &partial_template_nodes_slice)
	//     //corehelpers.PanicIf(myerr)

	//     //fmt.Println("TEST::: BEGIN ::: ")
	//     //fmt.Println(string(partial_template_nodes))
	//     //fmt.Println("THIS IS::: WEIRD!!!! ::: ")
	//     //fmt.Println(partial_template_nodes_slice)
	//     //fmt.Println("TEST::: END :::")

	//     contentNode := Node{node_id, node_path, node_created_by, node_name, node_type, &node_created_date, content_parent_node_id, nil, nil, false, "", nil, nil, ""}
	//     //templateNode := Node{template_node_id," ",0, template_node_name,0,&time.Time{}, 0, parent_template_nodes_final, nil, false, "", nil, nil, ""}
	//     //template := coremodulesettingsmodels.Template{template_id, template_node_id, template_alias, parent_template_node_id, "", nil, nil, nil, template_is_partial, &templateNode}
	//     template := coremodulesettingsmodels.Template{template_id,template_path, tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}
	//     //templateNode := Node{template_node_id," ",0, template_node_name,0,time.Time{},parent_template_nodes_final, nil, false}
	//     //template := &coremodulesettingsmodels.Template{}
	//     content := &Content{content_id, content_node_id, content_content_type_node_id, meta, contentNode, coremodulesettingsmodels.ContentType{}, &template, nil, content_url_str, content_domains,nil}
	//     contentSlice = append(contentSlice, content)

	//   }
	//   return
}

func (c *Content) GetProperty(name string, fromLvl, toLvl int) (value string) {
	// var m2 map[string]string
	// m2[name] = name
	db := coreglobals.Db

	var propertyValue string
	queryStr := `SELECT content.meta->>$1 as propertyValue
FROM content
WHERE path @> subpath($2,$3,$4)
--WHERE node.path @> subpath($2,$3,nlevel(node.path)-$4)`

	err := db.QueryRow(queryStr, name, c.Path, fromLvl, toLvl).Scan(
		&propertyValue)

	/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
	//corehelpers.PanicIf(err)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No property with that name.")
	case err != nil:
		log.Fatal(err)
	default:
		value = propertyValue
		// fmt.Printf("property name is: %v\n", propertyValue)
	}
	return
}

// Can be useful when we want to generate breadcrumbs and menus
//
func (c *Content) GetProperty2(name, ltreeQuery string) (properties coreglobals.StringSlice) {
	db := coreglobals.Db

	//   queryStr := `SELECT json_agg(props.propertyValue) AS properties
	// FROM
	// (
	//   SELECT content.meta->$1 as propertyValue
	//   FROM node
	//   JOIN content
	//   ON content.node_id = node.id
	//   WHERE node.path ~ $2 -- eg. '1.9.*'
	// ) props`

	//   err := db.QueryRow(queryStr, name, ltreeQuery).Scan(
	//         &properties)

	//   switch {
	//     case err == sql.ErrNoRows:
	//             log.Printf("No property with that name.")
	//     case err != nil:
	//             log.Fatal(err)
	//     default:
	//             fmt.Printf("properties are: %v\n", properties)
	//     }
	//   return

	queryStr := `
  SELECT content.meta->>$1 as propertyValue
  FROM content
  WHERE path ~ $2 -- eg. '1.9.*'`

	rows, err := db.Query(queryStr, name, ltreeQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var propertyValue string
		if err := rows.Scan(&propertyValue); err != nil {
			log.Fatal(err)
		}
		properties = append(properties, propertyValue)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func (c *Content) TemplateFunctionTest(param1 string) template.HTML {
	str := fmt.Sprintf("This is a function inside the content, accessible from the template.<br>The content id is: %d<br>and the parameter value we passed is: %s<br>It gives us a convenient way to fetch additional content, such as the information of the home node - site title, site description social options etc.", c.Id, param1)
	return template.HTML(str)
}

func (c *Content) Post() {
	var meta interface{} = nil
	var publicAccessMembers interface{} = nil
	var publicAccessMemberGroups interface{} = nil
	var userPermissions interface{} = nil
	var userGroupPermissions interface{} = nil

	if c.Meta != nil {
		j, _ := json.Marshal(c.Meta)
		meta = j
	}

	if c.PublicAccessMembers != nil {
		j, _ := json.Marshal(c.PublicAccessMembers)
		publicAccessMembers = j
	}

	if c.PublicAccessMemberGroups != nil {
		j, _ := json.Marshal(c.PublicAccessMemberGroups)
		publicAccessMemberGroups = j
	}

	if c.PublicAccessMemberGroups != nil {
		j, _ := json.Marshal(c.UserPermissions)
		userPermissions = j
	}

	if c.UserPermissions != nil {
		j, _ := json.Marshal(c.UserGroupPermissions)
		userGroupPermissions = j
	}

	// http://godoc.org/github.com/lib/pq
	// pq does not support the LastInsertId() method of the Result type in database/sql.
	// To return the identifier of an INSERT (or UPDATE or DELETE),
	// use the Postgres RETURNING clause with a standard Query or QueryRow call:

	db := coreglobals.Db

	var parentContent Content

	if c.ParentId != nil && *c.ParentId != 0{
		// Channel c, is for getting the parent template
		// We need to append the id of the newly created template to the path of the parent id to create the new path
		c1 := make(chan Content)

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			c1 <- GetContentById(*c.ParentId)
		}()

		go func() {
			for i := range c1 {
				fmt.Println(i)
				parentContent = i
			}
		}()

		wg.Wait()
	}

	if *c.ParentId == 0 {
		c.ParentId = nil
	}

	// This channel and WaitGroup is just to make sure the insert query is completed before we continue
	c2 := make(chan int)
	var id int64

	var wg1 sync.WaitGroup

	wg1.Add(1)

	go func() {
		defer wg1.Done()
		sqlStr := `INSERT INTO content ( 
			parent_id, name, created_by, content_type_id, template_id,  
			meta, public_access_members, public_access_member_groups, 
			user_permissions, user_group_permissions) 
			VALUES (
				$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
				) RETURNING id`
		err1 := db.QueryRow(sqlStr, c.ParentId, c.Name, c.CreatedBy, c.ContentTypeId, c.TemplateId,
			meta, publicAccessMembers, publicAccessMemberGroups,
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

	// fmt.Println(parentTemplate.Path + "." + strconv.FormatInt(id, 10))

	sqlStr := `UPDATE content 
    SET path=$1 
    WHERE id=$2`

	path := strconv.FormatInt(id, 10)
	if c.ParentId != nil {
		path = parentContent.Path + "." + strconv.FormatInt(id, 10)
	}

	_, err6 := db.Exec(sqlStr, path, id)
	corehelpers.PanicIf(err6)

	log.Println("content created successfully")

}

func (c *Content) Put() {

	db := coreglobals.Db

	//meta, _ := json.Marshal(c.Meta)

	var meta interface{} = nil
	var publicAccessMembers interface{} = nil
	var publicAccessMemberGroups interface{} = nil
	var userPermissions interface{} = nil
	var userGroupPermissions interface{} = nil

	if c.Meta != nil {
		j, _ := json.Marshal(c.Meta)
		meta = j
	}

	if c.PublicAccessMembers != nil {
		j, _ := json.Marshal(c.PublicAccessMembers)
		publicAccessMembers = j
	}

	if c.PublicAccessMemberGroups != nil {
		j, _ := json.Marshal(c.PublicAccessMemberGroups)
		publicAccessMemberGroups = j
	}

	if c.UserPermissions != nil {
		j, _ := json.Marshal(c.UserPermissions)
		userPermissions = j
	}

	if c.UserGroupPermissions != nil {
		j, _ := json.Marshal(c.UserGroupPermissions)
		userGroupPermissions = j
	}

	// publicAccessMemberGroups, _ := json.Marshal(c.PublicAccessMemberGroups)
	// userPermissions, _ := json.Marshal(c.UserPermissions)
	// userGroupPermissions, _ := json.Marshal(c.UserGroupPermissions)

	var parentContent Content

	if c.ParentId != nil && *c.ParentId != 0 {
		c1 := make(chan Content)

		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			c1 <- GetContentById(*c.ParentId)
		}()

		go func() {
			for i := range c1 {
				fmt.Println(i)
				parentContent = i
			}
		}()

		wg.Wait()
	}

	if *c.ParentId == 0 {
		c.ParentId = nil
	}

	sqlStr := `UPDATE content 
	SET path=$1, parent_id=$2, name=$3, created_by=$4, content_type_id=$5, template_id=$6, 
	meta=$7, public_access_members=$8, public_access_member_groups=$9, 
	user_permissions=$10, user_group_permissions=$11 
 	WHERE id=$12;`

	path := strconv.Itoa(c.Id)
	if c.ParentId != nil {
		path = parentContent.Path + "." + strconv.Itoa(c.Id)
	}

	_, err := db.Exec(sqlStr, path, c.ParentId, c.Name, c.CreatedBy, c.ContentTypeId, c.TemplateId,
		meta, publicAccessMembers, publicAccessMemberGroups, userPermissions, userGroupPermissions, c.Id)

	corehelpers.PanicIf(err)

	log.Println("content updated successfully")
}

func DeleteContent(id int) {
	db := coreglobals.Db

	sqlStr := `DELETE FROM content 
    WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	corehelpers.PanicIf(err)

	log.Printf("content with id %d was successfully deleted", id)
}

// func (t *Content) Post(){

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

//   _, err = tx.Exec("INSERT INTO content (node_id, content_type_node_id, meta) VALUES ($1, $2, $3)", node_id, t.ContentTypeId, meta)
//   corehelpers.PanicIf(err)
//   //defer r2.Close()

//   if(t.Node.NodeType == 2){
//     var fi FileInfo
//     var fin FileNode
//     if(t.ContentTypeId == 40){
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

//   // _, err = tx.Exec(`UPDATE content
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

// func (c *Content) Update(){

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

//   _, err = tx.Exec(`UPDATE content
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

//     // if content is of media type: folder
//     if(c.ContentTypeId == 16){

//       // check if node has children (SQL SELECT QUERY USING LTREE PATH)
//       rows, err101 := tx.Query(`SELECT content.node_id as node_id, meta as content_meta
//         FROM content
//         JOIN node
//         ON node.id = content.node_id
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
//         var content_meta_byte_arr []byte

//         if err := rows.Scan(&node_id, &content_meta_byte_arr); err != nil {
//           log.Fatal(err)
//         }

//         var content_meta map[string]interface{}
//         json.Unmarshal([]byte(string(content_meta_byte_arr)), &content_meta)

//         var path string = content_meta["path"].(string)
//         var newPath string = strings.Replace(path, folderNode.OldPath, folderNode.FullPath, -1)
//         // update node's content.meta.url part where substing equals oldurl - with c.Meta.url
//         fmt.Println("TEST ::: TEST ::: ERR3")

//         res = append(res,Lol{node_id, newPath})
//         // _, err102 := tx.Exec(`UPDATE content
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
//         _, err102 := tx.Exec(`UPDATE content
//           SET meta = json_object_update_key(meta::json, 'path', $1::text)::jsonb
//           WHERE node_id=$2`, string(res[i].NewPath), res[i].Id)
//         corehelpers.PanicIf(err102)
//       }

//     }
//   }

//   tx.Commit()
// }

func GetBackendContentById(id int) (content Content) {
	db := coreglobals.Db
	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
  modified_content_type.id AS ct_id, modified_content_type.path AS ct_path, modified_content_type.parent_id AS ct_parent_id, modified_content_type.name as ct_name, modified_content_type.alias AS ct_alias,
  modified_content_type.created_by as ct_created_by, modified_content_type.description AS ct_description, modified_content_type.icon AS ct_icon, modified_content_type.thumbnail AS ct_thumbnail, 
  modified_content_type.meta::json AS ct_meta, modified_content_type.ct_tabs AS ct_tabs, modified_content_type.parent_content_types AS ct_parent_content_types, 
  modified_content_type.composite_content_types AS ct_composite_content_types,
  modified_content_type.template_id as ct_template_id, modified_content_type.allowed_template_ids as ct_allowed_template_ids 
FROM content
JOIN
LATERAL
(
  SELECT ct.*,pct.*,cct.*,ct_tabs_with_dt.*
  FROM content_type AS ct,
  -- Parent content types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS parent_content_types
    FROM 
    (
      SELECT c.id, c.path, c.parent_id, c.name, c.alias, c.created_by, c.description, c.icon, c.thumbnail, c.meta, gf.* AS tabs
      FROM content_type AS c,
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
                    FROM content_type AS ct
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
    -- Composite content types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS composite_content_types
    FROM 
    (
      SELECT c.id, c.path, c.parent_id, c.name, c.alias, c.created_by, c.description, c.icon, c.thumbnail, c.meta, gf.* AS tabs
      FROM content_type AS c,
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
                    FROM content_type AS ct
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
      WHERE id = ANY(ct.composite_content_type_ids)
    )res1
  ) cct,
  -- Tabs
  LATERAL 
  (
    SELECT res2.tabs AS ct_tabs
    FROM 
    (
      SELECT c.id AS cid, gf.* AS tabs
      FROM content_type AS c,
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
                    FROM content_type AS ct
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
) modified_content_type
ON modified_content_type.id = content.content_type_id
WHERE content.id=$1`
	// queryStr :=
	// `SELECT my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date, my_node.parent_id as content_parent_id,
	//   content.id as content_id, content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
	//   res.id as ct_id, res.node_id as ct_node_id, res.parent_content_type_node_id as ct_parent_content_type_node_id, res.alias as ct_alias,
	//   res.description as ct_description, res.icon as ct_icon, res.thumbnail as ct_thumbnail, res.meta::json as ct_meta, res.ct_tabs as ct_tabs, res.parent_content_types as ct_parent_content_types
	//   FROM content
	//   JOIN node as my_node
	//   ON my_node.id = content.node_id
	//   JOIN
	//   LATERAL
	//   (
	//     SELECT my_content_type.*,ffgd.*,gf2.*
	//     FROM content_type as my_content_type, node as my_content_type_node,
	//     LATERAL
	//     (
	//         SELECT array_to_json(array_agg(okidoki)) as parent_content_types
	//         FROM (
	//           SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.thumbnail, c.parent_content_type_node_id, c.meta, gf.* as tabs
	//           FROM content_type as c, node,
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
	//                   FROM content_type as ct
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
	//           where path @> subpath(my_content_type_node.path,0,nlevel(my_content_type_node.path)-1) and c.node_id = node.id
	//         )okidoki
	//     ) ffgd,
	//     --
	//     LATERAL
	//     (
	//         SELECT okidoki.tabs as ct_tabs
	//         FROM (
	//           SELECT c.id as cid, gf.* as tabs
	//           FROM content_type as c, node,
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
	//           FROM content_type as ct
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
	//         WHERE c.id = my_content_type.id
	//         )okidoki
	//         limit 1
	//     ) gf2
	//     --
	//     WHERE my_content_type_node.id = my_content_type.node_id
	//   ) res
	//   ON res.node_id = content.content_type_node_id
	//   WHERE my_node.id=$1`

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte

	var ct_id, ct_created_by int
	var ct_parent_id, ct_template_id sql.NullInt64

	var ct_path, ct_name, ct_alias, ct_description, ct_icon, ct_thumbnail string
	var ct_tabs, ct_meta []byte
	var ct_parent_content_types, ct_composite_content_types []byte
	var ct_allowed_template_ids coreglobals.IntSlice

	row := db.QueryRow(queryStr, id)

	err := row.Scan(
		&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
		&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
		&content_user_permissions, &content_user_group_permissions,
		&ct_id, &ct_path, &ct_parent_id, &ct_name, &ct_alias, &ct_created_by,
		&ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_parent_content_types, &ct_composite_content_types,
		&ct_template_id, &ct_allowed_template_ids)

	corehelpers.PanicIf(err)

	var ctpid int
	if ct_parent_id.Valid {
		// use s.String
		ctpid = int(ct_parent_id.Int64)
	} else {
		// NULL value
	}

	var cpid int
	if content_parent_id.Valid {
		cpid = int(content_parent_id.Int64)
	}

	var contentTemplateId int
	if content_template_id.Valid {
		contentTemplateId = int(content_template_id.Int64)
	}

	var template_id int
	if ct_template_id.Valid {
		template_id = int(ct_template_id.Int64)
	}

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(content_user_permissions, &user_perm)
	json.Unmarshal(content_user_group_permissions, &user_group_perm)

	var parent_content_types, composite_content_types []coremodulesettingsmodels.ContentType
	var tabs []coremodulesettingsmodels.Tab
	var ct_metaMap map[string]interface{}
	var content_metaMap map[string]interface{}

	// var public_access *PublicAccess

	// json.Unmarshal(content_public_access, &public_access)

	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	json.Unmarshal(content_public_access_members, &public_access_members)
	json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

	json.Unmarshal(ct_parent_content_types, &parent_content_types)
	json.Unmarshal(ct_composite_content_types, &composite_content_types)
	json.Unmarshal(ct_tabs, &tabs)
	json.Unmarshal(ct_meta, &ct_metaMap)
	json.Unmarshal(content_meta, &content_metaMap)

	content_type := coremodulesettingsmodels.ContentType{ct_id, ct_path, &ctpid, ct_name, ct_alias, ct_created_by, &time.Time{}, ct_description, ct_icon, ct_thumbnail, ct_metaMap, tabs, parent_content_types, nil, false, false, false, nil, nil, composite_content_types, template_id, ct_allowed_template_ids}

	content = Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
		content_content_type_id, &contentTemplateId, content_metaMap, public_access_members, public_access_member_groups, user_perm, user_group_perm, "", nil, nil, nil, nil, &content_type}

	return
}

func GetFrontendContentById(paramId int) (content *Content) {
	db := coreglobals.Db
	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) > 1
  )oki1
) heh
ON heh.id = content.id 
WHERE content.id = $1;`

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
	var content_url sql.NullString
	var content_domains coreglobals.StringSlice

	var template_id, template_created_by int
	var template_path, template_name, template_alias string
	var template_parent_id sql.NullInt64
	var template_is_partial bool
	var template_parent_templates []byte

	err := db.QueryRow(queryStr, paramId).Scan(
		&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
		&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
		&content_user_permissions, &content_user_group_permissions,
		&content_url,
		&content_domains,
		&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
		&template_created_by, &template_is_partial, &template_parent_templates)

	/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
	//corehelpers.PanicIf(err)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No content with that url.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("content domains is %v\n", content_domains)
	}

	var cpid int
	if content_parent_id.Valid {
		cpid = int(content_parent_id.Int64)
	}

	var contentTemplateId int
	if content_template_id.Valid {
		contentTemplateId = int(content_template_id.Int64)
	}

	var content_url_str string
	if content_url.Valid {
		content_url_str = content_url.String
	}

	var tpid int
	if template_parent_id.Valid {
		tpid = int(template_parent_id.Int64)
	}

	var parent_templates_final []*coremodulesettingsmodels.Template
	var meta map[string]interface{}

	json.Unmarshal(template_parent_templates, &parent_templates_final)
	json.Unmarshal(content_meta, &meta)

	// var public_access *PublicAccess

	// json.Unmarshal(content_public_access, &public_access)

	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	json.Unmarshal(content_public_access_members, &public_access_members)
	json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(content_user_permissions, &user_perm)
	json.Unmarshal(content_user_group_permissions, &user_group_perm)

	template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

	content = &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
		content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
		content_url_str, content_domains, nil, nil, &template, nil}
	return
}

func GetFrontendContentByUrl(name, url string) (content *Content) {
	db := coreglobals.Db

	//   queryStr := `SELECT cn.id AS node_id, cn.path AS node_path, cn.created_by AS node_created_by, cn.name AS node_name, cn.node_type AS node_type,
	//   cn.created_date AS node_created_date, cn.parent_id AS content_parent_id,
	//   content.id AS content_id, content.node_id AS content_node_id, content.content_type_node_id AS content_content_type_node_id, content.meta AS content_meta, okidoki.content_url as content_url, content.public_access as content_public_access,
	//   tpl.parent_template_node_id AS parent_template_node_id, tpl.alias AS template_alias, tpl.partial_template_nodes,
	//   tn.id AS template_node_id, tn.parent_template_nodes AS parent_template_nodes, tn.name AS template_node_name,
	//   heh.domains
	// FROM content
	// JOIN node AS cn
	// ON content.node_id = cn.id
	// JOIN
	// (
	//   SELECT my_node.*, res1.*
	//   FROM node AS my_node,
	//   LATERAL
	//   (
	//     -- SELECT array_to_json(array_agg(node)) AS parent_template_nodes
	//     SELECT json_agg((SELECT x FROM (SELECT node.id, node.path, node.name, node.node_type, node.created_by, node.parent_id) x)) AS parent_template_nodes
	//     FROM node
	//     WHERE path @> subpath(my_node.path,0,nlevel(my_node.path)-1) AND node_type=3
	//     ORDER BY my_node.path ASC
	//   ) res1
	//   WHERE my_node.node_type = 3
	// ) AS tn
	// ON (content.meta->>'template_node_id')::int = tn.id
	// JOIN
	// (
	//   SELECT template.*, res2.*
	//   FROM template,
	//   LATERAL
	//   (
	//     SELECT json_agg((SELECT x FROM (SELECT node.id, node.path, node.name, node.node_type, node.created_by, node.parent_id) x)) AS partial_template_nodes
	//     FROM node
	//     WHERE node.id = ANY(template.partial_template_node_ids)
	//     --WHERE node.id IN (SELECT unnest(template.partial_template_node_ids))
	//     ORDER BY template.node_id ASC
	//   ) res2
	// ) AS tpl
	// ON tpl.node_id = tn.id
	// JOIN
	// (
	//   SELECT *
	//   FROM node as mynode,
	//   LATERAL
	//   (
	//     SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
	//     FROM node
	// --    JOIN "domain"
	// --    ON "domain".node_id = node.id
	//     WHERE path @> mynode.path AND nlevel(path)>2
	//   ) ok
	// )okidoki
	// ON okidoki.id = cn.id
	// -- JOIN domain
	// -- ON ltree2text(subpath(cn.path,1,1)) = domain.node_id::text
	// JOIN
	// (
	//   SELECT mynode.*, oki1.*
	//   FROM node as mynode,
	//   LATERAL
	//   (
	//     SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
	//     FROM content, jsonb_array_elements_text(meta->'domains') elem
	//     WHERE ltree2text(subpath(mynode.path,1,1)) = content.node_id::text and nlevel(mynode.path) > 1
	//     -- SELECT array_agg(name)domains
	// --    FROM domain
	// --    WHERE ltree2text(subpath(mynode.path,1,1)) = domain.node_id::text and nlevel(mynode.path) > 1
	//   )oki1
	// ) heh
	// ON heh.id = cn.id
	// WHERE lower(cn.name) = $1;`
	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
-- ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
JOIN 
(
  SELECT * 
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
    FROM content
--    JOIN "domain"
--    ON "domain".node_id = node.id
    WHERE path @> mycontent.path AND nlevel(path)>1
  ) ok
)okidoki
ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) > 1
  )oki1
) heh
ON heh.id = content.id 
WHERE lower(content.name) = $1;`

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
	var content_url sql.NullString
	var content_domains coreglobals.StringSlice

	var template_id, template_created_by int
	var template_path, template_name, template_alias string
	var template_parent_id sql.NullInt64
	var template_is_partial bool
	var template_parent_templates []byte

	err := db.QueryRow(queryStr, name).Scan(
		&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
		&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
		&content_user_permissions, &content_user_group_permissions,
		&content_url,
		&content_domains,
		&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
		&template_created_by, &template_is_partial, &template_parent_templates)

	/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
	//corehelpers.PanicIf(err)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No content with that url.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("content domains is %v\n", content_domains)
	}

	var cpid int
	if content_parent_id.Valid {
		cpid = int(content_parent_id.Int64)
	}

	var contentTemplateId int
	if content_template_id.Valid {
		contentTemplateId = int(content_template_id.Int64)
	}

	var content_url_str string
	if content_url.Valid {
		content_url_str = content_url.String
	}

	var tpid int
	if template_parent_id.Valid {
		tpid = int(template_parent_id.Int64)
	}

	var parent_templates_final []*coremodulesettingsmodels.Template
	var meta map[string]interface{}

	json.Unmarshal(template_parent_templates, &parent_templates_final)
	json.Unmarshal(content_meta, &meta)

	// var public_access *PublicAccess

	// json.Unmarshal(content_public_access, &public_access)

	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	json.Unmarshal(content_public_access_members, &public_access_members)
	json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)

	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(content_user_permissions, &user_perm)
	json.Unmarshal(content_user_group_permissions, &user_group_perm)

	template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

	content = &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
		content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm,
		content_url_str, content_domains, nil, nil, &template, nil}

	for i := 0; i < len(content.Domains); i++ {
		//fmt.Println("lol: " + content.Domains[i])
		fullUrl := content.Domains[i] + "/" + content.Url
		fmt.Println("Fullurl: " + fullUrl)
		if url == fullUrl {
			return
		}
	}
	// for _, value := range content.Domains{
	//   fmt.Println("lol: " + value)
	//     fullUrl := value + "/" + content.Url
	//     fmt.Println("Fullurl: " + fullUrl)
	//     if(url == fullUrl){
	//       return
	//     }
	// }
	fmt.Println("YOU SHOULDN't SEE THIS IF THE URL IS RIGHT1")
	fmt.Println("url: " + url)

	return nil

}

func GetFrontendContentByDomain(domain string) (content *Content) {
	db := coreglobals.Db

	queryStr := `SELECT content.id AS content_id, content.path AS content_path, content.parent_id AS content_parent_id,
content.name AS content_name,  content.created_by AS content_created_by, 
content.created_date AS content_created_date, content.content_type_id AS content_content_type_id, content.template_id AS content_template_id,
content.meta AS content_meta, content.public_access_members AS content_public_access_members, content.public_Access_member_groups AS content_public_access_member_groups,
content.user_permissions AS content_user_permissions, content.user_group_permissions AS content_user_group_permissions,
--okidoki.content_url as content_url,
heh.domains AS content_domains,
tpl.id AS template_id, tpl.path AS template_path, tpl.parent_id AS template_parent_id,
tpl.name AS template_name, tpl.alias AS template_alias, tpl.created_by AS template_created_by,
tpl.is_partial AS template_is_partial, tpl.parent_templates as template_parent_templates
FROM content
JOIN
(
  SELECT my_template.*, res1.*
  FROM template AS my_template,
  LATERAL 
  (
    SELECT json_agg((SELECT x FROM (SELECT template.id, template.path, template.parent_id, template.name, template.alias, template.created_by, template.is_partial) x)) AS parent_templates
    FROM template
    WHERE path @> subpath(my_template.path,0,nlevel(my_template.path)-1)
    ORDER BY my_template.path ASC
  ) res1
) AS tpl
--ON (content.meta->>'template_id')::int = tpl.id
ON content.template_id = tpl.id 
-- JOIN 
-- (
--   SELECT * 
--   FROM content as mycontent,
--   LATERAL
--   (
--     SELECT string_agg(replace(lower(name), ' ', '-'), '/' ORDER BY path)content_url
--     FROM content
-- --    JOIN "domain"
-- --    ON "domain".node_id = node.id
--     WHERE path @> mycontent.path AND nlevel(path)>1
--   ) ok
-- )okidoki
-- ON okidoki.id = content.id
JOIN
(
  SELECT mycontent.*, oki1.*
  FROM content as mycontent,
  LATERAL
  (
    SELECT string_to_array(string_agg(elem,', '),', ')::varchar[] as domains
    FROM content, jsonb_array_elements_text(meta->'domains') elem
    WHERE ltree2text(subpath(mycontent.path,0,1)) = content.id::text and nlevel(mycontent.path) = 1
  )oki1
) heh
ON heh.id = content.id 
WHERE $1 = ANY(heh.domains) and nlevel(content.path) = 1;`

	var content_id, content_created_by, content_content_type_id int
	var content_path, content_name string
	var content_parent_id, content_template_id sql.NullInt64
	var content_created_date *time.Time
	var content_meta, content_public_access_members, content_public_access_member_groups, content_user_permissions, content_user_group_permissions []byte
	//var content_url sql.NullString
	var content_domains coreglobals.StringSlice

	var template_id, template_created_by int
	var template_path, template_name, template_alias string
	var template_parent_id sql.NullInt64
	var template_is_partial bool
	var template_parent_templates []byte

	err := db.QueryRow(queryStr, domain).Scan(
		&content_id, &content_path, &content_parent_id, &content_name, &content_created_by,
		&content_created_date, &content_content_type_id, &content_template_id, &content_meta, &content_public_access_members, &content_public_access_member_groups,
		&content_user_permissions, &content_user_group_permissions,
		//&content_url,
		&content_domains,
		&template_id, &template_path, &template_parent_id, &template_name, &template_alias,
		&template_created_by, &template_is_partial, &template_parent_templates)

	/* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
	//corehelpers.PanicIf(err)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No content with that domain.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("content domains is %v\n", content_domains)
	}

	var cpid int
	if content_parent_id.Valid {
		cpid = int(content_parent_id.Int64)
	}

	var contentTemplateId int
	if content_template_id.Valid {
		contentTemplateId = int(content_template_id.Int64)
	}

	// var content_url_str string
	// if content_url.Valid {
	//   content_url_str = content_url.String
	// }

	var tpid int
	if template_parent_id.Valid {
		tpid = int(template_parent_id.Int64)
	}

	var parent_templates_final []*coremodulesettingsmodels.Template
	var meta map[string]interface{}

	json.Unmarshal(template_parent_templates, &parent_templates_final)
	json.Unmarshal(content_meta, &meta)

	// var public_access *PublicAccess

	// json.Unmarshal(content_public_access, &public_access)

	var public_access_members map[string]interface{}
	var public_access_member_groups map[string]interface{}

	json.Unmarshal(content_public_access_members, &public_access_members)
	json.Unmarshal(content_public_access_member_groups, &public_access_member_groups)
	var user_perm, user_group_perm map[string]*PermissionTest // map[string]PermissionsContainer
	user_perm = nil
	user_group_perm = nil
	json.Unmarshal(content_user_permissions, &user_perm)
	json.Unmarshal(content_user_group_permissions, &user_group_perm)

	template := coremodulesettingsmodels.Template{template_id, template_path, &tpid, template_name, template_alias, template_created_by, &time.Time{}, template_is_partial, "", parent_templates_final}

	content = &Content{content_id, content_path, &cpid, content_name, content_created_by, content_created_date,
		content_content_type_id, &contentTemplateId, meta, public_access_members, public_access_member_groups, user_perm, user_group_perm, "", content_domains, nil, nil, &template, nil}

	fmt.Println(content_domains)

	for _, value := range content.Domains {
		// fullUrl := value + "/" + content.Url
		// fmt.Println("Fullurl: " + fullUrl)
		// if(url == fullUrl){
		//   return
		// }
		if value == domain {
			return
		}
	}
	fmt.Println("YOU SHOULDN't SEE THIS IF THE URL IS RIGHT")
	//fmt.Println("url: " + url)

	return nil

}

// type content Content

// func (c *Content) UnmarshalJSON(b []byte) (err error) {
// 	j, i, ni := content{}, 0, 0
// 	var m map[string]interface{}

// 	if err = json.Unmarshal(b, &j); err == nil {
// 		*c = Content(j)
// 		return
// 	}
//   if err = json.Unmarshal(b, &i); err == nil {
//     c.Id = i
//     return
//   }
//   if err = json.Unmarshal(b, &ni); err == nil {
//     c.Id = ni
//     return
//   }
// 	// if err = json.Unmarshal(b, &n); err == nil {
// 	// 	d.Id = n
// 	// 	return
// 	// }
// 	if err = json.Unmarshal(b, &m); err == nil {
// 		c.Meta = m
// 		return
// 	}
// 	return
// }

// func (c *Content) MarshalJSON() ([]byte, error) {
//     if c.Id != 0 && c.Id != 0 {
//         return json.Marshal(map[string]interface{}{
//             "id": c.Id,
//             //"node_id": d.Id,
//             "node_id": c.Id,
//             "content_type_node_id": c.ContentTypeId,
//             "meta": c.Meta,
//         })
//     }
//     if c.Id != 0 {
//         return json.Marshal(c.Id)
//     }
//     if c.Id != 0 {
//         return json.Marshal(c.Id)
//     }
//     if c.ContentTypeId != 0 {
//         return json.Marshal(c.ContentTypeId)
//     }
//     return json.Marshal(nil)
// }
