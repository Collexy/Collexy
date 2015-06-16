package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/user/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PermissionApiController struct{}

func (this *PermissionApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	permissions := models.GetPermissions()
	res, err := json.Marshal(permissions)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *PermissionApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"permission_browse", "permission_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			permission := models.GetPermissionById(id)

			res, err := json.Marshal(permission)
			corehelpers.PanicIf(err)
			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse permissions")
		}
	}

}

func (this *PermissionApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"permission_create", "permission_all"})
		if hasPermission {

			permission := models.Permission{}

			err := json.NewDecoder(r.Body).Decode(&permission)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			permission.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create permissions")
		}
	}

}

func (this *PermissionApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"permission_update", "permission_all"})
		if hasPermission {

			permission := models.Permission{}

			err := json.NewDecoder(r.Body).Decode(&permission)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			permission.Update()
		} else {
			fmt.Fprintf(w, "You do not have permission to update permissions")
		}
	}
}

func (this *PermissionApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"permission_delete", "permission_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeletePermission(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete permissions")
		}

	}
}
