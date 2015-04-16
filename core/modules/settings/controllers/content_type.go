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

type ContentTypeApiController struct{}

func (this *ContentTypeApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryStrParams := r.URL.Query()

	contentTypes := models.GetContentTypes(queryStrParams)
	res, err := json.Marshal(contentTypes)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *ContentTypeApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	//user := coremoduleuser.GetLoggedInUser(r)

	contentTypes := models.GetContentTypesByIdChildren(id)

	res, err := json.Marshal(contentTypes)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

// func (this *ContentTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     contentType := models.ContentType{}

//     err := json.NewDecoder(r.Body).Decode(&contentType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     contentType.Post()

// }

// func (this *ContentTypeApiController) GetContentTypes(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     contentTypes := models.GetContentTypes()

//     res, err := json.Marshal(contentTypes)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)

// }

// func (this *ContentTypeApiController) GetContentTypeExtendedById(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     var idStr string = ""
//     idStr = r.URL.Query().Get(":id")

//     id, _ := strconv.Atoi(idStr)
//     content := models.GetContentTypeExtendedById(id)
//     res, err := json.Marshal(content)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }

func (this *ContentTypeApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]

	id, _ := strconv.Atoi(idStr)

	var extended bool = false
	extended, _ = strconv.ParseBool(r.URL.Query().Get("extended"))

	//extended, _ := strconv.Atoi(extendedStr)

	if !extended {
		content := models.GetContentTypeById(id)
		res, err := json.Marshal(content)
		corehelpers.PanicIf(err)

		fmt.Fprintf(w, "%s", res)
	} else {
		content := models.GetContentTypeExtendedById(id)
		res, err := json.Marshal(content)
		corehelpers.PanicIf(err)

		fmt.Fprintf(w, "%s", res)
	}

}

// func (this *ContentTypeApiController) GetContentTypeById(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     var idStr string = ""
//     idStr = r.URL.Query().Get(":id")

//     if(len(idStr)>0){
//         fmt.Println("lol1")
//         id, _ := strconv.Atoi(idStr)
//         content := models.GetContentTypeById(id)
//         res, err := json.Marshal(content)
//         corehelpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }else{
//         fmt.Println("lol2")
//         contentTypes := models.GetContentTypes()
//         res, err := json.Marshal(contentTypes)
//         corehelpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }
// }

// func (this *ContentTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     contentType := models.ContentType{}

//     err := json.NewDecoder(r.Body).Decode(&contentType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     // b, err := json.Marshal(contentType)
//     // if err != nil {
//     //     fmt.Println(err)
//     //     return
//     // }

//     contentType.Update()
// }
