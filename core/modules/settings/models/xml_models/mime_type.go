package xml_models

import (
	coreglobals "collexy/core/globals"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	"encoding/xml"
	"log"
)

type MimeType struct {
	XMLName     xml.Name `xml:"mimeType"`
	Id          *int
	Name        string `xml:"name"`
	MediaType   string `xml:"mediaType,omitempty"`
	MediaTypeId *int   `xml:"mediaTypeId,omitempty"`
}

func GetMimeTypes() (mimeTypes []MimeType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, name, media_type_id   
        FROM mime_type`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, media_type_id int
		var name string

		if err := rows.Scan(&id, &name, &media_type_id); err != nil {
			log.Fatal(err)
		}

		mimeType := MimeType{xml.Name{}, &id, name, "", &media_type_id}
		mimeTypes = append(mimeTypes, mimeType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func (this *MimeType) Post(mediaTypes []*coremodulesettingsmodels.MediaType) {
	// do DB POST
	for _, mt := range mediaTypes {
		if mt.Alias == this.MediaType {
			this.MediaTypeId = &mt.Id
		}
	}

	db := coreglobals.Db

	sqlStr := `INSERT INTO mime_type (name, media_type_id) 
	VALUES ($1, $2)`

	_, err1 := db.Exec(sqlStr, this.Name, this.MediaTypeId)
	if err1 != nil {
		log.Println(err1)
	}

}
