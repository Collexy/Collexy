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
	"os"
	//"github.com/gorilla/schema"
	"encoding/xml"
	"encoding/json"
	"log"
	"io/ioutil"
	//"path/filepath"
	coreglobals "collexy/core/globals"
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
	User  *coremoduleuser.User
	Media *models.Media
}

type TestDataMember struct {
	Data      *TestStructMember
	HasMember bool
}

type TestStructMember struct {
	Member *coremodulemembermodels.Member
	Media  *models.Media
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
		if hasPermission {
			media := models.Media{}

			err := json.NewDecoder(r.Body).Decode(&media)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			media.Update()
		} else {
			fmt.Fprintf(w, "You do not have permission to update media")
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

	media := models.GetBackendMediaById(id,GetProtectedMedia(w,r,id))

	res, err := json.Marshal(media)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func buildMap(mySlice ...*coreglobals.MediaAccessItem) (myMap map[string]*coreglobals.MediaAccessItem) {
	myMap = make(map[string]*coreglobals.MediaAccessItem)
	for _, item := range mySlice {
	
		myMap[item.Url] = item
		itemKeyIdStr := strconv.Itoa(item.MediaId)
		myMap[itemKeyIdStr] = item
	}
	return
}

func GetProtectedMedia(w http.ResponseWriter, r *http.Request, id int) (protectedItem *coreglobals.MediaAccessItem) {
	fmt.Println(*r.URL)


	if _, err := os.Stat("./config/media-access.xml"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("media-access.xml config file does not exist")
		} else {
			// other error
		}
	} else {

		configFile, err1 := os.Open("./config/media-access.xml")
		defer configFile.Close()
		if err1 != nil {
			log.Println("Error opening media-access.xml config file")
			//printError("opening config file", err1.Error())
		}

		XMLdata, err2 := ioutil.ReadAll(configFile)

		fmt.Println(string(XMLdata))

		if err2 != nil {
			log.Println("Error reading from media-access.xml config file")
			fmt.Printf("error: %v", err2)
		}

		var v coreglobals.MediaAccessItems
		err := xml.Unmarshal(XMLdata, &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}


		

		//fmt.Printf("%#v\n", v)

		coreglobals.MediaAccessConf = buildMap(v.Items...)

		// urlStr = r.URL.Path

		// log.Println("urlStr: " + urlStr)
		// log.Println(coreglobals.MediaAccessConf[urlStr])

	}

	for _, value := range coreglobals.MediaAccessConf{
		if(value.MediaId == id){
			protectedItem = value
			// isProtected = true
			// break
			return
		}
	}

	return
}

// func GetProtectedMedia(w http.ResponseWriter, r *http.Request, id int) (protectedItem *coreglobals.MediaAccessItem) {
// 	fmt.Println(*r.URL)
// 	var urlStr string = ""

// 	if _, err := os.Stat("./config/media-access.json"); err != nil {
// 		if os.IsNotExist(err) {
// 			// file does not exist
// 			log.Println("media-access.json config file does not exist")
// 		} else {
// 			// other error
// 		}
// 	} else {

// 		configFile, err1 := os.Open("./config/media-access.json")
// 		defer configFile.Close()
// 		if err1 != nil {
// 			log.Println("Error opening media-access.json config file")
// 			//printError("opening config file", err1.Error())
// 		}

// 		jsonParser := json.NewDecoder(configFile)
// 		if err1 = jsonParser.Decode(&coreglobals.MediaAccessConf); err1 != nil {
// 			log.Println("Error parsing media-access.json config file")
// 			log.Println(err1.Error())
// 			//printError("parsing config file", err1.Error())
// 		}

// 		urlStr = r.URL.Path

// 		log.Println("urlStr: " + urlStr)
// 		log.Println(coreglobals.MediaAccessConf[urlStr])

// 	}

// 	for _, value := range coreglobals.MediaAccessConf{
// 		if(value.MediaId == id){
// 			protectedItem = value
// 			// isProtected = true
// 			// break
// 			return
// 		}
// 	}

// 	return
// }