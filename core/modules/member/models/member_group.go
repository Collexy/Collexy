package models

import (
	coreglobals "collexy/core/globals"
	"database/sql"
	"log"
	"time"
)

type MemberGroup struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	CreatedBy   int       `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
}

func (this *MemberGroup) Post() {
	db := coreglobals.Db

	_, err := db.Exec(`INSERT INTO member_group (name, alias, created_by) 
        VALUES ($1, $2, $3)`, this.Name, this.Alias, this.CreatedBy)

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

func DeleteMemberGroup(id int) {

	db := coreglobals.Db

	sqlStr := `delete FROM "member_group" 
	WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	if err != nil {
		panic(err)
	}

	log.Printf("member group with id %d was successfully deleted", id)
}

func GetMemberGroups() (memberGroups []*MemberGroup) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, name, alias, created_by, created_date 
        FROM member_group`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, created_by int
		var name, alias string
		var created_date time.Time

		if err := rows.Scan(&id, &name, &alias, &created_by, &created_date); err != nil {
			log.Fatal(err)
		}

		memberGroup := &MemberGroup{id, name, alias, created_by, created_date}
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
	var name, alias string
	var created_date time.Time

	err := db.QueryRow(`SELECT id, name, alias, created_by, created_date 
        FROM member_group WHERE id=$1`, id).Scan(&id, &name, &alias, &created_by, &created_date)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No member group with that ID.")
	case err != nil:
		log.Fatal(err)
	default:

		memberGroup = &MemberGroup{id, name, alias, created_by, created_date}
	}
	return
}
