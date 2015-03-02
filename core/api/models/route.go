package models

import
(
	coreglobals "collexy/core/globals"
	"database/sql"
	corehelpers "collexy/core/helpers"
	"log"
	"encoding/json"
)

type Route struct {
	Id int `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
	ParentId int `json:"parent_id,omitempty"`
	Url string `json:"url,omitempty"`
	Components []RouteComponent `json:"components,omitempty"`
	IsAbstract bool `json:"is_abstract"`
}

type RouteComponent struct {
	Single string `json:"single"`
}

func GetRoutes() (routes []Route){
	db := coreglobals.Db

    querystr := `SELECT * FROM route`

    var id int
    var path, name string
    var components []byte
    var parent_id sql.NullInt64
    var url sql.NullString
    var is_abstract bool

	rows, err := db.Query(querystr)
	corehelpers.PanicIf(err)
	defer rows.Close()

	for rows.Next(){
		err:= rows.Scan(
			&id, &path, &name, &parent_id, &url, &components, &is_abstract)

		var parent_id_int int
		if parent_id.Valid {
			// use s.String
			parent_id_int = int(parent_id.Int64)
		} else {
			// NULL value
		}

		var url_str string
		if url.Valid {
			// use s.String
			url_str = url.String
		} else {
			// NULL value
		}

		var routeComponents []RouteComponent
		err1 := json.Unmarshal(components, &routeComponents)
		if(err1 != nil){
			routeComponents = nil
		}
		//corehelpers.PanicIf(err1)

		switch {
			case err == sql.ErrNoRows:
				log.Printf("No node with that ID.")
			case err != nil:
				log.Fatal(err)
			default:
				
				
				

				route := Route{id, path, name, parent_id_int, url_str, routeComponents, is_abstract}
				routes = append(routes,route)
		}
	}

	return
}