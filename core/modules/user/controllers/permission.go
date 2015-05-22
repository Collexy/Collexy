package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/user/models"
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"net/http"
	//"strconv"
)

type PermissionApiController struct{}

func (this *PermissionApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	permissions := models.GetPermissions()
	res, err := json.Marshal(permissions)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

// func (this *PermissionApiController) GetById(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	idStr := params["id"]
// 	id, _ := strconv.Atoi(idStr)

// 	user := models.GetLoggedInUser(r)

// 	userGroup := models.GetPermissionById(id, user)

// 	res, err := json.Marshal(userGroup)
// 	corehelpers.PanicIf(err)

// 	fmt.Fprintf(w, "%s", res)
// }
