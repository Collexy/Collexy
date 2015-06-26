package models

import (
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	//"encoding/json"
	//"time"
	//"fmt"
	//"net/http"
	//"html/template"
	"database/sql"
	"log"
	//"strconv"
)

type MimeType struct {
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	MediaTypeId *int   `json:"media_type_id,omitempty"`
}

func GetMimeTypes() (mimeTypes []*MimeType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, name, media_type_id  
        FROM mime_type`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var media_type_id sql.NullInt64

		if err := rows.Scan(&id, &name, &media_type_id); err != nil {
			log.Fatal(err)
		}

		var media_type_id_int int
		var media_type_id_int_pointer *int = nil
		if media_type_id.Valid {
			media_type_id_int = int(media_type_id.Int64)
			media_type_id_int_pointer = &media_type_id_int
		}

		mimeType := &MimeType{id, name, media_type_id_int_pointer}
		mimeTypes = append(mimeTypes, mimeType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetMimeTypeById(id int) (mimeType *MimeType) {
	db := coreglobals.Db

	var name string
	var media_type_id sql.NullInt64

	err := db.QueryRow(`SELECT name, media_type_id
        FROM mime_type WHERE id=$1`, id).Scan(&name, &media_type_id)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No mime type with that ID.")
	case err != nil:
		log.Fatal(err)
	default:

		var media_type_id_int int
		var media_type_id_int_pointer *int = nil
		if media_type_id.Valid {
			media_type_id_int = int(media_type_id.Int64)
			media_type_id_int_pointer = &media_type_id_int
		}

		mimeType = &MimeType{id, name, media_type_id_int_pointer}
	}
	return
}

func (m *MimeType) Update() {

	db := coreglobals.Db

	sqlStr := `UPDATE mime_type 
	SET name=$1, media_type_id=$2  
	WHERE id=$3`

	_, err1 := db.Exec(sqlStr, m.Name, m.MediaTypeId, m.Id)

	corehelpers.PanicIf(err1)

	log.Println("mime type updated successfully")
}

func (m *MimeType) Post() {

	db := coreglobals.Db

	sqlStr := `INSERT INTO mime_type (name, media_type_id) 
	VALUES ($1, $2)`
	_, err1 := db.Exec(sqlStr, m.Name, m.MediaTypeId)

	corehelpers.PanicIf(err1)

	log.Println("mime type created successfully")
}
