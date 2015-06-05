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

type DataTypeApiController struct{}

func (this *DataTypeApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"data_type_browse", "data_type_all"})
		if hasPermission {

			dataTypes := models.GetDataTypes()

			res, err := json.Marshal(dataTypes)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse data types")
		}
	}

}

func (this *DataTypeApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"data_type_browse", "data_type_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]

			id, _ := strconv.Atoi(idStr)

			dataType := models.GetDataTypeById(id)

			res, err := json.Marshal(dataType)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse data types")
		}
	}

}

func (this *DataTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"data_type_create", "data_type_all"})
		if hasPermission {

			dataType := models.DataType{}

			err := json.NewDecoder(r.Body).Decode(&dataType)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			dataType.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create data types")
		}
	}

}

func (this *DataTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleuser.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"data_type_update", "data_type_all"})
		if hasPermission {

			dataType := models.DataType{}

			err := json.NewDecoder(r.Body).Decode(&dataType)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			dataType.Update()
		} else {
			fmt.Fprintf(w, "You do not have permission to update data types")
		}
	}
}
