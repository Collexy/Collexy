package models

import (
	coreglobals "collexy/core/globals"
	"database/sql"
	"log"
	"time"
)

type MemberGroup struct {
	Id          int       `json:"id"`
	Path        string    `json:"path"`
	ParentId    int       `json:"parent_id,omitempty"`
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	CreatedBy   int       `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
}

func (this *MemberGroup) Post() {
	db := coreglobals.Db

	_, err := db.Exec(`INSERT INTO member_group (path, parent_id, name, alias, created_by) 
        VALUES ($1, $2, $3, $4, $5)`, this.Path, this.ParentId, this.Name, this.Alias, this.CreatedBy)

	if err != nil {
		log.Fatal(err)
	}
}

func (this *MemberGroup) Put() {
	db := coreglobals.Db

	_, err := db.Exec(`UPDATE member_group set name=$1, alias=$2 WHERE id=$3`, this.Name, this.Alias, this.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func GetMemberGroups() (memberGroups []*MemberGroup) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, path, parent_id, name, alias, created_by, created_date 
        FROM member_group`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, created_by int
		var path, name, alias string
		var created_date time.Time
		var parent_id sql.NullInt64

		if err := rows.Scan(&id, &path, &parent_id, &name, &alias, &created_by, &created_date); err != nil {
			log.Fatal(err)
		}

		var pid int

		if parent_id.Valid {
			pid = int(parent_id.Int64)
		}

		memberGroup := &MemberGroup{id, path, pid, name, alias, created_by, created_date}
		memberGroups = append(memberGroups, memberGroup)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMemberGroupById(id int) (memberGroup *MemberGroup) {
	db := coreglobals.Db

	var created_by int
	var path, name, alias string
	var created_date time.Time
	var parent_id sql.NullInt64

	err := db.QueryRow(`SELECT id, path, parent_id, name, alias, created_by, created_date 
        FROM member_group WHERE id=$1`, id).Scan(&id, &path, &parent_id, &name, &alias, &created_by, &created_date)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No member group with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		var pid int

		if parent_id.Valid {
			pid = int(parent_id.Int64)
		}

		memberGroup = &MemberGroup{id, path, pid, name, alias, created_by, created_date}
	}
	return
}
