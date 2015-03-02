package controllers

import
(
	"fmt"
	"collexy/core/api/models"
	"encoding/json"
	"net/http"
	corehelpers "collexy/core/helpers"
)

type AngularRouteApiController struct {}

func (this *AngularRouteApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    queryStrParams := r.URL.Query()

    user := models.GetLoggedInUser(r)

    angularRoutes := models.GetAngularRoutes(queryStrParams,user)

    res, err := json.Marshal(angularRoutes)
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}