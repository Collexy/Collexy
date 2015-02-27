package models

import
(
	coreglobals "collexy/core/globals"
	"database/sql"
	"collexy/helpers"
	"log"
	"encoding/json"
	"collexy/globals"
	// "fmt"
)

type MenuLink struct {
	Id int `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
	ParentId int `json:"parent_id,omitempty"`
	RouteId int `json:"route_id,omitempty"`
	Icon string `json:"icon,omitempty"`
	Atts map[string]interface{} `json:"atts,omitempty"`
	Type int `json:"type"`
	Menu string `json:"menu"`
	Permissions []string `json:"permissions,omitempty"`
	//RoutePath string `json:"route_path,omitempty"`
	Route *Route `json:"route,omitempty"`
}

func GetMenuLinks(menuName string) (menuLinks []MenuLink){
	db := coreglobals.Db

    querystr := `SELECT menu_link.*,
route.path as route_path, route.name as route_name, route.parent_id as route_parent_id, route.url as route_url, route.components as route_components, route.is_abstract as route_is_abstract 
FROM menu_link
JOIN route
ON route.id = menu_link.route_id
WHERE menu_link.menu = $1 
-- WHERE user_group_ids @> '1'
order by menu_link.path,menu_link.id`

    var id, link_type int
    var path, name, menu string
    var atts []byte
    var parent_id, route_id sql.NullInt64
    var icon sql.NullString
    //var permissions globals.StringSlice
    var permissions globals.StringSlice

    //var route_route_id int
    var route_path, route_name string
    var route_parent_id sql.NullInt64
    var route_url sql.NullString
    var route_is_abstract bool
    var route_components []byte

	rows, err := db.Query(querystr, menuName)
	helpers.PanicIf(err)
	defer rows.Close()

	for rows.Next(){
		err:= rows.Scan(
			&id, &path, &name, &parent_id, &route_id, &icon, &atts, &link_type, &menu, &permissions, 
			&route_path, &route_name, &route_parent_id, &route_url, &route_components, &route_is_abstract)

		var parent_id_int int
		if parent_id.Valid {
			// use s.String
			parent_id_int = int(parent_id.Int64)
		} else {
			// NULL value
		}

		var route_id_int int
		if route_id.Valid {
			// use s.String
			route_id_int = int(route_id.Int64)
		} else {
			// NULL value
		}

		var icon_str string
		if icon.Valid {
			// use s.String
			icon_str = icon.String
		} else {
			// NULL value
		}

		var attsMap map[string]interface{}
		json.Unmarshal(atts, &attsMap)
		//helpers.PanicIf(err1)

		// ROUTE

		var route_parent_id_int int
		if route_parent_id.Valid {
			// use s.String
			route_parent_id_int = int(route_parent_id.Int64)
		} else {
			// NULL value
		}

		var route_url_str string
		if route_url.Valid {
			// use s.String
			route_url_str = route_url.String
		} else {
			// NULL value
		}

		var routeComponentsMap []RouteComponent
		json.Unmarshal(route_components, &routeComponentsMap)

		switch {
			case err == sql.ErrNoRows:
				log.Printf("No node with that ID.")
			case err != nil:
				log.Fatal(err)
			default:
				// var permArr []Permission
				// errLol := json.Unmarshal(permissions, &permArr)
				// log.Println(errLol)

				route := Route{route_id_int, route_path, route_name, route_parent_id_int, route_url_str, routeComponentsMap, route_is_abstract}

				menu_link := MenuLink{id, path, name, parent_id_int, route_id_int, icon_str, attsMap, link_type, menu, permissions, &route}
				menuLinks = append(menuLinks,menu_link)

		}
	}
	return
}