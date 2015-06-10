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
	Alias           string   `json:"alias,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
}

func GetUserGroups(user *User) (userGroups []UserGroup) {
	db := coreglobals.Db

	querystr := `SELECT id, name, alias, permissions FROM user_group`

	rows, err := db.Query(querystr)
	if err != nil {
		log.Println("lol")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, alias string
		var permissions coreglobals.StringSlice

		err := rows.Scan(&id, &name, &alias, &permissions)

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

			userGroup := UserGroup{id, name, alias, permissions}
			userGroups = append(userGroups, userGroup)
		}
	}

	return
}

func GetUserGroupById(id int, user *User) (userGroup UserGroup) {
	db := coreglobals.Db

	querystr := `SELECT name, alias, permissions FROM user_group WHERE id=$1`

	row := db.QueryRow(querystr, id)

	var name, alias string
	var permissions coreglobals.StringSlice

	err := row.Scan(&name, &alias, &permissions)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No usergroup with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		userGroup = UserGroup{id, name, alias, permissions}

	}

	return
}

func (u *UserGroup) Post() {

	//meta, err := json.Marshal(d.Meta)
	//corehelpers.PanicIf(err)

	db := coreglobals.Db

	permissions, _ := coreglobals.StringSlice(u.Permissions).Value()

	// sqlStr := `INSERT INTO data_type (name, alias, created_by, html, editor_alias, meta)
	// VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// err1 := db.QueryRow(sqlStr, d.Name, d.Alias, d.CreatedBy, d.Html, d.EditorAlias, meta).Scan(&id)
	sqlStr := `INSERT INTO user_group (name, alias, permissions) 
		VALUES ($1, $2, $3)`
	_, err1 := db.Exec(sqlStr, u.Name, u.Alias, permissions)

	if err1 != nil {
		panic(err1)
	}


	log.Println("user group created successfully")
}

func (u *UserGroup) Put() {

	permissions, _ := coreglobals.StringSlice(u.Permissions).Value()

	db := coreglobals.Db

	sqlStr := `UPDATE user_group 
	SET name=$1, alias=$2, permissions=$3 
		WHERE id=$4`

	_, err1 := db.Exec(sqlStr, u.Name, u.Alias, permissions, u.Id)

	if err1 != nil {
		panic(err1)
	}

	log.Println("user group updated successfully")
}

func DeleteUserGroup(id int) {

	db := coreglobals.Db

	sqlStr := `delete FROM user_group 
	WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	if err != nil {
		panic(err)
	}

	log.Printf("user group with id %d was successfully deleted", id)
}