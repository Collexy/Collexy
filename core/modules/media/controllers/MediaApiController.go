package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	_ "github.com/lib/pq"
	//"collexy/helpers"
	"collexy/core/modules/media/models"
	"strconv"
	//"log"
	//"github.com/gorilla/schema"
	"encoding/json"
	//"log"
	//"io/ioutil"
	//"path/filepath"
	//coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"github.com/gorilla/mux"
	//"html/template"
	//"strings"
	//"github.com/gorilla/context"
	coremodulemembermodels "collexy/core/modules/member/models"
	coremoduleuser "collexy/core/modules/user/models"
)

type MediaApiController struct{}

func (this *MediaApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryStrParams := r.URL.Query()

	user := coremoduleuser.GetLoggedInUser(r)

	media := models.GetMedia(queryStrParams, user)

	res, err := json.Marshal(media)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *MediaApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	//user := coremoduleuser.GetLoggedInUser(r)

	media := models.GetMediaById(id)

	res, err := json.Marshal(media)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *MediaApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	user := coremoduleuser.GetLoggedInUser(r)

	media := models.GetMediaByIdChildren(id, user)

	res, err := json.Marshal(media)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *MediaApiController) GetByIdParents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	user := coremoduleuser.GetLoggedInUser(r)

	media := models.GetMediaByIdParents(id, user)

	res, err := json.Marshal(media)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

type TestData struct {
	Data    *TestStruct
	HasUser bool
}

type TestStruct struct {
	User    *coremoduleuser.User
	Media *models.Media
}

type TestDataMember struct {
	Data      *TestStructMember
	HasMember bool
}

type TestStructMember struct {
	Member  *coremodulemembermodels.Member
	Media *models.Media
}



// func (this *MediaApiController) Post(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     if user := coremoduleuser.GetLoggedInUser(r); user != nil {
//         var hasPermission bool = false
//         hasPermission = user.HasPermissions([]string{"media_create"})
//         if(hasPermission){

//             media := models.Media{}

//             err := json.NewDecoder(r.Body).Decode(&media)

//             if err != nil {
//                 http.Error(w, "Bad Request", 400)
//             }

//             media.Post()
//         } else {
//             fmt.Fprintf(w,"You do not have permission to create media")
//         }
//     }
// }

func (this *MediaApiController) Put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if user := coremoduleuser.GetLoggedInUser(r); user != nil {
        var hasPermission bool = false
        hasPermission = user.HasPermissions([]string{"media_update"})
        if(hasPermission){
            media := models.Media{}

            err := json.NewDecoder(r.Body).Decode(&media)

            if err != nil {
                http.Error(w, "Bad Request", 400)
            }

            media.Update()
        } else {
            fmt.Fprintf(w,"You do not have permission to update media")
        }

    }
}

// func (this *MediaApiController) Delete(w http.ResponseWriter, r *http.Request){
//     w.Header().Set("Content-Type", "application/json")
//     if user := coremoduleuser.GetLoggedInUser(r); user != nil {
//         var hasPermission bool = false
//         hasPermission = user.HasPermissions([]string{"media_delete"})
//         if(hasPermission){
//             params := mux.Vars(r)

//             idStr := params["id"]
//             id, _ := strconv.Atoi(idStr)

//             models.DeleteMedia(id)
//         } else {
//             fmt.Fprintf(w,"You do not have permission to delete media")
//         }

//     }
// }



func (this *MediaApiController) GetBackendMediaById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]

	id, _ := strconv.Atoi(idStr)

	// Media object including the media' Node, the media type object.
	// Note: Inside the media type object is an array of parent MediaTypes
	media := models.GetBackendMediaById(id)

	res, err := json.Marshal(media)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}
