package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/member/models"
	coremoduleusermodels "collexy/core/modules/user/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type MemberGroupApiController struct{}

func (this *MemberGroupApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_group_browse", "member_all"})
		if hasPermission {
			memberGroups := models.GetMemberGroups()
			res, err := json.Marshal(memberGroups)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse member groups")
		}
	}
}

func (this *MemberGroupApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_group_browse", "member_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]

			id, _ := strconv.Atoi(idStr)

			memberGroup := models.GetMemberGroupById(id)
			res, err := json.Marshal(memberGroup)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse member groups")
		}
	}
}

func (this *MemberGroupApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_group_create", "member_all"})
		if hasPermission {

			memberGroup := models.MemberGroup{}

			err := json.NewDecoder(r.Body).Decode(&memberGroup)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			memberGroup.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create member groups")
		}
	}

}

func (this *MemberGroupApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_group_update", "member_all"})
		if hasPermission {

			memberGroup := models.MemberGroup{}

			err := json.NewDecoder(r.Body).Decode(&memberGroup)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			memberGroup.Put()
		} else {
			fmt.Fprintf(w, "You do not have permission to update member groups")
		}
	}
}

func (this *MemberGroupApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_group_delete", "member_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeleteMemberGroup(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete member groups")
		}

	}
}
