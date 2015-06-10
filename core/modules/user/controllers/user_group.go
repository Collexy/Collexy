package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/user/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserGroupApiController struct{}

func (this *UserGroupApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_group_create", "user_group_all"})
		if hasPermission {

			user := models.GetLoggedInUser(r)

			userGroups := models.GetUserGroups(user)

			res, err := json.Marshal(userGroups)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse user groups")
		}
	}	
}

func (this *UserGroupApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_group_create", "user_group_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			user := models.GetLoggedInUser(r)

			userGroup := models.GetUserGroupById(id, user)

			res, err := json.Marshal(userGroup)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse user groups")
		}
	}
}

func (this *UserGroupApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_group_create", "user_group_all"})
		if hasPermission {

			userGroup := models.UserGroup{}

			err := json.NewDecoder(r.Body).Decode(&userGroup)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			userGroup.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create user groups")
		}
	}

}

func (this *UserGroupApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_group_update", "user_group_all"})
		if hasPermission {

			userGroup := models.UserGroup{}

			err := json.NewDecoder(r.Body).Decode(&userGroup)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			userGroup.Put()
		} else {
			fmt.Fprintf(w, "You do not have permission to update user groups")
		}
	}
}

func (this *UserGroupApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_group_delete", "user_group_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeleteUserGroup(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete user groups")
		}

	}
}