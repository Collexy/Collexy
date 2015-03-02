package models

import
(
	corehelpers "collexy/core/helpers"
	coreglobals "collexy/core/globals"
	"net/url"
	"database/sql"
	"encoding/json"
)

type AngularRoute struct {
  Id int `json:"id,omitempty"`
  Name string `json:"name,omitempty"`
  Alias string `json:"alias,omitempty"`
  Path string `json:"path,omitempty"`  
  ParentId *int `json:"parent_id,omitempty"`
  Type int `json:"type,omitempty"`
  Icon string `json:"icon,omitempty"`
  Url string `json:"url,omitempty"`  
  Components []map[string]interface{} `json:"components,omitempty"`  
  RedirectTo string `json:"redirect_to,omitempty"`  
  Data map[string]interface{} `json:"data,omitempty"`  
  Ref string `json:"ref,omitempty"`  
  Parent string `json:"parent,omitempty"`
}

func GetAngularRoutes(queryStringParams url.Values, user *User) (routes []AngularRoute){

	db := coreglobals.Db
  queryStr := `SELECT adm_route.id, adm_route.name, adm_route.alias, adm_route.path, adm_route.parent_id, adm_route.type, adm_route.icon, adm_route.url, adm_route.components, 
  ar2.alias as parent 
  FROM adm_route 
  LEFT JOIN adm_route as ar2 
  ON adm_route.parent_id = ar2.id`

  // if ?node-type=x&levels=x(,x..)
  // else if ?node-type=x
  // else if ?levels=x(,x..)
  if(queryStringParams.Get("type") != "" && queryStringParams.Get("levels") != ""){
      queryStr = queryStr + ` WHERE adm_route.type=` + queryStringParams.Get("type") + ` and adm_route.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
  } else if(queryStringParams.Get("type") != "" && queryStringParams.Get("levels")==""){
      queryStr = queryStr + ` WHERE adm_route.type=` + queryStringParams.Get("type")
  } else if(queryStringParams.Get("type") == "" && queryStringParams.Get("levels") != ""){
      queryStr = queryStr + ` WHERE adm_route.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
  }

  queryStr = queryStr + ` ORDER BY path,id asc`
  
  rows, err := db.Query(queryStr)
  corehelpers.PanicIf(err)
  defer rows.Close()

  var id, ttype int
  var path, name, alias string
  var components []byte

  var icon, url, parent sql.NullString
  var parent_id sql.NullInt64

  

  

  //var userId = strconv.Itoa(user.Id)


  for rows.Next(){

	  err := rows.Scan(&id, &name, &alias, &path, &parent_id, &ttype, &icon, &url, &components, &parent)
	  corehelpers.PanicIf(err)

	  var icon_str, url_str, parent_str string
	  var parent_id_int int

	  if(icon.Valid){
	  	icon_str=icon.String
	  } else {
	  	icon_str = ""
	  }

	  if(url.Valid){
	  	url_str=url.String
	  } else {
	  	url_str = ""
	  }

	  if(parent.Valid){
	  	parent_str=parent.String
	  } else {
	  	parent_str = ""
	  }

	  if(parent_id.Valid){
	  	parent_id_int = int(parent_id.Int64)
	  	} 
	  var componentsUm []map[string]interface{}
	  json.Unmarshal(components,&componentsUm)
	  if(user.UserGroups != nil){
	  	for i:=0; i< len(user.UserGroups); i++{
		  	for j:=0; j<len(user.UserGroups[i].AngularRouteIds); j++{
		  		if(id==user.UserGroups[i].AngularRouteIds[j]){
		  			route := AngularRoute{id,name, alias, path, &parent_id_int, ttype, icon_str, url_str, componentsUm, "", nil, "", parent_str}
		  			routes = append(routes, route)
		  		}
		  	}
		  }
	  }
	  
  }
  return
}