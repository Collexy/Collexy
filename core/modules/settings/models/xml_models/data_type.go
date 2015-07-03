package xml_models

import
(
	"encoding/xml"
	coreglobals "collexy/core/globals"
	"log"
)

type DataType struct {
	XMLName xml.Name `xml:"dataType" json:"-"`
	Id int	 `xml"dataTypeId,omitempty`
	Name    string   `xml:"name"`	
}

func GetDataTypes() (dataTypes []DataType) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, name  
        FROM data_type`)
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

		

		dataType := DataType{xml.Name{},id, name}
		dataTypes = append(dataTypes, dataType)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}