package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/member/models"
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
	coremoduleusermodels "collexy/core/modules/user/models"
	"github.com/gorilla/mux"
)

type MemberTypeApiController struct{}

func (this *MemberTypeApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_type_browse", "member_all"})
		if hasPermission {
			queryStrParams := r.URL.Query()
			memberTypes := models.GetMemberTypes(queryStrParams)
			res, err := json.Marshal(memberTypes)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to create member types")
		}
	}
}

func (this *MemberTypeApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_type_browse", "member_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]

			id, _ := strconv.Atoi(idStr)

			var extended bool = false
			extended, _ = strconv.ParseBool(r.URL.Query().Get("extended"))

			//extended, _ := strconv.Atoi(extendedStr)

			if !extended {
				memberType := models.GetMemberTypeById(id)
				res, err := json.Marshal(memberType)
				corehelpers.PanicIf(err)

				fmt.Fprintf(w, "%s", res)
			} else {
				memberType := models.GetMemberTypeExtendedById(id)
				res, err := json.Marshal(memberType)
				corehelpers.PanicIf(err)

				fmt.Fprintf(w, "%s", res)
			}
		} else {
			fmt.Fprintf(w, "You do not have permission to create member types")
		}
	}
}

func (this *MemberTypeApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_type_browse", "member_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			//user := coremoduleuser.GetLoggedInUser(r)

			memberTypes := models.GetMemberTypesByIdChildren(id)

			res, err := json.Marshal(memberTypes)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)

		} else {
			fmt.Fprintf(w, "You do not have permission to create member types")
		}
	}
}

func (this *MemberTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_type_create", "member_all"})
		if hasPermission {

			memberType := models.MemberType{}

			err := json.NewDecoder(r.Body).Decode(&memberType)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			memberType.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create member types")
		}
	}

}

func (this *MemberTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_type_update", "member_all"})
		if hasPermission {

			memberType := models.MemberType{}

			err := json.NewDecoder(r.Body).Decode(&memberType)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			memberType.Put()
		} else {
			fmt.Fprintf(w, "You do not have permission to update member types")
		}
	}
}

func (this *MemberTypeApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_type_delete", "member_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeleteMemberType(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete member types")
		}

	}
}
