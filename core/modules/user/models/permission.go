package models

import (
	coreglobals "collexy/core/globals"
	//"database/sql"
	"log"
)

type Permission struct {
	Name string `json:"name"`
}

// func (this *Permission) Post() {
// 	db := coreglobals.Db

// 	_, err := db.Exec(`INSERT INTO member_group (path, parent_id, name, alias, created_by)
//         VALUES ($1, $2, $3, $4, $5)`, this.Path, this.ParentId, this.Name, this.Alias, this.CreatedBy)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (this *Permission) Put() {
// 	db := coreglobals.Db

// 	_, err := db.Exec(`UPDATE member_group set name=$1, alias=$2 WHERE id=$3`, this.Name, this.Alias, this.Id)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GetPermissions() (permissions []*Permission) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT name 
        FROM "permission"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}

		permission := &Permission{name}
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
