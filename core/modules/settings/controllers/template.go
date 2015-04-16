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
	"github.com/gorilla/mux"
)

type TemplateApiController struct{}

// func (this *TemplateApiController) PostTemplate(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     template := models.Template{}

//     err := json.NewDecoder(r.Body).Decode(&template)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     template.Post()

// }

// func (this *TemplateApiController) PutTemplate(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     template := models.Template{}

//     err := json.NewDecoder(r.Body).Decode(&template)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     template.Update()

// }

func (this *TemplateApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryStrParams := r.URL.Query()

	templates := models.GetTemplates(queryStrParams)

	res, err := json.Marshal(templates)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)

}

func (this *TemplateApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	//user := coremoduleuser.GetLoggedInUser(r)

	templates := models.GetTemplatesByIdChildren(id)

	res, err := json.Marshal(templates)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *TemplateApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]

	id, _ := strconv.Atoi(idStr)

	template := models.GetTemplateById(id)

	res, err := json.Marshal(template)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)

}
