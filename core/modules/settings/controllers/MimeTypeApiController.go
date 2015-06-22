package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/settings/models"
	coremoduleuser "collexy/core/modules/user/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MimeTypeApiController struct{}

func (this *MimeTypeApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"mime_type_browse", "mime_type_all"})
		if hasPermission {

			mimeTypes := models.GetMimeTypes()

			res, err := json.Marshal(mimeTypes)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse mime types")
		}
	}

}

func (this *MimeTypeApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"mime_type_browse", "mime_type_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]

			id, _ := strconv.Atoi(idStr)

			mimeType := models.GetMimeTypeById(id)

			res, err := json.Marshal(mimeType)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse mime types")
		}
	}

}

func (this *MimeTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"mime_type_update", "mime_type_all"})
		if hasPermission {

			mimeType := models.MimeType{}

			err := json.NewDecoder(r.Body).Decode(&mimeType)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			mimeType.Update()
		} else {
			fmt.Fprintf(w, "You do not have permission to update mime types")
		}
	}
}

func (this *MimeTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"mime_type_create", "mime_type_all"})
		if hasPermission {

			mimeType := models.MimeType{}

			err := json.NewDecoder(r.Body).Decode(&mimeType)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			mimeType.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create mime types")
		}
	}

}
