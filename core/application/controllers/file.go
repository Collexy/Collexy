package controllers

import (
	"fmt"
	//"collexy/core/api/models"
	//corehelpers "collexy/core/helpers"
	//"encoding/json"
	"net/http"
	// "collexy/globals"
	//coreglobals "collexy/core/globals"
	//"github.com/gorilla/mux"
	"os"
)

type FileApiController struct{}

func (this *FileApiController) Delete(w http.ResponseWriter, r *http.Request) {

	queryStrParams := r.URL.Query()

	path := queryStrParams.Get("path")
	fileName := queryStrParams.Get("fileName")

	str := path + "\\" + fileName

	err := os.Remove(str)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Sprintf("file with path: %s, and filename: %s is successfully deleted", path, fileName)
}

// func (this *DataTypeEditorApiController) Get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	res, err := json.Marshal(coreglobals.DataTypeEditors)
// 	corehelpers.PanicIf(err)

// 	fmt.Fprintf(w, "%s", res)
// }

// func (this *DataTypeEditorApiController) GetByAlias(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	alias := params["alias"]

// 	for _, dte := range coreglobals.DataTypeEditors {
// 		if dte.Alias == alias {
// 			res, err := json.Marshal(dte)
// 			corehelpers.PanicIf(err)

// 			fmt.Fprintf(w, "%s", res)
// 		}
// 	}

// }
