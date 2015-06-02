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

type MediaTypeApiController struct{}

func (this *MediaTypeApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Media-Type", "application/json")

	queryStrParams := r.URL.Query()

	mediaTypes := models.GetMediaTypes(queryStrParams)
	res, err := json.Marshal(mediaTypes)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *MediaTypeApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Media-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	//user := coremoduleuser.GetLoggedInUser(r)

	mediaTypes := models.GetMediaTypesByIdChildren(id)

	res, err := json.Marshal(mediaTypes)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

// func (this *MediaTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Media-Type", "application/json")

//     mediaType := models.MediaType{}

//     err := json.NewDecoder(r.Body).Decode(&mediaType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     mediaType.Post()

// }

// func (this *MediaTypeApiController) GetMediaTypes(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Media-Type", "application/json")

//     mediaTypes := models.GetMediaTypes()

//     res, err := json.Marshal(mediaTypes)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)

// }

// func (this *MediaTypeApiController) GetMediaTypeExtendedById(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Media-Type", "application/json")

//     var idStr string = ""
//     idStr = r.URL.Query().Get(":id")

//     id, _ := strconv.Atoi(idStr)
//     media := models.GetMediaTypeExtendedById(id)
//     res, err := json.Marshal(media)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }

func (this *MediaTypeApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Media-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]

	id, _ := strconv.Atoi(idStr)

	var extended bool = false
	extended, _ = strconv.ParseBool(r.URL.Query().Get("extended"))

	//extended, _ := strconv.Atoi(extendedStr)

	if !extended {
		media := models.GetMediaTypeById(id)
		res, err := json.Marshal(media)
		corehelpers.PanicIf(err)

		fmt.Fprintf(w, "%s", res)
	} else {
		media := models.GetMediaTypeExtendedById(id)
		res, err := json.Marshal(media)
		corehelpers.PanicIf(err)

		fmt.Fprintf(w, "%s", res)
	}

}

// func (this *MediaTypeApiController) GetMediaTypeById(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Media-Type", "application/json")

//     var idStr string = ""
//     idStr = r.URL.Query().Get(":id")

//     if(len(idStr)>0){
//         fmt.Println("lol1")
//         id, _ := strconv.Atoi(idStr)
//         media := models.GetMediaTypeById(id)
//         res, err := json.Marshal(media)
//         corehelpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }else{
//         fmt.Println("lol2")
//         mediaTypes := models.GetMediaTypes()
//         res, err := json.Marshal(mediaTypes)
//         corehelpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }
// }

// func (this *MediaTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Media-Type", "application/json")

//     mediaType := models.MediaType{}

//     err := json.NewDecoder(r.Body).Decode(&mediaType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     // b, err := json.Marshal(mediaType)
//     // if err != nil {
//     //     fmt.Println(err)
//     //     return
//     // }

//     mediaType.Update()
// }
