package controllers

import
(
	"fmt"
	"collexy/core/api/models"
	"encoding/json"
	"net/http"
	"collexy/helpers"
    // "collexy/globals"
)

type RouteApiController struct {}

// func (this *RouteApiController) Get(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
    
//     res, err := json.Marshal(globals.Routes)
//     helpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }

func (this *RouteApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    routes := models.GetRoutes()
    res, err := json.Marshal(routes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}