package controllers

import
(
	"fmt"
	"collexy/core/api/models"
	"encoding/json"
	"net/http"
	"collexy/helpers"
    "collexy/globals"
    "github.com/gorilla/mux"
)

type MenuApiController struct {}

func (this *MenuApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    res, err := json.Marshal(globals.Menus)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *MenuApiController) GetByName(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    name := params["name"]

    var menu models.AdminMenu

    for i := 0; i <= len(globals.Menus); i++ {
    	var temp models.AdminMenu = globals.Menus[i].(models.AdminMenu)
	    if(temp.Name == name){
	    	menu = temp
	    	res, err := json.Marshal(menu)
		    helpers.PanicIf(err)

		    fmt.Fprintf(w,"%s",res)
		    break
	    }
	}
    
    // res, err := json.Marshal(globals.Menus)
    // helpers.PanicIf(err)

    // fmt.Fprintf(w,"%s",res)
}