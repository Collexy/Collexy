package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/settings/models"
	_ "github.com/lib/pq"
	"strconv"
	//"log"
	//"github.com/gorilla/schema"
	"encoding/json"
	//"log"
	//"io/ioutil"
	//"path/filepath"
	//"strings"
	//"html/template"
	coremoduleuser "collexy/core/modules/user/models"
	"github.com/gorilla/mux"
)

type TemplateApiController struct{}

func (this *TemplateApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"template_browse", "template_all"})
		if hasPermission {

			queryStrParams := r.URL.Query()

			templates := models.GetTemplates(queryStrParams)

			res, err := json.Marshal(templates)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse templates")
		}
	}

}

func (this *TemplateApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"template_browse", "template_all"})
		if hasPermission {
			params := mux.Vars(r)
			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			//user := coremoduleuser.GetLoggedInUser(r)

			templates := models.GetTemplatesByIdChildren(id)

			res, err := json.Marshal(templates)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse templates")
		}
	}
}

func (this *TemplateApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"template_browse", "template_all"})
		if hasPermission {
			params := mux.Vars(r)
			idStr := params["id"]

			id, _ := strconv.Atoi(idStr)

			template := models.GetTemplateById(id)

			res, err := json.Marshal(template)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse templates")
		}
	}

}

func (this *TemplateApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"template_create", "template_all"})
		if hasPermission {

			template := models.Template{}

			err := json.NewDecoder(r.Body).Decode(&template)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			template.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create templates")
		}
	}

}

func (this *TemplateApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"template_update", "template_all"})
		if hasPermission {

			template := models.Template{}

			err := json.NewDecoder(r.Body).Decode(&template)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			template.Update()

		} else {
			fmt.Fprintf(w, "You do not have permission to update templates")
		}
	}
}

func (this *TemplateApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"template_delete", "template_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeleteTemplate(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete templates")
		}

	}
}
