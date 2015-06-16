package models

import (
	coreglobals "collexy/core/globals"
	// corehelpers "collexy/core/helpers"
	"database/sql"
	"log"
)

type Permission struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetPermissions() (permissions []*Permission) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, name 
        FROM "permission"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		permission := &Permission{id, name}
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetPermissionById(id int) (permission *Permission) {
	db := coreglobals.Db

	var name string

	err := db.QueryRow(`SELECT name from permission WHERE id=$1`, id).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No permission with that name.")
	case err != nil:
		log.Fatal(err)
	default:
		permission = &Permission{id, name}
	}
	return
}

func (d *Permission) Post() {

	db := coreglobals.Db

	sqlStr := `INSERT INTO permission (name) 
	VALUES ($1)`
	_, err1 := db.Exec(sqlStr, d.Name)

	if err1 != nil {
		panic(err1)
	}

	log.Println("permission created successfully")
}

func (d *Permission) Update() {

	db := coreglobals.Db

	sqlStr := `UPDATE data_type 
	SET name=$1
	WHERE id=$2`

	_, err1 := db.Exec(sqlStr, d.Name, d.Id)

	if err1 != nil {
		panic(err1)
	}

	log.Println("permission updated successfully")
}

func DeletePermission(id int) {

	db := coreglobals.Db

	sqlStr := `delete FROM permission 
	WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	if err != nil {
		panic(err)
	}

	log.Printf("permission with id %d was successfully deleted", id)
}
