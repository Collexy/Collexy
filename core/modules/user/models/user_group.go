package models

import (
	coreglobals "collexy/core/globals"
	"database/sql"
	"log"
	// "database/sql/driver"
	// "bytes"
	//"encoding/json"
)

type UserGroup struct {
	Id              int      `json:"id"`
	Name            string   `json:"name,omitempty"`
	PermissionIds   []int    `json:"permission_ids,omitempty"`
	AngularRouteIds []int    `json:"angular_route_ids,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
}

func GetUserGroups(user *User) (userGroups []UserGroup) {
	db := coreglobals.Db

	querystr := `SELECT id, name, permissions FROM user_group`

	rows, err := db.Query(querystr)
	if err != nil {
		log.Println("lol")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var permissions coreglobals.StringSlice

		err := rows.Scan(&id, &name, &permissions)

		switch {
		case err == sql.ErrNoRows:
			log.Printf("No usergroup with that ID.")
		case err != nil:
			log.Fatal(err)
		default:

			// var perm []string

			// err1 := json.Unmarshal(permissions, &perm)
			// if(err1 != nil){
			// 	panic(err1)
			// }

			userGroup := UserGroup{id, name, nil, nil, permissions}
			userGroups = append(userGroups, userGroup)
		}
	}

	return
}

func GetUserGroupById(id int, user *User) (userGroup UserGroup) {
	db := coreglobals.Db

	querystr := `SELECT name, permissions FROM user_group WHERE id=$1`

	row := db.QueryRow(querystr, id)

	var name string
	var permissions coreglobals.StringSlice

	err := row.Scan(&name, &permissions)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No usergroup with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		userGroup = UserGroup{id, name, nil, nil, permissions}

	}

	return
}
